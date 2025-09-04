/*||------------------------------------------------------------------------------------------------||
//|| Donate Page
//|| src/routes/donate.tsx
//||------------------------------------------------------------------------------------------------||*/

/*||------------------------------------------------------------------------------------------------||
//|| React & Vendor
//||------------------------------------------------------------------------------------------------||*/

import React, { useState, useEffect } from "react";
import NavMain                        from "../../components/nav/NavMain";
import FooterMain                     from "../../components/footer/FooterMain";
import Donation                       from "../../components/dynamic/Donate";
import StripePaymentForm              from "../../components/dynamic/StripePaymentForm";
import ProgressBar                    from "../../components/base/ProgressBar";
import { CheckCircle, CreditCard, Bitcoin, Mail as MailIcon } from "lucide-react";
import { Elements }                   from "@stripe/react-stripe-js";
import { loadStripe }                 from "@stripe/stripe-js";

/*||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||*/

import type { DonateApiResponse, DonateCryptoOption, DonateAddress } from "../../interfaces/donate/donate";

/*||------------------------------------------------------------------------------------------------||
//|| Stripe Config
//||------------------------------------------------------------------------------------------------||*/

const stripePromise = loadStripe("pk_live_YOUR_PUBLIC_KEY");

/*||------------------------------------------------------------------------------------------------||
//|| Donate Page Component
//||------------------------------------------------------------------------------------------------||*/

