//||------------------------------------------------------------------------------------------------||
//|| ID Verification
//||------------------------------------------------------------------------------------------------||

import React, { useState, useRef, useEffect }   from "react";
import MembersLayout                            from "../../../layouts/MembersLayout";

//||------------------------------------------------------------------------------------------------||
//|| VerificationProcess
//||------------------------------------------------------------------------------------------------||

import { useVerificationProcess }               from "../../../utils/useVerificationProcess";

//||------------------------------------------------------------------------------------------------||
//|| Steps
//||------------------------------------------------------------------------------------------------||

import IDVerificationStep0                      from "./IDVerification.Step0";
import IDVerificationStep1                      from "./IDVerification.Step1";
import IDVerificationStep2                      from "./IDVerification.Step2";
import IDVerificationStep3                      from "./IDVerification.Step3";

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function IDVerification() {
      
      //||------------------------------------------------------------------------------------------------||
      //|| Handle File Selection
      //||------------------------------------------------------------------------------------------------||
      
      const [step, setStep]         = useState(0);
      const videoRef                = useRef<HTMLVideoElement>(null);
      const canvasRef               = useRef<HTMLCanvasElement>(null);

      //||------------------------------------------------------------------------------------------------||
      //|| onUpload
      //||------------------------------------------------------------------------------------------------||

      const onUpload = async (file: File | Blob, which: "front" | "back" | "selfie") => {
            try {
                  if (!process?.identifier) throw new Error("Missing verification process identifier");

                  const formData = new FormData();
                  formData.append("media", file);

                  const res = await fetch(`/v1/api/verify/upload?identifier=${process.identifier}&which=${which}`, {
                        method            : "POST",
                        body              : formData,
                        credentials       : "include"
                  });

                  const json = await res.json();
                  console.log(`[UPLOAD] ${which}`, json);

                  if (!res.ok || !json.success) {
                        throw new Error(json.message || "Upload failed");
                  }
            } catch (err: any) {
                  console.error("Upload error:", err.message || err);
                  alert(`Upload failed: ${err.message || err}`);
            }
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Verification Process
      //||------------------------------------------------------------------------------------------------||

      const { process, error, loading } = useVerificationProcess();

      //||------------------------------------------------------------------------------------------------||
      //|| Calculate Step
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            if (!process) return;
            if (process.front && !process.back && step < 2) setStep(2)
            if (process.front && process.back && step < 3) setStep(3)
            //||------------------------------------------------------------------------------------------------||
            //|| Check if the process is complete
            //||------------------------------------------------------------------------------------------------||

            if (process?.status === "complete") {
                  alert("Verification process is complete!");
                  window.location.href = "/verification/success";
            } else if (process?.status === "failed") {
                  alert("Verification process failed. Please try again.");
                  window.location.href = "/verification/init?type=IDEN";
            }
      }, [process, step]);      

      //||------------------------------------------------------------------------------------------------||
      //|| Title
      //||------------------------------------------------------------------------------------------------||

      let title = "Verify Your Age/Government ID";
      switch(step) { 
            case 1 : title = "Step 1. Upload your Passport or Government ID Front"; break;
            case 2 : title = "Step 2. Upload the back of your passport"; break;
            case 3 : title = "Step 3. Upload your selfie"; break;
      }
      
      //||------------------------------------------------------------------------------------------------||
      //|| Loading
      //||------------------------------------------------------------------------------------------------||

      if (loading)            return <MembersLayout title="Loading...">Loading...</MembersLayout>;
      if (error)              return <MembersLayout title="Error">{error}<button className="btn btn-primary" onClick={() => {window.location.href='/verification/init?type=IDEN'}}>Try again</button></MembersLayout>;
      if (!process)           return <MembersLayout title="Error">Could not load process</MembersLayout>;

      //||------------------------------------------------------------------------------------------------||
      //|| Handle File Selection
      //||------------------------------------------------------------------------------------------------||
   
      return (
            <MembersLayout title={ title } >
                  {step === 0 && ( <IDVerificationStep0 updateStep={ ( step : number ) => { setStep(step) } } process={process} /> )}
                  {step === 1 && ( <IDVerificationStep1 onUpload={ (file : File | Blob ) => { onUpload(file, "front") } } process={process} /> )}                  
                  {step === 2 && ( <IDVerificationStep2 onUpload={ (file : File | Blob ) => { onUpload(file, "back") } } process={process} /> )}                  
                  {step === 3 && ( <IDVerificationStep3 onUpload={ (file : File | Blob ) => { onUpload(file, "back") } } process={process} /> )}
            </MembersLayout>
      );
}
