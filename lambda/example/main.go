package example

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kyani-inc/kms-lambda-template/lambda/example/handlers/example"
)

func main() {
	lambda.Start(example.HandleRequest)
}
