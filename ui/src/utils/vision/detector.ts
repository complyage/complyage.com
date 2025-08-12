//||------------------------------------------------------------------------------------------------||
//|| Detector
//|| utils/vision/detector.ts
//||------------------------------------------------------------------------------------------------||

import * as tf                            from "@tensorflow/tfjs";
import * as faceLandmarksDetection        from "@tensorflow-models/face-landmarks-detection";
import "@tensorflow/tfjs-backend-webgl";

//||------------------------------------------------------------------------------------------------||
//|| Export Detector
//||------------------------------------------------------------------------------------------------||

export type Detector = Awaited<ReturnType<typeof faceLandmarksDetection.createDetector>>;

//||------------------------------------------------------------------------------------------------||
//|| Initialize Detector
//||------------------------------------------------------------------------------------------------||

export async function initDetector(): Promise<Detector> {
	await tf.setBackend("webgl");
	await tf.ready();
	const detector = await faceLandmarksDetection.createDetector(faceLandmarksDetection.SupportedModels.MediaPipeFaceMesh, {
		runtime: "tfjs",
		refineLandmarks: true,
		maxFaces: 1,
	});
	return detector;
}

//||------------------------------------------------------------------------------------------------||
//|| Wait for Video
//||------------------------------------------------------------------------------------------------||

export async function waitForVideoReady(video: HTMLVideoElement): Promise<void> {
	await video.play().catch(() => {});
	await new Promise<void>((resolve) => {
		const tryReady = () => {
			if (video.readyState >= 2 && video.videoWidth > 0 && video.videoHeight > 0) resolve();
			else requestAnimationFrame(tryReady);
		};
		tryReady();
	});
}
