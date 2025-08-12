//||------------------------------------------------------------------------------------------------||
//|| OAuthSettingsSection
//|| Component for configuring OAuth redirect URL, private key, and verification permissions
//||------------------------------------------------------------------------------------------------||

import React, {useEffect, useRef, useState} from "react";
import {Camera, X, Upload} from "lucide-react";

declare global {
	interface Window {
		cv: any;
	}
}

interface IDVerificationProps {
	onNext: (file: File | Blob) => void;
}

function isImageBlurry(ctx: CanvasRenderingContext2D, width: number, height: number): boolean {
	if (!ctx || width < 10 || height < 10) return true;
	const data = ctx.getImageData(0, 0, width, height).data;
	const gray: number[] = [];
	for (let i = 0; i < data.length; i += 4) gray.push((data[i] + data[i + 1] + data[i + 2]) / 3);
	let variance = 0;
	for (let i = 1; i < gray.length - 1; i++) {
		const diff = gray[i - 1] - 2 * gray[i] + gray[i + 1];
		variance += diff * diff;
	}
	variance /= gray.length;
	return variance < 100;
}

export default function IDVerification({onNext}: IDVerificationProps) {
	const videoRef = useRef<HTMLVideoElement>(null);
	const canvasRef = useRef<HTMLCanvasElement>(null);
	const overlayRef = useRef<HTMLCanvasElement>(null);
	const fileInputRef = useRef<HTMLInputElement>(null);

	const [cameraAvailable, setCameraAvailable] = useState(false);
	const [errorMessage, setErrorMessage] = useState("");
	const [docType, setDocType] = useState<"id" | "passport">("id");
	const [isBlurry, setIsBlurry] = useState(false);
	const [cvReady, setCvReady] = useState(false);
	const [capturedImage, setCapturedImage] = useState<string | null>(null);

	// ✅ Load OpenCV
	useEffect(() => {
		if (window.cv && window.cv.Mat) {
			setCvReady(true);
			return; // ✅ Already loaded, skip
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

	// ✅ Initialize Camera
	useEffect(() => {
		const init = async () => {
			try {
				const stream = await navigator.mediaDevices.getUserMedia({video: {facingMode: "environment"}});
				setCameraAvailable(true);
				if (videoRef.current) videoRef.current.srcObject = stream;
			} catch {
				setCameraAvailable(false);
				setErrorMessage("Camera not available.");
			}
		};
		init();
		return () => {
			if (videoRef.current?.srcObject) (videoRef.current.srcObject as MediaStream).getTracks().forEach((t) => t.stop());
		};
	}, []);

	// ✅ Real-Time Blur Detection
	useEffect(() => {
		if (!cvReady) return;
		const interval = setInterval(() => {
			if (!videoRef.current || !canvasRef.current || capturedImage) return;
			const video = videoRef.current;
			const canvas = canvasRef.current;
			const ctx = canvas.getContext("2d", {willReadFrequently: true});
			if (!ctx || !video.videoWidth) return;
			canvas.width = video.videoWidth / 2;
			canvas.height = video.videoHeight / 2;
			ctx.drawImage(video, 0, 0, canvas.width, canvas.height);
			setIsBlurry(isImageBlurry(ctx, canvas.width, canvas.height));
		}, 500);
		return () => clearInterval(interval);
	}, [cvReady, capturedImage]);

	// ✅ Capture from Camera
	const handleCapture = () => {
		if (isBlurry) return alert("Image too blurry. Please retake.");
		const video = videoRef.current;
		if (!video) return;
		const canvas = document.createElement("canvas");
		canvas.width = video.videoWidth;
		canvas.height = video.videoHeight;
		const ctx = canvas.getContext("2d", {willReadFrequently: true});
		if (!ctx) return;
		ctx.drawImage(video, 0, 0);
		const imgData = canvas.toDataURL("image/jpeg");
		setCapturedImage(imgData);
		canvas.toBlob((blob) => blob && onNext(blob), "image/jpeg");
	};

	// ✅ Manual Upload (Fallback)
	const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		if (e.target.files && e.target.files[0]) {
			const file = e.target.files[0];
			const url = URL.createObjectURL(file);
			setCapturedImage(url);
			onNext(file);
		}
	};

	// ✅ Reset (Retake)
	const handleReset = () => {
		setCapturedImage(null);
		// ✅ if camera is available, restart it
		if (videoRef.current?.srcObject === null) {
			navigator.mediaDevices.getUserMedia({video: {facingMode: "environment"}}).then((stream) => {
				if (videoRef.current) videoRef.current.srcObject = stream;
			});
		}
	};

	return (
		<div className="w-full max-w-2xl mx-auto text-center p-4">
			{/* DocType Selector */}
			<div className="flex gap-4 mb-3 justify-center">
				<label className="flex items-center gap-2 bg-black/40 p-2 rounded-lg px-5">
					<input type="radio" name="docType" value="id" checked={docType === "id"} onChange={() => setDocType("id")} />
					Government ID
				</label>
				<label className="flex items-center gap-2 bg-black/40 p-2 rounded-lg px-5">
					<input type="radio" name="docType" value="passport" checked={docType === "passport"} onChange={() => setDocType("passport")} />
					Passport
				</label>
			</div>

			{/* Camera / Preview */}
			{cameraAvailable ? (
				<div className="relative w-full max-w-md aspect-video rounded-lg overflow-hidden mx-auto">
					{/* Live video or captured image */}
					<video
						ref={videoRef}
						autoPlay
						playsInline
						className={`absolute inset-0 w-full h-full object-cover ${capturedImage ? "hidden" : ""}`}
					/>
					{capturedImage && <img src={capturedImage} alt="Captured" className="absolute inset-0 w-full h-full object-cover z-10" />}
					<canvas ref={overlayRef} className={`absolute inset-0 w-full h-full pointer-events-none ${capturedImage ? "hidden" : ""}`} />
					<canvas ref={canvasRef} className="hidden" />

					{/* Blur Indicator */}
					{!capturedImage && (
						<div
							className={`absolute top-2 left-1/2 -translate-x-1/2 px-3 py-1 rounded text-white ${
								isBlurry ? "bg-red-600" : "bg-green-600"
							}`}>
							{isBlurry ? "❌ Blurry" : "✅ Clear"}
						</div>
					)}

					{/* Capture / Reset Button */}
					{cameraAvailable && (
						<button
							type="button"
							onClick={capturedImage ? handleReset : handleCapture}
							className="absolute bottom-4 left-1/2 -translate-x-1/2 bg-white rounded-full p-4 shadow-lg hover:bg-gray-200 transition z-20">
							{capturedImage ? <X className="w-6 h-6 text-black" /> : <Camera className="w-6 h-6 text-black" />}
						</button>
					)}
				</div>
			) : (
				<div className="flex bg-black w-full min-h-[300px] items-center justify-center rounded-lg">
					<p className="bg-white/10 text-red-500 p-4 rounded-lg">{errorMessage}</p>
				</div>
			)}

			{/* OR Manual Upload */}
			<p className="mt-4 text-gray-200">or</p>
			<button type="button" className="btn btn-secondary mt-5 px-5" onClick={() => fileInputRef.current?.click()}>
				<Upload className="inline-block mr-2 w-5 h-5" /> Upload ID Front
			</button>
			<input type="file" accept="image/*" ref={fileInputRef} onChange={handleFileChange} className="hidden" />
		</div>
	);
}
