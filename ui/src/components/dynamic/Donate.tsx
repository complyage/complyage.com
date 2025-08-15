//||------------------------------------------------------------------------------------------------||
//|| Donation
//|| src/components/payments/Donation.tsx
//||------------------------------------------------------------------------------------------------||

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, {useEffect, useRef, useState}                    from "react";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

type DonationProps = {
	donation                : number;
	setDonation             : (amount: number) => void;
};

const clamp = (v: number, min: number, max: number) => Math.max(min, Math.min(max, v));

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function Donation({donation, setDonation}: DonationProps) {

      //||------------------------------------------------------------------------------------------------||
      //|| Const
      //||------------------------------------------------------------------------------------------------||

      const presets   = [0, 1, 5, 10, 50, 100];
	const presetSet = React.useMemo(() => new Set(presets), [presets]);
	const isCustom = !presetSet.has(donation);

      //||------------------------------------------------------------------------------------------------||
      //|| Preset
      //||------------------------------------------------------------------------------------------------||

	const setPreset = (amt: number) => setDonation(clamp(Math.floor(amt), 0, 1_000_000));

      //||------------------------------------------------------------------------------------------------||
      //|| Custom Change
      //||------------------------------------------------------------------------------------------------||

	const onCustomChange: React.ChangeEventHandler<HTMLInputElement> = (e) => {
		const raw = Number(e.target.value);
		const val = Number.isFinite(raw) ? Math.floor(raw) : 0;
		setPreset(val);
	};

      //||------------------------------------------------------------------------------------------------||
      //|| Custom Change
      //||------------------------------------------------------------------------------------------------||

      return (
		<div className="">
			<div className="flex items-center gap-2 flex-wrap" role="group" aria-label="Choose donation amount">
				{presets.map((amt) => {
					const active = donation === amt;
					return (
						<button
							key={amt}
							type="button"
							aria-pressed={active}
							onClick={() => setPreset(amt)}
							className={[
								"px-3 py-2 rounded-2xl text-sm font-bold transition-colors",
								active ? "bg-orange-500 text-white" : "bg-gray-600 text-white/90 hover:bg-gray-500",
							].join(" ")}>
							{amt === 0 ? "No Donation" : `$${amt}`}
						</button>
					);
				})}

				{/* Custom amount */}
				<div
					className={[
						"flex items-center gap-2 px-2 py-1 rounded-2xl",
						isCustom ? "ring-2 ring-orange-400/60" : "ring-1 ring-transparent",
					].join(" ")}>
					<span className="text-sm">Custom</span>
					<input
						type="number"
						min={0}
						step={1}
						value={donation}
						onChange={onCustomChange}
						className="input input-bordered w-24 h-9"
						aria-label="Custom donation amount in dollars"
					/>
				</div>
			</div>
		</div>
	);
}
