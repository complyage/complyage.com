//||------------------------------------------------------------------------------------------------||
//|| OAuthSettingsSection
//|| Component for configuring OAuth redirect URL, private key, and verification permissions
//||------------------------------------------------------------------------------------------------||

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, {useEffect, useRef, useState}                    from "react";
import {Camera as CameraIcon, X as XIcon, Check, RefreshCcw}   from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Commponents
//||------------------------------------------------------------------------------------------------||

import SpinnerCircle                                from "../../components/base/SpinnerCircle";

//||------------------------------------------------------------------------------------------------||
//|| Helper
//||------------------------------------------------------------------------------------------------||

import isImageBlurry                                  from "../../utils/isImageBlurry";

//||------------------------------------------------------------------------------------------------||
//|| Global
//||------------------------------------------------------------------------------------------------||

declare global {
	interface Window {
		cv: any;
	}
}

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface CameraProps {
	onUpload          : (file: File | Blob) => void;
      onReset?          : () => void;
      onClose?          : () => void;
      errorMessage?     : string | null;
}


//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

export default function Camera({ onUpload, onReset, onClose, errorMessage }: CameraProps) {

      //||------------------------------------------------------------------------------------------------||
      //|| Ref
      //||------------------------------------------------------------------------------------------------||
      
      const videoRef          = useRef<HTMLVideoElement>(null);
      const canvasRef         = useRef<HTMLCanvasElement>(null);
	const overlayRef        = useRef<HTMLCanvasElement>(null);
	const fileInputRef      = useRef<HTMLInputElement>(null);
      
      //||------------------------------------------------------------------------------------------------||
      //|| Effect
      //||------------------------------------------------------------------------------------------------||
      
	const [cameraAvailable, setCameraAvailable]           = useState(false);
	const [isBlurry, setIsBlurry]                         = useState(false);
	const [cvReady, setCvReady]                           = useState(false);
      const [initError, setInitError]                       = useState<string | null>(null);
      const [actionError, setActionError]                   = useState<string | null>(null);
      const [cornersVisible, setCornersVisible]             = useState(false);
	const [capturedImage, setCapturedImage]               = useState<string | null>(null);

      //||------------------------------------------------------------------------------------------------||
      //|| Load OpenCV
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            if (window.cv && window.cv.Mat) {
                  setCvReady(true);
                  return;
            }
            if (!document.querySelector('script[src*="opencv.js"]')) {
                  const script = document.createElement("script");
                  script.src = "https://docs.opencv.org/4.x/opencv.js";
                  script.async = true;
                  document.body.appendChild(script);
            }
            const check = setInterval(() => {
                  if (window.cv && window.cv.Mat) {
                        setCvReady(true);
                        clearInterval(check);
                  }
            }, 300);
            return () => clearInterval(check);
      }, []);

      //||------------------------------------------------------------------------------------------------||
      //|| Initiate Camera
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            const init = async () => {
                  try {
                        const constraints = { video: { facingMode: "environment" } };
                        let stream: MediaStream;
                        try {
                              stream = await navigator.mediaDevices.getUserMedia(constraints);
                        } catch {
                              stream = await navigator.mediaDevices.getUserMedia({ video: true });
                        }
                        setCameraAvailable(true);
                        if (videoRef.current) {
                              videoRef.current.srcObject = stream;
                              await videoRef.current.play().catch(() => {});
                        }
                  } catch (err) {
                        console.error("Camera init failed:", err);
                        setCameraAvailable(false);
                        setInitError("Camera not available.");
                  }
            };
            init();
            return () => {
                  const tracks = (videoRef.current?.srcObject as MediaStream)?.getTracks();
                  tracks?.forEach((t) => t.stop());
            };
      }, []);

      //||------------------------------------------------------------------------------------------------||
      //|| SetAlert
      //||------------------------------------------------------------------------------------------------||
      
      const setAlert = (message: string) => {
            setActionError(message);
            setTimeout(() => setActionError(null), 3000);
      };

      //||------------------------------------------------------------------------------------------------||
      //|| OpenCV Corner Detection Helper
      //||------------------------------------------------------------------------------------------------||

      const checkCornersVisible = (ctx: CanvasRenderingContext2D, width: number, height: number): boolean => {
            if (!window.cv) return false;
            const cv = window.cv;
            const imgData = ctx.getImageData(0, 0, width, height);
            const src = cv.matFromImageData(imgData);
            const gray = new cv.Mat();
            const edges = new cv.Mat();
            cv.cvtColor(src, gray, cv.COLOR_RGBA2GRAY);
            cv.Canny(gray, edges, 50, 150);
        
            const contours = new cv.MatVector();
            const hierarchy = new cv.Mat();
            cv.findContours(edges, contours, hierarchy, cv.RETR_EXTERNAL, cv.CHAIN_APPROX_SIMPLE);
        
            let foundQuad = false;
            let maxArea = 0;
        
            for (let i = 0; i < contours.size(); i++) {
                const cnt = contours.get(i);
                const approx = new cv.Mat();
                cv.approxPolyDP(cnt, approx, 0.02 * cv.arcLength(cnt, true), true);
        
                if (approx.rows === 4) {
                    const rect = cv.boundingRect(cnt);
                    const area = rect.width * rect.height;
                    const aspect = rect.width / rect.height;
        
                    // ✅ Looser thresholds
                    const minArea = width * height * 0.05;           // 15% of frame
                    const ratioOK = aspect > 0.8 && aspect < 3.0; // allow more shapes
                    const sizeOK = area > minArea;
        
                    // ✅ Relaxed center check
                    const centerX = rect.x + rect.width / 2;
                    const centerY = rect.y + rect.height / 2;
                    const centerOK = Math.abs(centerX - width / 2) < width / 2 && Math.abs(centerY - height / 2) < height / 2;
        
                    if (sizeOK && ratioOK && centerOK && area > maxArea) {
                        foundQuad = true;
                        maxArea = area;
                    }
                }
        
                approx.delete();
                cnt.delete();
            }
        
            src.delete(); gray.delete(); edges.delete(); contours.delete(); hierarchy.delete();
            return foundQuad;
      };        

      //||------------------------------------------------------------------------------------------------||
      //|| Blur Detection
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            if (!cvReady) return;
            const interval = setInterval(() => {
                  if (!videoRef.current || !canvasRef.current || capturedImage) return;
                  const video = videoRef.current;
                  const canvas = canvasRef.current;
                  const ctx = canvas.getContext("2d", { willReadFrequently: true });
                  if (!ctx || !video.videoWidth) return;

                  canvas.width = video.videoWidth / 2;
                  canvas.height = video.videoHeight / 2;
                  ctx.drawImage(video, 0, 0, canvas.width, canvas.height);

                  setIsBlurry(isImageBlurry(ctx, canvas.width, canvas.height));
                  setCornersVisible(checkCornersVisible(ctx, canvas.width, canvas.height));
            }, 600);
            return () => clearInterval(interval);            
      }, [cvReady, capturedImage]);

      //||------------------------------------------------------------------------------------------------||
      //|| Start Camera
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            startCamera();
            return () => {
                const tracks = (videoRef.current?.srcObject as MediaStream)?.getTracks();
                tracks?.forEach((t) => t.stop());
            };
      }, []);  

      //||------------------------------------------------------------------------------------------------||
      //|| Handle Capture
      //||------------------------------------------------------------------------------------------------||

      const handleCapture = () => {
            if (isBlurry) return setAlert("Image too blurry. Please retake.");
            if (!cornersVisible) return setAlert("Make sure all 4 corners of the ID are visible.");
            const video = videoRef.current;
            if (!video) return;
            const canvas = document.createElement("canvas");
            canvas.width = video.videoWidth;
            canvas.height = video.videoHeight;
            const ctx = canvas.getContext("2d", { willReadFrequently: true });
            if (!ctx) return;
            ctx.drawImage(video, 0, 0);
            const imgData = canvas.toDataURL("image/png");
            setCapturedImage(imgData);
            canvas.toBlob((blob) => blob && onUpload(blob), "image/png");
            onUpload(new Blob([imgData], { type: "image/png" }));
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Start Camera
      //||------------------------------------------------------------------------------------------------||

      const startCamera = async () => {
            try {
                const constraints = { video: { facingMode: "environment" } };
                let stream: MediaStream;
                try {
                    stream = await navigator.mediaDevices.getUserMedia(constraints);
                } catch {
                    stream = await navigator.mediaDevices.getUserMedia({ video: true });
                }
                setCameraAvailable(true);
                setInitError(null);
                if (videoRef.current) {
                    videoRef.current.srcObject = stream;
                    await videoRef.current.play().catch(() => {});
                }
            } catch (err) {
                console.error("Camera init failed:", err);
                setCameraAvailable(false);
                setInitError("Camera not available.");
            }
      };      

      //||------------------------------------------------------------------------------------------------||
      //|| Reset Capture
      //||------------------------------------------------------------------------------------------------||

      const handleReset = () => {
            setCapturedImage(null);
            if (onReset) onReset();

            const tracks = (videoRef.current?.srcObject as MediaStream)?.getTracks();
            const inactive = !tracks || tracks.every((t) => t.readyState === "ended");
            if (inactive) {
                  navigator.mediaDevices.getUserMedia({ video: true }).then((stream) => {
                        if (videoRef.current) {
                              videoRef.current.srcObject = stream;
                              videoRef.current.play().catch(() => {});
                        }
                  });
            }
      };    

      //||------------------------------------------------------------------------------------------------||
      //|| Reset Capture
      //||------------------------------------------------------------------------------------------------||

	return (
            <div className="flex flex-col w-full relative items-center justify-center mx-auto">
                  
                  <button onClick={onClose} className="absolute top-2 right-2 rounded-xl bg-black/90 hover:text-gray-700 transition p-2 cursor-pointer">
                        <XIcon className="w-8 h-8 text-white" />
                  </button>

                  {/* Camera / Preview */}
                  {cameraAvailable ? (
                        <div className="relative w-full aspect-video rounded-lg overflow-hidden mx-auto border-2 ">
                              {/* Live video or captured image */}
                              <video
                                    ref={videoRef}
                                    autoPlay
                                    playsInline
                                    className={`absolute inset-0 w-full h-auto object-cover ${capturedImage ? "hidden" : ""}`}
                              />
                              {capturedImage && <img src={capturedImage} alt="Captured" className="absolute inset-0 w-full h-auto z-10" />}
                              <canvas ref={overlayRef} className={`absolute inset-0 w-full h-full pointer-events-none ${capturedImage ? "hidden" : ""}`} />
                              <canvas ref={canvasRef} className="hidden" />

                              {/* Capture / Reset Button */}
                              {cameraAvailable && (
                                     errorMessage ? (
                                          <button
                                                type="button"
                                                onClick={handleReset}
                                                className="absolute bottom-4 left-1/2 -translate-x-1/2 bg-white rounded-full p-4 shadow-lg hover:bg-gray-200 transition z-20"
                                          >
                                                <XIcon className="w-6 h-6 text-black" />
                                          </button>
                                    ) : (
                                          <button
                                                type="button"
                                                onClick={capturedImage ? handleReset : handleCapture}
                                                className="absolute bottom-4 left-1/2 -translate-x-1/2 bg-white rounded-full p-4 shadow-lg hover:bg-gray-200 transition z-20"
                                          >
                                                {capturedImage ? <SpinnerCircle /> : <CameraIcon className="w-6 h-6 text-black" />}
                                          </button>
                                    )
                              )}                

                              {!capturedImage && (
                                    <div className="absolute top-2 items-center w-[50%] left-1/2 -translate-x-1/2 flex flex-row gap-4">
                                          <div className={`text-xs font-bold p-2 w-1/2 rounded-md text-white ${isBlurry ? "bg-red-500/70" : "bg-black/80"}`}>
                                                {isBlurry ? ( <><XIcon className="inline w-4 h-4 mr-1" /> Hold Still</> ) : ( <><Check className="inline w-4 h-4 mr-1" />Clear</>)}
                                          </div>

                                          <div className={`text-xs items-center font-bold p-2 w-1/2 rounded-md text-white ${!cornersVisible ? "bg-red-500/70" : "bg-black/70"}`}>
                                                {cornersVisible ? ( <><Check className="inline w-4 h-4 mr-1" /> ID Corners Visible</> ) : ( <><XIcon className="inline w-4 h-4 mr-1" /> Corners Missing</> )}
                                          </div>
                                    </div>
                              )}
                        </div>
			) : (
				<div className="flex bg-black w-full min-h-[300px] items-center justify-center rounded-lg">
                              <div className="flex flex-col bg-black w-full min-h-[300px] items-center justify-center rounded-lg">
                                    <p className="bg-white/10 text-red-500 p-4 rounded-lg">{initError}</p>
                                    <button onClick={startCamera} className="mt-4 px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700 transition"><RefreshCcw /></button>
                              </div>
				</div>
			)}
                  { actionError && (
                        <div className="mt-4 text-red-500">
                              <p className="bg-black/50 p-3 rounded-lg text-white mt-5">{actionError}</p>
                        </div>
                  )}
                  { errorMessage && (
                        <div className="mt-4 text-red-500">
                              <p className="bg-red-500/50 p-3 rounded-lg text-white mt-5">{errorMessage}</p>
                        </div>
                  )}
		</div>
	);
}
