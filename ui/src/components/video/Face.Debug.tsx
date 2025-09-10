import React, { useRef, useEffect, useState } from "react";

interface FaceDebugProps {
      onUpload: (file: File | Blob) => void;
}

export default function FaceDebug({ onUpload }: FaceDebugProps) {
      const videoRef = useRef<HTMLVideoElement | null>(null);
      const [stream, setStream] = useState<MediaStream | null>(null);
      const [capturing, setCapturing] = useState(false);

      // Start camera on mount
      useEffect(() => {
            (async () => {
                  try {
                        const mediaStream = await navigator.mediaDevices.getUserMedia({ video: true });
                        setStream(mediaStream);
                        if (videoRef.current) {
                              videoRef.current.srcObject = mediaStream;
                        }
                  } catch (e) {
                        // Handle error (optional)
                  }
            })();
            return () => {
                  if (stream) stream.getTracks().forEach(t => t.stop());
            };
      }, []);

      // Assign stream to video element (in case ref changes)
      useEffect(() => {
            if (videoRef.current && stream) {
                  videoRef.current.srcObject = stream;
            }
      }, [videoRef, stream]);

      // Take a snapshot and upload as Blob
      const handleCapture = async () => {
            if (!videoRef.current) return;
            setCapturing(true);
            const video = videoRef.current;
            const canvas = document.createElement("canvas");
            canvas.width = video.videoWidth || 480;
            canvas.height = video.videoHeight || 480;
            const ctx = canvas.getContext("2d");
            if (ctx) {
                  ctx.drawImage(video, 0, 0, canvas.width, canvas.height);
                  canvas.toBlob(blob => {
                        if (blob) onUpload(blob);
                        setCapturing(false);
                  }, "image/jpeg", 0.92);
            } else {
                  setCapturing(false);
            }
      };

      return (
            <div className="flex flex-col items-center justify-center gap-4 p-6 bg-gray-900 rounded-xl">
                  <video
                        ref={videoRef}
                        autoPlay
                        playsInline
                        muted
                        className="rounded-lg w-full max-w-xs aspect-square bg-black"
                        style={{ objectFit: "cover" }}
                  />
                  <button
                        className="px-6 py-3 rounded bg-green-700 text-white font-bold text-xl hover:bg-green-500 disabled:opacity-50"
                        onClick={handleCapture}
                        disabled={capturing}
                  >
                        {capturing ? "Capturing..." : "Take a Pic"}
                  </button>
            </div>
      );
}
