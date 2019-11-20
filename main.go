package main

import (
	"fmt"
	"github.com/kyani-inc/kms-lambda-template/local/helpers"
	"github.com/kyani-inc/kms-lambda-template/local/helpers/environments"
	"net/http"
	"os"
)

// Only for testing locally.
func main() {
	fmt.Println("Starting...")
	environments.LoadEnvironmentVariablesFromYml(environments.Staging)

	http.HandleFunc("/example", helpers.Example)

	fmt.Println(fmt.Sprintf("Listening on %s", os.Getenv("PORT")))
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
}
