//||------------------------------------------------------------------------------------------------||
//|| Get Face Area Percentage from Face Object
//|| utils/vision/faceMetrics.ts
//||------------------------------------------------------------------------------------------------||

 export function getFaceAreaPct(face: any, canvasSize = 500): number {
		const box = face?.box;
		if (!box) return 0;
		const areaPct = ((box.width * box.height) / (canvasSize * canvasSize)) * 100;
		return Math.max(0, Math.min(100, areaPct));
 }