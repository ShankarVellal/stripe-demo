package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "os"

  stripe "github.com/stripe/stripe-go/v72"
  "github.com/stripe/stripe-go/v72/paymentintent"
)

type CheckoutData struct {
  ClientSecret string `json:"client_secret"`
}

func main() {
  stripe.Key = "sk_test_51HyURQAgpXBKSE2oh2iOeS0hBsMuo3GHgTaqT5PtAy40AYd6DNt1psFlLWcnPMDAEYzZOULWGPjmVrqW1XcUxx8B00xxjF4LMm"

  http.HandleFunc("/secret", func(w http.ResponseWriter, r *http.Request) {
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

  http.HandleFunc("/webhook", func(w http.ResponseWriter, req *http.Request) {
    const MaxBodyBytes = int64(65536)
    req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)
    payload, err := ioutil.ReadAll(req.Body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
        w.WriteHeader(http.StatusServiceUnavailable)
        return
    }

    event := stripe.Event{}

    if err := json.Unmarshal(payload, &event); err != nil {
        fmt.Fprintf(os.Stderr, "Failed to parse webhook body json: %v\n", err.Error())
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // Unmarshal the event data into an appropriate struct depending on its Type
    switch event.Type {
    case "payment_intent.succeeded":
        var paymentIntent stripe.PaymentIntent
        err := json.Unmarshal(event.Data.Raw, &paymentIntent)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
            w.WriteHeader(http.StatusBadRequest)
            return
        }
        fmt.Println("PaymentIntent was successful!")
    case "payment_method.attached":
        var paymentMethod stripe.PaymentMethod
        err := json.Unmarshal(event.Data.Raw, &paymentMethod)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
            w.WriteHeader(http.StatusBadRequest)
            return
        }
        fmt.Println("PaymentMethod was attached to a Customer!")
    // ... handle other event types
    default:
        fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
    }

    w.WriteHeader(http.StatusOK)
  })


  http.ListenAndServe(":8080", nil)
}