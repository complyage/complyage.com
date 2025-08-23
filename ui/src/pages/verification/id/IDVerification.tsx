//||------------------------------------------------------------------------------------------------||
//|| IDVerification (Container)
//|| src/pages/verification/id/IDVerification.tsx
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||

      import React, {useRef, useState, useEffect}           from "react";
      import {useNavigate, useSearchParams}                 from "react-router-dom";
      import { Home, IdCard }                                     from "lucide-react";
      import { getEnv }                                     from "../../../data/getEnv";

      //||------------------------------------------------------------------------------------------------||
      //|| Interfaces
      //||------------------------------------------------------------------------------------------------||

      import { VerificationID, VerificationMedia }          from "../../../interfaces/verify/id/process";

      //||------------------------------------------------------------------------------------------------||
      //|| Steps
      //||------------------------------------------------------------------------------------------------||

      import IDVerificationStep1                            from "./IDVerification.Initial";
      import IDVerificationSubmit                           from "./IDVerification.Submit";
      import IDVerificationMedia                            from "./IDVerification.Media";

      //||------------------------------------------------------------------------------------------------||
      //|| Components
      //||------------------------------------------------------------------------------------------------||

      import InlineAlert                                    from "../../../components/base/InlineAlert";
      import MembersLayout                                  from "../../../layouts/MembersLayout";
      import ProgressSteps, { ProgessStep }                 from "../../../components/base/ProgressSteps";

      //||------------------------------------------------------------------------------------------------||
      //|| Step Props
      //||------------------------------------------------------------------------------------------------||

      export interface StepProps {
            which             : "front" | "back" | "selfie" | "other";
            process           : VerificationID;
            updateProcess     : (process: VerificationID) => void;
            onUpload          : (file: File | Blob | null, which: "front" | "back" | "selfie") => Promise<void>;
            getUpload         : (which: "front" | "back" | "selfie" | "other") => Promise<VerificationMedia | null>;
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Page
      //||------------------------------------------------------------------------------------------------||

      export default function IDVerification() {

            //||------------------------------------------------------------------------------------------------||
            //|| Verification
            //||------------------------------------------------------------------------------------------------||

            const [searchParams]                = useSearchParams();
            const verificationUUID              = searchParams.get("identifier") || "";
            const [mainError, setMainError]     = useState<string | null>(null);

            //||------------------------------------------------------------------------------------------------||
            //|| Navigate
            //||------------------------------------------------------------------------------------------------||

            const navigate                     = useNavigate();  

            //||------------------------------------------------------------------------------------------------||
            //|| Process
            //||------------------------------------------------------------------------------------------------||

            const [process, setProcess] = useState<VerificationID>({
                  step                : 0,
                  status              : "PEND",
                  verificationUUID    : verificationUUID,
            });

            //||------------------------------------------------------------------------------------------------||
            //|| onUpload
            //||------------------------------------------------------------------------------------------------||

            const onUpload = async (file: File | Blob | null, which: "front" | "back" | "selfie") => {
                  try {
                        if (!process.verificationUUID) throw new Error("Missing verification process identifier");

                        //||------------------------------------------------------------------------------------------------||
                        //|| Form Data / Request
                        //||------------------------------------------------------------------------------------------------||
                        const formData = new FormData();
                        formData.append("media", file || "");
                        const res = await fetch(`/v1/api/verify/id/media/upload?identifier=${process.verificationUUID}&which=${which}`, {
                              method            : "POST",
                              body              : formData,
                              credentials       : "include"
                        });
                        //||------------------------------------------------------------------------------------------------||
                        //|| Response
                        //||------------------------------------------------------------------------------------------------||
                        const json = await res.json();
                        if (!res.ok || !json.success) {
                              throw new Error(json.message || "Upload failed");
                        }
                        //||------------------------------------------------------------------------------------------------||
                        //|| Response
                        //||------------------------------------------------------------------------------------------------||
                        const media = json.data as VerificationMedia || null;
                        updateProcess({ [which]: media });
                  } catch (err: any) {
                        console.error("Upload error:", err.message || err);
                        alert(`Upload failed: ${err.message || err}`);
                  }
            };       

            //||------------------------------------------------------------------------------------------------||
            //|| getUpload
            //||------------------------------------------------------------------------------------------------||

            const getUpload = async ( which: "front" | "back" | "selfie" | "other" ) : Promise<VerificationMedia | null> => {
                  try {
                        if (!process.verificationUUID) throw new Error("Missing verification process identifier");
                        //||------------------------------------------------------------------------------------------------||
                        //|| Request
                        //||------------------------------------------------------------------------------------------------||
                        const res = await fetch(`/v1/api/verify/id/media/fetch?identifier=${process.verificationUUID}&which=${which}`, {
                              method            : "GET",
                              credentials       : "include"
                        });
                        //||------------------------------------------------------------------------------------------------||
                        //|| Response
                        //||------------------------------------------------------------------------------------------------||
                        const json = await res.json();
                        console.log(`[FETCH=UPLOAD] ${which}`, json);
                        //||------------------------------------------------------------------------------------------------||
                        //|| Request
                        //||------------------------------------------------------------------------------------------------||
                        if (!res.ok || !json.success) {
                              throw new Error(json.message || "Upload failed");
                              updateProcess({ [which]: null });
                        }

                        console.log("Fetched upload:", json.data);
                        return json.data as Media;
                  } catch (err: any) {
                        console.error("Upload error:", err.message || err);
                        return null;
                  }
            };                 

            //||------------------------------------------------------------------------------------------------||
            //|| Continue
            //||------------------------------------------------------------------------------------------------||

            const handleContinue = () => {
                  updateProcess({
                        step: (process.step || 0) + 1
                  });
            };

            //||------------------------------------------------------------------------------------------------||
            //|| Set Process
            //||------------------------------------------------------------------------------------------------||

            const updateProcess = (update: Partial<VerificationID>) => {
                  setProcess((prev) => ({
                        ...prev,
                        ...update
                  }));
            }

            //||------------------------------------------------------------------------------------------------||
            //|| Handle Returning
            //||------------------------------------------------------------------------------------------------||

            useEffect(() => {
                  if (!process.verificationUUID) return;
                  (async () => {
                        try {
                              const res = await fetch(`/v1/api/verify/id/status?identifier=${process.verificationUUID}`);
                              if (!res.ok) {
                                    throw new Error("Failed to load verification process");
                              }
                              if (res.ok) {
                                    const data = await res.json();                                    
                                    console.log("/ID/LOAD/", data);
                                    const newProcess = process;
                                    if (data.data.type !== "IDEN") setMainError("Invalid verification type. Whatcha trying to do?");
                                    newProcess.status = data.data.status;
                                    switch(data.data.status) {
                                          case "PEND": break;
                                          case "INPG": break;
                                          case "PEVF":
                                                setMainError("This verification is pending approval. Please wait..");
                                                break;
                                          case "APPR":
                                                setMainError("This verification is already approved. Please view your dashboard for more information.");
                                                break;
                                          case "REJT":
                                                setMainError("This verification has been rejected. Please view your dashboard for more information.");
                                                break;
                                          default:
                                                setMainError("Invalid Status.");
                                                newProcess.step = 1;
                                    }
                                    updateProcess({...process, ...newProcess });
                              }
                        } catch {}
                  })();
            }, [process.verificationUUID]);

            //||------------------------------------------------------------------------------------------------||
            //|| Stepper
            //||------------------------------------------------------------------------------------------------||

            const steps: ProgessStep[] = [
                  { label: "How it works",      description: "How to get your ID and Age verified" },
                  { label: "Upload ID Front",   description: "Upload the front of your ID." },
                  { label: "Upload ID Back",    description: "Upload the back of your ID" },
                  { label: "Take a Selfie",     description: "Take a Selfie" },
                  { label: "Submit for Approval",     description: "Submit for Approval" },
            ];    

            //||------------------------------------------------------------------------------------------------||
            //|| Main
            //||------------------------------------------------------------------------------------------------||

            if (mainError !== null ) {
                  return (
                        <MembersLayout title="ID Verification" icon={ Home }>
                              <div className="w-full max-w-2xl mx-auto text-center p-6">
                                    <InlineAlert message={ mainError } isError={false} />
                                    {process.verificationUUID !== "" && (<button onClick={() => navigate(`/verification/status?identifier=${ process.verificationUUID}`)} className="mt-4 btn bg-blue-600 text-2xl h-auto py-2">View Status</button>)}
                              </div>
                              <div className="w-full max-w-2xl mx-auto text-center p-6 border-t border-gray-500">
                                    <button onClick={() => navigate("/members/verification")} className="btn btn-primary text-2xl bg-black/40">Go Back</button>                              
                              </div>
                        </MembersLayout>
                  );
            }      

            //||------------------------------------------------------------------------------------------------||
            //|| Step
            //||------------------------------------------------------------------------------------------------||

            return (
                  <MembersLayout title="ID Verification" icon={ IdCard }>
                        <div className="w-full max-w-2xl mx-auto">                              
                              <ProgressSteps steps={ steps } currentStep={ process.step } className="mb-6" /> 
                              { process.step <= 1 && <IDVerificationStep1 updateProcess={ updateProcess } process={process} />}
                              { process.step === 2 && <IDVerificationMedia which="front" updateProcess={ updateProcess } process={process} onUpload={ onUpload } getUpload={ getUpload } />}
                              { process.step === 3 && <IDVerificationMedia which="back" updateProcess={ updateProcess } process={process}  onUpload={ onUpload } getUpload={ getUpload } />}
                              { process.step === 4 && <IDVerificationMedia which="selfie" updateProcess={ updateProcess } process={process} onUpload={ onUpload } getUpload={ getUpload } />}
                              { process.step === 5 && <IDVerificationSubmit which="other" updateProcess={ updateProcess } process={process} onUpload={ onUpload } getUpload={ getUpload } />}
                              {process.step > 0 && process.step < 5 && (                              
                                    <div className="flex flex-row items-center justify-between w-full border-t border-gray-500 mt-3 mx-auto pt-4">
                                          <button className="btn btn-primary text-2xl py-3 px-10 h-auto" onClick={() => updateProcess({ step: process.step - 1 })} >Back</button>
                                          <button className="btn btn-secondary text-2xl py-3 px-10 h-auto" onClick={() => handleContinue()}>Continue</button>
                                    </div>
                              )}
                              <div className="text-gray-500 text-sm block text-center pt-5">
                                    {process.verificationUUID} {process.step}
                              </div>

                        </div>

                  </MembersLayout>                  
            );
      }
