//||------------------------------------------------------------------------------------------------||
//|| StripePaymentForm
//|| src/components/payments/StripePaymentForm.tsx
//||------------------------------------------------------------------------------------------------||

      import React, {useState}                                                                  from "react";
      import {Lock, CreditCard, ChevronLeft}                                                    from "lucide-react";
      import {CardNumberElement, CardExpiryElement, CardCvcElement, useStripe, useElements}     from "@stripe/react-stripe-js";

      //||------------------------------------------------------------------------------------------------||
      //|| Data
      //||------------------------------------------------------------------------------------------------||

      import { formatStripeAmount }                        from "../../data/getCurrencies";

      //||------------------------------------------------------------------------------------------------||
      //|| Component
      //||------------------------------------------------------------------------------------------------||

      import InlineAlert                                                                        from "../base/InlineAlert";

      //||------------------------------------------------------------------------------------------------||
      //|| Props
      //||------------------------------------------------------------------------------------------------||

      type StripePaymentFormProps = {
            clientSecret            : string;
            baseAmount              : number;
            donationAmount          : number;
            chargeAmount            : number;
            currency                : string;
            onBack?                 : () => void;
            onSuccess               : (paymentIntentId: string) => void;
            setZip?                 : (zip: string) => void;
            billingZip?             : string;
      };

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      export default function StripePaymentForm({clientSecret, baseAmount, donationAmount, chargeAmount, currency, onBack, onSuccess,  billingZip = ""}: StripePaymentFormProps) {

            //||------------------------------------------------------------------------------------------------||
            //|| Hooks
            //||------------------------------------------------------------------------------------------------||
            const stripe = useStripe();
            const elements = useElements();

            //||------------------------------------------------------------------------------------------------||
            //|| State
            //||------------------------------------------------------------------------------------------------||

            const [zip, setZip] = useState(billingZip);
            const [error, setError] = useState<string | null>(null);
            const [busy, setBusy] = useState(false);

            //||------------------------------------------------------------------------------------------------||
            //|| Submit
            //||------------------------------------------------------------------------------------------------||

            const handleSubmit = async (e: React.FormEvent) => {
                  e.preventDefault();
                  setError(null);
                  setBusy(true);

                  if (!stripe || !elements) {
                        setError("Stripe is not loaded.");
                        setBusy(false);
                        return;
                  }

                  if (!clientSecret) {
                        setError("Client secret is missing.");
                        setBusy(false);
                        return;
                  }

                  const cardElement = elements.getElement(CardNumberElement);
                  if (!cardElement) {
                        setError("Card input missing.");
                        setBusy(false);
                        return;
                  }

                  const result = await stripe.confirmCardPayment(clientSecret, {
                        payment_method: {
                              card: cardElement,
                              billing_details: {
                                    address: {postal_code: zip.trim()},
                              },
                        },
                  });
                  if (result.error) {
                        setError(result.error.message || "Payment failed.");
                        setBusy(false);
                        return;
                  }
                  if (result.paymentIntent && result.paymentIntent.status === "succeeded") {
                        onSuccess(result.paymentIntent.id);
                  } else {
                        setError("Payment did not complete.");
                        setBusy(false);
                  }
                  setBusy(false);
            };

            //||------------------------------------------------------------------------------------------------||
            //|| JSX
            //||------------------------------------------------------------------------------------------------||

            return (
                  <form onSubmit={handleSubmit} className="">
                        
                        {error && <InlineAlert message={error} isError={true} />}

                        {!error && ( 

                              <div className="flex items-center bg-black/60 rounded-lg mb-4 p-3 text-gray-300 text-sm w-full">
                                    <CreditCard className="w-5 h-5 mr-3 shrink-0 text-yellow-400" />

                                    <div className="flex flex-col sm:flex-row sm:items-center w-full justify-between text-white gap-2">
                                          <div>
                                                <span>
                                                      We’ll place a charge for&nbsp;
                                                      { baseAmount > 0 && (
                                                            <span className="font-bold text-white">
                                                                  {formatStripeAmount(baseAmount, currency)}{" "}
                                                                  {currency}
                                                            </span>
                                                      )}
                                                </span>
                                                {donationAmount > 0 && (
                                                      <span>
                                                            {baseAmount > 0 ? "&nbsp;+ a donation of&nbsp;" : ""}                                                            
                                                            <span className="font-bold">
                                                                  {formatStripeAmount(donationAmount, currency)}{" "}
                                                                  {currency}
                                                            </span>
                                                      </span>
                                                )}
                                          </div>

                                          <span className="sm:ml-auto font-bold block pt-2 sm:pt-0 text-yellow-400">
                                                Total:&nbsp;{formatStripeAmount(chargeAmount, currency)} { currency }
                                          </span>
                                    </div>
                              </div>
                        )}                              

                        <div className="grid grid-cols-2 gap-4 mt-8">
                              <label className="col-span-2">
                                    <span className="block text-sm mb-1 text-white">Card number</span>
                                    <div className="rounded-lg border border-white/20 bg-[#181818] px-4 py-3 text-2xl">
                                          <CardNumberElement
                                                options={{
                                                      style: {
                                                            base: {
                                                                  fontSize: "30px",
                                                                  color: "#fefeff",
                                                                  backgroundColor: "transparent",
                                                                  iconColor: "#cfcfcf",
                                                                  "::placeholder": {color: "#cfcfcf"},
                                                            },
                                                            invalid: {
                                                                  color: "#ff5a5f",
                                                                  iconColor: "#ff5a5f",
                                                            },
                                                      },
                                                }}
                                          />
                                    </div>
                              </label>
                              <label>
                                    <span className="block text-sm mb-1 text-white">Expiration</span>
                                    <div className="rounded-lg border border-white/20 bg-[#181818] px-4 py-3 text-2xl">
                                          <CardExpiryElement
                                                options={{
                                                      style: {
                                                            base: {
                                                                  fontSize: "30px",
                                                                  color: "#fefeff",
                                                                  backgroundColor: "transparent",
                                                                  iconColor: "#cfcfcf",
                                                                  "::placeholder": {color: "#cfcfcf"},
                                                            },
                                                            invalid: {
                                                                  color: "#ff5a5f",
                                                                  iconColor: "#ff5a5f",
                                                            },
                                                      },
                                                }}
                                          />
                                    </div>
                              </label>
                              <label>
                                    <span className="block text-sm mb-1 text-white">CVV/CVC</span>
                                    <div className="rounded-lg border border-white/20 bg-[#181818] px-4 py-3 text-2xl">
                                          <CardCvcElement
                                                options={{
                                                      style: {
                                                            base: {
                                                                  fontSize: "30px",
                                                                  color: "#fefeff",
                                                                  backgroundColor: "transparent",
                                                                  iconColor: "#cfcfcf",
                                                                  "::placeholder": {color: "#cfcfcf"},
                                                            },
                                                            invalid: {
                                                                  color: "#ff5a5f",
                                                                  iconColor: "#ff5a5f",
                                                            },
                                                      },
                                                }}
                                          />
                                    </div>
                              </label>
                              <label>
                                    <span className="block text-sm mb-1 text-white">Billing ZIP / Postal</span>
                                    <input
                                          placeholder="ZIP / Postal"
                                          value={zip}
                                          onChange={(e) => setZip(e.target.value)}
                                          className="input input-bordered w-full text-2xl h-16 bg-[#181818] border-white/20 text-white"
                                    />
                              </label>
                        </div>
                        <hr className="mb-5 mt-5 border-gray-500" />
                        <div className="flex items-center justify-between w-full">

                              {onBack ? (
                                    <div>
                                          <button
                                                type="button"
                                                onClick={onBack}
                                                className="inline-flex items-center gap-2 px-5 py-2.5 rounded-2xl bg-black/40 shadow-sm cursor-pointer hover:bg-black/60"
                                          >
                                                <ChevronLeft className="w-4 h-4" />Back
                                          </button>
                                    </div>
                              ) : <div />}

                              <div>
                                    <button
                                          type="submit"
                                          disabled={!stripe || !elements || busy}
                                          className={`inline-flex items-center gap-2 px-5 py-2.5 rounded-2xl border shadow-sm ${
                                                !busy
                                                      ? "bg-orange-400 cursor-pointer font-bold text-white border-0 opacity-80 hover:opacity-100"
                                                      : "bg-gray-200 text-gray-500 cursor-not-allowed"
                                          }`}>
                                          <Lock className="w-4 h-4" /> Pay Now
                                    </button>
                              </div>
                        </div>

                  </form>
            );
      }
