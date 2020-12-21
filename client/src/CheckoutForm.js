import React, { useEffect, useState } from "react";
import {useStripe, useElements, CardElement} from '@stripe/react-stripe-js';

import CardSection from './CardSection';

export default function CheckoutForm() {
  const [clientSecret, setClientSecret] = useState(null);
  const [processing, setProcessing] = useState(false);
  const [succeeded, setSucceeded] = useState(false);
  const stripe = useStripe();
  const elements = useElements();

  useEffect(() => {
    fetch('/secret').then(function(response) {
      return response.json();
    }).then((responseJson) => {
      setClientSecret(responseJson.client_secret)
    });
  }, []);

  const handleSubmit = async (event) => {
    // We don't want to let default form submission happen here,
    // which would refresh the page.
    event.preventDefault();
    setProcessing(true);

    if (!stripe || !elements) {
      // Stripe.js has not yet loaded.
      // Make sure to disable form submission until Stripe.js has loaded.
      return;
    }

    const result = await stripe.confirmCardPayment(clientSecret, {
      payment_method: {
        card: elements.getElement(CardElement),
        billing_details: {
          name: 'Jenny Rosen',
        },
      }
    });

    if (result.error) {
      // Show error to your customer (e.g., insufficient funds)
      console.log(result.error.message);
    } else {
      // The payment has been processed!
      if (result.paymentIntent.status === 'succeeded') {
        // Show a success message to your customer
        // There's a risk of the customer closing the window before callback
        // execution. Set up a webhook or plugin to listen for the
        // payment_intent.succeeded event that handles any business critical
        // post-payment actions.
        setSucceeded(true);
        console.log("payment succeeded");
      }
    }
  };

  const renderForm = () => {
    return(
      <form onSubmit={handleSubmit}>
        <CardSection />
        <button disabled={processing || !clientSecret || !stripe}>Confirm order</button>
      </form>
    );
  };

  const renderSuccess = () => {
    return (
      <div>
        <h1>Your test payment succeeded</h1>
      </div>
    );
  };

  if (succeeded) {
    return renderSuccess();
  } else {
    return renderForm();
  }
}
