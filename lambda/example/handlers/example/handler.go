package example

import "github.com/kyani-inc/kms-lambda-template/types/aws"

func HandleRequest(req aws.LambdaRequest) (resp aws.LambdaResponse, err error) {

	resp.AddHeader("Content-Type", "application/json")

	request := struct {
		Message string
	}{}

	// Bind body to MakePaymentRequest
	err = req.Bind(&request)

	if err != nil {
		return
	}

	response := struct {
		Message string
	}{"Hello, you said " + request.Message}

	err = resp.SetBody(response)

	resp.SetHttpStatusCode(200)
	return
}
