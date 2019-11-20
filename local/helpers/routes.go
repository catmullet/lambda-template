package helpers

import (
	"fmt"
	"github.com/kyani-inc/kms-lambda-template/lambda/example/handlers/example"
	"github.com/kyani-inc/kms-lambda-template/types/aws"
	"io/ioutil"
	"net/http"
)

func Example(w http.ResponseWriter, r *http.Request) {
	resp, _ := example.HandleRequest(ConvertHttpRequestToLambdaRequest(r))
	ConvertLambdaResponseToHttpResponse(resp, w)
	return
}

func ConvertHttpRequestToLambdaRequest(r *http.Request) (req aws.LambdaRequest) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	req.Body = string(body)
	req.Headers = make(map[string]string)

	for name, values := range r.Header {
		for _, value := range values {
			req.Headers[name] = value
		}
	}

	req.IsBase64Encoded = false
	req.Path = r.URL.Path
	return
}

func ConvertLambdaResponseToHttpResponse(res aws.LambdaResponse, w http.ResponseWriter) {
	for name, value := range res.Headers {
		w.Header().Set(name, value)
	}

	if res.StatusCode >= 200 {
		w.WriteHeader(res.StatusCode)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	_, _ = fmt.Fprintf(w, res.Body)
	return
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
