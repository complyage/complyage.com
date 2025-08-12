//||------------------------------------------------------------------------------------------------||
//|| Step 1
//|| ID Verification
//||------------------------------------------------------------------------------------------------||

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import { Camera as CameraIcon, Upload as UploadIcon }                             from "lucide-react";
import React, {useEffect, useRef, useState}           from "react";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import Camera                                      from "../../../components/dynamic/Camera";
import Upload                                      from "../../../components/dynamic/Upload";
import { VerificationProcessID }                   from "../../../interfaces/verification.process.id";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface IDVStep2Props {
      onNext                  : (file: File | Blob) => void;
      onUpload                : (file: File | Blob) => void;
      process                 : VerificationProcessID;
}

//||------------------------------------------------------------------------------------------------||
//|| ID Verification
//||------------------------------------------------------------------------------------------------||

export default function IDVStep2({ onNext, onUpload }: IDVStep2Props) {
      //||------------------------------------------------------------------------------------------------||
      //|| Const
      //||------------------------------------------------------------------------------------------------||
      const [captureMode, setCaptureMode]  = useState<"camera" | "upload" | null>(null);
      //||------------------------------------------------------------------------------------------------||
      //|| Handler
      //||------------------------------------------------------------------------------------------------||
      const handleUpload = async (file: File | Blob) => {
            onUpload(file);
      }
      //||------------------------------------------------------------------------------------------------||
      //|| Upload
      //||------------------------------------------------------------------------------------------------||
      if (captureMode === "upload") return ( <Upload onUpload={ handleUpload }  /> );
      //||------------------------------------------------------------------------------------------------||
      //|| Capture
      //||------------------------------------------------------------------------------------------------||
      if (captureMode === "camera") return ( <Camera onUpload={ handleUpload } onClose={ () => setCaptureMode(null) } /> );
      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||
      return (
            <div className="w-full mx-auto text-center p-4 flex flex-row rounded-lg gap-5">
                  {/* Camera Button */}
                  <div className="flex w-1/2 justify-center">
                        <button
                              onClick={() => setCaptureMode("camera")}
                              className="flex flex-col items-center justify-center p-4 w-full bg-gray-800 text-white rounded-l-lg hover:bg-gray-600 transition"
                        >
                              <CameraIcon className="w-8 h-8 mb-2" />
                              <span className="text-sm">Use Camera</span>
                        </button>
                  </div>

                  {/* Upload Button */}
                  <div className="flex w-1/2 justify-center">
                        <button
                              onClick={() => setCaptureMode("upload")}
                              className="flex flex-col items-center justify-center p-4 w-full bg-gray-800 text-white rounded-r-lg hover:bg-gray-600 transition"
                        >
                              <UploadIcon className="w-8 h-8 mb-2" />
                              <span className="text-sm">Upload File</span>
                        </button>
                  </div>
            </div>
	);
}
