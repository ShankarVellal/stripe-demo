package main

import (
  "encoding/json"
  "net/http"

  stripe "github.com/stripe/stripe-go/v72"
  "github.com/stripe/stripe-go/v72/paymentintent"
)

type CheckoutData struct {
  ClientSecret string `json:"client_secret"`
}

func main() {
  http.HandleFunc("/secret", func(w http.ResponseWriter, r *http.Request) {
    stripe.Key = "sk_test_51HyURQAgpXBKSE2oh2iOeS0hBsMuo3GHgTaqT5PtAy40AYd6DNt1psFlLWcnPMDAEYzZOULWGPjmVrqW1XcUxx8B00xxjF4LMm"

  	params := &stripe.PaymentIntentParams{
    		Amount: stripe.Int64(1099),
    		Currency: stripe.String(string(stripe.CurrencyUSD)),
  	}
  	// Verify your integration in this guide by including this parameter
  	params.AddMetadata("integration_check", "accept_a_payment")

  	intent, _ := paymentintent.New(params)
	
    data := CheckoutData{
      ClientSecret: intent.ClientSecret,
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(data)
  })

  http.ListenAndServe(":8080", nil)
}