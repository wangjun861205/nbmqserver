package main

import (
	"log"
	"nbmq"
)

func main() {
	l, err := nbmq.NewListener("127.0.0.1:9999")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Stop()
	l.Run()
}
