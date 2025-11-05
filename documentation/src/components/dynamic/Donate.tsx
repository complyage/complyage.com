//||------------------------------------------------------------------------------------------------||
//|| Donation
//|| src/components/payments/Donation.tsx
//||------------------------------------------------------------------------------------------------||

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React from "react";
import { formatStripeAmount, getCurrencySymbol } from "../../data/getCurrencies";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

type DonationProps = {
      donation    : number;
      baseAmount  : number;
      setDonation : (amount: number) => void;
      currency    : string;
      size?       : "sm" | "md" | "lg";
};

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function Donation({ baseAmount, donation, setDonation, currency, size = "sm" }: DonationProps) {
      const presets   = (baseAmount === 0) ? [100, 300, 500, 1000, 5000] : [0, 100, 300, 500, 1000, 5000];
      const presetSet = new Set(presets);
      const isCustom  = !presetSet.has(donation);

      // For input, show dollars (not pennies)
      const inputValue = formatStripeAmount(donation, currency, true);

      // When user changes input, multiply by 100 to get pennies/cents
      const onCustomChange: React.ChangeEventHandler<HTMLInputElement> = (e) => {
            const raw = Number(e.target.value);
            const cents = Number.isFinite(raw) && raw > 0 ? Math.round(raw * 100) : 0;
            setDonation(cents);
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Size Variants
      //||------------------------------------------------------------------------------------------------||
      const sizeClasses = {
            sm: {
                  button: "px-3 py-2 text-xs",
                  custom: "px-2 py-1 text-xs",
                  input : "w-16 h-9 text-sm",
            },
            md: {
                  button: "px-4 py-2 text-sm",
                  custom: "px-3 py-2 text-sm",
                  input : "w-20 h-10 text-base",
            },
            lg: {
                  button: "px-6 py-3 text-base",
                  custom: "px-4 py-3 text-base",
                  input : "w-24 h-12 text-lg",
            },
      }[size];

      return (
            <div className="mx-auto">
                  <div className="flex justify-center items-center gap-2 flex-wrap" role="group" aria-label="Choose donation amount">
                        {presets.map((amt) => {
                              const active = donation === amt;
                              return (
                                    <button
                                          key={amt}
                                          type="button"
                                          aria-pressed={active}
                                          onClick={() => setDonation(amt)}
                                          className={[
                                                sizeClasses.button,
                                                "rounded-2xl font-bold transition-colors",
                                                active
                                                      ? "bg-orange-500 text-white"
                                                      : "bg-gray-600 text-white/90 hover:bg-gray-500",
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
                                    "flex items-center gap-2 rounded-2xl font-bold",
                                    sizeClasses.custom,
                                    isCustom ? "bg-orange-400/60" : "ring-1 ring-transparent",
                              ].join(" ")}
                        >
                              <span>Custom</span>
                              {getCurrencySymbol(currency)}
                              <input
                                    type="number"
                                    min={0}
                                    step={0.01}
                                    value={inputValue}
                                    onChange={onCustomChange}
                                    className={`input input-bordered text-center bg-black text-yellow-400 font-bold ${sizeClasses.input}`}
                                    aria-label="Custom donation amount in dollars"
                              />
                              {currency}
                        </div>
                  </div>
            </div>
      );
}
