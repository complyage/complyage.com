//||------------------------------------------------------------------------------------------------||
//|| CCVerification (Container)
//|| src/pages/verification/CCVerification.tsx
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||

      import React, {useRef, useState, useEffect}           from "react";
      import {useNavigate, useSearchParams}                 from "react-router-dom";
      import {CheckCircle, Home}                            from "lucide-react";

      //||------------------------------------------------------------------------------------------------||
      //|| Data
      //||------------------------------------------------------------------------------------------------||

      import type { VerificationTypes }                     from "../../../interfaces/models/model.verify";
      import { getVerificationIcon, getVerificationType }   from "../../../data/getVerificationData";

      //||------------------------------------------------------------------------------------------------||
      //|| Fail
      //||------------------------------------------------------------------------------------------------||

      import VerificationCheckFail                          from "./VerificationCheck.Fail";

      //||------------------------------------------------------------------------------------------------||
      //|| Components
      //||------------------------------------------------------------------------------------------------||

      import MembersLayout                                  from "../../../layouts/MembersLayout";
      import ProgressSteps, { ProgessStep }                 from "../../../components/base/ProgressSteps";
      import InlineAlert                                    from "../../../components/base/InlineAlert";
      import SpinnerCircle                                  from "../../../components/base/SpinnerCircle";

      //||------------------------------------------------------------------------------------------------||
      //|| Page
      //||------------------------------------------------------------------------------------------------||

      export default function VerificationCheck() {

            //||------------------------------------------------------------------------------------------------||
            //|| Verification
            //||------------------------------------------------------------------------------------------------||

            const [searchParams] = useSearchParams();
            const verificationId = searchParams.get("identifier") || "";
      
            //||------------------------------------------------------------------------------------------------||
            //|| Navigate
            //||------------------------------------------------------------------------------------------------||

            const navigate                     = useNavigate();

            //||------------------------------------------------------------------------------------------------||
            //|| useState
            //||------------------------------------------------------------------------------------------------||

            const [code,      setCode]                    = useState("");
            const [loaded,    setLoaded]                  = useState(false);
            const [verificationType, setVerificationType] = useState<VerificationTypes>("UNKN");
            const [busy,      setBusy]                    = useState(false);
            const [error,     setError]                   = useState<string | null>(null);
            const [loadError, setLoadError]               = useState<string | null>(null);
            const [complete,  setComplete]                = useState(false);

            //||------------------------------------------------------------------------------------------------||
            //|| Handle Returning
            //||------------------------------------------------------------------------------------------------||

            useEffect(() => {
                  if (!verificationId) {
                        setError("Verification Identifier was not provided.");
                        return;
                  }
                  (async () => {
                        try {
                              const res = await fetch(`/v1/api/verify/load?identifier=${verificationId}`);
                              if (res.ok) {
                                    const data = await res.json();                                    
                                    console.log("Verification Data:", data);
                                    setLoaded(true);
                                    setVerificationType(data.data?.type || "UNKN");
                              } else {
                                    console.error("Failed to load verification data:", res.status, res.statusText);
                                    const data = await res.json();
                                    setLoadError(data?.message || "Failed to load verification data.");
                                    setLoaded(true);
                              }
                        } catch(e) {
                              console.log(e);
                              console.error("Error fetching verification data:", verificationId);
                              setLoadError("Failed to load verification data.");
                              setLoaded(false);
                        }
                  })();
            }, [verificationId]);

            //||------------------------------------------------------------------------------------------------||
            //|| Props
            //||------------------------------------------------------------------------------------------------||

            const submitCode = async () => {
                  setBusy(true);
                  setError(null);
                  const payload = {
                        "uuid" : verificationId,
                        "code" : code                    
                  }
                  try {
                        const res = await fetch(`/v1/api/verify/check`, {
                              method      : "POST",
                              headers     : { "Content-Type": "application/json" },
                              body        : JSON.stringify(payload),
                        });
                        const response = await res.json();
                        console.log("Verification response:", response);
                        if (!res.ok) {
                              setError(response.message || "Verification failed.");
                              setBusy(false);
                              return;
                        }
                        console.log("DATA STATUS : " , response.data);
                        if (response.data && typeof(response.data.status) !== "undefined" && response.data.status === "VERF") {
                              setComplete(true);
                              setBusy(false);
                              return;
                        }
                  } catch (err) {
                        console.error("Error submitting verification code:", err);
                        setLoadError("Error submitting verification code.");
                  } finally {
                        setBusy(false);
                  }
            };

            //||------------------------------------------------------------------------------------------------||
            //|| Steps
            //||------------------------------------------------------------------------------------------------||

            const steps: ProgessStep[] = [
                  { label: "Code Entry", description: "Enter the code" },
                  { label: "Verified!", description: "Verification Complete!" },
            ];            

            //||------------------------------------------------------------------------------------------------||
            //|| Step
            //||------------------------------------------------------------------------------------------------||

            return (
			<MembersLayout title={getVerificationType(verificationType)} icon={getVerificationIcon(verificationType)}>
				<div className="relative min-h-[90vh]">
					<div className="absolute inset-0 pointer-events-none z-0" />
					<div className="relative z-10 max-w-2xl mx-auto">
						<div className="py-10 flex flex-col gap-4 items-center">
							<ProgressSteps steps={steps} currentStep={1} />
						</div>

						{(!loaded || busy) && (
							<div className="w-full max-w-2xl mx-auto p-8 bg-black/30 rounded-3xl mt-5 flex flex-col items-center shadow-2xl">
								<SpinnerCircle />
								<h2 className="text-gray-100 font-bold p-2 text-center">Loading...</h2>
							</div>
						)}

						{loadError && <VerificationCheckFail error={loadError} verificationId={verificationId} onBack={ () => { setLoadError(null) } } />}

						<div className="w-full max-w-2xl mx-auto">
							{!loadError && loaded && !complete && (
								<div className="space-y-7 p-10 bg-black/20">
									<p className="text-gray-50 text-center text-2xl font-bold tracking-tight mb-3">
										Enter your 6-digit verification code
									</p>
									<p className="text-gray-400 text-center mb-6">
										We sent a unique code to your bank statement. Enter it below to verify your card and complete setup.
									</p>
                                                      { error && (
                                                            <InlineAlert
                                                                  isError
                                                                  message={error}
                                                            />
                                                      )}
									<div className="flex flex-col items-center gap-3">
										<input
											type="text"
											inputMode="numeric"
											maxLength={6}
											placeholder="_ _ _ _ _ _"
											className="input font-bold input-bordered w-56 sm:w-72 text-3xl h-16 text-center tracking-widest focus:border-orange-500 shadow-inner bg-black/20"
											value={code}
											onChange={(e) => setCode(e.target.value.replace(/[^0-9]/g, ""))}
										/>
									</div>
									<div className="w-full flex justify-center mt-6">
										<button
											onClick={submitCode}
											disabled={busy || code.length !== 6}
											className={`btn text-2xl h-auto py-3 px-10 ${ busy || code.length !== 6 ? " btn-primary" : " btn-secondary" }
                                                            `}>
											{busy ? (
												<>
													<SpinnerCircle size={22} className="inline mr-2 align-middle" />
													Verifying...
												</>
											) : (
												<>Submit Code</>
											)}
										</button>
									</div>
                                                      <span className="pt-2 block text-center text-sm text-gray-500 mt-2 select-all">{verificationId}</span>
								</div>
							)}

							{complete && !loadError && (
								<div className="space-y-8 p-10 bg-green-900/10 rounded-3xl shadow-2xl border border-green-400/30 mt-12 text-center">
									<CheckCircle className="w-16 h-16 text-green-400 mx-auto mb-2 animate-bounce" />
									<p className="text-2xl text-green-300 font-extrabold mb-2">Verification Complete!</p>
									<p className="text-gray-100 mb-8">
										Your verification is complete!
									</p>
									<button
										onClick={() => navigate("/members")}
										className="btn btn-primary bg-green-500/80 hover:bg-green-400 text-xl px-10 py-7 rounded-2xl shadow-lg font-bold">
										<CheckCircle className="inline w-6 h-6 mr-2 align-middle" /> Go to Dashboard
									</button>
								</div>
							)}
						</div>
					</div>
				</div>
			</MembersLayout>
		);

      }

