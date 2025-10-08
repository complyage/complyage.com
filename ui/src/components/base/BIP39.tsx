/*||------------------------------------------------------------------------------------------------||/
/|| BIP39Six.tsx — `suggest` only controls auto-generation on mount; autocomplete always on
/||------------------------------------------------------------------------------------------------||*/

import React, {useEffect, useState} from "react";
import {X} from "lucide-react";

import BIP39Text from "./BIP39.Text";

/*||------------------------------------------------------------------------------------------------||/
/|| Props
/||------------------------------------------------------------------------------------------------||*/

interface BIP39Props {
	mode: "CREATE" | "VERIFY";
	wordCount: 6 | 12 | 18 | 24;
	setValue: (words: string[]) => void;
}

/*||------------------------------------------------------------------------------------------------||/
/|| Helpers
/||------------------------------------------------------------------------------------------------||*/

function colsForCount(count: number) {
	switch (count) {
		case 6:
			return 3;
		case 12:
			return 4;
		case 18:
			return 4;
		case 24:
			return 4;
		default:
			return 3;
	}
}

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function BIP39({mode, wordCount, setValue}: BIP39Props) {
	const makeEmpty = (n: number) => Array.from({length: n}, () => "");
	const [words, setWords] = useState<string[]>(() => makeEmpty(wordCount));
      //||------------------------------------------------------------------------------------------------||
      //|| Component
      //||------------------------------------------------------------------------------------------------||
	useEffect(() => {
		setWords((prev) => {
			if (prev.length === wordCount) return prev;
			const next = makeEmpty(wordCount);
			for (let i = 0; i < Math.min(prev.length, next.length); i++) next[i] = prev[i];
			setValue(next);
			return next;
		});
	}, [wordCount]);
      //||------------------------------------------------------------------------------------------------||
      //|| Update Word List
      //||------------------------------------------------------------------------------------------------||
	const updateWord = (index: number, value: string) => {
		setWords((prev) => {
			const next = [...prev];
			next[index] = value;
			// bubble up full array to consumer
			setValue(next);
			return next;
		});
	};
      //||------------------------------------------------------------------------------------------------||
      //|| Update Word List
      //||------------------------------------------------------------------------------------------------||
	const clearAll = () => {
            setWords(makeEmpty(wordCount));
	};      
      //||------------------------------------------------------------------------------------------------||
      //|| Grid Layout
      //||------------------------------------------------------------------------------------------------||
	const cols                           = colsForCount(wordCount);
	const gridStyle: React.CSSProperties = {gridTemplateColumns: `repeat(${cols}, minmax(0, 1fr))`};
      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||
	return (
		<div className="bg-gray-800 rounded-lg">
                  {mode === "CREATE" && (
                        <div className="flex justify-between items-center p-4 border-b border-gray-600">
                              <span className="font-semibold">Your Recovery Phrase</span>
                              <button
                                    className="btn btn-sm btn-secondary"
                                    type="button"
                                    onClick={() => clearAll()}
                                    title="Regenerate"
                              >
                                    Clear all
                                    <X className="inline ml-2 mb-1" />
                              </button>
                        </div>
                  )}
			<div className="grid gap-4" style={gridStyle}>
				{words.map((w, i) => (
					<BIP39Text key={i} index={i} value={w} setValue={updateWord} mode={mode} />
				))}
			</div>
		</div>
	);
}

/*||------------------------------------------------------------------------------------------------||/
/|| EOC
/||------------------------------------------------------------------------------------------------||*/
