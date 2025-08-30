//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, {useEffect, useState} from "react";
import {useNavigate} from "react-router-dom";

import {CheckCircle, ArrowRight, CreditCard, MapPin, Mail, Phone, IdCard, Smile, Image as ImageIcon, Shield} from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

import {VerificationStatus} from "../../interfaces/verificationStatus";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import MembersLayout from "../../layouts/MembersLayout";

//||------------------------------------------------------------------------------------------------||
//|| API Verification Interface
//||------------------------------------------------------------------------------------------------||

interface ApiVerification {
	vType: string;
	vStatus: string;
}

//||------------------------------------------------------------------------------------------------||
//|| Location Interfaces
//||------------------------------------------------------------------------------------------------||

interface LocationInfo {
	ipAddress?: string;
	city?: string;
	state?: string;
	country?: string;
	allowed?: string[]; // e.g., ['IDEN','CRCD'] (VerifyType codes)
}

interface ZoneRecord {
	id: number;
	state?: string | null;
	country?: string | null;
	requirements?: string | null; // CSV like "IDEN,CRCD" (after your migration)
}

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function Dashboard() {
	//||------------------------------------------------------------------------------------------------||
	//|| Const
	//||------------------------------------------------------------------------------------------------||

	const [verifications, setVerifications] = useState<VerificationStatus[]>([]);
	const [location, setLocation] = useState<LocationInfo | null>(null);
	const navigate = useNavigate();

	//||------------------------------------------------------------------------------------------------||
	//|| Base Verification
	//||------------------------------------------------------------------------------------------------||

	const baseVerifications: VerificationStatus[] = [
		{type: "MAIL", label: "Email", blurb: "Confirm your email to receive account updates.", complete: false, icon: <Mail />},
		{type: "IDEN", label: "ID / Age", blurb: "Verify your age with a government ID.", complete: false, icon: <IdCard />},
		{type: "PHNE", label: "Phone", blurb: "Add and confirm your phone number.", complete: false, icon: <Phone />},
		{type: "ADDR", label: "Address", blurb: "Verify your billing or home address.", complete: false, icon: <MapPin />},
		{type: "CRCD", label: "Credit Card", blurb: "Secure your account with a valid card on file.", complete: false, icon: <CreditCard />},
		{type: "FACE", label: "Facial Age Estimation", blurb: "Complete a quick facial age scan.", complete: false, icon: <Smile />},
	];

	//||------------------------------------------------------------------------------------------------||
	//|| Helpers: Map code -> label
	//||------------------------------------------------------------------------------------------------||

	const codeToLabel = (code: string) => {
		const found = baseVerifications.find((v) => v.type === code);
		return found ? found.label : code;
	};

	//||------------------------------------------------------------------------------------------------||
	//|| Get Statuses
	//||------------------------------------------------------------------------------------------------||

	useEffect(() => {
		const fetchStatuses = async () => {
			try {
				const res = await fetch("/auth/me", {credentials: "include"});
				const json = await res.json();
				if (json.success) {
					// Identity parsing
					let identity = json.data?.identity;
					if (typeof identity === "string") {
						try {
							identity = JSON.parse(identity);
						} catch {
							identity = {};
						}
					}
					// Approved array
					let approved: string[] = [];
					const rawApproved = identity?.approved;
					if (Array.isArray(rawApproved)) {
						approved = rawApproved as string[];
					} else if (typeof rawApproved === "string" && rawApproved.trim() !== "" && rawApproved !== "undefined") {
						try {
							approved = JSON.parse(rawApproved);
						} catch {
							// ignore
						}
					}

					// Set verification completion
					const newVerifications = baseVerifications.map((v) => {
						v.complete = approved.includes(v.type);
						return v;
					});
					setVerifications(newVerifications);

					// Try to pick up location if API returns it
					const loc = json.data?.location;
					if (loc && (loc.ipAddress || loc.city || loc.state || loc.country)) {
						setLocation((prev) => ({
							...prev,
							ipAddress: loc.ipAddress ?? prev?.ipAddress,
							city: loc.city ?? prev?.city,
							state: loc.state ?? prev?.state,
							country: loc.country ?? prev?.country,
						}));
					}
				} else {
					setVerifications(baseVerifications);
				}
			} catch {
				setVerifications(baseVerifications);
			}
		};
		fetchStatuses();
	}, []);

	//||------------------------------------------------------------------------------------------------||
	//|| Fetch: Location + Zones (best-effort)
	//||------------------------------------------------------------------------------------------------||

	useEffect(() => {
		const fetchLocationAndZones = async () => {
			// Try a public location endpoint in your API (adjust as needed)
			let ip: string | undefined;
			let city: string | undefined;
			let state: string | undefined;
			let country: string | undefined;

			try {
				const r = await fetch("/public/location", {credentials: "include"});
				if (r.ok) {
					const j = await r.json();
					ip = j?.data?.ipAddress ?? j?.ipAddress;
					city = j?.data?.city ?? j?.city;
					state = j?.data?.state ?? j?.state;
					country = j?.data?.country ?? j?.country;
				}
			} catch {
				// ignore; we'll still render UI with whatever we have
			}

			// Fetch zones so we can infer allowed verification types by state/country
			let zones: ZoneRecord[] = [];
			try {
				let zr = await fetch("/public/zones", {credentials: "include"});
				if (!zr.ok) {
					// fallback route name if your handler is mounted differently
					zr = await fetch("/zones", {credentials: "include"});
				}
				if (zr.ok) {
					const zjson = await zr.json();
					if (zjson?.success && Array.isArray(zjson.data)) {
						zones = zjson.data.map((z: any) => ({
							id: z.id ?? z.ID ?? z.id_zone ?? 0,
							state: z.state ?? z.zone_state ?? null,
							country: z.country ?? z.zone_country ?? null,
							requirements: z.requirements ?? z.zone_requirements ?? null,
						})) as ZoneRecord[];
					}
				}
			} catch {
				// ignore; allowed list will be generic if we can't match a zone
			}

			// Use the freshest location values (effect may merge with /auth/me results)
			setLocation((prev) => {
				const merged: LocationInfo = {
					ipAddress: ip ?? prev?.ipAddress,
					city: city ?? prev?.city,
					state: state ?? prev?.state,
					country: country ?? prev?.country,
				};

				// Derive allowed from zones (match on state first, then country)
				const match =
					zones.find((z) => merged.state && z.state && z.state.toUpperCase() === merged.state.toUpperCase()) ||
					zones.find((z) => merged.country && z.country && z.country.toUpperCase() === merged.country.toUpperCase());

				if (match?.requirements) {
					const req = String(match.requirements).trim();
					const codes = req
						.split(",")
						.map((s) => s.trim())
						.filter(Boolean);
					merged.allowed = Array.from(new Set(codes));
				} else {
					// Fallback: allow everything by default if no zone match
					merged.allowed = baseVerifications.map((v) => v.type);
				}

				return merged;
			});
		};

		fetchLocationAndZones();
		// eslint-disable-next-line react-hooks/exhaustive-deps
	}, []);

	//||------------------------------------------------------------------------------------------------||
	//|| UI: Location Card
	//||------------------------------------------------------------------------------------------------||

	const LocationCard: React.FC<{loc: LocationInfo}> = ({loc}) => {
		const allowed = loc.allowed ?? baseVerifications.map((v) => v.type);

		return (
			<div className="col-span-1 lg:col-span-3">
				<div className="bg-base-100 rounded-2xl shadow-xl p-6 border border-base-200">
					<div className="flex items-center gap-3 mb-4">
						<div className="text-primary">
							<MapPin className="w-6 h-6" />
						</div>
						<h2 className="text-xl font-bold">Your Location</h2>
					</div>

					<div className="grid grid-cols-1 md:grid-cols-2 gap-6">
						{/* Left: Location + Allowed */}
						<div className="space-y-4">
							<div className="grid grid-cols-2 gap-3">
								<InfoPill label="IP Address" value={loc.ipAddress ?? "—"} />
								<InfoPill label="City" value={loc.city ?? "—"} />
								<InfoPill label="State" value={loc.state ?? "—"} />
								<InfoPill label="Country" value={loc.country ?? "—"} />
							</div>

							<div className="mt-4">
								<div className="text-sm font-semibold text-base-content/70 mb-2">Verification Types Allowed</div>
								<div className="flex flex-wrap gap-2">
									{allowed.map((c) => (
										<span key={c} className="badge badge-primary badge-outline px-3 py-3 text-sm rounded-lg">
											{codeToLabel(c)}
										</span>
									))}
								</div>
							</div>
						</div>

						{/* Right: VPN Ad / Notice */}
						<div className="mt-5relative overflow-hidden rounded-xl border border-base-200 bg-gradient-to-br from-base-200 to-base-100 p-5">
							<div className="flex items-start gap-4">
								<div className="shrink-0 text-primary">
									<Shield className="w-10 h-10" />
								</div>
								<div className="flex-1">
									<h3 className="text-lg font-bold leading-tight">Using a VPN?</h3>
									<p className="text-sm text-base-content/70 mt-1">
										VPNs and proxies can change your detected location and may limit which verification options are
										available in your region.
									</p>
									<div className="flex flex-wrap gap-2 mt-4">
										<button className="btn btn-primary btn-sm">Recommended VPNs</button>
										<button className="btn btn-ghost btn-sm">I’m not on a VPN</button>
									</div>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		);
	};

	//||------------------------------------------------------------------------------------------------||
	//|| UI: Info Pill
	//||------------------------------------------------------------------------------------------------||

	const InfoPill: React.FC<{label: string; value: string}> = ({label, value}) => (
		<div className="w-full">
			<div className="text-xs font-semibold uppercase tracking-wide text-base-content/60 mb-1">{label}</div>
			<div className="bg-base-200 rounded-lg px-3 py-2 text-sm font-medium border border-base-300">{value}</div>
		</div>
	);

	//||------------------------------------------------------------------------------------------------||
	//|| Default
	//||------------------------------------------------------------------------------------------------||

	return (
		<MembersLayout>
			<div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
				<LocationCard loc={location ?? {allowed: baseVerifications.map((v) => v.type)}} />
			</div>

			<div className="mt-5 grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
				{verifications.map((v, idx) => (
					<div key={idx} className="flex flex-col justify-between p-4 bg-base-100 rounded-lg shadow border border-base-200">
						{/* Top row: icon + text */}
						<div className="flex items-start gap-3">
                                          { v.complete ? (
							      <div className="p-2 rounded-full bg-success flex-grow-5">{React.cloneElement(v.icon as React.ReactElement, {size: 28})}</div>
                                          ) : (
                                                <div className="p-2 rounded-full bg-white/20 flex-shrink-5ds">{React.cloneElement(v.icon as React.ReactElement, {size: 28})}</div>
                                          )}

							<div className="flex flex-col">
								<h3 className="text-base font-semibold">{v.label}</h3>
								<p className="text-sm text-base-content/70">{v.blurb}</p>
							</div>
						</div>

						{/* Bottom row: button */}
						<div className="flex justify-end mt-3">
							{v.complete ? (
								<button className="btn btn-success btn-sm">
									<CheckCircle className="w-4 h-4 mr-1" /> Verified
								</button>
							) : (
								<button className="btn btn-primary btn-sm" onClick={() => navigate(`/verification/init/?type=${v.type}`)}>
									Complete <ArrowRight className="w-4 h-4 ml-1" />
								</button>
							)}
						</div>
					</div>
				))}
			</div>
		</MembersLayout>
	);
}
