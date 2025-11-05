//||------------------------------------------------------------------------------------------------||
//|| Compute Direction
//|| Compute head direction from MediaPipe FaceMesh landmarks.
//||------------------------------------------------------------------------------------------------||

export type Dir =
	| "left-center"
	| "right-center"
	| "center-center"
	| "center-top"
	| "left-top"
	| "right-top"
	| "center-bottom"
	| "left-bottom"
	| "right-bottom"
	| "unknown-unknown";

export type DirVector = {
	horizontal: "left" | "center" | "right" | "unknown";
	vertical: "top" | "center" | "bottom" | "unknown";
};

//||------------------------------------------------------------------------------------------------||
//|| Compute head direction from MediaPipe FaceMesh landmarks.
//||------------------------------------------------------------------------------------------------||

type Opts = {
	mirrored?         : boolean; // your preview is mirrored; flip yaw
	yawCenter?        : number; // |yaw| <= yawCenter -> center
	pitchUpEdge?      : number; // ratio threshold for "up"  (lower = more up)
	pitchDownEdge?    : number; // ratio threshold for "down" (higher = more down)
};

//||------------------------------------------------------------------------------------------------||
//|| Compute head direction from MediaPipe FaceMesh landmarks.
//|| 
//|| Yaw: normalized nose x offset vs. eye span
//|| Pitch: ratio of nose y between eyes and mouth (~0.50 neutral)
//|| Returns a Dir string like "center-top"
//||------------------------------------------------------------------------------------------------||

export function computeDirection(face: any, opts: Opts = {}): Dir {
	const {
		mirrored          = true,
		yawCenter         = 0.22, // was ~0.15–0.20; a bit looser
		pitchUpEdge       = 0.47, // consider "up" if ratio < 0.47 (was ~0.38 which is too strict)
		pitchDownEdge     = 0.62, // consider "down" if ratio > 0.62
	} = opts;

	const P = (i: number) => face?.keypoints?.[i];

      //||------------------------------------------------------------------------------------------------||
      //|| Required landmarks (MediaPipe canonical indices)
      //||------------------------------------------------------------------------------------------------||

	const L_OUT = P(33),
		L_IN = P(133),
		L_UP = P(159),
		L_DN = P(145);
	const R_OUT = P(263),
		R_IN = P(362),
		R_UP = P(386),
		R_DN = P(374);
	const M_LEFT = P(61),
		M_RIGHT = P(291);
	const NOSE = P(4);

	if (!L_OUT || !L_IN || !L_UP || !L_DN || !R_OUT || !R_IN || !R_UP || !R_DN || !M_LEFT || !M_RIGHT || !NOSE) {
		return "unknown-unknown";
	}

      //||------------------------------------------------------------------------------------------------||
      //|| Eye centers & midpoints
      //||------------------------------------------------------------------------------------------------||

	const leftEyeCenter = {
		x: (L_OUT.x + L_IN.x + L_UP.x + L_DN.x) / 4,
		y: (L_OUT.y + L_IN.y + L_UP.y + L_DN.y) / 4,
	};
	const rightEyeCenter = {
		x: (R_OUT.x + R_IN.x + R_UP.x + R_DN.x) / 4,
		y: (R_OUT.y + R_IN.y + R_UP.y + R_DN.y) / 4,
	};
	const eyesMid = {
		x: (leftEyeCenter.x + rightEyeCenter.x) / 2,
		y: (leftEyeCenter.y + rightEyeCenter.y) / 2,
	};
	const mouthMid = {
		x: (M_LEFT.x + M_RIGHT.x) / 2,
		y: (M_LEFT.y + M_RIGHT.y) / 2,
	};

      //||------------------------------------------------------------------------------------------------||
      //|| YAW (left/right)
      //||------------------------------------------------------------------------------------------------||

      const eyeSpan           = R_IN.x - L_IN.x || 1e-6;
	let yaw                 = (NOSE.x - eyesMid.x) / eyeSpan; // <0 left, >0 right in image coords
	if (mirrored) yaw = -yaw; // flip because preview is mirrored
	let h: "left" | "center" | "right";
	if (yaw < -yawCenter) h = "left"; else if (yaw > yawCenter) h = "right"; else h = "center";

      //||------------------------------------------------------------------------------------------------||
      //|| PITCH (up/down)
      //||------------------------------------------------------------------------------------------------||

	const denom             = mouthMid.y - eyesMid.y || 1e-6; // face height
	const pitchRatio        = (NOSE.y - eyesMid.y) / denom; // ~0.50 neutral; lower = up
	let v: "top" | "center" | "bottom";
	if (pitchRatio < pitchUpEdge) v = "top"; else if (pitchRatio > pitchDownEdge) v = "bottom"; else v = "center";

      //||------------------------------------------------------------------------------------------------||
      //|| Create the Dir
      //||------------------------------------------------------------------------------------------------||

	const dir = `${h}-${v}` as Dir;

      //||------------------------------------------------------------------------------------------------||
      //|| Dir
      //||------------------------------------------------------------------------------------------------||

      if (dir === "left-top")             return "left-center";
	if (dir === "right-top")            return "right-center";
	if (dir === "left-bottom")          return "left-center";
	if (dir === "right-bottom")         return "right-center";
	if (dir === "center-bottom")        return "center-center";

      //||------------------------------------------------------------------------------------------------||
      //|| Return
      //||------------------------------------------------------------------------------------------------||

      return dir;
}
