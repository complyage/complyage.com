//||------------------------------------------------------------------------------------------------||
//|| AddressVerification.Step3
//|| src/components/security/AddressVerification.Step3.tsx
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||

      import React                              from "react";
      import {CreditCard, Lock}                 from "lucide-react";
      import InlineAlert                        from "../../../components/base/InlineAlert";
      import {CardData}                         from "../../../interfaces/verification.location";

      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||

      type Props = {
            card              : CardData;
            setCard           : React.Dispatch<React.SetStateAction<CardData>>;
            cardOk            : boolean;
            busy              : boolean;
            serverMsg         : string;
            onBack            : () => void;
            onSubmit          : () => void;
      };

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      export default function AddressVerificationStep3({card, setCard, cardOk, busy, serverMsg, onBack, onSubmit}: Props) {
            return (
                  <div className="space-y-6 p-6">
                        <div className="flex items-center gap-2 text-gray-300">
                              <CreditCard className="w-5 h-5" />
                              <span>We’ll place a $1.00 verification hold. Enter your billing details.</span>
                        </div>

                        {!!serverMsg && <div className="text-xs mt-1 opacity-80">{serverMsg}</div>}                        

                        <div className="grid grid-cols-2 gap-4">
                              <label className="col-span-2">
                                    <span className="block text-sm mb-1">Card number</span>
                                    <input
                                          placeholder="4242 4242 4242 4242"
                                          inputMode="numeric"
                                          maxLength={19}
                                          value={card.number}
                                          onChange={(e) => setCard((c) => ({...c, number: digitsAndSpaces(e.target.value)}))}
                                          className="input input-bordered w-full text-2xl h-16"
                                    />
                              </label>

                              <label>
                                    <span className="block text-sm mb-1">Exp. MM</span>
                                    <input
                                          placeholder="MM"
                                          inputMode="numeric"
                                          maxLength={2}
                                          value={card.expMonth}
                                          onChange={(e) => setCard((c) => ({...c, expMonth: onlyDigits(e.target.value)}))}
                                          className="input input-bordered w-full text-2xl h-16"
                                    />
                              </label>

                              <label>
                                    <span className="block text-sm mb-1">Exp. YY</span>
                                    <input
                                          placeholder="YY"
                                          inputMode="numeric"
                                          maxLength={4}
                                          value={card.expYear}
                                          onChange={(e) => setCard((c) => ({...c, expYear: onlyDigits(e.target.value)}))}
                                          className="input input-bordered w-full text-2xl h-16"
                                    />
                              </label>

                              <label>
                                    <span className="block text-sm mb-1">CVV/CVC</span>
                                    <input
                                          placeholder="CVC"
                                          inputMode="numeric"
                                          maxLength={4}
                                          value={card.cvc}
                                          onChange={(e) => setCard((c) => ({...c, cvc: onlyDigits(e.target.value)}))}
                                          className="input input-bordered w-full text-2xl h-16"
                                    />
                              </label>

                              <label>
                                    <span className="block text-sm mb-1">Billing ZIP / Postal</span>
                                    <input
                                          placeholder="ZIP / Postal"
                                          value={card.postal}
                                          onChange={(e) => setCard((c) => ({...c, postal: e.target.value}))}
                                          className="input input-bordered w-full text-2xl h-16"
                                    />
                              </label>
                        </div>

                        {card.number || card.expMonth || card.expYear || card.cvc || card.postal ? (
                              cardOk ? (
                                    <InlineAlert message="Card looks good." isError={false} />
                              ) : (
                                    <InlineAlert message="Enter a valid card, expiry, CVC and billing postal/ZIP." isError={true} />
                              )
                        ) : null}

                        <div className="flex items-center justify-between">
                              <button onClick={onBack} className="px-4 py-2 rounded-2xl font-bold bg-black/30 opacity-70 hover:opacity-100">
                                    Back
                              </button>
                              <button
                                    onClick={onSubmit}
                                    disabled={!cardOk || busy}
                                    className={`inline-flex items-center gap-2 px-5 py-2.5 rounded-2xl border shadow-sm ${
                                          cardOk && !busy
                                                ? "bg-orange-400 cursor-pointer font-bold text-white border-0 opacity-80 hover:opacity-100"
                                                : "bg-gray-200 text-gray-500 cursor-not-allowed"
                                    }`}>
                                    <Lock className="w-4 h-4" /> Verify Address ($1.00)
                              </button>
                        </div>

                  </div>
            );
      }
