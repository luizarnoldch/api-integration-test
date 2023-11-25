package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.DynamoDBEvent) (error) {
	eventJSON, err := json.Marshal(event)
	if err != nil {
		log.Fatalf("Error al convertir el objeto event a JSON: %v", err)
		return err
	}
	fmt.Println("EVENT---")
	fmt.Println(event)
	fmt.Println("----------------------------------------------")
	fmt.Println("Event JSON:", string(eventJSON))
	fmt.Println("----------------------------------------------")
	
	return nil
}

func main() {
	lambda.Start(handler)
}

