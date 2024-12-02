package main

import (
	"context"
	"log"

	"github.com/A-Victory/Go-Micro/logger/data"
)

type RPCServer struct{}

type RPCPayload struct {
	Name string
	Data string
}

func (r *RPCServer) LogInfo(payload *RPCPayload, resp *string) error {
	collection := client.Database("logs").Collection("logs")

	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name: payload.Name,
		Data: payload.Data,
	})
	if err != nil {
		log.Println("Error communicating with mongo: ", err)
		return err
	}

	*resp = "Proccessed payload via RPC: " + payload.Name
	return nil
}
