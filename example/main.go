package main

import (
	"fmt"
	"log"
	"os"

	"github.com/LUSHDigital/core-mage/env"
	coresql "github.com/LUSHDigital/core-sql"
	"github.com/streadway/amqp"

	_ "github.com/lib/pq"
)

func main() {
	// Because of a change in go 1.13 of `go list`, mage is broken if this file doesn't exist.
	// We need to keep a regular go file without build tags in the same directory as the magefile.
	// TODO: Remove this file after this issue is fixed: https://github.com/magefile/mage/issues/262

	env.LoadDefault()

	if err := checkRabbit(os.Getenv("RABBITMQ_URL")); err != nil {
		log.Fatal(err)
	}

	if err := checkCockroach(os.Getenv("COCKROACH_URL")); err != nil {
		log.Fatal(err)
	}
}

func checkRabbit(url string) error {
	log.Printf("testing RabbitMQ connection: %q", url)

	cn, err := amqp.Dial(url)
	if err != nil {
		return fmt.Errorf("error creating RabbitMQ connection: %w", err)
	}
	defer cn.Close()

	ch, err := cn.Channel()
	if err != nil {
		return fmt.Errorf("error creating RabbitMQ channel: %w", err)
	}
	defer ch.Close()

	return nil
}

func checkCockroach(url string) error {
	log.Printf("testing CockroachDB connection: %q", url)

	db, err := coresql.Open("postgres", url)
	if err != nil {
		return fmt.Errorf("error creating CockroachDB connection: %w", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		return fmt.Errorf("error pinging CockroachDB connection: %w", err)
	}

	return nil
}
