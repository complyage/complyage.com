//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useState }                                    from "react";
import { useNavigate, useLocation }                           from "react-router-dom";
import { Lock }                                               from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Utils
//||------------------------------------------------------------------------------------------------||

import { makeBIPList }                                        from "../../utils/bip39";
import keyLevelToWordCount                                    from "../../utils/keyLevelToWordCount";
import { validateBIP39 }                                      from "../../utils/validateBIP39";

//||------------------------------------------------------------------------------------------------||
//|| API
//||------------------------------------------------------------------------------------------------||

import apiURL                                                 from "../../utils/apiURL";


//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import SpinnerCircle                                          from "../../components/base/SpinnerCircle";
import BIP39                                                  from "../../components/base/BIP39";
import KeyPairGenerator                                       from "../../components/base/KeyPairGenerator";
import InlineAlert                                            from "../../components/base/InlineAlert";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

import { BIPList }                                            from "../../interfaces/auth/bip.list";

//||------------------------------------------------------------------------------------------------||
//|| Default Component
//||------------------------------------------------------------------------------------------------||

export default function AuthComplete() {

      const navigate                                  = useNavigate();
      const [isAdvanced, setAdvanced]                 = useState<boolean>(false);
      const [wordList, setWordList]                   = useState<BIPList>(makeBIPList());
      const [password, setPassword]                   = useState("tester44");
      const [encryptionLevel, setEncryptionLevel]     = useState<number>(1);
      const [statusMessage, setStatusMessage]         = useState("");
      const [privateKey, setPrivateKey]               = useState<string>("");
      const [publicKey, setPublicKey]                 = useState<string>("");
      const [loading, setLoading]                     = useState(false);

      const location                                  = useLocation();
      const params                                    = new URLSearchParams(location.search);
      const prefix                                    = location.pathname.startsWith("/overlay") ? "/overlay" : "";

      const levels = [
            { id: 1, name: "Standard [Recommended]", description: "We securely manage your keys with enterprise-grade encryption, so you never have to worry about losing access." },
            { id: 2, name: "BIP39 (6 words)",        description: "Protect your account with a short 6-word passphrase – easy to remember, strong enough for advanced users." },
            { id: 3, name: "BIP39 (12 words)",       description: "Use a 12-word mnemonic for stronger security." },
            { id: 4, name: "BIP39 (18 words)",       description: "Stronger 18-word mnemonic for advanced users." },
            { id: 5, name: "BIP39 (24 words)",       description: "Maximum mnemonic length and security with 24 words." },
            { id: 6, name: "Generate Key Pair",      description: "Bring your own RSA private/public key pair." }
      ];

      //||------------------------------------------------------------------------------------------------||
      //|| Handle Submit
      //||------------------------------------------------------------------------------------------------||

      const handleSubmit = async (e: React.FormEvent) => {
            e.preventDefault();
            const payload = new URLSearchParams({
                  password        : password,
                  wordList        : JSON.stringify(wordList),
                  encryptionLevel : encryptionLevel.toString(),
                  privateKey      : privateKey || "",
                  publicKey       : publicKey || ""
            }).toString();

            // Levels: 2..5 => BIP39 (6/12/18/24), 6 => raw private key
            if (encryptionLevel > 1 && encryptionLevel < 6) {
                  // normalize list and validate
                  const words = (wordList || []).map(w => (w || "").trim().toLowerCase());

                  // quick UX guard
                  if (words.some(w => w === "")) {
                        setStatusMessage("Please fill every word in your recovery phrase.");
                        return;
                  }

                  // validate (dictionary + checksum where applicable)
                  const result = await validateBIP39(words);

                  if (!result.valid) {
                        setStatusMessage(`Invalid mnemonic: ${result.reason || "Unknown error"}`);
                        return;
                  }
            }

            try {
                  setLoading(true);
                  const res  = await fetch(apiURL("/auth/complete"), {
                        method   : "POST",
                        credentials : "include",                        
                        headers  : { "Content-Type": "application/x-www-form-urlencoded" },
                        body     : payload,
                  });
                  const json = await res.json();
                  if (json.success) {
                        if (prefix === "/overlay") {
                              if (window.self !== window.top) {
                                    window.parent.postMessage({ action: "refreshParent" }, "*");
                              } else {
                                    navigate("/members/");
                              }
                        } else {
                              navigate("/members/");
                        }
                        return;
                  }
                  setStatusMessage(`${json.error || "Save failed."}`);
            } catch (err) {
                  console.error(err);
                  setStatusMessage("Something went wrong. Please try again.");
            } finally {
                  setLoading(false);
            }
      };

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      return (
            <div className="relative z-10 flex items-center justify-center min-h-screen px-4 pt-24 pb-12">
                  <div className="bg-base-100 bg-opacity-90 rounded-lg shadow-lg w-full max-w-3xl p-8 text-center mx-auto">
                        <h1 className="text-3xl font-bold mb-6 text-gray-200">Complete Your Signup</h1>
                        <p className="mb-6 text-base-content/70">Set your password and encryption level.</p>
                        <InlineAlert message={statusMessage} isError={true} />

                        <form onSubmit={handleSubmit} className="flex flex-col gap-4 text-left">
                              <label className="font-semibold">Password</label>
                              <input
                                    type="password"
                                    placeholder="Enter your password"
                                    value={password}
                                    onChange={(e) => setPassword(e.target.value)}
                                    className="input input-bordered w-full py-5 text-xl h-12"
                                    required
                              />

                              <div className="flex items-center mt-4">
                                    <input
                                          type="checkbox"
                                          className="checkbox checkbox-sm mr-2"
                                          checked={isAdvanced}
                                          onChange={() => setAdvanced(!isAdvanced)}
                                    />
                                    <span className="text-sm text-gray-300">
                                          Show advanced encryption options
                                    </span>
                              </div>

                              
                        {isAdvanced && (
                              <div className="p-4 bg-gray-800 rounded-lg mt-4">
                                    
                                    <label className="font-semibold  mb-4 block"><Lock className="inline mr-2 mb-1" /> Advanced Security Options</label>

                                    <div className="flex flex-col gap-6">

                                          {/* Clickable Cards (dark mode, with highlighting + clean pros/cons) */}
                                          <div className="grid grid-cols-1 md:grid-cols-3 gap-4">

                                                {/* Level 1: Custodial */}
                                                <div
                                                      key="LEVEL.1"
                                                      onClick={() => setEncryptionLevel(1)}
                                                      className={`cursor-pointer p-4 rounded-lg text-center border-2 text-sm transition ${
                                                            encryptionLevel === 1
                                                                  ? "border-secondary bg-gray-700/80 shadow-lg ring-2 ring-primary"
                                                                  : "border-gray-600 bg-gray-800 hover:border-primary hover:bg-gray-700/60"
                                                      }`}
                                                >
                                                      <h4 className={ encryptionLevel === 1 ? "font-semibold text-orange-300 mb-3" : "font-semibold text-white mb-3"}>We Manage Your Key</h4>
                                                      <p className="text-xs text-gray-300 mb-3 h-16">
                                                            Your private key is stored and managed by us. Simple and low effort, but we can be subpoena'd for your data.
                                                      </p>
                                                      <div className="text-xs text-left space-y-2 bg-black/20 p-2 rounded">
                                                            <div><span className="font-bold text-green-200">Pros:</span> Easy setup, no backups needed</div>
                                                            <div><span className="font-bold text-red-400">Cons:</span> You rely fully on our security</div>
                                                      </div>
                                                </div>

                                                {/* Level 2-5: BIP39 */}
                                                <div
                                                      key="LEVEL.2"
                                                      onClick={() => setEncryptionLevel(3)}
                                                      className={`cursor-pointer p-4 rounded-lg text-center border-2 text-sm transition ${
                                                            encryptionLevel >= 2 && encryptionLevel <= 5
                                                                  ? "border-secondary bg-gray-700/80 shadow-lg ring-2 ring-primary"
                                                                  : "border-gray-600 bg-gray-800 hover:border-primary hover:bg-gray-700/60"
                                                      }`}
                                                >
                                                      <h4
                                                            className={
                                                                  encryptionLevel >= 2 && encryptionLevel <= 5
                                                                        ? "font-semibold text-orange-300 mb-3"
                                                                        : "font-semibold text-white mb-3"
                                                            }
                                                      >
                                                            BIP39 Mnemonic
                                                      </h4>
                                                      <p className="text-xs text-gray-300 mb-3 h-16">
                                                            Generate a recovery phrase (seed words) that only you control. You’re responsible for keeping it safe.
                                                      </p>
                                                      <div className="text-xs text-left space-y-2 mb-4 bg-black/20 p-2 rounded">
                                                            <div>
                                                                  <span className="font-semibold text-green-400">Pros:</span> Standard, portable, widely supported
                                                            </div>
                                                            <div>
                                                                  <span className="font-semibold text-red-400">Cons:</span> Must back up and secure the phrase
                                                            </div>
                                                      </div>

                                                      {/* Word Count Buttons (horizontal, dark mode) */}
                                                      <div className="mt-3">
                                                            <div className="text-xs font-semibold text-gray-300 mb-2 text-left">
                                                                  Number of Words
                                                            </div>
                                                            <div className="grid grid-cols-4 gap-2">
                                                                  {[6, 12, 18, 24].map((count, idx) => (
                                                                        <button
                                                                              key={count}
                                                                              onClick={(e) => {
                                                                                    e.stopPropagation(); // ✅ prevents parent div from firing
                                                                                    setEncryptionLevel(idx + 2);
                                                                              }}
                                                                              className={`cursor-pointer px-3 py-2 rounded-md border text-xs transition ${
                                                                                    encryptionLevel === idx + 2
                                                                                          ? "bg-secondary text-white border-primary shadow-md"
                                                                                          : "bg-gray-700 border-gray-600 text-gray-200 hover:border-primary hover:bg-gray-600"
                                                                              }`}
                                                                        >
                                                                              {count}
                                                                        </button>
                                                                  ))}
                                                            </div>
                                                      </div>
                                                </div>


                                                {/* Level 6: Full Self-Managed */}
                                                <div
                                                      key="LEVEL.3"
                                                      onClick={() => setEncryptionLevel(6)}
                                                      className={`cursor-pointer p-4 rounded-lg text-center border-2 text-sm transition ${
                                                            encryptionLevel === 6
                                                                  ? "border-secondary bg-gray-700/80 shadow-lg ring-2 ring-primary"
                                                                  : "border-gray-600 bg-gray-800 hover:border-primary hover:bg-gray-700/60"
                                                      }`}
                                                >
                                                      <h4 className={ encryptionLevel === 6 ? "font-semibold text-orange-300 mb-3" : "font-semibold text-white mb-3"}>Fully Self-Managed</h4>
                                                      <p className="text-xs text-gray-300 mb-3 h-16">
                                                            You create and secure your own cryptographic keys (e.g., hardware wallet).
                                                      </p>
                                                      <div className="text-xs text-left space-y-2 bg-black/20 p-2 rounded">
                                                            <div><span className="font-semibold text-green-400">Pros:</span> Maximum control, no third-party risk</div>
                                                            <div><span className="font-semibold text-red-400">Cons:</span> Highest responsibility, risk of total loss</div>
                                                      </div>
                                                </div>

                                          </div>


                                          {/* Show BIP39 if levels 2–5 */}
                                          {(encryptionLevel >= 2 && encryptionLevel <= 5) && (
                                                <>
                                                <div className="text-center text-gray-400">
                                                      You may choose your own or use our pre-generated list below.
                                                </div>
                                                      <BIP39
                                                            mode="CREATE"
                                                            wordCount={keyLevelToWordCount(encryptionLevel)}
                                                            setValue={setWordList}
                                                      />
                                                </>
                                          )}

                                          {/* Show RSA Key Generator if level 6 */}
                                          {encryptionLevel === 6 && (
                                                <KeyPairGenerator
                                                      defaultPrivate={privateKey}
                                                      defaultPublic={publicKey}
                                                      setValue={({ privateKey, publicKey }) => {
                                                            setPrivateKey(privateKey);
                                                            setPublicKey(publicKey);
                                                      }}
                                                />
                                          )}
                                    </div>
                              </div>
                        )}


                              <button className="btn btn-secondary btn-xl" disabled={loading}>
                                    {loading ? <SpinnerCircle /> : "Complete Signup"}
                              </button>
                        </form>
                  </div>
            </div>
      );
}
