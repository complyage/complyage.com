//||------------------------------------------------------------------------------------------------||
//|| ID Verification Status
//|| src/pages/verification/id/IDVerification.Status.tsx
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||

      import React, { useState, useEffect }                       from "react";
      import { Clock, AlertCircle, UserPlus, CheckCircle }        from "lucide-react";
      import { useSearchParams }                                  from "react-router-dom";

      //||------------------------------------------------------------------------------------------------||
      //|| Data
      //||------------------------------------------------------------------------------------------------||
      
      import { getVerificationStatus, getVerificationIcon }       from "../../../data/getVerificationData";
      
      //||------------------------------------------------------------------------------------------------||
      //|| Interfaces
      //||------------------------------------------------------------------------------------------------||

      import { VerificationID }                                         from "../../../interfaces/verify/id/process";
      import { VerificationStatuses }                                   from "../../../interfaces/models/model.verify";
      import { VerificationIDStatusProcess, VerificationIDStatusStep }  from "../../../interfaces/verify/status/process";

      //||------------------------------------------------------------------------------------------------||
      //|| Components
      //||------------------------------------------------------------------------------------------------||

      import MembersLayout                                        from "../../../layouts/MembersLayout";
      import ProgressSteps, { ProgressStep }                      from "../../../components/base/ProgressSteps";
      import SpinnerCircle                                        from "../../../components/base/SpinnerCircle";
      import InlineAlert                                          from "../../../components/base/InlineAlert";

      //||------------------------------------------------------------------------------------------------||
      //|| Props
      //||------------------------------------------------------------------------------------------------||

      interface IDVerificationStatusProps {
            process        : VerificationID;
            onTryAgain?    : () => void;
            onEscalate?    : () => void;
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Status Icon Map
      //||------------------------------------------------------------------------------------------------||

      const StatusIcon: Record<VerificationStatuses, React.ComponentType<any>> = {
            PEND: Clock,
            PEVF: Clock,
            VERF: CheckCircle,
            RJCT: AlertCircle,
            ESCL: UserPlus,
            EXPD: AlertCircle,
      };

      const StatusColor: Record<VerificationStatuses, string> = {
            PEND: "text-yellow-500",
            PEVF: "text-blue-500",
            VERF: "text-green-600",
            RJCT: "text-red-600",
            ESCL: "text-orange-600",
            INPR: "text-green-300",   
            EXPD: "text-gray-500",
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Steps
      //||------------------------------------------------------------------------------------------------||

      const steps: ProgressStep[] = [
            { label: "Awaiting Agent",          description: "Awaiting Agent" },
            { label: "Parsing ID Data",         description: "Parsing ID Data" },
            { label: "Verifying DOB",           description: "Verifying DOB" },
            { label: "Verifying Photo",         description: "Upload the back of your ID" },
            { label: "Matching Photo",          description: "Matching ID to Selfie" },
            { label: "Encrypting Data",         description: "Encrypting your photos" },
      ];

      //||------------------------------------------------------------------------------------------------||
      //|| Main Component
      //||------------------------------------------------------------------------------------------------||

      const IDVerificationStatus: React.FC<IDVerificationStatusProps> = () => {

            //||------------------------------------------------------------------------------------------------||
            //|| Verification
            //||------------------------------------------------------------------------------------------------||

            const [searchParams] = useSearchParams();
            const verificationId = searchParams.get("identifier") || "";

            //||----------------------------------------------------------------------------------------||
            //|| State
            //||----------------------------------------------------------------------------------------||

            const [loading, setLoading]   = useState<boolean>(true);
            const [reload, setReload]     = useState<boolean>(false);
            const [error , setError]      = useState<string | null>(null);
            const [process, setProcess]   = useState<VerificationIDStatusProcess>({
                  status            : "UNKN" as VerificationStatuses,
                  step              : 0,
                  verificationUUID  : verificationId,
                  steps             : [],
            });

            //||----------------------------------------------------------------------------------------||
            //|| Data
            //||----------------------------------------------------------------------------------------||

            const fetchStatus = () => {
                  fetch(`/v1/api/verify/status?identifier=${verificationId}`).then(resp => {
                        if (!resp.ok) {
                              throw new Error("Failed to fetch verification status.");
                        }
                        return resp.json();
                  }).then(response => {
                        console.log("Verification Status Data:", response);
                        if (response.success !== true) {
                              throw new Error(response.message || "Unknown error fetching verification status.");
                        }
                        setProcess({
                              status            : response.data.status,
                              step              : response.data.step,
                              steps             : response.data.steps,
                              verificationUUID  : response.data.verificationUUID,
                        } as VerificationIDStatusProcess);
                        setLoading(false);
                        window.setTimeout(() => setReload(false), 2000);
                  }).catch(err => {
                        setError(err.message);
                        setLoading(false);
                        window.setTimeout(() => setReload(false), 2000);
                  });
            }

            //||----------------------------------------------------------------------------------------||
            //|| Data
            //||----------------------------------------------------------------------------------------||

            const testReset = () => {
                  fetch(`/v1/api/verify/id/reset?identifier=${verificationId}`).then(resp => {
                        if (!resp.ok) {
                              alert("Failed to reset verification status.");
                        }
                        return resp.json();
                  }).then(response => {
                        console.log("Verification Status Data:", response);
                  }).catch(err => {
                        alert(err.message);
                  });
            }            

            //||----------------------------------------------------------------------------------------||
            //|| Poll Status Every 5 Seconds
            //||----------------------------------------------------------------------------------------||

            useEffect(() => {
			if (!verificationId) return;			
			const interval = setInterval(() => { 
                        setReload(true); 
                        fetchStatus(); 
                  }, 2000);
			fetchStatus();
			return () => clearInterval(interval);
		}, [verificationId]);

            //||----------------------------------------------------------------------------------------||
            //|| Data
            //||----------------------------------------------------------------------------------------||

            const status      = process.status || "MISS" as VerificationStatuses;
            const Icon        = StatusIcon[status] || Clock;
            const color       = StatusColor[status] || "text-gray-700";
            const label       = getVerificationStatus(status || "MISS");
            
            //||----------------------------------------------------------------------------------------||
            //|| Show Final / Escalate
            //||----------------------------------------------------------------------------------------||

            const isFinal = ["VERF", "RJCT", "EXPD", "CNCL"].includes(status);

            //||----------------------------------------------------------------------------------------||
            //|| Loading
            //||----------------------------------------------------------------------------------------||

            useEffect(() => {
                  if (!verificationId) setError("No verification ID provided.");
            }, [verificationId]);
            
            //||----------------------------------------------------------------------------------------||
            //|| Use Load
            //||----------------------------------------------------------------------------------------||

            useEffect(() => {
                  fetchStatus();
            }, []);
                        
            //||----------------------------------------------------------------------------------------||
            //|| Loading
            //||----------------------------------------------------------------------------------------||

            // if (error !== null) {
            //       return (
            //             <MembersLayout>
            //                   <div className="w-full max-w-2xl mx-auto flex flex-col items-center justify-center p-8 bg-black/20 rounded-xl mt-10">
            //                         <InlineAlert message={error} isError />
            //                         <div className="flex justify-center mt-10">
            //                               <button
            //                                     className="btn btn-primary text-lg px-10 py-2"
            //                                     onClick={() => { window.location.href = "/members"; }}
            //                               >
            //                                     Go Home
            //                               </button>
            //                         </div>                                    
            //                   </div>
            //             </MembersLayout>
            //       );
            // }

            //||----------------------------------------------------------------------------------------||
            //|| Loading
            //||----------------------------------------------------------------------------------------||

            if (loading) {
                  return (
                        <MembersLayout>
                              <div className="w-full max-w-2xl mx-auto flex flex-col items-center justify-center p-8 bg-black/20 rounded-xl mt-10">
                                    <SpinnerCircle className="w-16 h-16 text-gray-500 animate-spin" />
                                    <span className="text-lg text-gray-600 mt-4">Please wait...</span>
                              </div>
                        </MembersLayout>
                  );
            }

            //||----------------------------------------------------------------------------------------||
            //|| JSX
            //||----------------------------------------------------------------------------------------||

            return (
                  <MembersLayout>
                        <div className="w-full max-w-full mx-auto flex flex-col items-center justify-center rounded-lg">


                              <div className="w-full mb-8">
                                    { process.step } 
                                    <textarea defaultValue={JSON.stringify(process)} className="w-full h-24 bg-black/20 text-xs p-2 rounded-lg text-gray-400 font-mono" readOnly />
                                    <ProgressSteps steps={steps} currentStep={process.step || 1} />
                              </div>

                              <div className="bg-black/20 pt-4 w-full max-w-2xl rounded-xl relative">
                                    <div className="absolute top-3 right-5 z-10">
                                          { reload && ( <SpinnerCircle className="inline-block w-4 h-4" /> )}   
                                    </div>

                                    <div className="flex flex-col items-center mb-5">
                                          <button className="btn btn-sm btn-secondary absolute top-3 left-5" onClick={testReset}>Reset (Test Only)</button>
                                          <Icon className={`w-24 h-24 ${color} mx-auto`} />
                                    </div>
                                    
                                    <span className={`text-3xl py-5 text-center block font-bold tracking-wide items-center gap-2 bg-white/80 text-gray-800 shadow-xl border-black mx-[-20px]` }>
                                          {label}
                                    </span>

                                    <div className="p-4">
                                    { process.steps && process.steps.length > 0 && (
                                          <table className="table table-auto w-full mt-4">
                                                <thead className="bg-gray-800">
                                                      <tr className="text-sm font-bold text-gray-200">
                                                            <th className="w-1/3">Step</th>
                                                            <th className="w-1/3">Details</th>
                                                            <th className="w-1/3  text-right">Timestamp</th>
                                                      </tr>
                                                </thead>
                                                <tbody>
                                                {process.steps.map((step, index) => {
                                                      console.log("Rendering step:", step);
                                                      return (
                                                      <tr key={index}>
                                                            <td className={`text-xs text-gray-400"}`}>
                                                                  {step.type}
                                                            </td>
                                                            <td className={`text-xs text-gray-200"}`}>
                                                                  {step.details}
                                                            </td>
                                                            <td className="text-xs text-yellow-400 text-right">
                                                                  {new Date(step.timestamp).toLocaleString()}
                                                            </td>
                                                      </tr>
                                                )})}
                                                </tbody>
                                          </table>
                                    )}
                                    </div>
                              </div>

                              {/* Escalate Option */}
                              {status === "ESCL" && onEscalate && (
                                    <div className="mt-2 flex flex-col items-center">
                                          <div className="text-orange-700 mb-2 text-lg font-medium">
                                                Automatic verification couldn't complete.<br />
                                                <span className="text-base text-gray-600">You can request a manual review.</span>
                                          </div>
                                          <button
                                                className="btn btn-primary text-lg px-8 py-2"
                                                onClick={onEscalate}
                                          >
                                                Escalate to Human Verification
                                          </button>
                                    </div>
                              )}

                              {/* Final Statuses */}
                              {isFinal && (
                                    <div className="mt-6 flex flex-col items-center">
                                          <div className="mb-2 text-lg font-semibold">
                                                {status === "VERF" && (
                                                      <span className="text-green-700">✅ Your ID has been verified!</span>
                                                )}
                                                {status === "RJCT" && (
                                                      <span className="text-red-700">❌ Your ID verification was rejected.</span>
                                                )}
                                                {status === "EXPD" && (
                                                      <span className="text-gray-500">⏰ Your verification has expired.</span>
                                                )}
                                          </div>
                                          <button
                                                className="btn btn-secondary text-lg px-10 py-2 mt-4"
                                          >
                                                {status === "VERF" ? "Verify Another ID" : "Try Again"}
                                          </button>
                                    </div>
                              )}

                        </div>

                        <div className="block text-xs text-gray-500 pt-4 text-center mt-1 tracking-widest uppercase">
                              {process.verificationUUID}             
                        </div>

                  </MembersLayout>
            );
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Export
      //||------------------------------------------------------------------------------------------------||

      export default IDVerificationStatus;
