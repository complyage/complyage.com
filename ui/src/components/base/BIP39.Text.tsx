//||------------------------------------------------------------------------------------------------||
//|| components/base/BIP39Text.tsx
//|| Small self-contained input + suggestions + randomize component
//||------------------------------------------------------------------------------------------------||

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, {useState, useEffect, useMemo, useRef}              from "react";
import {RefreshCcw}                                   from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Data
//||------------------------------------------------------------------------------------------------||

import getBIP39                                       from "../../data/getBIP39";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface BIP39TextProps {
	index             : number;
      value             : string;
	setValue          : (index: number, value: string) => void;
	mode              : "CREATE" | "VERIFY";
}

//||------------------------------------------------------------------------------------------------||
//|| JSX
//||------------------------------------------------------------------------------------------------||

export default function BIP39Text({index, value, setValue, mode}: BIP39TextProps) {
      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||
      const [focused, setFocused]               = useState<boolean>(false);
      const [valid, setValid]                   = useState<boolean>(false);
      const [text, setText]                     = useState<string>(value || "");
	const containerRef                        = useRef<HTMLDivElement | null>(null);
      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||
      const bip39 = useMemo(() => getBIP39(), []);
      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||
	const q           = (text || "").trim().toLowerCase();
	const suggestions = text.length ? bip39.filter((w) => w.startsWith(q)).slice(0, 10) : [];
      useEffect(() => {
            if (mode === "CREATE") {
                  const r = getRandom();
                  setText(r);
                  setValid(true);
                  setValue(index, r);
            }
      }, []);
      //||------------------------------------------------------------------------------------------------||
      //|| Handle the Select
      //||------------------------------------------------------------------------------------------------||
      useEffect(() => {
            setText(value || "");
            setValid(value !== "" && bip39.includes(value));
      }, [value]);      
      //||------------------------------------------------------------------------------------------------||
      //|| Handle the Select
      //||------------------------------------------------------------------------------------------------||
	const handleSelect = (suggestion: string) => {
            setValid(bip39.includes(suggestion));
            setText(suggestion);
            if (bip39.includes(suggestion) && suggestion !== "") setValue(index, suggestion);
		window.requestAnimationFrame(() => {
			const el = containerRef.current?.querySelector("input") as HTMLInputElement | null;
			el?.focus();
		});
	};
      //||------------------------------------------------------------------------------------------------||
      //|| Handle
      //||------------------------------------------------------------------------------------------------||
      const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
            const x = e.target.value;
            setText(x);
            setValid(bip39.includes(x));
            if (bip39.includes(x) && x !== "") setValue(index, x);
      }
      //||------------------------------------------------------------------------------------------------||
      //|| Get Random
      //||------------------------------------------------------------------------------------------------||
      const getRandom = () => {     
            return bip39[Math.floor(Math.random() * bip39.length)];
      }
      //||------------------------------------------------------------------------------------------------||
      //|| Set Random
      //||------------------------------------------------------------------------------------------------||
      const setRandom = () => {
            const r = getRandom();
            setText(r);
            setValid(true);
            setValue(index, r);
      }
      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||
	return (
		<div ref={containerRef} className="relative bg-black p-2 rounded flex flex-col gap-2 overflow-visible">
			<label className="block text-center font-bold text-gray-300">{index + 1}</label>

			<div className="flex gap-2">
				<input
					value={text}
					onChange={(e) => handleChange(e)}
					onFocus={() => setFocused(true)}
					onBlur={() => setTimeout(() => setFocused(false), 150)}
					className={`input input-bordered w-full text-xl text-center ${ (text === "") ? "" : !valid ? "text-red-400 input-error" : "input-primary" }`}
					placeholder={`Word ${index + 1}`}
					autoComplete="off"
				/>

				{ ( mode === "CREATE" ) && (
					<button
						type="button"
						onClick={ setRandom }
						className="btn btn-xs btn-accent text-white bg-black/60 border-0 shadow-none mt-3"
						title="Generate random word">
						<RefreshCcw />
					</button>
				)}
			</div>

			{focused && suggestions.length > 0 && (
				<ul className="absolute z-50 w-full bg-black border border-gray-600 shadow max-h-40 overflow-y-auto rounded text-sm text-white">
					{suggestions.map((s) => (
						<li key={s} onMouseDown={() => handleSelect(s)} className="px-3 py-2 hover:bg-gray-600 cursor-pointer">
							{s}
						</li>
					))}
				</ul>
			)}
		</div>
	);
}
