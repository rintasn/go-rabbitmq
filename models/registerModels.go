package models

type User struct {
	IDUSER    string `json:"iduser"`
	NAME_USER string `json:"name_user"`
	EMAIL     string `json:"email"`
	PHONE     string `json:"phone"`
	STATUS    string `json:"status"`
	USERNAME  string `json:"username"`
	PASSWORD  string `json:"password"`
}

type Account struct {
	ID_USER        string `json:"id_user"`
	SALDO          string `default:"0"`
	STATUS_ACCOUNT string `default:"1"`
}

//
//func RegisterAccountModels(user User) int64 {
//	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
//	util_rabbitmq.FailOnError(err, "Failed to connect to RabbitMQ")
//	defer conn.Close()
//
//	ch, err := conn.Channel()
//	util_rabbitmq.FailOnError(err, "Failed to open a channel")
//	defer ch.Close()
//
//	q, err := ch.QueueDeclare(
//		"register", // name
//		false,      // durable
//		false,      // delete when unused
//		false,      // exclusive
//		false,      // no-wait
//		nil,        // arguments
//	)
//	util_rabbitmq.FailOnError(err, "Failed to declare a queue")
//
//	messages, err := ch.Consume(
//		q.Name, // queue
//		"",     // consumer
//		true,   // auto-ack
//		false,  // exclusive
//		false,  // no-local
//		false,  // no-wait
//		nil,    // args
//	)
//	util_rabbitmq.FailOnError(err, "Failed to register a consumer")
//
//	var forever chan struct{}
//
//	go func() {
//		for data := range messages {
//			db := config.CreateConnection()
//			log.Printf("Received a message: %s", data.Body)
//			var jsonData = []byte(data.Body)
//
//			var data User
//
//			var err = json.Unmarshal(jsonData, &data)
//			if err != nil {
//				fmt.Println(err.Error())
//				return
//			}
//
//			//=====================================//
//			defer db.Close()
//			sqlStatement := `insert into tbusers (iduser,name_user,email,phone,status) values ($1,$2,$3,$4,$5) returning iduser`
//
//			var id int64
//
//			err2 := db.QueryRow(sqlStatement, data.IDUSER, data.NAME_USER, data.EMAIL, data.PHONE, data.STATUS).Scan(&id)
//			if err2 != nil {
//				log.Fatalf("Tidak bisa akses query. %v", err)
//			}
//			fmt.Printf("Insert data single record %v", id)
//		}
//	}()
//	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
//	<-forever
//
//	//return id
//	return 1
//}
