## Stripe Payments Demo

This is a demo of integrating Stripe Payments using a custom payment flow (vs. prebuilt Stripe Checkout page). This uses a Golang server and React client.

### Instructions

- Start the server first using:
```bash
cd server
STRIPE_SECRET_KEY=<your_stripe_test_secret_key> go run main.go
```

- Start the  client using:
```bash
cd client
npm start
```

- Run `stripe listen --forward-to localhost:8080/webhook` to forward post-payment events to the webhook handler in the server. Please see [here](https://stripe.com/docs/webhooks/test) for more information installing the Stripe CLI

- Note that this uses test keys, so only the test cards on [Stripe Payments Testing](https://stripe.com/docs/testing) page will work.

- Successful payments are logged in ```successful_payments.log``` file in the ```server``` folder
