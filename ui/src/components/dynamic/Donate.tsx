//||------------------------------------------------------------------------------------------------||
//|| Donation
//|| src/components/payments/Donation.tsx
//||------------------------------------------------------------------------------------------------||

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, {useEffect, useRef, useState}                    from "react";
import {formatStripeAmount, getCurrencySymbol}                 from "../../data/getCurrencies";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

type DonationProps = {
	donation                : number;
	setDonation             : (amount: number) => void;
      currency                : string;
};

const presets = [0, 100, 300, 500, 1000, 5000];
const presetSet = new Set(presets);

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function Donation({ donation, setDonation, currency }: DonationProps) {
      const isCustom = !presetSet.has(donation);

      // For input, show dollars (not pennies)
      const inputValue = formatStripeAmount(donation, currency, true);

      // When user changes input, multiply by 100 to get pennies/cents
      const onCustomChange: React.ChangeEventHandler<HTMLInputElement> = (e) => {
            const raw = Number(e.target.value);
            // ignore NaN or negatives
            const cents = Number.isFinite(raw) && raw > 0 ? Math.round(raw * 100) : 0;
            setDonation(cents);
      };

      return (
            <div className="mx-auto">
                  <div className="flex justify-center items-center gap-2" role="group" aria-label="Choose donation amount">
                        {presets.map((amt) => {
                              const active = donation === amt;
                              return (
                                    <button
                                          key={amt}
                                          type="button"
                                          aria-pressed={active}
                                          onClick={() => setDonation(amt)}
                                          className={[
                                                "px-3 py-2 rounded-2xl text-xs font-bold transition-colors",
                                                active ? "bg-orange-500 text-white" : "bg-gray-600 text-white/90 hover:bg-gray-500",
                                          ].join(" ")}
                                    >
                                          {amt === 0
                                                ? "No Donation"
                                                : `${formatStripeAmount(amt, currency)}`}
                                    </button>
                              );
                        })}

                        {/* Custom amount in dollars */}
                        <div
                              className={[
                                    "flex items-center gap-2 px-2 py-1 rounded-2xl",
                                    isCustom ? "bg-orange-400/60" : "ring-1 ring-transparent",
                              ].join(" ")}
                        >
                              <span className="text-xs font-bold">Custom</span>
                              { getCurrencySymbol(currency) }
                              <input
                                    type="number"
                                    min={0}
                                    step={0.01}
                                    value={inputValue}
                                    onChange={onCustomChange}
                                    className="input input-bordered w-16 h-9 font-bold text-center bg-black text-yellow-400"
                                    aria-label="Custom donation amount in dollars"
                              />
                              { currency }
                        </div>
                  </div>
            </div>
      );
}