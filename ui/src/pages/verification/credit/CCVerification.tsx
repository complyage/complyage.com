//||------------------------------------------------------------------------------------------------||
//|| CCValidation
//|| src/components/payments/CCValidation.tsx
//||------------------------------------------------------------------------------------------------||

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, {useRef, useState, useMemo} from "react";
import {useNavigate} from "react-router-dom";
import {CreditCard, Edit3, DollarSign, Lock, RefreshCw} from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

type CardData = {
	nameOnCard: string;
	number: string;
	expMonth: string; // MM
	expYear: string; // YY or YYYY
	cvc: string;
};

type SubmitPayload = {
	card: CardData;
	descriptor: string; // statement descriptor to post
	includeDonation: boolean; // did the user include a donation
};

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import MembersLayout from "../../../layouts/MembersLayout";

//||------------------------------------------------------------------------------------------------||
//|| Utils
//||------------------------------------------------------------------------------------------------||

const clamp = (v: number, min: number, max: number) => Math.max(min, Math.min(max, v));

const isLetter = (ch: string) => (ch >= "A" && ch <= "Z") || (ch >= "a" && ch <= "z");
const isDigit = (ch: string) => ch >= "0" && ch <= "9";
const isAllowedDescriptorChar = (ch: string) => isLetter(ch) || isDigit(ch) || ch === " " || ch === "." || ch === "," || ch === "-" || ch === "_" || ch === "*";

const onlyDigits = (s: string) => {
	let out = "";
	for (const ch of s) if (isDigit(ch)) out += ch;
	return out;
};

