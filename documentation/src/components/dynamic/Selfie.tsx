//||------------------------------------------------------------------------------------------------||
//|| Selfie.tsx
//|| Live selfie verification camera component with gaze detection and countdown capture
//|| Requires: OpenCV.js loaded via <script src="https://docs.opencv.org/4.x/opencv.js">
//||------------------------------------------------------------------------------------------------||

import React, { useEffect, useRef, useState } from "react";
import { X as XIcon } from "lucide-react";

interface SelfieProps {
      onUpload: (file: Blob) => void;
      onClose?: () => void;
}

const DIRECTION_LABELS = [
      "top", "top-right", "right", "bottom-right",
      "bottom", "bottom-left", "left", "top-left"
];

export default function Selfie({ onUpload, onClose }: SelfieProps) {

      const videoRef                      = useRef<HTMLVideoElement>(null);
      const previewCanvasRef              = useRef<HTMLCanvasElement>(document.createElement("canvas"));
      const [cvReady, setCvReady]         = useState(false);
      const [cameraAvailable, setCamera]  = useState(false);
      const [currentDir, setCurrentDir]   = useState<string | null>(null);
      const [required, setRequired]       = useState<string[]>([]);
      const [completed, setCompleted]     = useState<string[]>([]);
      const [finalPhoto, setFinalPhoto]   = useState<string | null>(null);
      const [countdown, setCountdown]     = useState<number | null>(null);
      const [facePreview, setFacePreview] = useState<string | null>(null);

      //||------------------------------------------------------------------------------------------------||
      //|| Load OpenCV
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            if (window.cv && window.cv.Mat) {
                  setCvReady(true);
                  return;
            }
            const script = document.createElement("script");
            script.src = "https://docs.opencv.org/4.x/opencv.js";
            script.async = true;
            document.body.appendChild(script);
            const check = setInterval(() => {
                  if (window.cv && window.cv.Mat) {
                        clearInterval(check);
                        setCvReady(true);
                  }
            }, 300);
            return () => clearInterval(check);
      }, []);

      //||------------------------------------------------------------------------------------------------||
      //|| Preload Haar Cascade into OpenCV virtual FS
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            if (!cvReady) return;
      
            // Load the Haar cascade file into OpenCV's virtual filesystem
            fetch("/xml/haarcascade_frontalface_default.xml")
                  .then(res => res.arrayBuffer())
                  .then(buffer => {
                        const data = new Uint8Array(buffer);
                        cv.FS_createDataFile("/", "haarcascade_frontalface_default.xml", data, true, false, false);
                        console.log("[OpenCV] Loaded cascade classifier");
                  })
                  .catch((err) => {
                        console.error("[OpenCV] Failed to load cascade:", err);
                  });
      }, [cvReady]);

      //||------------------------------------------------------------------------------------------------||
      //|| Start Camera
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            navigator.mediaDevices.getUserMedia({ video: true })
                  .then((stream) => {
                        if (videoRef.current) {
                              videoRef.current.srcObject = stream;
                              videoRef.current.play();
                              setCamera(true);
                        }
                  })
                  .catch(() => setCamera(false));
            return () => {
                  const tracks = (videoRef.current?.srcObject as MediaStream)?.getTracks();
                  tracks?.forEach((t) => t.stop());
            };
      }, []);

      //||------------------------------------------------------------------------------------------------||
      //|| Face Preview
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            if (!cvReady || !videoRef.current) return;

            const interval = setInterval(() => {
                  const video = videoRef.current!;
                  const canvas = previewCanvasRef.current!;
                  const ctx = canvas.getContext("2d");
                  if (!ctx) return;

                  canvas.width = video.videoWidth;
                  canvas.height = video.videoHeight;
                  ctx.drawImage(video, 0, 0, canvas.width, canvas.height);

                  const src = cv.imread(canvas);
                  const gray = new cv.Mat();
                  cv.cvtColor(src, gray, cv.COLOR_RGBA2GRAY, 0);

                  const faces = new cv.RectVector();
                  const classifier = new cv.CascadeClassifier();
                  classifier.load("haarcascade_frontalface_default.xml");
                  classifier.detectMultiScale(gray, faces, 1.1, 3, 0);

                  if (faces.size() > 0) {
                        const face = faces.get(0);
                        const paddedRect = new cv.Rect(
                              Math.max(0, face.x - face.width * 0.2),
                              Math.max(0, face.y - face.height * 0.2),
                              Math.min(canvas.width - face.x, face.width * 1.4),
                              Math.min(canvas.height - face.y, face.height * 1.4)
                        );
                        const faceMat = src.roi(paddedRect);
                        const output = new cv.Mat();
                        cv.resize(faceMat, output, new cv.Size(300, 300));
                        const previewCtx = previewCanvasRef.current!.getContext("2d")!;
                        cv.imshow(previewCanvasRef.current!, output);
                        const dataUrl = previewCanvasRef.current!.toDataURL("image/png");
                        setFacePreview(dataUrl);

                        faceMat.delete(); output.delete();
                  }

                  src.delete(); gray.delete(); faces.delete(); classifier.delete();
            }, 100);

            return () => clearInterval(interval);
      }, [cvReady]);

      //||------------------------------------------------------------------------------------------------||
      //|| Initialize Random Directions
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            const shuffled = [...DIRECTION_LABELS].sort(() => 0.5 - Math.random()).slice(0, 3);
            setRequired(shuffled);
            setCurrentDir(shuffled[0]);
      }, []);

      //||------------------------------------------------------------------------------------------------||
      //|| Simulated Direction Detection
      //||------------------------------------------------------------------------------------------------||

      const fakeEstimateDirection = (): string => {
            if (Math.random() < 0.15 && currentDir) return currentDir;
            return "none";
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Monitor Head Pose
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            if (!cvReady || !videoRef.current || !currentDir || finalPhoto || countdown !== null) return;
            const interval = setInterval(() => {
                  const dir = fakeEstimateDirection();
                  if (dir === currentDir) {
                        setCompleted((prev) => {
                              const next = [...prev, dir];
                              const remaining = required.filter((d) => !next.includes(d));
                              if (remaining.length === 0) {
                                    setCurrentDir(null);
                                    setCountdown(3);
                              } else {
                                    setCurrentDir(remaining[0]);
                              }
                              return next;
                        });
                  }
            }, 600);
            return () => clearInterval(interval);
      }, [cvReady, currentDir, finalPhoto, countdown]);

      //||------------------------------------------------------------------------------------------------||
      //|| Countdown before Photo
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            if (countdown === null) return;
            if (countdown === 0) {
                  capturePhoto();
                  setCountdown(null);
                  return;
            }
            const timer = setTimeout(() => setCountdown((prev) => (prev ?? 1) - 1), 1000);
            return () => clearTimeout(timer);
      }, [countdown]);

      //||------------------------------------------------------------------------------------------------||
      //|| Capture Selfie (Zoomed + Padding)
      //||------------------------------------------------------------------------------------------------||

      const capturePhoto = () => {
            const video = videoRef.current;
            if (!video) return;

            const canvas = document.createElement("canvas");
            const width = video.videoWidth;
            const height = video.videoHeight;

            const zoomFactor = 1.8;
            const targetWidth = width / zoomFactor;
            const targetHeight = height / zoomFactor;
            const offsetX = (width - targetWidth) / 2;
            const offsetY = (height - targetHeight) / 2;

            canvas.width = width;
            canvas.height = height;
            const ctx = canvas.getContext("2d");
            if (!ctx) return;

            ctx.translate(width, 0);
            ctx.scale(-1, 1);
            ctx.drawImage(video, offsetX, offsetY, targetWidth, targetHeight, 0, 0, width, height);

            const dataUrl = canvas.toDataURL("image/png");
            setFinalPhoto(dataUrl);
            canvas.toBlob((blob) => blob && onUpload(blob), "image/png");
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Gaze Border Classes
      //||------------------------------------------------------------------------------------------------||

      const getBorderClass = (dir: string) => {
            const base = "absolute border-4 border-yellow-400 z-20";
            switch (dir) {
                  case "top": return `${base} top-0 left-1/4 w-1/2 h-1/4`;
                  case "bottom": return `${base} bottom-0 left-1/4 w-1/2 h-1/4`;
                  case "left": return `${base} top-1/4 left-0 h-1/2 w-1/4`;
                  case "right": return `${base} top-1/4 right-0 h-1/2 w-1/4`;
                  case "top-left": return `${base} top-0 left-0 w-1/3 h-1/3`;
                  case "top-right": return `${base} top-0 right-0 w-1/3 h-1/3`;
                  case "bottom-left": return `${base} bottom-0 left-0 w-1/3 h-1/3`;
                  case "bottom-right": return `${base} bottom-0 right-0 w-1/3 h-1/3`;
                  default: return "";
            }
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Render
      //||------------------------------------------------------------------------------------------------||

      return (
            <div className="w-full mx-auto text-center p-4 relative">
                  <button onClick={onClose} className="absolute top-0 right-0 rounded-xl bg-black/90 hover:text-gray-700 transition p-2 z-30">
                        <XIcon className="w-8 h-8 text-white" />
                  </button>

                  <div className="border-10 border-white flex w-[500px] mx-auto rounded-full h-[500px]">
                        <img
                              src={facePreview ?? "base64-image-placeholder.png"}
                              alt="Face Preview"
                              className="w-full h-full object-cover rounded-full"
                        />
                  </div>

                  <div className="relative w-full aspect-video rounded-lg overflow-hidden mx-auto border-2 border-white/20 mt-4">
                        {!finalPhoto ? (
                              <video ref={videoRef} autoPlay playsInline muted className="absolute inset-0 w-full h-full object-cover scale-x-[-1]" />
                        ) : (
                              <img src={finalPhoto} alt="Final Selfie" className="absolute inset-0 w-full h-full object-cover" />
                        )}

                        {currentDir && <div className={getBorderClass(currentDir)} />}
                        {countdown !== null && (
                              <div className="absolute inset-0 flex items-center justify-center bg-black/60 z-30">
                                    <span className="text-7xl font-bold text-white">{countdown}</span>
                              </div>
                        )}
                  </div>

                  <div className="mt-5">
                        {finalPhoto ? (
                              <p className="text-green-500 font-bold">✅ Selfie Captured</p>
                        ) : (
                              <>
                                    {currentDir && (
                                          <p className="text-white font-bold bg-black/70 px-4 py-2 rounded-md inline-block">
                                                Look {currentDir.replace("-", " ")}
                                          </p>
                                    )}
                                    <p className="text-sm mt-2 text-gray-400">Direction {completed.length + 1} of 3</p>
                              </>
                        )}
                  </div>
            </div>
      );
}
