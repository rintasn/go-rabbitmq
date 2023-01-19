package main

import (
	"encoding/json"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"main/configAccount"
	util_rabbitmq "main/util-rabbitmq"
)

type User struct {
	IDACCOUNT      string `json:"id_account"`
	ID_USER        string `json:"iduser"`
	SALDO          string `default:"0"`
	STATUS_ACCOUNT string `default:"1"`
}

func main() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	util_rabbitmq.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	util_rabbitmq.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"accounts", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	util_rabbitmq.FailOnError(err, "Failed to declare a queue")

	messages, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	util_rabbitmq.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for data := range messages {
			db := configAccount.CreateConnection()
			log.Printf("Received a message: %s", data.Body)
			var jsonData = []byte(data.Body)

			var data User

			var err = json.Unmarshal(jsonData, &data)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			//=====================================//
			defer db.Close()
			sqlStatement := `insert into accounts (iduser,saldo,status_account) values ($1,$2,$3) returning id_account`

			var id int64

			err2 := db.QueryRow(sqlStatement, data.ID_USER, 0, 1).Scan(&id)

			if err2 != nil {
				log.Fatalf("Tidak bisa akses query. %v", err2)
			}
			fmt.Printf("Insert data single record %v", id)

			//=====================================================//
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
