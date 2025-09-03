//||------------------------------------------------------------------------------------------------||
//|| AddressVerification.Step1
//|| src/components/security/AddressVerification.Step1.tsx
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||

      import React, { useMemo }                 from "react";
      import {CheckCircle2}                     from "lucide-react";

      //||------------------------------------------------------------------------------------------------||
      //|| Step Props
      //||------------------------------------------------------------------------------------------------||

      import { StepProps }                      from "./AddressVerification";

      //||------------------------------------------------------------------------------------------------||
      //|| Interface
      //||------------------------------------------------------------------------------------------------||

      import { VerificationAddress }            from "../../../interfaces/verify/address/process";

      //||------------------------------------------------------------------------------------------------||
      //|| Data
      //||------------------------------------------------------------------------------------------------||

      import { getAllCountries }                   from "../../../data/getCountries";

      //||------------------------------------------------------------------------------------------------||
      //|| Default
      //||------------------------------------------------------------------------------------------------||

      export default function AddressVerificationStep1({ process, updateProcess }: StepProps) {

            //||------------------------------------------------------------------------------------------------||
            //|| Addr
            //||------------------------------------------------------------------------------------------------||

            const addr        = process.verifyAddress || {};
            const countries   = useMemo(() => getAllCountries(), []);

            //||------------------------------------------------------------------------------------------------||
            //|| Default
            //||------------------------------------------------------------------------------------------------||

            return (
                  <div className="space-y-4 p-5">
                        <div className="grid grid-cols-3 gap-3">
                              <div className="col-span-3">
                                    <span className="block text-sm mb-1">Country</span>
                                    <select
                                          value={addr.country}
                                          onChange={e => { updateProcess({ ...process, verifyAddress: { ...addr, country: e.target.value } }); }}
                                          className="select select-bordered w-full text-2xl h-16"
                                    >
                                          {countries.map(c => (
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
                                          value={addr.line1 || ""}
                                          onChange={e => updateProcess({ ...process, verifyAddress: { ...addr, line1: e.target.value } }) }
                                          className="input input-bordered w-full text-2xl h-16"
                                    />
                              </label>

                              <label className="col-span-3">
                                    <span className="block text-sm mb-1">Address line 2 (optional)</span>
                                    <input
                                          placeholder="Apt, suite, etc."
                                          value={addr.line2 || ""}
                                          onChange={e => updateProcess({ ...process, verifyAddress: { ...addr, line2: e.target.value } }) }
                                          className="input input-bordered w-full text-2xl h-16"
                                    />
                              </label>

                              <label className="col-span-2">
                                    <span className="block text-sm mb-1">City</span>
                                    <input
                                          placeholder="City"
                                          value={addr.city || ""}
                                          onChange={e => updateProcess({ ...process, verifyAddress: { ...addr, city: e.target.value } }) }
                                          className="input input-bordered w-full text-2xl h-16"
                                    />
                              </label>

                              <label className="col-span-1">
                                    <span className="block text-sm mb-1">State / Province</span>
                                    <input
                                          placeholder="State/Prov"
                                          value={addr.state || ""}
                                          onChange={e => updateProcess({ ...process, verifyAddress: { ...addr, state: e.target.value } }) }
                                          className="input input-bordered w-full text-2xl h-16"
                                    />
                              </label>

                              <label className="col-span-1">
                                    <span className="block text-sm mb-1">ZIP / Postal</span>
                                    <input
                                          placeholder="Postal"
                                          value={addr.postal || ""}
                                          onChange={e => updateProcess({ ...process, verifyAddress: { ...addr, postal: e.target.value } }) }
                                          className="input input-bordered w-full text-2xl h-16"
                                    />
                              </label>
                        </div>

                        <div className="flex items-center justify-end pt-2">
                              <button
                                    onClick={() => updateProcess({ ...process, step: 2 })}
                                    disabled={!addr.line1 || !addr.city || !addr.state || !addr.postal || !addr.country}
                                    className="inline-flex items-center gap-2 px-5 py-2.5 rounded-2xl shadow-sm bg-orange-400 font-bold text-white border-0 opacity-80 hover:opacity-100 disabled:opacity-50 disabled:cursor-not-allowed"
                              >
                                    <CheckCircle2 className="w-4 h-4" /> Continue
                              </button>
                        </div>
                  </div>
            );
      }