const digitsAndSpaces = (s: string) => {
	let out = "";
	for (const ch of s) if (isDigit(ch) || ch === " ") out += ch;
	return out;
};

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function CCValidation() {
	//||------------------------------------------------------------------------------------------------||
	//|| Var
	//||------------------------------------------------------------------------------------------------||
	const navigate = useNavigate();
	const currentStep = useRef<1 | 2>(2);

	// Card
	const [card, setCard] = useState<CardData>({nameOnCard: "", number: "", expMonth: "", expYear: "", cvc: ""});

	// Amounts
	const [customDonation, setCustom]         = useState<string>("");
      const [donation, setDonation]             = useState<number>(0); // in cents
	const [presets, setPresets]               = useState<number[]>([0, 1, 5, 10, 50, 100]);

	// Descriptor
	const [descriptor, setDescriptor] = useState<string>("CompanyXYZ -114112");

	// UI State
	const [busy, setBusy] = useState<boolean>(false);

	//||------------------------------------------------------------------------------------------------||
	//|| Derived
	//||------------------------------------------------------------------------------------------------||

	const cardValid = useMemo(() => {
		const num = onlyDigits(card.number);
		const month = Number(card.expMonth);
		const yearN = Number(card.expYear.length === 2 ? `20${card.expYear}` : card.expYear);
		const now = new Date();
		const exp = new Date(yearN || 0, (month || 0) - 1, 1);
		const inFuture = !!month && !!yearN && exp >= new Date(now.getFullYear(), now.getMonth(), 1);
		const luhn = (digits: string) => {
			let sum = 0;
			let alt = false;
			for (let i = digits.length - 1; i >= 0; i--) {
				let n = Number(digits[i]);
				if (alt) {
					n *= 2;
					if (n > 9) n -= 9;
				}
				sum += n;
				alt = !alt;
			}
			return sum % 10 === 0;
		};
		const cvcOk = (card.cvc.length === 3 || card.cvc.length === 4) && onlyDigits(card.cvc).length === card.cvc.length;
		return card.nameOnCard.trim().length >= 2 && num.length >= 12 && luhn(num) && inFuture && cvcOk;
	}, [card]);

	//||------------------------------------------------------------------------------------------------||
	//|| Handlers
	//||------------------------------------------------------------------------------------------------||

	const goStep = (s: 1 | 2) => {
		currentStep.current = s;
		forceRender();
	};

	const selectDonation = (dollars: number) => {
		setDonation(dollars);
		setCustom("");
	};

	const applyCustom = (v: string) => {
		setCustom(v);
		let clean = "";
		for (const ch of v) if (isDigit(ch) || ch === ".") clean += ch;
		const parts = clean.split(".");
		const whole = parts[0] || "0";
		const frac = (parts[1] || "").slice(0, 2);
		setDonation(Number.isFinite(dollars) ? clamp(dollars, 0, 1_000_000) : 0);
	};

	const handleSubmit = async () => {
		if (!cardValid) return;
		setBusy(true);
		const payload: SubmitPayload = {
			card,
			descriptor: descriptor.trim()
		};
		try {
			// Wire up to your API here (e.g., create PaymentIntent with statement descriptor)
			// await fetch("/v1/payments/validate", { method: "POST", headers: {"Content-Type":"application/json"}, credentials: "include", body: JSON.stringify(payload) });
			console.log("PAY SUBMIT", payload);
			navigate("/members/thanks");
		} catch (e) {
			console.error(e);
			alert("Payment submit failed");
		} finally {
			setBusy(false);
		}
	};

	// minimal force render helper to sync ref-based step with UI
	const [, setTick] = useState<number>(0);
	const forceRender = () => setTick((t) => t + 1);

	//||------------------------------------------------------------------------------------------------||
	//|| JSX
	//||------------------------------------------------------------------------------------------------||

	return (
		<MembersLayout title="Credit Card Validation">
			<div className="w-full max-w-2xl mx-auto">
				{/* Stepper */}
				<div className="flex items-center gap-3 text-sm mb-6 select-none">
					<div className={`flex items-center gap-2 ${currentStep.current === 1 ? "font-semibold" : "opacity-60"}`}>
						<div className="w-6 h-6 rounded-full flex items-center justify-center border">1</div>
						<span>Card Details</span>
					</div>
					<div className="h-px flex-1 bg-gray-300" />
					<div className={`flex items-center gap-2 ${currentStep.current === 2 ? "font-semibold" : "opacity-60"}`}>
						<div className="w-6 h-6 rounded-full flex items-center justify-center border">2</div>
						<span>Descriptor & Amount</span>
					</div>
				</div>

				{/* Step 1: Card */}
				{currentStep.current === 1 && (
					<div className="space-y-4 p-5">
						<div className="flex items-center gap-2 text-gray-700">
							<CreditCard className="w-5 h-5" />
							<span>Enter your card details. Hook to Stripe Elements in production.</span>
						</div>

						<label className="block">
							<span className="block text-sm mb-1">Name on card</span>
							<input
								placeholder="Jane Q. Customer"
								value={card.nameOnCard}
								onChange={(e) => setCard((c) => ({...c, nameOnCard: e.target.value}))}
                                                className="input input-bordered input-primary  w-full"
                                                />
						</label>

						<label className="block">
							<span className="block text-sm mb-1">Card number</span>
							<input
								placeholder="4242 4242 4242 4242"
								inputMode="numeric"
								maxLength={19}
								value={card.number}
								onChange={(e) => setCard((c) => ({...c, number: digitsAndSpaces(e.target.value)}))}
                                                className="input input-bordered w-full"
                                                />
						</label>

						<div className="grid grid-cols-3 gap-4">
							<label className="block">
								<span className="block text-sm mb-1">Exp. MM</span>
								<input
									placeholder="MM"
									inputMode="numeric"
									maxLength={2}
									value={card.expMonth}
									onChange={(e) => setCard((c) => ({...c, expMonth: onlyDigits(e.target.value)}))}
                                                      className="input input-bordered w-full"
                                                      />
							</label>
							<label className="block">
								<span className="block text-sm mb-1">Exp. YY</span>
								<input
									placeholder="YY"
									inputMode="numeric"
									maxLength={4}
									value={card.expYear}
									onChange={(e) => setCard((c) => ({...c, expYear: onlyDigits(e.target.value)}))}
                                                      className="input input-bordered w-full"
                                                      />
							</label>
							<label className="block">
								<span className="block text-sm mb-1">CVC</span>
								<input
									placeholder="CVC"
									inputMode="numeric"
									maxLength={4}
									value={card.cvc}
									onChange={(e) => setCard((c) => ({...c, cvc: onlyDigits(e.target.value)}))}
                                                      className="input input-bordered w-full"
                                                      />
							</label>
						</div>

						<div className="flex items-center justify-between pt-2">
							<div className={`text-sm ${cardValid ? "text-green-700" : "text-red-700"}`}>
								{cardValid ? "Looks good." : "Enter a valid card, expiry, and CVC."}
							</div>
							<button
								onClick={() => goStep(2)}
								disabled={!cardValid}
								className={`inline-flex items-center gap-2 px-4 py-2 rounded-2xl border shadow-sm ${
									cardValid ? "bg-black text-white" : "bg-gray-200 text-gray-500 cursor-not-allowed"
								}`}>
								Continue
							</button>
						</div>
					</div>
				)}

				{/* Step 2: Descriptor & Amount */}
				{currentStep.current === 2 && (
					<div className="space-y-6 p-6">
				            <div className="flex items-center justify-between border-b border-gray-400 py-2 mb-4">
								<div className="flex items-center gap-2">
									<span className="font-bold">How we support the site and YOU</span>
								</div>
						</div>

                                    <p className="mt-2 mb-4 text-sm text-gray-100 leading-loose rounded-md">
                                          We believe in privacy, transparency, and independence. This small validation charge helps us keep the lights on without selling your data or bombarding you with ads. 
                                          We never store your full card number, and the statement descriptor shows exactly what your bank will see. 
                                          Your support — whether it’s just this verification or an optional donation — keeps our service free, secure, and working for you, not against you.
                                    </p>						

                                    <div className="flex items-center justify-between border-b border-gray-400 py-2 mb-4">
								<div className="flex items-center gap-2">
									<span className="font-bold">How would you like the charge to appear?</span>
								</div>
						</div>

                                    <select
                                          value={descriptor}
                                          className="w-full border border-gray-400 p-2 outline-none focus:ring-2 focus:ring-black/10"
                                    >
                                          <option className="bg-black text-white p-2" value="ComplyAge.com - [XXXX]">ComplyAge.com - [XXXX]</option>
                                          <option className="bg-black text-white p-2" value="The Privacy Foundation - [XXXX]">The Privacy Foundation - [XXXX]</option>
                                          <option className="bg-black text-white p-2" value="The Fredom Federation - [XXXX]">The Fredom Federation - [XXXX]</option>
                                    </select>

						<div>
							<div className="flex items-center justify-between border-b border-gray-400 py-2 mb-4">
								<div className="flex items-center gap-2">
									<span className="font-bold">Optional donation</span>
								</div>
							</div>

							<div className="flex flex-wrap gap-2">
								{presets.map((amount) => (
									<button
										key={amount}
										onClick={() => selectDonation(amount)}
										className={`px-3 py-2 rounded-2xl text-sm bg-gray-600 cursor-pointer font-bold ${donation === amount ? " bg-orange-500" : ""}`}>
										{amount === 0 ? "None" : "$" + amount}
									</button>
								))}
								<div className="flex items-center gap-2">
									<span className="text-sm ml-4">Custom</span>
									<input
										value={donation}
										onChange={(e) => setCustom(Math.floor(e.target.value))}
										placeholder="1549.00"
										inputMode="decimal"
										className="w-28 bg-black/40 ml-2 px-3 py-2 text-white rounded-md"
									/>
								</div>
							</div>
						</div>

                                    <div className="flex items-center justify-between border-b border-gray-400 py-2 mb-4">
                                          <div className="flex items-center gap-2">
                                                <span className="font-bold">Final total</span>
                                          </div>
						</div>                                    

						<div className="bg-black/20 p-4">
							<div className="flex items-center justify-between">
								<div className="text-xs">Base amount</div>
								<div className="text-xs font-medium">$0.XX</div>
							</div>
							<div className="flex items-center justify-between mt-1">
								<div className="text-xs">Donation</div>
								<div className="text-xs font-medium">${donation}.XX</div>
							</div>
							<div className="h-px border my-3 border-gray-500" />
							<div className="flex items-center justify-between font-semibold">
								<div className="text-md">Total</div>
								<div className="text-md">${donation}.XX</div>
							</div>
						</div>

						<div className="flex items-center justify-between">
							<button onClick={() => goStep(1)} className="px-4 py-2 rounded-2xl font-bold bg-black/30 opacity-70 hover:opacity-100">Back</button>
							<button
								onClick={handleSubmit}
								disabled={busy}
								className={`inline-flex items-center gap-2 px-5 py-2.5 rounded-2xl border shadow-sm ${
									!busy ? "bg-orange-400 cursor-pointer font-bold text-white border-0 opacity-80 hover:opacity-100" : "bg-gray-200 text-gray-500 cursor-not-allowed"
								}`}>
								<Lock className="w-4 h-4" /> Verify your card
							</button>
						</div>
					</div>
				)}
			</div>
		</MembersLayout>
	);
}
