
export default function checkCornersVisible(ctx: CanvasRenderingContext2D, width: number, height: number): boolean => {
      const imgData = ctx.getImageData(0, 0, width, height);
      const src = cv.matFromImageData(imgData);
      const gray = new cv.Mat();
      const edges = new cv.Mat();
      cv.cvtColor(src, gray, cv.COLOR_RGBA2GRAY);
      cv.Canny(gray, edges, 50, 150);
  
      const contours = new cv.MatVector();
      const hierarchy = new cv.Mat();
      cv.findContours(edges, contours, hierarchy, cv.RETR_EXTERNAL, cv.CHAIN_APPROX_SIMPLE);
  
      let found = false;
      for (let i = 0; i < contours.size(); i++) {
          const cnt = contours.get(i);
          const approx = new cv.Mat();
          cv.approxPolyDP(cnt, approx, 0.02 * cv.arcLength(cnt, true), true);
  
          // ✅ Look for quadrilateral (4 points)
          if (approx.rows === 4) {
              found = true;
          }
          approx.delete();
          cnt.delete();
      }
  
      src.delete(); gray.delete(); edges.delete(); contours.delete(); hierarchy.delete();
      return found;
  };