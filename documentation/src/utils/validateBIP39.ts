//||------------------------------------------------------------------------------------------------||
//|| utils/bip39.validate.ts
//|| Validate BIP39 mnemonics no checksum, just basic word list and count checks
//||------------------------------------------------------------------------------------------------||

import getBIP39                                        from "../data/getBIP39";

//||------------------------------------------------------------------------------------------------||
//|| Strength
//||------------------------------------------------------------------------------------------------||


export type BIPStrength = "weak" | "standard" | "strong" | "max";

export interface BIPValidateResult {
  valid     : boolean;
  reason?   : string;
  strength? : BIPStrength;
  wordCount : 6 | 12 | 18 | 24;
}

function normalizeWords(words: string[]): string[] {
  return words.map(w => (w || "").trim().toLowerCase());
}

function strengthForCount(n: number): BIPStrength {
  switch (n) {
    case 6: return "weak";
    case 12: return "standard";
    case 18: return "strong";
    case 24: return "max";
    default: return "weak";
  }
}

export async function validateBIP39(wordsIn: string[]): Promise<BIPValidateResult> {
  const list   = getBIP39();
  const words  = normalizeWords(wordsIn) as (string[]);
  const count  = words.length as 6 | 12 | 18 | 24;

  if (![6, 12, 18, 24].includes(count)) {
    return { valid: false, reason: "Invalid word count", wordCount: 6 };
  }

  const set = new Set(list);
  if (!words.every(w => set.has(w))) {
    return { valid: false, reason: "One or more words not in list", wordCount: count };
  }

  return { valid: true, strength: strengthForCount(count), wordCount: count };
}