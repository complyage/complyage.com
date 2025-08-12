/*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
//|| /components/base/KeyPairGenerator.tsx
//|| Generates and displays private/public keys from /auth/generate with 5s reload cooldown
//||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

import React, { useState, useEffect } from "react";

interface KeyPairGeneratorProps {
      setValue?                     : (keys: { privateKey: string; publicKey: string }) => void;
      defaultPublic                 : string;
      defaultPrivate                : string;
}

/*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
//|| Fetch Keys
//||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

export default function KeyPairGenerator({ setValue, defaultPublic, defaultPrivate }: KeyPairGeneratorProps) {

      /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
      //|| Var
      //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

      const [privateKey, setPrivateKey]         = useState<string>(defaultPrivate);
      const [publicKey, setPublicKey]           = useState<string>(defaultPublic);
      const [loading, setLoading]               = useState<boolean>(true);
      const [cooldown, setCooldown]             = useState<number>(0);

      /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
      //|| Fetch Keys
      //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

      const fetchKeys = async () => {
            try {
                  setLoading(true);
                  const res = await fetch("/auth/generate", { method: "GET" });
                  const json = await res.json();
                  if (json.data?.privateKey && json.data?.publicKey) {
                        setPrivateKey(json.data.privateKey);
                        setPublicKey(json.data.publicKey);
                        setValue?.({ privateKey: json.data.privateKey, publicKey: json.data.publicKey });
                  }
            } catch (err) {
                  console.error("Failed to fetch keys", err);
            } finally {
                  setLoading(false);
                  setCooldown(5); // start 5-second cooldown
            }
      };

      /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
      //|| Countdown Timer
      //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

      useEffect(() => {
            if (cooldown <= 0) return;
            const timer = setInterval(() => setCooldown(c => c - 1), 1000);
            return () => clearInterval(timer);
      }, [cooldown]);

      /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
      //|| Initial Load
      //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/

      useEffect(() => {
            if (!defaultPrivate || !defaultPublic) fetchKeys(); else setLoading(false);
      }, []);

      /*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
      //|| JSX
      //||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/
      return (
            <div className="bg-gray-800 p-4 rounded-lg">
                  <div className="flex justify-between items-center mb-4">
                        <h3 className="text-2xl font-bold text-white">Key Pair</h3>
                        <button
                              type="button"
                              onClick={fetchKeys}
                              className="btn btn-sm btn-secondary"
                              disabled={loading || cooldown > 0}
                        >
                              {loading ? "Loading..." : cooldown > 0 ? `Reload (${cooldown}s)` : "Reload"}
                        </button>
                  </div>

                  <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                        <div>
                              <label className="block text-gray-300 font-bold mb-2">Private Key</label>
                              <textarea
                                    value={privateKey}
                                    readOnly
                                    className="textarea textarea-bordered w-full h-128 font-mono text-xs bg-black text-green-300"
                              />
                        </div>

                        <div>
                              <label className="block text-gray-300 font-bold mb-2">Public Key</label>
                              <textarea
                                    value={publicKey}
                                    readOnly
                                    className="textarea textarea-bordered w-full h-128 font-mono text-xs bg-black text-green-300"
                              />
                        </div>
                  </div>
            </div>
      );
}

/*||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||
//|| EOC
//||=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-||*/
