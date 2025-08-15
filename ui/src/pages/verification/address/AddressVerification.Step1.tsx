//||------------------------------------------------------------------------------------------------||
//|| AddressVerification.Step1
//|| src/components/security/AddressVerification.Step1.tsx
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||

      import React                              from "react";
      import {CheckCircle2}                     from "lucide-react";

      //||------------------------------------------------------------------------------------------------||
      //|| Interface
      //||------------------------------------------------------------------------------------------------||

      import {Address, Country}                 from "../../../interfaces/verification.location";

      //||------------------------------------------------------------------------------------------------||
      //|| Props
      //||------------------------------------------------------------------------------------------------||

      type Props = {
            countries: Country[];
            country: Country;
            setCountry: React.Dispatch<React.SetStateAction<Country>>;
            addr: Address;
            setAddr: React.Dispatch<React.SetStateAction<Address>>;
            addrOk: boolean;
            busy: boolean;
            serverMsg: string;
            onStandardize: () => void;
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Default
      //||------------------------------------------------------------------------------------------------||

      export default function AddressVerificationStep1({countries, country, setCountry, addr, setAddr, addrOk, busy, serverMsg, onStandardize}: Props) {
            return (
                  <div className="space-y-4 p-5">
                        <div className="grid grid-cols-3 gap-3">
                              <div className="col-span-3">
                                    <span className="block text-sm mb-1">Country</span>
                                    <select
                                          value={country.code}
                                          onChange={(e) => {
                                                const c = countries.find((x) => x.code === e.target.value) || countries[0];
                                                setCountry(c);
                                                setAddr((a) => ({...a, country: c.code}));
                                          }}
                                          className="select select-bordered w-full text-2xl h-16">
                                          {countries.map((c) => (
                                                <option key={c.code} value={c.code}>
                                                      {c.flag} {c.name}
                                                </option>
                                          ))}
                                    </select>
                              </div>

                              <label className="col-span-3">
                                    <span className="block text-sm mb-1">Address line 1</span>
                                    <input
                                          placeholder="123 Main St"
                                          value={addr.line1}
                                          onChange={(e) => setAddr((a) => ({...a, line1: e.target.value}))}
                                          className="input input-bordered w-full text-2xl h-16"
                                    />
                              </label>

                              <label className="col-span-3">
                                    <span className="block text-sm mb-1">Address line 2 (optional)</span>
                                    <input
                                          placeholder="Apt, suite, etc."
                                          value={addr.line2}
                                          onChange={(e) => setAddr((a) => ({...a, line2: e.target.value}))}
                                          className="input input-bordered w-full text-2xl h-16"
                                    />
                              </label>

                              <label className="col-span-2">
                                    <span className="block text-sm mb-1">City</span>
                                    <input
                                          placeholder="City"
                                          value={addr.city}
                                          onChange={(e) => setAddr((a) => ({...a, city: e.target.value}))}
                                          className="input input-bordered w-full text-2xl h-16"
                                    />
                              </label>

                              <label className="col-span-1">
                                    <span className="block text-sm mb-1">State / Province</span>
                                    <input
                                          placeholder="State/Prov"
                                          value={addr.state}
                                          onChange={(e) => setAddr((a) => ({...a, state: e.target.value}))}
                                          className="input input-bordered w-full text-2xl h-16"
                                    />
                              </label>

                              <label className="col-span-1">
                                    <span className="block text-sm mb-1">ZIP / Postal</span>
                                    <input
                                          placeholder="Postal"
                                          value={addr.postal}
                                          onChange={(e) => setAddr((a) => ({...a, postal: e.target.value}))}
                                          className="input input-bordered w-full text-2xl h-16"
                                    />
                              </label>
                        </div>

                        <div className="flex items-center justify-end pt-2">
                              <button
                                    onClick={onStandardize}
                                    disabled={!addrOk || busy}
                                    className={`inline-flex items-center gap-2 px-5 py-2.5 rounded-2xl shadow-sm bg-orange-400 font-bold text-white border-0 opacity-80 hover:opacity-100 ${
                                          addrOk && !busy ? "bg-black text-white" : "bg-gray-200 text-gray-500 cursor-not-allowed"
                                    }`}>
                                    <CheckCircle2 className="w-4 h-4" /> Continue
                              </button>
                        </div>

                        {!!serverMsg && <div className="text-xs mt-1 opacity-80">{serverMsg}</div>}
                  </div>
            );
      }
