/*||------------------------------------------------------------------------------------------------||
//|| DonateCard Component
//|| src/components/donate/DonateCard.tsx
//||------------------------------------------------------------------------------------------------||*/

import React, { useState, useEffect } from "react";
import { Elements } from "@stripe/react-stripe-js";
import { loadStripe } from "@stripe/stripe-js";

import Donation          from "../dynamic/Donate";
import StripePaymentForm from "../dynamic/StripePaymentForm";
import Call              from "../../classes/call";

/*||------------------------------------------------------------------------------------------------||
//|| Stripe Config
//||------------------------------------------------------------------------------------------------||*/

const stripePromise = loadStripe(import.meta.env.VITE_MERCHANT_PUBLIC || "pk_test_12345");

export default function DonateCard() {
      const [donation, setDonation]         = useState(500);
      const [currency, setCurrency]         = useState("USD");
      const [clientSecret, setClientSecret] = useState<string>("");

      const [step, setStep] = useState<1 | 2 | 3>(1);

      //||------------------------------------------------------------------------------------------------||
      //|| Create Payment Intent (only once user moves to step 2)
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            if (step !== 2) return;
            (async () => {
                  const chirp = new Call("/v1/api/donate/intent", { amount: donation, currency });
                  chirp.method = "POST";
                  chirp.debug  = true;
                  await chirp.execute();
                  if (chirp.ok()) {
                        setClientSecret(chirp.data("clientSecret") || "");
                  }
            })();
      }, [step, donation, currency]);

      //||------------------------------------------------------------------------------------------------||
      //|| Step 1: Select Donation Amount
      //||------------------------------------------------------------------------------------------------||

      if (step === 1) {
            return (
                  <div className="flex flex-col items-center bg-black/20 p-5 rounded-lg">
                        <h1 className="text-2xl font-bold text-left w-full max-w-3xl px-3 pb-5">Please select your donation amount:</h1>
                        <Donation baseAmount={0} donation={donation} setDonation={setDonation} currency={currency} size="lg" />
                        <div className="flex w-full mt-10 justify-end">
                              <button
                                    onClick={() => setStep(2)}
                                    className="mt-6 px-6 py-3 bg-orange-500 text-white font-bold rounded-lg hover:bg-orange-600">
                                    Continue to Payment
                              </button>
                        </div>
                  </div>
            );
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Step 2: Enter Credit Card Info
      //||------------------------------------------------------------------------------------------------||

      if (step === 2) {
            return (
                  <div className="w-full flex flex-col items-center">
                        {clientSecret ? (
                              <Elements
                                    stripe={stripePromise}
                                    options={{ clientSecret, appearance: { theme: "night" } }}>
                                    <StripePaymentForm
                                          clientSecret={clientSecret}
                                          baseAmount={0}
                                          donationAmount={donation}
                                          chargeAmount={donation}
                                          currency={currency}
                                          onSuccess={async (paymentIntentId: string) => {
                                                // Verify and log on the backend
                                                const chirp = new Call("/v1/api/donate/complete", {
                                                      method: "CARD",
                                                      merchant: "Stripe",
                                                      amount: donation / 100, // cents → dollars
                                                      txID: paymentIntentId,
                                                });
                                                chirp.method = "POST";
                                                await chirp.execute();

                                                // Move to step 3 regardless
                                                setStep(3);
                                          }}
                                    />
                              </Elements>
                        ) : (
                              <div className="text-gray-400">Loading payment form…</div>
                        )}
                  </div>
            );
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Step 3: Thank You Screen
      //||------------------------------------------------------------------------------------------------||
      if (step === 3) {
            return (
                  <div className="flex flex-col items-center text-center p-8 bg-black/20 rounded-xl">
                        <h2 className="text-4xl font-bold text-green-500 mb-4">Thank You!</h2>
                        <p className="text-gray-200 text-2xl">
                              Your donation has been received successfully.
                              <br /> We appreciate your support in keeping this project alive.
                        </p>
                        <button
                              onClick={() => setStep(1)}
                              className="mt-6 px-6 py-3 bg-gray-700 text-white font-bold rounded-lg hover:bg-gray-600">
                              Make Another Donation
                        </button>
                  </div>
            );
      }

      return null;
}
