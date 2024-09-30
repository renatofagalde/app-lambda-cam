package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler é o método que será chamado pela Lambda
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Obter o diretório atual
	currentDir, err := os.Getwd()
	if err != nil {
		// Caso ocorra algum erro ao obter o diretório
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf("Erro ao obter o diretório corrente: %s", err.Error()),
		}, nil
	}

	// Criar a resposta com o nome do diretório
	response := fmt.Sprintf("O diretório corrente é: %s", filepath.Base(currentDir))

	// Retornar a resposta para o API Gateway
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       response,
	}, nil
}

func main() {
	// Iniciar o handler da Lambda
	lambda.Start(Handler)
}
