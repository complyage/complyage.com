import React from "react";

export type VerificationStatus = {
	type: string;
	label: string;
	blurb: string;
	complete: boolean;
	icon: React.ReactNode;
};