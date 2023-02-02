package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main(){
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil{
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil{
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"paid-notif",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil{
		panic(err)
	}

	msg, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil{
		panic(err)
	}

	message := make(chan string)

	go func(){
		for d := range msg{
			message <- string(d.Body)
		}
	}()
	fmt.Println("waiting for message")
	email := <-message

	sendMail(email)
}