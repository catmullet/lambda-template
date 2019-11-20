package main

import (
	"github.com/kyani-inc/cms/payment_service"
	"github.com/kyani-inc/cms/types"
	"github.com/kyani-inc/kms-api-payments/app/types/aws"
	"github.com/kyani-inc/kms-api-payments/app/types/cms_types"
	"github.com/kyani-inc/kms-api-payments/lambda/cms/handlers/cms"
	"github.com/kyani-inc/kms-api-payments/lambda/cms/handlers/token"
	"github.com/kyani-inc/kms-api-payments/local/helpers/environments"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	environments.LoadEnvironmentVariablesFromYml(environments.Staging)
	code := m.Run()

	os.Exit(code)
}

var Token *string

func TestCmsToken(t *testing.T) {
	card := types.Card{}
	card.EncryptedNumber = `bGiokMefG81j/iR1PG1ery7+GvDeAAHQ6fabPS4UGm9DDDUnrJJozQWhHr8iDVk8du+kDo2oSXchyCNz6jObIVQvVMzs4dgp14ARyNQxv1Gvjv7t/l7XiGLuHVSQ4Wb2aTOd3Ziq30D758pKkLWz605J17PxLDiljT6iBO9k+X6V5vcX/rtm2sxzlGn9m7xMvCJJw5Tp8upTfL1GTlHogXxnZvUSrPRNdqvUA7iRt+Lw3/ERwheVZQEBUvFZtfCd4proMK6vriEJGPVUr/hjt58si4EP85Sue8AXHf6Pyvx9mbA2KkIFDqCoMAl/nAVqULI4KR9CO+vOkfnY0qb8vw==`
	card.CardHolderName = "Kyani Test"
	card.ExpirationMonth = "02"
	card.ExpirationYear = "2023"
	card.SecurityCode = "123"
	card.LastFour = "1111"

	req := aws.LambdaRequest{}

	req.HTTPMethod = http.MethodPost
	req.BindToRequestBody(card)
	req.SetJsonTypeToRequest()

	resp := payment_service.SaveCardResponse{}

	res, err := token.HandleRequest(req)

	if err != nil {
		t.Error(err.Error())
	}

	err = res.Bind(&resp)

	if err != nil {
		t.Error(err.Error())
	}

	Token = &resp.Token.Token
}

func TestOneCmsPaymentTestCard(t *testing.T) {
	success := t.Run("TestCmsToken", TestCmsToken)

	if !success {
		t.Fail()
	}

	payReq := cms_types.MakePaymentRequest{
		OnlineOrderID: 123456789,
		Currency:      "USD",
		Tokens: []cms_types.Token{
			{
				Token:               "1f61c59a-e616-4c85-8d16-bc6cd1a5e724",
				CardExpirationMonth: "1",
				CardExpirationYear:  "20",
				LastFour:            "1111",
				PayAmount:           35,
			},
		},
		Merchant:     "100176",
		CustomFields: nil,
		System:       "join",
	}

	req := aws.LambdaRequest{}

	req.HTTPMethod = http.MethodPost
	req.BindToRequestBody(payReq)
	req.SetJsonTypeToRequest()

	resp, err := cms.HandleRequest(req)

	if err != nil {
		t.Error(err)
	}

	res := cms_types.MakePaymentResponse{}

	resp.Bind(&res)

	log.Println(res.Result)

	if res.Result != "Approved" {
		log.Println("Failed Payment")
		t.Fail()
	}
}

func TestMultipleCmsPaymentTestCards(t *testing.T) {
	success := t.Run("TestCmsToken", TestCmsToken)

	if !success {
		t.Fail()
	}

	payReq := cms_types.MakePaymentRequest{
		OnlineOrderID: 123456789,
		Currency:      "USD",
		Tokens: []cms_types.Token{
			{
				Token:               "1f61c59a-e616-4c85-8d16-bc6cd1a5e724",
				CardExpirationMonth: "1",
				CardExpirationYear:  "20",
				LastFour:            "1111",
				PayAmount:           15,
			},
			{
				Token:               "1f61c59a-e616-4c85-8d16-bc6cd1a5e724",
				CardExpirationMonth: "1",
				CardExpirationYear:  "20",
				LastFour:            "1111",
				PayAmount:           20,
			},
		},
		Merchant:     "100176",
		CustomFields: nil,
		System:       "join",
	}

	req := aws.LambdaRequest{}

	req.HTTPMethod = http.MethodPost
	req.BindToRequestBody(payReq)
	req.SetJsonTypeToRequest()

	resp, err := cms.HandleRequest(req)

	if err != nil {
		t.Error(err)
	}

	res := cms_types.MakePaymentResponse{}

	resp.Bind(&res)

	log.Println(res.Result)

	if res.Result != "Approved" {
		log.Println("Failed Payment: Should be Approved")
		t.Fail()
	}
}

func TestOneCmsPaymentTestBadCard(t *testing.T) {
	success := t.Run("TestCmsToken", TestCmsToken)

	if !success {
		t.Fail()
	}

	payReq := cms_types.MakePaymentRequest{
		OnlineOrderID: 123456789,
		Currency:      "USD",
		Tokens: []cms_types.Token{
			{
				Token:               "1f61c59a-e616-4c85-8d16-bc6cd1a5e72b",
				CardExpirationMonth: "1",
				CardExpirationYear:  "20",
				LastFour:            "1111",
				PayAmount:           35,
			},
		},
		Merchant:     "100176",
		CustomFields: nil,
		System:       "join",
	}

	req := aws.LambdaRequest{}

	req.HTTPMethod = http.MethodPost
	req.BindToRequestBody(payReq)
	req.SetJsonTypeToRequest()

	resp, err := cms.HandleRequest(req)

	if err != nil {
		t.Error(err)
	}

	res := cms_types.MakePaymentResponse{}

	resp.Bind(&res)

	log.Println(res.Result)

	if res.Result != "Failed" {
		log.Println("Approved Payment: Should be Failed")
		t.Fail()
	}
}

func TestMultipleCmsPaymentTestBadCard(t *testing.T) {
	success := t.Run("TestCmsToken", TestCmsToken)

	if !success {
		t.Fail()
	}

	payReq := cms_types.MakePaymentRequest{
		OnlineOrderID: 123456789,
		Currency:      "USD",
		Tokens: []cms_types.Token{
			{
				Token:               "1f61c59a-e616-4c85-8d16-bc6cd1a5e724",
				CardExpirationMonth: "1",
				CardExpirationYear:  "20",
				LastFour:            "1111",
				PayAmount:           15,
			},
			{
				Token:               "1f61c59a-e616-4c85-8d16-bc6cd1a5e72b",
				CardExpirationMonth: "1",
				CardExpirationYear:  "20",
				LastFour:            "1111",
				PayAmount:           20,
			},
		},
		Merchant:     "100176",
		CustomFields: nil,
		System:       "join",
	}

	req := aws.LambdaRequest{}

	req.HTTPMethod = http.MethodPost
	req.BindToRequestBody(payReq)
	req.SetJsonTypeToRequest()

	resp, err := cms.HandleRequest(req)

	if err != nil {
		t.Error(err)
	}

	res := cms_types.MakePaymentResponse{}

	resp.Bind(&res)

	log.Println(res.Result)

	if res.Result != "Failed" {
		log.Println("Approved Payment: Should be Failed")
		t.Fail()
	}
}
