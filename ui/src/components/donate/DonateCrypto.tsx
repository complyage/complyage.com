/*||------------------------------------------------------------------------------------------------||
//|| DonateCrypto Component
//|| src/components/donate/DonateCrypto.tsx
//||------------------------------------------------------------------------------------------------||*/

import React, { useState, useEffect } from "react";
import { Copy, Check } from "lucide-react";

import Call from "../../classes/call";
import type { DonateApiResponse } from "../../interfaces/donate/donate";

export default function DonateCrypto() {
      const [crypto, setCrypto] = useState<DonateApiResponse["crypto"]>([]);
      const [cryptoIdx, setCryptoIdx] = useState(0);
      const [copied, setCopied] = useState(false);

      //||------------------------------------------------------------------------------------------------||
      //|| Fetch crypto donation options
      //||------------------------------------------------------------------------------------------------||
      useEffect(() => {
            (async () => {
                  const chirp = new Call("/v1/api/donate");
                  chirp.method = "GET";
                  await chirp.execute();
                  if (chirp.ok()) {
                        setCrypto((chirp.responsePayload.data as DonateApiResponse).crypto || []);
                  }
            })();
      }, []);

      if (!crypto || crypto.length === 0) return null;

      const selected = crypto[cryptoIdx];

      return (
            <div className="w-full mt-7 flex flex-col md:flex-row gap-6">
                  {/* Left: Vertical crypto list */}
                  <div className="flex flex-col w-full md:w-1/3 gap-2">
                        {crypto.map((c, i) => (
                              <button
                                    key={c.address}
                                    onClick={() => setCryptoIdx(i)}
                                    className={`flex flex-col items-center justify-center p-2 font-bold transition-all ${
                                          cryptoIdx === i
                                                ? "bg-yellow-400 text-black border-yellow-600 shadow-lg scale-105"
                                                : "bg-black/10 border-none hover:bg-yellow-500/40"
                                    }`}>
                                    <span className="uppercase text-sm">{c.name}</span>
                                    <span className="mt-1 text-xs opacity-70">{c.symbol || ""}</span>
                              </button>
                        ))}
                  </div>

                  {/* Right: QR + address */}
                  <div className="flex flex-col items-center rounded-2xl shadow-lg p-6 w-full md:w-2/3">
                        <div className="mb-3 text-2xl font-bold tracking-wide uppercase">{selected.name}</div>
                        <img
                              src={`${import.meta.env.VITE_COMPLYAGE_API_URL}/public/img/qr?data=${btoa(
                                    selected.prefix + selected.address
                              )}`}
                              alt={`${selected.name} QR`}
                              className="rounded-xl w-44 h-44 mx-auto border-4 border-white shadow-lg"
                        />

                        <div className="flex items-center gap-2 mt-3 bg-white/10 p-3 rounded text-sm font-mono break-all w-full">
                              <span className="flex-1 text-center text-lg">{selected.address}</span>
                              <button
                                    onClick={() => {
                                          navigator.clipboard.writeText(selected.address);
                                          setCopied(true);
                                          setTimeout(() => setCopied(false), 2000);
                                    }}
                                    className="p-1 hover:bg-white/20 rounded"
                                    title="Copy to clipboard">
                                    {copied ? <Check className="w-4 h-4 text-green-400" /> : <Copy className="w-4 h-4" />}
                              </button>
                        </div>

                        <div className="text-xs mt-2 opacity-60">Scan QR or copy address</div>
                  </div>
            </div>
      );
}
