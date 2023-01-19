package controller

import (
	"context"
	"encoding/json"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"main/models"
	util_rabbitmq "main/util-rabbitmq"
	"net/http"
	"time"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type Response struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []models.User `json:"data"`
}

func RegisterAccount(w http.ResponseWriter, r *http.Request) {
	var userStruct models.User

	// decode data json request
	err := json.NewDecoder(r.Body).Decode(&userStruct)

	if err != nil {
		log.Fatalf("Tidak bisa mendecode dari request body.  %v", err)
	}

	// panggil modelsnya lalu insert
	//insertID := models.RegisterAccountModels(userStruct)

	// format response objectnya
	//res := response{
	//	ID:      insertID,
	//	Message: "Data telah ditambahkan",
	//}

	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	util_rabbitmq.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	util_rabbitmq.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"users", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	u, err2 := ch.QueueDeclare(
		"accounts", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	util_rabbitmq.FailOnError(err, "Failed to declare a queue")
	util_rabbitmq.FailOnError(err2, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dataJson, _ := json.Marshal(userStruct)

	body := (dataJson)
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	util_rabbitmq.FailOnError(err, "Failed to publish a message")

	err2 = ch.PublishWithContext(ctx,
		"",     // exchange
		u.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	util_rabbitmq.FailOnError(err2, "Failed to publish a message")

	log.Printf(" [x] Sent %s\n", body)

	// kirim response
	json.NewEncoder(w).Encode(body)
}
