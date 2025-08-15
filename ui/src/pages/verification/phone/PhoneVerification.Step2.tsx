//||------------------------------------------------------------------------------------------------||
//|| PhoneVerification.Step1
//|| src/components/payments/PhoneVerification.Step1.tsx
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| React
      //||------------------------------------------------------------------------------------------------||

      import React                                          from "react";
      import { Key, Lock }                                  from "lucide-react";
      
      //||------------------------------------------------------------------------------------------------||
      //|| Interfacss
      //||------------------------------------------------------------------------------------------------||

      import { VerificationPhone }                          from "../../../interfaces/verify/phone/process";

      //||------------------------------------------------------------------------------------------------||
      //|| React
      //||------------------------------------------------------------------------------------------------||

      interface PhoneStep1Props {
            process           : VerificationPhone;
            setProcess        : (process: Partial<VerificationPhone>) => void;
      }

      //||------------------------------------------------------------------------------------------------||
      //|| React
      //||------------------------------------------------------------------------------------------------||

      export default function PhoneVerificationStep1({ process, setProcess }: PhoneStep1Props) {
            return (
                  <div className="space-y-5 p-5">
                        <div className="flex items-center gap-2 text-gray-300">
                              <Key className="w-5 h-5" />
                              <span>
                                    Enter the 6-digit code we sent to {country.dial} {localPhone}
                              </span>
                        </div>

                        <label className="block">
                              <span className="block text-sm mb-1">Verification code</span>
                              <input
                                    placeholder="••••••"
                                    inputMode="numeric"
                                    maxLength={6}
                                    value={onlyDigits(code)}
                                    onChange={(e) => setCode(onlyDigits(e.target.value))}
                                    className="input input-bordered w-full tracking-widest text-center h-16 text-2xl"
                              />
                        </label>

                        <div className="flex items-center justify-between">
                              <button
                                    onClick={() => goStep(1)}
                                    className="px-4 py-2 rounded-2xl font-bold bg-black/30 opacity-70 hover:opacity-100">
                                    Back
                              </button>
                              <div className="flex items-center gap-2">
                                    <button
                                          onClick={handleResend}
                                          disabled={busy || cooldown > 0}
                                          className={`inline-flex items-center gap-2 px-4 py-2 rounded-2xl border shadow-sm ${
                                                cooldown === 0 && !busy
                                                      ? "bg-black/10 hover:bg-black/20"
                                                      : "bg-gray-200 text-gray-500 cursor-not-allowed"
                                          }`}>
                                          <RefreshCw className="w-4 h-4" /> Resend {cooldown > 0 ? `(${cooldown})` : ""}
                                    </button>
                                    <button
                                          onClick={handleVerify}
                                          disabled={!codeValid || busy}
                                          className={`inline-flex items-center gap-2 px-5 py-2.5 rounded-2xl border shadow-sm ${
                                                codeValid && !busy
                                                      ? "bg-orange-400 cursor-pointer font-bold text-white border-0 opacity-80 hover:opacity-100"
                                                      : "bg-gray-200 text-gray-500 cursor-not-allowed"
                                          }`}>
                                          <Lock className="w-4 h-4" /> Verify
                                    </button>
                              </div>
                        </div>
                  </div>	
            );
      }
