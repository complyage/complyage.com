

export default function isImageBlurry(ctx: CanvasRenderingContext2D, width: number, height: number): boolean {
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

