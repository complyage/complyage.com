/*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
//|| /components/base/BIP39Six.tsx
//|| Fixed 6 BIP39 word input component with dropdown auto-complete
//||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

      /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
      //|| Import
      //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

      import React, {useState, useEffect}                   from "react";
      import {bip39}                                        from "../../utils/bip39";
      import { RefreshCcw }                                    from "lucide-react";

      /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
      //|| Interfaces
      //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

      import { BIPList }                                    from "../../interfaces/bip.list";

      /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
      //|| Props
      //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

      interface BIP39Props {
            initialWords                 : BIPList;
            setValue                     : (bipList : BIPList) => void;
      }

      /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
      //|| Component
      //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

      export default function BIP39({initialWords, setValue}: BIP39Props) {

            /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
            //|| Var
            //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

            const [words, setWords]                   = useState<BIPList>(initialWords);
            const [focusedIndex, setFocusedIndex]     = useState<number | null>(null);
            const wordlist                            = bip39();

            /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
            //|| Update Word
            //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

            const updateWord = (index: number, value: string) => {
                  // create a new copy so React detects the change
                  const updatedWords: BIPList = { ...words };
            
                  const key = `word${index + 1}` as keyof BIPList;
                  updatedWords[key] = value;
            
                  setWords(updatedWords);
                  setValue(updatedWords);
            };

            /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
            //|| Get Suggestions
            //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

            const getSuggestions = (input: string) => (input ? wordlist.filter((w) => w.startsWith(input.toLowerCase())).slice(0, 10) : []);

            /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
            //|| IsValid
            //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

            const isValid = (word: string) => wordlist.includes(word);

            /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
            //|| JSX
            //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/
            
            return (
                  <div className="bg-gray-800 p-4 rounded-lg">
                        <h3 className="text-2xl font-bold mb-4 text-white text-center">Your 6-Word Passphrase</h3>

                        {/* Fixed 2x3 grid */}
                        <div className="grid grid-cols-3 grid-rows-2 gap-4">
                              {Object.values(words).map((word, index) => {
                                    const suggestions = getSuggestions(word);

                                    // function to generate a random word for this index
                                    const randomizeWord = () => {
                                          const random = wordlist[Math.floor(Math.random() * wordlist.length)];
                                          updateWord(index, random);
                                    };

                                    return (
                                          <div key={index} className="relative bg-black p-2 rounded flex flex-col gap-2">
                                                <label className="block text-center font-bold text-gray-300">{index + 1}</label>

                                                <div className="flex gap-2">
                                                      <input
                                                            value={word}
                                                            onChange={(e) => updateWord(index, e.target.value)}
                                                            onFocus={() => setFocusedIndex(index)}
                                                            onBlur={() => setTimeout(() => setFocusedIndex(null), 120)}
                                                            className={`input input-bordered w-full text-xl text-center text-yellow-500 ${
                                                                  word && !isValid(word) ? "input-error" : "input-primary"
                                                            }`}
                                                            placeholder={`Word ${index + 1}`}
                                                            autoComplete="off"
                                                      />
                                                      <button
                                                            type="button"
                                                            onClick={randomizeWord}
                                                            className="btn btn-xs btn-accent text-white bg-black/60 border-0 shadow-none mt-3"
                                                            title="Generate random word"
                                                      >
                                                            <RefreshCcw />
                                                      </button>
                                                </div>

                                                {focusedIndex === index && suggestions.length > 0 && (
                                                      <ul className="absolute z-20 w-full bg-black border border-gray-600 shadow max-h-40 overflow-y-auto rounded text-sm text-white">
                                                            {suggestions.map((suggestion) => (
                                                                  <li
                                                                        key={suggestion}
                                                                        onMouseDown={() => updateWord(index, suggestion)}
                                                                        className="px-3 py-2 hover:bg-gray-200 dark:hover:bg-neutral-700 cursor-pointer"
                                                                  >
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
