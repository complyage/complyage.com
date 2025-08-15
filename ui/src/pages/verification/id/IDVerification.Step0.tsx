//||------------------------------------------------------------------------------------------------||
//|| Step 1
//|| ID Verification
//||------------------------------------------------------------------------------------------------||

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import { Camera as CameraIcon, Upload as UploadIcon }      from "lucide-react";
import React, {useEffect, useRef, useState}                from "react";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

import { VerificationProcessID }                           from "../../../interfaces/verify/id/process";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface IDVStep0Props {
      updateStep        : (step: number) => void;
      process           : VerificationProcessID;
}

//||------------------------------------------------------------------------------------------------||
//|| ID Verification
//||------------------------------------------------------------------------------------------------||

export default function IDVStep0({ updateStep, process }: IDVStep0Props) {
      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||
      return (
            <div className="w-full max-w-3xl mx-auto text-center p-6 flex flex-col items-center justify-center gap-6 rounded-xl shadow-lg bg-black/50">

                  <h1 className="text-2xl font-semibold">Verify Your Identity</h1>

                  <p className="text-white max-w-xl">
                        This takes about <strong>3 minutes</strong>. You’ll need a valid 
                        <strong> government-issued ID or passport</strong> and a 
                        <strong> high-resolution webcam</strong>. If you're on a desktop, 
                        you can also scan the QR code to switch to your phone.
                  </p>

                  {/* QR Code Block */}
                  <div className="flex flex-col items-center gap-2 mt-4">
                        <div className="w-40 h-40 bg-gray-100 border border-gray-300 rounded">
                              {/* Replace this with actual QR code */}
                              <div className="w-full h-full flex items-center justify-center text-sm text-gray-400">
                                    <img src={`/v1/api/verify/qr/generate?identifier=${process.identifier}&type=IDEN`} />
                              </div>
                        </div>
                        <p className="text-sm text-gray-500">Scan with your mobile camera</p>
                  </div>

                  {/* Continue Button */}
                  <button onClick={ () => updateStep(1) } className="mt-6 px-6 py-3 bg-black text-white text-lg font-medium rounded-lg hover:bg-gray-800 transition">Continue</button>
            </div>
      );
}
