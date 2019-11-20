package example

import (
	"encoding/json"
	"github.com/kyani-inc/kms-lambda-template/types/aws"
	"testing"
)

func TestHandleRequest(t *testing.T) {

	request := struct {
		Message string
	}{"Hi Billy!"}

	body, _ := json.Marshal(request)

	req := aws.LambdaRequest{
		Body:    string(body),
		Headers: map[string]string{"Content-Type": "application/json"},
	}

	resp, err := HandleRequest(req)

	if err != nil {
		t.Error(err.Error())
	}

	if resp.StatusCode == 500 || resp.StatusCode == 400 {
		t.Fail()
	}
}
