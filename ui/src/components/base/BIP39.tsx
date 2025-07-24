import React, {useState} from "react";
import wordlist from "bip39/wordlists/english.json";

export default function BIP39MnemonicInput() {
	const [words, setWords] = useState(Array(12).fill(""));

	const handleWordChange = (index: number, value: string) => {
		const newWords = [...words];
		newWords[index] = value.trim().toLowerCase();
		setWords(newWords);
	};

	const isWordValid = (word: string) => {
		return wordlist.includes(word);
	};

	return (
		<div className="grid grid-cols-2 md:grid-cols-3 gap-4">
			{words.map((word, index) => (
				<div key={index} className="flex flex-col">
					<input
						list="bip39-words"
						value={word}
						onChange={(e) =>
							handleWordChange(
								index,
								e.target.value
							)
						}
						className={`input input-bordered ${
							word && !isWordValid(word)
								? "input-error"
								: "input-primary"
						}`}
						placeholder={`Word ${index + 1}`}
					/>
				</div>
			))}

			<datalist id="bip39-words">
				{wordlist.map((w) => (
					<option key={w} value={w} />
				))}
			</datalist>
		</div>
	);
}
