//||------------------------------------------------------------------------------------------------||
//|| Handle the Steps of the Selfie
//|| utils/vision/steps.ts
//||------------------------------------------------------------------------------------------------||

import type { Dir }                         from "./computeDirection";

//||------------------------------------------------------------------------------------------------||
//|| Allowed Steps
//||------------------------------------------------------------------------------------------------||

export const ALLOWED_STEPS : Dir[]  = ["center-top", "left-center", "right-center"];
export type StepId                  = (typeof ALLOWED_STEPS)[number];
export type { Dir };

//||------------------------------------------------------------------------------------------------||
//|| Pick Random Steps
//||------------------------------------------------------------------------------------------------||

export function pickRandomSteps(count: number): Dir[] {
	const pool              = [...ALLOWED_STEPS];
	const chosen: Dir[]     = [];
	while (chosen.length < count && pool.length > 0) {
		const idx = Math.floor(Math.random() * pool.length);
		chosen.push(pool.splice(idx, 1)[0]);
	}
	return chosen;
}

//||------------------------------------------------------------------------------------------------||
//|| Humanize Direction for Instructions
//||------------------------------------------------------------------------------------------------||

export function humanizeDir(dir: Dir): string {
	const map: Record<Dir, string> = {
		"center-top"            : "Look Up",
		"left-center"           : "Look Left",
		"right-center"          : "Look Right",
		"center-center"         : "Look Straight",
            "center-bottom"         : "Look Down",
            "left-top"              : "Look Up Left",
            "right-top"             : "Look Up Right",
            "left-bottom"           : "Look Down Left",
            "right-bottom"          : "Look Down Right",
		"unknown-unknown"       : "Unknown",
	};
	return map[dir] || dir;
}

//||------------------------------------------------------------------------------------------------||
//|| Returns a UI anchor point inside 500x500 circle for a dot/marker
//||------------------------------------------------------------------------------------------------||

export function getTargetPoint(step: StepId) {
	switch (step) {
		case "left-center":
			return {x: 125, y: 250};
		case "right-center":
			return {x: 375, y: 250};
		case "center-top":
			return {x: 250, y: 100};
	}
      return {x: 250, y: 250};
}
