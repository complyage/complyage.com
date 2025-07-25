/*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
//|| /components/base/BIP39.tsx
//|| BIP39 word input component with dropdown auto-complete
//||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

      import React, { useState, useEffect }                 from "react";
      import { bip39 }                                      from "../../utils/bip39";

      /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
      //|| Props
      //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

      interface TextBIP39Props {
            count                         : number;
            initialWords?                 : string[];
            setValue?                     : (words: string[]) => void;
      }

      /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
      //|| Component
      //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

      export default function TextBIP39({count, initialWords, setValue}: TextBIP39Props) {
		const wordlist = bip39();

		const getRandomWords = (n: number): string[] => Array.from({length: n}, () => wordlist[Math.floor(Math.random() * wordlist.length)]);

		const [words, setWords] = useState<string[]>(initialWords?.slice(0, count) || getRandomWords(count));
		const [focusedIndex, setFocusedIndex] = useState<number | null>(null);

		useEffect(() => {
			const newWords = initialWords?.slice(0, count) || getRandomWords(count);
			setWords(newWords);
			if (setValue) setValue(newWords);
		}, [count]);

		const updateWord = (index: number, value: string) => {
			const newWords = [...words];
			newWords[index] = value.trim().toLowerCase();
			setWords(newWords);
			if (setValue) setValue(newWords);
		};

		const getSuggestions = (input: string) => (input ? wordlist.filter((w) => w.startsWith(input.toLowerCase())).slice(0, 10) : []);

		const isValid = (word: string) => wordlist.includes(word);

		return (
			<div className="bg-gray-800 p-4 rounded-lg">
				<h3 className="text-2xl font-bold my-3">Choose your word list</h3>
				<div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-6 gap-4 min-h-[300px]">
					{words.map((word, index) => {
						const suggestions = getSuggestions(word);

						return (
							<div key={index} className="relative bg-black p-2">
								<label className="text-center block font-bold mb-2">{`${index + 1}`}</label>
								<input
									value={word}
									onChange={(e) => updateWord(index, e.target.value)}
									onFocus={() => setFocusedIndex(index)}
									onBlur={() => setTimeout(() => setFocusedIndex(null), 120)}
									className={`input input-bordered w-full text-xl text-center ${
										word && !isValid(word) ? "input-error" : "input-primary"
									}`}
									placeholder={`Word ${index + 1}`}
									autoComplete="off"
								/>

								{focusedIndex === index && suggestions.length > 0 && (
									<ul className="absolute z-20 w-full bg-white border border-gray-300 shadow max-h-40 overflow-y-auto rounded text-sm">
										{suggestions.map((suggestion) => (
											<li
												key={suggestion}
												onMouseDown={() => updateWord(index, suggestion)}
												className="px-3 py-2 bg-gray-800 hover:bg-gray-400 cursor-pointer text-white">
												{suggestion}
											</li>
										))}
									</ul>
								)}
							</div>
						);
					})}
				</div>
			</div>
		);
	}

      /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
      //|| EOC
      //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/
