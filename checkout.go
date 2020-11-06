package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

type CheckoutData struct {
	ClientSecret string `json:"client_secret"`
}

func init() {
	stripe.Key = "vApyWmKKftXTMEMUZNOjhWqH0sj413ku"
}

func checkout(c *gin.Context) {

	values := c.Request.URL.RawQuery
	pinsQ := strings.Split(values, ",")

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(1200 * len(pinsQ))),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
	}

	// Verify your integration in this guide by including this parameter
	params.AddMetadata("integration_check", "accept_a_payment")

	pi, _ := paymentintent.New(params)

	c.HTML(http.StatusOK, "checkout.html", gin.H{
		"client_secret": pi.ClientSecret,
	})
}

func paymentFinished(c *gin.Context) {
	fmt.Println("Payment intent seems to be ok. pi = " + c.Param("pi"))
	//wait for confirmation on webhooks!
}
