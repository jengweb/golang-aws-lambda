package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type MockedPutItem struct {
	dynamodbiface.DynamoDBAPI
	Response dynamodb.PutItemOutput
}

func (d MockedPutItem) PutItem(in *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return &d.Response, nil
}

func TestHandler(t *testing.T) {
	t.Run("Successful Request", func(t *testing.T) {
		m := MockedPutItem{
			Response: dynamodb.PutItemOutput{},
		}

		d := deps{
			ddb:   m,
			table: "test_table",
		}

		_, err := d.handler(events.APIGatewayProxyRequest{})
		if err != nil {
			t.Fatalf("Everythink should be ok")
		}
	})
}