export default function DonatePage() {

      //||------------------------------------------------------------------------------------------------||
      //|| State
      //||------------------------------------------------------------------------------------------------||

      const [donation, setDonation]           = useState(500);    // in cents (e.g. $5.00)
      const [currency, setCurrency]           = useState("USD");
      const [clientSecret, setClientSecret]   = useState<string>(""); // Fill in with backend logic as needed
      const [chargeComplete, setChargeComplete] = useState(false);

      const [donateData, setDonateData]       = useState<DonateApiResponse | null>(null);
      const [activeTab, setActiveTab]         = useState<"card" | "crypto" | "cash">("card");
      const [cryptoIdx, setCryptoIdx]         = useState<number>(0);

      //||------------------------------------------------------------------------------------------------||
      //|| Fetch donation data (crypto + address)
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            fetch("/v1/api/donate", { credentials: "include" })
                  .then(res => res.json())
                  .then((data: DonateApiResponse) => setDonateData(data));
      }, []);

      // Optionally: Fetch clientSecret when donation/currency changes
      // useEffect(() => { fetchClientSecret(); }, [donation, currency]);

      //||------------------------------------------------------------------------------------------------||
      //|| Render
      //||------------------------------------------------------------------------------------------------||

      return (
            <main className="min-h-screen flex flex-col bg-base-100 text-base-content">
                  <NavMain />

                  {/* Hero */}
                  <section className="relative flex flex-col items-center justify-center text-center py-20 bg-primary text-primary-content px-4">
                        <h1 className="text-5xl font-extrabold mb-4 mt-[120px]">Support Our Mission</h1>
                        <p className="text-lg max-w-2xl">
                              Your donation keeps this project open, ad-free, and independent.<br />
                              We never sell data, never show tracking ads, and your gift directly funds real compliance & privacy tools for everyone.
                        </p>
                  </section>

                  {/* Progress */}
                  <section className="py-8 px-4 bg-gray-800 text-base-content text-center">
                        <h3 className="text-2xl font-bold mb-4">Monthly Support Raised</h3>
                        <ProgressBar current={1272} goal={4314} prefix="$" label="of goal" />
                  </section>
                  
                  {/* Donation Method Buttons - OUTSIDE */}
                  <section className="w-full max-w-5xl mx-auto mt-10 mb-0 flex flex-col items-center">
                        <div className="flex justify-center gap-2 w-full mb-7">
                              <button
                                    className={`flex-1 flex items-center justify-center gap-2 px-5 py-4 rounded-2xl font-bold transition-colors text-xl
                                          ${activeTab === "card" ? "bg-orange-500 text-white shadow-lg" : "bg-gray-700 text-white/80 hover:bg-orange-500/70"}`}
                                    onClick={() => setActiveTab("card")}
                              >
                                    <CreditCard className="w-6 h-6" /> Credit Card
                              </button>
                              <button
                                    className={`flex-1 flex items-center justify-center gap-2 px-5 py-4 rounded-2xl font-bold transition-colors text-xl
                                          ${activeTab === "crypto" ? "bg-yellow-400 text-black shadow-lg" : "bg-gray-700 text-white/80 hover:bg-yellow-400/80"}`}
                                    onClick={() => setActiveTab("crypto")}
                              >
                                    <Bitcoin className="w-6 h-6" /> Crypto
                              </button>
                              <button
                                    className={`flex-1 flex items-center justify-center gap-2 px-5 py-4 rounded-2xl font-bold transition-colors text-xl
                                          ${activeTab === "cash" ? "bg-lime-500 text-black shadow-lg" : "bg-gray-700 text-white/80 hover:bg-lime-400/80"}`}
                                    onClick={() => setActiveTab("cash")}
                              >
                                    <MailIcon className="w-6 h-6" /> Cash/Check
                              </button>
                        </div>
                  </section>

                  {/* Donation UI */}
                  <section className="py-8 px-4 w-full max-w-5xl mx-auto flex flex-col gap-10">
                        <div className="bg-base-200 rounded-xl shadow-lg p-8 flex flex-col items-center">

                              <h2 className="text-2xl font-bold mb-2 text-center">Make a Donation</h2>
                              <p className="mb-6 text-base-content/70 text-center">
                                    Choose an amount and method below.<br />
                                    <span className="text-sm text-orange-400 font-bold">All donations are anonymous and non-refundable.</span>
                              </p>

                              {/* Stripe Credit Card Donation */}
                              {activeTab === "card" && (
                                    <>
                                          <Donation donation={donation} setDonation={setDonation} currency={currency} />
                                          {!chargeComplete && (
                                                <div className="w-full mt-7">
                                                      <Elements stripe={stripePromise}>
                                                            <StripePaymentForm
                                                                  clientSecret={clientSecret}
                                                                  baseAmount={0}
                                                                  donationAmount={donation}
                                                                  chargeAmount={donation}
                                                                  currency={currency}
                                                                  onSuccess={() => setChargeComplete(true)}
                                                            />
                                                      </Elements>
                                                </div>
                                          )}
                                          {chargeComplete && (
                                                <div className="flex flex-col items-center gap-2 mt-6 text-green-400 text-lg">
                                                      <CheckCircle className="w-10 h-10" />
                                                      Thank you for your donation!
                                                </div>
                                          )}
                                    </>
                              )}

                              {/* Crypto Donation */}
                              {activeTab === "crypto" && donateData && (
                                    <div className="w-full mt-7">
                                          {/* Select dropdown for currency */}
                                          <div className="flex flex-col md:flex-row items-center gap-3 mb-6">
                                                <label className="font-bold">Choose a currency:</label>
                                                <select
                                                      value={cryptoIdx}
                                                      onChange={e => setCryptoIdx(Number(e.target.value))}
                                                      className="select select-bordered font-bold bg-black text-yellow-400"
                                                >
                                                      {donateData.crypto.map((c, i) => (
                                                            <option value={i} key={c.symbol}>{c.name} ({c.symbol})</option>
                                                      ))}
                                                </select>
                                          </div>
                                          {/* Show selected crypto option */}
                                          {donateData.crypto.length > 0 && (
                                                <div className={`flex flex-col items-center rounded-2xl shadow-lg p-6 w-full max-w-md mx-auto ${donateData.crypto[cryptoIdx].color || ""}`}>
                                                      <div className="mb-3 text-2xl font-bold tracking-wide uppercase">{donateData.crypto[cryptoIdx].name}</div>
                                                      <img src={donateData.crypto[cryptoIdx].qr} alt={`${donateData.crypto[cryptoIdx].name} QR`} className="rounded-xl w-44 h-44 mx-auto border-4 border-white shadow-lg" />
                                                      <div className="break-all mt-3 bg-white/10 p-3 rounded text-sm font-mono">
                                                            {donateData.crypto[cryptoIdx].address}
                                                      </div>
                                                      <div className="text-xs mt-2 opacity-60">Scan QR or copy address</div>
                                                </div>
                                          )}
                                    </div>
                              )}

                              {/* Cash/Check Donation */}
                              {activeTab === "cash" && donateData && (
                                    <div className="w-full mt-7">
                                          <div className="flex flex-col items-center rounded-2xl shadow-lg p-6 w-full max-w-md mx-auto bg-lime-600/70 text-white">
                                                <div className="mb-3 text-xl font-bold tracking-wide uppercase">Mail a Cash/Check Donation</div>
                                                <div className="text-lg font-semibold">{donateData.address.name}</div>
                                                <div className="mt-2 mb-2">
                                                      <div>{donateData.address.address1}</div>
                                                      {donateData.address.address2 && <div>{donateData.address.address2}</div>}
                                                      <div>
                                                            {donateData.address.city}, {donateData.address.state} {donateData.address.postal}
                                                      </div>
                                                      <div>{donateData.address.country}</div>
                                                </div>
                                                <div className="text-xs mt-2 opacity-90">
                                                      Please make checks payable to <b>{donateData.address.name}</b>.<br />
                                                      <span className="opacity-70">Include your email if you’d like a receipt.</span>
                                                </div>
                                          </div>
                                    </div>
                              )}
                        </div>
                  </section>

                  <FooterMain />
            </main>
      );
}
