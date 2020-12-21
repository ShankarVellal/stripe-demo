## Stripe Payments Demo

This is a demo of integrating Stripe Payments using a custom payment flow (vs. prebuilt Stripe Checkout page). This uses a Golang server and React client.

### Instructions

- Start the server first using:
```bash
cd server
go run main.go
```

- Start the  client using:
```bash
cd client
npm start
```

- Note that this uses test keys, so only the test cards on [Stripe Payments Testing](https://stripe.com/docs/testing) page will work.

- Successful payments are logged in ```successful_payments.log``` file in the ```server``` folder
