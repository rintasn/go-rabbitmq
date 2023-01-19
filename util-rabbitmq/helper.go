package util_rabbitmq

import "log"

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %v", msg, err)
	}
}
