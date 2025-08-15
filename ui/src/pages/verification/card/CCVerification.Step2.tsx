//||------------------------------------------------------------------------------------------------||
//|| CCVerification.Step2
//|| src/components/payments/CCVerification.Step2.tsx
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| React
      //||------------------------------------------------------------------------------------------------||

      import React, { useState }                          from "react";
      import { CreditCard }                               from "lucide-react";

      //||------------------------------------------------------------------------------------------------||
      //|| Stripe
      //||------------------------------------------------------------------------------------------------||

      import {
            CardNumberElement,
            CardExpiryElement,
            CardCvcElement,
            useElements,
            useStripe,
      }                                                   from "@stripe/react-stripe-js";

      //||------------------------------------------------------------------------------------------------||
      //|| Interfaces
      //||------------------------------------------------------------------------------------------------||

      import { VerificationCard }                         from "../../../interfaces/verify/card/process";
      import InlineAlert                                  from "../../../components/base/InlineAlert";

      //||------------------------------------------------------------------------------------------------||
      //|| Data
      //||------------------------------------------------------------------------------------------------||

      import { formatStripeAmount }                        from "../../../data/getCurrencies";
      
      //||------------------------------------------------------------------------------------------------||
      //|| Props
      //||------------------------------------------------------------------------------------------------||

      import { CCStepProps }                                from "./CCVerification";

      //||------------------------------------------------------------------------------------------------||
      //|| React
      //||------------------------------------------------------------------------------------------------||

      export default function CCVerificationStep2({ process, updateProcess }: CCStepProps) {

            //||------------------------------------------------------------------------------------------------||
            //|| React
            //||------------------------------------------------------------------------------------------------||

            const stripe = useStripe();
            const elements = useElements();

            //||------------------------------------------------------------------------------------------------||
            //|| OnSubmit
            //||------------------------------------------------------------------------------------------------||

            const [ processError, setProcessError ]         = useState<string | null>(null);
            const [ busy, setBusy ]                         = useState(false);
            const [ zip, setZip ]                           = useState(process.billingZip || "");

            //||------------------------------------------------------------------------------------------------||
            //|| OnSubmit
            //||------------------------------------------------------------------------------------------------||

            const onSubmit = async () => {
                  //||------------------------------------------------------------------------------------------------||
                  //|| Reset State
                  //||------------------------------------------------------------------------------------------------||
                  setProcessError(null);
                  setBusy(true);
                  //||------------------------------------------------------------------------------------------------||
                  //|| Ensure Stripe is Loaded
                  //||------------------------------------------------------------------------------------------------||

                  if (!stripe || !elements) {
                        setProcessError("Stripe has not loaded.");
                        setBusy(false);
                        return;
                  }
                  //||------------------------------------------------------------------------------------------------||
                  //|| Get Card Element
                  //||------------------------------------------------------------------------------------------------||
                  const cardElement = elements.getElement(CardNumberElement);
                  if (!cardElement) {
                        setProcessError("Card input is missing.");
                        setBusy(false);
                        return;
                  }
                  //||------------------------------------------------------------------------------------------------||
                  //|| Confirm Card Payment
                  //||------------------------------------------------------------------------------------------------||
                  try {
                        const result = await stripe.confirmCardPayment(process.clientSecret, {
                              payment_method: {
                                    card            : cardElement,
                                    billing_details : {
                                          address: { postal_code: zip.trim() },
                                    },
                              },
                        });
                        //||------------------------------------------------------------------------------------------------||
                        //|| Handle Stripe Error
                        //||------------------------------------------------------------------------------------------------||
                        if (result.error) {
                              setProcessError(result.error.message || "Payment failed.");
                              setBusy(false);
                              return;
                        }
                        //||------------------------------------------------------------------------------------------------||
                        //|| Extract Card Details (if available)
                        //||------------------------------------------------------------------------------------------------||
                        const intent = result.paymentIntent;
                        //||------------------------------------------------------------------------------------------------||
                        //|| Update Process
                        //||------------------------------------------------------------------------------------------------||
                        const update: Partial<VerificationCard> = {
                              step           : 3,
                              transactionId  : intent?.id || "",
                              billingZip     : zip.trim(),
                        };
                        updateProcess({ ...process, ...update });
                  } catch (error) {
                        setProcessError("An error occurred while processing the payment.");
                        setBusy(false);
                        return;
                  }
                  //||------------------------------------------------------------------------------------------------||
                  //|| Reset Busy State
                  //||------------------------------------------------------------------------------------------------||
                  setBusy(false);
            };

            
            //||------------------------------------------------------------------------------------------------||
            //|| JSX
            //||------------------------------------------------------------------------------------------------||

           return (
			<div className="space-y-6 p-6">
                        <textarea defaultValue={JSON.stringify(process, null, 2)} className="h-96 w-full" readOnly />
				<form
					data-payment-user-email="test@example.com"
					onSubmit={(e) => {
						e.preventDefault();
						onSubmit();
					}}>
					<div className="flex items-center gap-2 text-gray-300 bg-black/60 text-center w-full p-3 rounded-lg mb-4 text-sm">
						<CreditCard className="w-5 h-5" />
                                    We’ll place a charge for
                                    <span className="text-yellow-400 ml-1">{ formatStripeAmount(process.chargeAmount.amount, process.chargeAmount.currency)} {process.chargeAmount.currency}</span>. The initial charge will be refunded upon verification.
					</div>

					{processError && <InlineAlert message={processError} isError={true} />}

					<div className="grid grid-cols-2 gap-4">
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
												"::placeholder": {
													color: "#cfcfcf",
												},
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
												"::placeholder": {
													color: "#cfcfcf",
												},
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
												"::placeholder": {
													color: "#cfcfcf",
												},
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

					<div className="flex items-center justify-between mt-6">
						<button
							type="button"
							onClick={() => setStep(1)}
							className="px-4 py-2 rounded-2xl font-bold bg-black/30 opacity-70 hover:opacity-100 text-white">
							Back
						</button>
						<button
							type="submit"
							disabled={!stripe || !elements || busy}
							className={`inline-flex items-center gap-2 px-5 py-2.5 rounded-2xl border shadow-sm ${
								!busy
									? "bg-orange-400 cursor-pointer font-bold text-white border-0 opacity-80 hover:opacity-100"
									: "bg-gray-200 text-gray-500 cursor-not-allowed"
							}`}>
							<CreditCard className="w-4 h-4" /> Confirm Verification
						</button>
					</div>
				</form>
			</div>
		);



      }
