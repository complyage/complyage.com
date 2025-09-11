/*||------------------------------------------------------------------------------------------------||
//|| Donate Page
//|| src/routes/donate.tsx
//||------------------------------------------------------------------------------------------------||*/

/*||------------------------------------------------------------------------------------------------||
//|| React & Vendor
//||------------------------------------------------------------------------------------------------||*/

import React, { useState, useEffect } from "react";
import { Copy, Check, CheckCircle, CreditCard, Bitcoin, Mail as MailIcon } from "lucide-react";
import { Elements }                   from "@stripe/react-stripe-js";
import { loadStripe }                 from "@stripe/stripe-js";

/*||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||*/

import MonthlySupportHero             from "../../components/heroes/MonthlySupport";
import NavMain                        from "../../components/nav/NavMain";
import FooterMain                     from "../../components/footer/FooterMain";
import DonateCard                     from "../../components/donate/DonateCard";
import DonateCrypto                   from "../../components/donate/DonateCrypto";
import DonateCheck                    from "../../components/donate/DonateCheck";
import MainHero                       from "../../components/heroes/Main";

/*||------------------------------------------------------------------------------------------------||
//|| Classes
//||------------------------------------------------------------------------------------------------||*/

import Call                           from "../../classes/call";

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

      const [donation, setDonation]                   = useState(500);    // in cents (e.g. $5.00)
      const [currency, setCurrency]                   = useState("USD");
      const [clientSecret, setClientSecret]           = useState<string>(""); // Fill in with backend logic as needed
      const [chargeComplete, setChargeComplete]       = useState(false);
      const [donateData, setDonateData]               = useState<DonateApiResponse | null>(null);
      const [activeTab, setActiveTab]                 = useState<"card" | "crypto" | "cash">("card");
      const [cryptoIdx, setCryptoIdx]                 = useState<number>(0);
      const [copied, setCopied]                       = useState(false);

      //||------------------------------------------------------------------------------------------------||
      //|| Create Intent
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            (async () => {
                  const chirp = new Call("/v1/api/donate/intent", {
                        amount: donation,
                        currency: currency,
                  });
                  chirp.method = "POST";
                  await chirp.execute();
                  if (chirp.ok()) {
                        setClientSecret(chirp.data("clientSecret") || "");
                  }
            })();
      }, [donation, currency]);   

      //||------------------------------------------------------------------------------------------------||
      //|| Fetch donation data (crypto + address)
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {            
            (async() => {
                  const chirp = new Call("/v1/api/donate")
                  chirp.method = "GET";
                  chirp.debug = true;
                  await chirp.execute();
                  setDonateData(chirp.responsePayload.data as DonateApiResponse);
            })();
      }, []);

      //||------------------------------------------------------------------------------------------------||
      //|| Render
      //||------------------------------------------------------------------------------------------------||

      return (
		<main className="min-h-screen flex flex-col bg-base-100 text-base-content">
			<NavMain />

			{/* Hero */}
                  <MainHero
                        image="/img/hero/donate.webp"
                        title="Support Our Mission"
                        description="Your donation keeps this project open, ad-free, and independent. We never sell data, never show tracking ads, and your gift directly funds real compliance & privacy tools for everyone."
                  />

			{/* Donation Method Buttons - OUTSIDE */}
			<section className="w-full max-w-5xl mx-auto mt-10 mb-0 flex flex-col items-center">
				<div className="flex justify-center gap-2 w-full mb-7">
					<button
						className={`flex-1 flex items-center justify-center gap-2 px-5 py-4 rounded-2xl font-bold transition-colors text-xl
                                          ${activeTab === "card" ? "bg-blue-500/60 text-white shadow-lg" : "bg-gray-700 text-white/80 hover:bg-blue-500/70"}`}
						onClick={() => setActiveTab("card")}>
						<CreditCard className="w-6 h-6" /> Credit Card
					</button>
					<button
						className={`flex-1 flex items-center justify-center gap-2 px-5 py-4 rounded-2xl font-bold transition-colors text-xl
                                          ${ activeTab === "crypto" ? "bg-blue-500/60 text-white shadow-lg" : "bg-gray-700 text-white/80 hover:bg-blue-500/70"}`}
						onClick={() => setActiveTab("crypto")}>
						<Bitcoin className="w-6 h-6" /> Crypto
					</button>
					<button
						className={`flex-1 flex items-center justify-center gap-2 px-5 py-4 rounded-2xl font-bold transition-colors text-xl
                                          ${activeTab === "cash" ? "bg-blue-500/60 text-white shadow-lg" : "bg-gray-700 text-white/80 hover:bg-blue-500/70"}`}
						onClick={() => setActiveTab("cash")}>
						<MailIcon className="w-6 h-6" /> Cash/Check
					</button>
				</div>
			</section>

			{/* Donation UI */}
			<section className="py-8 px-4 w-full max-w-5xl mx-auto flex flex-col gap-10">
                        {activeTab === "card"   && <DonateCard />}
                        {activeTab === "crypto" && <DonateCrypto />}
                        {activeTab === "cash"   && <DonateCheck />}
			</section>

			{/* Progress */}
			<MonthlySupportHero cta={ false } />

			<FooterMain />
		</main>
	);
}
