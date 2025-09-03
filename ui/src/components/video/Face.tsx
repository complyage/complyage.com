//||------------------------------------------------------------------------------------------------||
//|| Face
//|| Component for Face/Selfie
//||------------------------------------------------------------------------------------------------||

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, {useState, useCallback, useRef, useEffect}    from "react";
import { Target }                                           from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Hooks
//||------------------------------------------------------------------------------------------------||

import { useCamera }                                        from "../../hooks/useCamera";
import { useLuma }                                          from "../../hooks/useLuma";
import { useFacePercent }                                   from "../../hooks/useFacePercent";
import { useFaceDirection }                                 from "../../hooks/useFaceDirection";
import { useCapture }                                       from "../../hooks/useCapture";

//||------------------------------------------------------------------------------------------------||
//|| Utils
//||------------------------------------------------------------------------------------------------||

import { statusLuma, iconLuma, boolLuma }                   from "../../utils/face/statusLuma"; 
import { statusFace, iconFace, boolFace }                   from "../../utils/face/statusFace";
import { statusFaceDir, iconFaceDir, 
         statusFaceCentered, boolFaceCentered, 
         iconFaceCentered }                                 from "../../utils/face/statusFaceDir";
import { statusStep }                                       from "../../utils/face/statusStep";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

import { FaceDirection }                                    from "../../interfaces/verify/id/types";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import ProgressSteps, { ProgressStep }                       from "../base/ProgressSteps";
import ProgressBar                                           from "../base/ProgressBar";

//||------------------------------------------------------------------------------------------------||
//|| Const
//||------------------------------------------------------------------------------------------------||

