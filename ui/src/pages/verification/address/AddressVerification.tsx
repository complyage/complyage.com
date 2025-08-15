//||------------------------------------------------------------------------------------------------||
//|| AddressVerification (Container)
//|| src/components/security/AddressVerification.tsx
//||------------------------------------------------------------------------------------------------||

import React, {useMemo, useRef, useState}                         from "react";
import {useNavigate}                                              from "react-router-dom";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import MembersLayout                                              from "../../../layouts/MembersLayout";
import InlineAlert                                                from "../../../components/base/InlineAlert";
import ProgressSteps, { ProgessStep }                             from "../../../components/base/ProgressSteps";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

import {Address, CardData, Country}                               from "../../../interfaces/verification.location";

//||------------------------------------------------------------------------------------------------||
//|| Utils
//||------------------------------------------------------------------------------------------------||

import {isAddressValid}                                           from "../../../utils/validate";

//||------------------------------------------------------------------------------------------------||
//|| Steps
//||------------------------------------------------------------------------------------------------||

import Step1 from "./AddressVerification.Step1";
import Step2 from "./AddressVerification.Step2";
import Step3 from "./AddressVerification.Step3";

//||------------------------------------------------------------------------------------------------||
//|| Data
//||------------------------------------------------------------------------------------------------||

const countries: Country[] = [
	{code: "US", name: "United States", flag: "🇺🇸"},
	{code: "CA", name: "Canada", flag: "🇨🇦"},
	{code: "GB", name: "United Kingdom", flag: "🇬🇧"},
	{code: "AU", name: "Australia", flag: "🇦🇺"},
	{code: "DE", name: "Germany", flag: "🇩🇪"},
];

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function AddressVerification() {

      //||------------------------------------------------------------------------------------------------||
      //|| Navigate
      //||------------------------------------------------------------------------------------------------||

      const navigate          = useNavigate();
	const currentStep       = useRef<1 | 2 | 3>(1);

      //||------------------------------------------------------------------------------------------------||
      //|| Recieved
      //||------------------------------------------------------------------------------------------------||

      const [country, setCountry]               = useState<Country>(countries[0]);
	const [addr, setAddr]                     = useState<Address>({
		line1       : "",
		line2       : "",
		city        : "",
		state       : "",
		postal      : "",
		country     : countries[0].code,
	});

      //||------------------------------------------------------------------------------------------------||
      //|| Standardized
      //||------------------------------------------------------------------------------------------------||

	const [std, setStd]                       = useState<Address | null>(null);
	const [notes, setNotes]                   = useState<string>("");

      //||------------------------------------------------------------------------------------------------||
      //|| Credit Card
      //||------------------------------------------------------------------------------------------------||

      const [card, setCard]                     = useState<CardData>({
		number                  : "",
		expMonth                : "",
		expYear                 : "",
		cvc                     : "",
		postal                  : "",
	});

      //||------------------------------------------------------------------------------------------------||
      //|| Status
      //||------------------------------------------------------------------------------------------------||
	
      const [busy, setBusy]                     = useState<boolean>(false);
	const [serverMsg, setServerMsg]           = useState<string>("");

      //||------------------------------------------------------------------------------------------------||
      //|| Derived
      //||------------------------------------------------------------------------------------------------||

	const addrOk = useMemo(() => isAddressValid(addr), [addr]);
	const cardOk = true;

      //||------------------------------------------------------------------------------------------------||
      //|| Tick
      //||------------------------------------------------------------------------------------------------||

      const [, setTick] = useState(0);
	const goStep = (s: 1 | 2 | 3) => {
		currentStep.current = s;
		setTick((t) => t + 1);
	};

      //||------------------------------------------------------------------------------------------------||
      //|| Standardize Address
      //||------------------------------------------------------------------------------------------------||

      const handleStandardize = async () => {
		if (!addrOk) return;
		setBusy(true);
		setServerMsg("");
		try {
			// TODO: Call your standardization service here:
			// const resp = await fetch("/v1/api/address/standardize", {...});
			// const data = await resp.json(); setStd(data.address); setNotes(data.notes ?? "");

			// Demo normalization:
			const normalized: Address = {
				...addr,
				line1: addr.line1.trim().toUpperCase(),
				line2: (addr.line2 || "").trim().toUpperCase(),
				city: addr.city.trim().toUpperCase(),
				state: addr.state.trim().toUpperCase(),
				postal: addr.postal.trim().toUpperCase(),
				country: addr.country,
			};
			setStd(normalized);
			setNotes("Standardized via demo normalizer");
			goStep(2);
		} catch {
			setServerMsg("Failed to standardize address.");
		} finally {
			setBusy(false);
		}
	};

      //||------------------------------------------------------------------------------------------------||
      //|| Continue
      //||------------------------------------------------------------------------------------------------||

	const handleConfirmAddress = () => {
		if (!std) return;
		goStep(3);
	};

      //||------------------------------------------------------------------------------------------------||
      //|| Submit Payment
      //||------------------------------------------------------------------------------------------------||

	const handleSubmitPayment = async () => {
		if (!cardOk || !std) return;
		setBusy(true);
		setServerMsg("");
		try {
			// TODO: Call your $1 verification endpoint here with { addressOriginal: addr, addressConfirmed: std, card }
			// await fetch("/v1/api/address/verify-payment", {...});
			await new Promise((r) => setTimeout(r, 400)); // demo
			navigate("/members/thanks");
		} catch {
			setServerMsg("Card verification failed. Please check your details and try again.");
		} finally {
			setBusy(false);
		}
	};


      //||------------------------------------------------------------------------------------------------||
      //|| Stepper
      //||------------------------------------------------------------------------------------------------||

      const steps: ProgessStep[] = [
            { label: "Enter Address", description: "Provide your physical address details." },
            { label: "Confirm Address", description: "Review and confirm your address." },
            { label: "Card $1 Check", description: "Verify your card for $1." },
      ]

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

	return (
		<MembersLayout title="Physical Address Verification">
			<div className="w-full max-w-2xl mx-auto">
				{/* Stepper */}
                        <ProgressSteps steps={ steps } currentStep={ currentStep.current } className="mb-6" />				

				{/* Inline hint on step 1 input */}
				{currentStep.current === 1 && (addr.line1 || addr.city || addr.state || addr.postal) ? (
					addrOk ? (
						<InlineAlert message="Looks good." isError={false} />
					) : (
						<InlineAlert message="Please complete all required address fields." isError={true} />
					)
				) : null}

                        { /* Step 1. */ }

				{currentStep.current === 1 && (
					<Step1
						countries={countries}
						country={country}
						setCountry={setCountry}
						addr={addr}
						setAddr={setAddr}
						addrOk={addrOk}
						busy={busy}
						serverMsg={serverMsg}
						onStandardize={handleStandardize}
					/>
				)}

                        { /* Step 2. */ }

				{currentStep.current === 2 && std && (
					<Step2 addr={addr} std={std} notes={notes} onBack={() => goStep(1)} onConfirm={handleConfirmAddress} />
				)}

                        { /* Step 3 . */ }

				{currentStep.current === 3 && (
					<Step3
						card={card}
						setCard={setCard}
						cardOk={cardOk}
						busy={busy}
						serverMsg={serverMsg}
						onBack={() => goStep(2)}
						onSubmit={handleSubmitPayment}
					/>
				)}
			</div>
		</MembersLayout>
	);

      //||------------------------------------------------------------------------------------------------||
      //|| EOF
      //||------------------------------------------------------------------------------------------------||

}
