package aws

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type LambdaResponse events.APIGatewayProxyResponse
type LambdaRequest events.APIGatewayProxyRequest

func (lr *LambdaRequest) GetBody() (body []byte) {
	return []byte(lr.Body)
}

func (lr *LambdaRequest) Bind(obj interface{}) (err error) {
	err = json.Unmarshal([]byte(lr.Body), &obj)
	return
}

func (lr *LambdaRequest) BindToRequestBody(obj interface{}) (err error) {
	body, err := json.Marshal(obj)
	lr.Body = string(body)
	return
}

func (lr *LambdaRequest) SetJsonTypeToRequest() {
	if len(lr.Headers) == 0 {
		lr.Headers = make(map[string]string)
	}

	lr.Headers["Content-Type"] = "application/json"
	return
}

func (lr *LambdaResponse) Bind(obj interface{}) (err error) {
	err = json.Unmarshal([]byte(lr.Body), &obj)
	return
}

func (lr *LambdaResponse) SetBody(body interface{}) (err error) {
	bodyBytes, err := json.Marshal(body)
	lr.Body = string(bodyBytes)
	return
}

func (lr *LambdaResponse) AddHeader(key, value string) {
	if len(lr.Headers) == 0 {
		lr.Headers = make(map[string]string)
	}

	lr.Headers[key] = value
}

func (lr *LambdaResponse) SetHttpStatusCode(code int) {
	lr.StatusCode = code
}

func (lr *LambdaResponse) SetIsBase64Encoded(isBase bool) {
	lr.IsBase64Encoded = isBase
}

func (lr *LambdaResponse) Error(status int, err error) (LambdaResponse, error) {
	lr.SetHttpStatusCode(status)
	lr.AddHeader("Content-Type", "application/json")
	lr.SetBody(map[string]string{
		"error": err.Error(),
	})
	return *lr, nil
}