const DEBUG = true;

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface FaceProps {
      onUpload    : (file: File | Blob) => void;
}

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function Face({ onUpload } : FaceProps) {

      //||------------------------------------------------------------------------------------------------||
	//|| Steps
      //|| 0 = Load Camera
      //|| 1 = Face In Frame
      //|| 2 = Face Movement
      //|| 3 = Face Capture
      //|| 4 = Face Upload
      //|| 5 = Complete
	//||------------------------------------------------------------------------------------------------||

      const [currentStep, setCurrentStep]                   = useState(0);

      //||------------------------------------------------------------------------------------------------||
	//|| State
	//||------------------------------------------------------------------------------------------------||

      const [currentLuma, setCurrentLuma]                   = useState(0);
      const [currentFace, setCurrentFace]                   = useState(0);
      const [currentDir,  setCurrentDir]                    = useState<FaceDirection>("unknown-unknown");

      //||------------------------------------------------------------------------------------------------||
	//|| Face Live
	//||------------------------------------------------------------------------------------------------||

      const faceMoveMin                                     = 3;
      const [faceMoveCount, setFaceMoveCount]               = useState<number>(0);
      const [faceInFrame, setFaceInFrame]                   = useState<boolean>(false);
      
      //||------------------------------------------------------------------------------------------------||
	//|| useRef
	//||------------------------------------------------------------------------------------------------||

	const canvasRef                                       = useRef<HTMLCanvasElement | null>(null);
      const videoRef                                        = useRef<HTMLVideoElement | null>(null);

      //||------------------------------------------------------------------------------------------------||
	//|| Hooks
	//||------------------------------------------------------------------------------------------------||

      const { cameraError, cameraLoading, cameraReady }     = useCamera( videoRef );
      const { luma }                                        = useLuma(videoRef, { sampleSize: 96, ema: 0.2, enabled: true });
      const { facePct, faceReady, faceError, facesCount }   = useFacePercent(videoRef, { ema: 0.2, enabled: cameraReady });
      const { faceDir, faceDirReady, faceDirError }         = useFaceDirection(videoRef, { enabled: cameraReady });
      const {
            blob,
            previewUrl,
            hasPreview,
            capturing,
            error: captureError,
            captureNow,
            retake,
            upload,
      } = useCapture(videoRef, { quality: 0.92, maxSize: 1080, unmirror: true, format: "image/jpeg" });

      //||------------------------------------------------------------------------------------------------||
	//|| Handle Camera Load
	//||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            if (cameraReady && currentStep === 0) setCurrentStep(1);
      }, [cameraReady, currentStep]);

      //||------------------------------------------------------------------------------------------------||
	//|| Handle Luma
	//||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            const myLuma = Math.floor(luma);
            if (currentLuma !== myLuma) setCurrentLuma( myLuma );
      }, [ luma ]);

      //||------------------------------------------------------------------------------------------------||
	//|| Handle Face Percentage
	//||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            if (!faceReady || faceError || facesCount === 0) return setCurrentFace(-1);
            const pct = Math.max(0, Math.min(100, Math.round(facePct)));
            if (pct !== currentFace) setCurrentFace(pct);
            if (facesCount !== 1) return setFaceInFrame(false);
            setFaceInFrame(true);
      }, [ faceReady, facesCount, faceError, facePct ]);

      //||------------------------------------------------------------------------------------------------||
	//|| Handle Face Direction
	//||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            if (faceDir !== currentDir) {
                  setCurrentDir(faceDir);
                  if (currentStep === 2) {
                        setFaceMoveCount(faceMoveCount + 1);            
                  }
            }
      }, [ faceDir ]);
      

      //||------------------------------------------------------------------------------------------------||
	//|| Handle the Steps
	//||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            if (!faceInFrame && currentStep > 1) {
                  console.log("Face out of frame");
            }else switch(currentStep) {
                  //||------------------------------------------------------------------------------------------------||
                  //|| Step 1 
                  //||------------------------------------------------------------------------------------------------||
                  case 1 : if (faceInFrame) setCurrentStep(2); break;
                  //||------------------------------------------------------------------------------------------------||
                  //|| Step 2
                  //||------------------------------------------------------------------------------------------------||
                  case 2 : if (faceMoveCount > faceMoveMin) setCurrentStep(3); break;
                  //||------------------------------------------------------------------------------------------------||
                  //|| Step 3
                  //||------------------------------------------------------------------------------------------------||
                  case 3 : break;
                  //||------------------------------------------------------------------------------------------------||
                  //|| Step 3
                  //||------------------------------------------------------------------------------------------------||
            }
      }, [currentStep, faceInFrame, currentFace, currentDir, currentLuma]);

      //||------------------------------------------------------------------------------------------------||
	//|| CaptureWrapper
	//||------------------------------------------------------------------------------------------------||

      const handleCapture = useCallback( async () => {
            await captureNow();
            setCurrentStep(4);
      }, [ captureNow ]);

      //||------------------------------------------------------------------------------------------------||
	//|| HandleRetake
	//||------------------------------------------------------------------------------------------------||

      const handleRetake = useCallback( () => {
            retake();
            setCurrentStep(3);
      }, [ retake ]);

      //||------------------------------------------------------------------------------------------------||
	//|| HandleUpload
	//||------------------------------------------------------------------------------------------------||

      const handleUpload = useCallback( async() => {
            await upload((file: Blob) => Promise.resolve(onUpload(file)));
            setCurrentStep(5);
      }, [ upload, onUpload ]);

      //||------------------------------------------------------------------------------------------------||
	//|| Spot 
	//||------------------------------------------------------------------------------------------------||

      const spotClass = "w-6 h-6 absolute z-30 text-yellow-400 p-1 border bg-black border-yellow-400"

      //||------------------------------------------------------------------------------------------------||
      //|| Stepper
      //||------------------------------------------------------------------------------------------------||

      const steps: ProgressStep[] = [
            { label: "Detect Face",       description: "Please put your face on camera" },
            { label: "Live Detection",    description: "Slowly Look around in different directions" },
            { label: "Take a Selfie",     description: "Pose for selfie" },
            { label: "Complete",          description: "All Done!" },
      ]

      //||------------------------------------------------------------------------------------------------||
	//|| JSX
	//||------------------------------------------------------------------------------------------------||

      return (
		<div className="w-full my-5 px-2">
			<div className="mx-auto max-w-5xl rounded-lg bg-black/30 p-4">
				<div className="grid grid-cols-1 md:grid-cols-2 gap-4">
					{/* Left: Camera Preview */}
					<div className="relative aspect-square overflow-hidden rounded-xl">
						{hasPreview && <img src={previewUrl!} alt="Preview" className="absolute inset-0 z-30 h-full w-full object-cover" />}
						<video ref={videoRef} playsInline muted autoPlay className="absolute inset-0 z-0 h-full w-full -scale-x-100 object-cover" />
						<canvas ref={canvasRef} className="absolute inset-0 z-10 h-full w-full -scale-x-100" />
						{/* Face Oval Guide */}
						<div className="pointer-events-none absolute inset-0 z-20 flex items-center justify-center">
							<svg viewBox="0 0 500 500" className="h-full w-full">
								<ellipse cx="250" cy="250" rx="180" ry="230" fill="none" stroke="white" strokeWidth="8" />
								<ellipse cx="250" cy="250" rx="180" ry="230" fill="none" stroke="rgba(0,0,0,0.25)" strokeWidth="1" />
							</svg>
						</div>
					</div>

					{/* Right: Status / Instructions */}
					<div className="flex flex-col gap-4 rounded-xl p-2 w-full">
						{currentStep > 0 && currentStep < 4 && (
							<div className="rounded-md bg-black/30 p-3 text-sm flex justify-center gap-4 w-[95%] mx-auto">
								<div className="flex flex-col items-center w-[32%] bg-white/20 rounded-lg p-3">
									{iconLuma(currentLuma)}
									{statusLuma(currentLuma)}
									{DEBUG && <span className="text-xs text-gray-400">Luma: {currentLuma}</span>}
								</div>
								<div className="flex flex-col items-center w-[32%] bg-white/20 rounded-lg p-3">
									{iconFace(currentFace)}
									{statusFace(currentFace)}
									{DEBUG && <span className="text-xs text-gray-400">Face: {currentFace}</span>}
								</div>
								{currentStep < 3 && (
									<div className="flex flex-col items-center w-[32%] bg-white/20 rounded-lg p-3">
										{iconFaceDir(currentDir)}
										{statusFaceDir(currentDir)}
										{DEBUG && <span className="text-xs text-gray-400">Face: {faceDir}</span>}
									</div>
								)}
								{currentStep === 3 && (
									<div className="flex flex-col items-center w-[32%] bg-white/20 rounded-lg p-3">
										{iconFaceCentered(currentDir)}
										{statusFaceCentered(currentDir)}
										{DEBUG && <span className="text-xs text-gray-400">Face: {faceDir}</span>}
									</div>
								)}
							</div>
						)}

						<div id="statusText" className="mt-auto bg-black/40 font-bold text-center py-2 text-2xl">
							{cameraReady ? (
								<>
									{currentStep === 0 ? (
										<span className="flex flex-col px-2 py-1 text-xl font-bold text-gray-400">Waiting for Camera</span>
									) : currentStep === 1 ? (
										<span className="flex flex-col px-2 py-1 text-xl font-bold text-yellow-400">
											Position Your Face in Center
										</span>
									) : currentStep === 2 ? (
										<span className="flex flex-col px-2 py-1 text-xl font-bold text-green-400">
											Move your head {faceMoveMin} times to capture
											<br />
											<ProgressBar current={faceMoveCount} goal={faceMoveMin} label="complete" />
										</span>
									) : currentStep === 3 ? (
										!boolLuma(currentLuma || 0) ? (
											<span className="flex flex-col px-2 py-1 text-xl font-bold text-yellow-300">
												Adjust your lighting
											</span>
										) : !boolFace(currentFace || 0) ? (
											<span className="flex flex-col px-2 py-1 text-xl font-bold text-yellow-300">
												Move closer or farther from the camera
											</span>
										) : !boolFaceCentered(currentDir) ? (
											<span className="flex flex-col px-2 py-1 text-xl font-bold text-yellow-300">
												Look into the camera
											</span>
										) : (
											<button
												onClick={handleCapture}
												className="rounded-md bg-green-600 px-10 py-3 text-2xl font-bold text-white hover:bg-green-500 disabled:opacity-60"
												disabled={capturing}>
												{capturing ? "Capturing..." : "Take a Photo"}
											</button>
										)
									) : currentStep === 4 ? (
										<div className="flex items-center justify-center gap-3">
											<button type="button" className="btn btn-primary" onClick={handleRetake}>
												Retake Photo
											</button>
											<button type="button" className="btn btn-secondary text-2xl h-auto p-4" onClick={handleUpload}>
												Looks Good
											</button>
										</div>
									) : currentStep === 5 ? (
										<span className="flex flex-col px-2 py-1 text-xl font-bold text-green-400">Complete</span>
									) : (
										<span className="flex flex-col px-2 py-1 text-xl font-bold text-gray-400">Preparing…</span>
									)}
								</>
							) : (
								"Initiating Camera"
							)}
						</div>
					</div>
				</div>

				{/* Bottom: Progress */}
				{currentStep > 0 && (
					<div className="mt-4">
						<ProgressSteps steps={steps} currentStep={currentStep} />
					</div>
				)}
			</div>
		</div>
	);
}
