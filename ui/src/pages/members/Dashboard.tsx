//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, {useEffect, useState} from "react";
import {useNavigate} from "react-router-dom";

//||------------------------------------------------------------------------------------------------||
//|| Lucide
//||------------------------------------------------------------------------------------------------||

import {CheckCircle, BadgeInfo, ArrowRight, CreditCard, MapPin, Mail, Phone, IdCard, Smile, Image as ImageIcon, Shield} from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

import {VerificationStatus}   from "../../interfaces/verification/status";
import {Identity}             from "../../interfaces/identity/identity";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import MembersLayout from "../../layouts/MembersLayout";

//||------------------------------------------------------------------------------------------------||
//|| API Verification Interface
//||------------------------------------------------------------------------------------------------||

interface ApiVerification {
	vType       : string;
	vStatus     : string;
}

//||------------------------------------------------------------------------------------------------||
//|| Location Interfaces
//||------------------------------------------------------------------------------------------------||

interface LocationInfo {
	ipAddress?        : string;
	city?             : string;
	region?           : string;
	country?          : string;
	types?            : string[];
      minAge?           : number;
}

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function Dashboard() {
	//||------------------------------------------------------------------------------------------------||
	//|| Const
	//||------------------------------------------------------------------------------------------------||

	const [verifications, setVerifications] = useState<VerificationStatus[]>([]);
	const [location, setLocation] = useState<LocationInfo | null>({
            minAge: 18,
      });
      const [identity, setIdentity] = useState<Identity>({
            verified : false, 
            verifiedAge: 0,
      });
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
                                    setIdentity(identity);
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
                  const res = await fetch("/user/location", {credentials: "include"})
                  try {
                        const json = await res.json();
                        if (json.success && json.data) {
                              const loc = json.data as LocationInfo;
                              json.data.types = json.data.types ?? json.data.types.split(',').map((s: string) => s.trim());
                              setLocation(json.data);
                        }
                  } catch {
                        // ignore
                  }
            };

		fetchLocationAndZones();

	}, []);

      //||------------------------------------------------------------------------------------------------||
      //|| Allowed: build a set from loc.types (CSV or array), case-insensitive
      //||------------------------------------------------------------------------------------------------||

      const allowedSet = React.useMemo(() => {
            const raw = (location?.types ?? "");
            const parts = Array.isArray(raw) ? raw : String(raw).split(",");
            return new Set(parts.map(s => s.trim().toUpperCase()).filter(Boolean));
      }, [location?.types]);      

	//||------------------------------------------------------------------------------------------------||
	//|| UI: Location Card
	//||------------------------------------------------------------------------------------------------||

	const LocationCard: React.FC<{loc: LocationInfo}> = ({loc}) => {

		return (
			<div className="col-span-1 lg:col-span-3">
				<div className="bg-base-100 rounded-2xl shadow-xl p-6 border border-base-200">
					<div className="grid grid-cols-1 md:grid-cols-2 gap-6">
						{/* Left: Location + Allowed */}
						<div className="space-y-4">
							{/* Heading inside left column */}
							<div className="flex items-center gap-3 mb-4 border-b border-gray-600 pb-2">
								<div className="text-primary">
									<MapPin className="w-6 h-6" />
								</div>
								<h2 className="text-xl font-bold">Your Location </h2>
                                                <div className="ml-auto text-center p-2 rounded-lg bg-black/20 shadow-2xl border border-gray-600">
                                                      <span className="block text-xs text-gray-400">Min. Age</span>
                                                      <span className="block font-bold text-white">{loc.minAge ?? "—"}</span>
                                                </div>
							</div>

							<div className="grid grid-cols-2 gap-3">
								<InfoPill label="IP Address" value={loc.ipAddress ?? "—"} />
								<InfoPill label="City" value={loc.city ?? "—"} />
								<InfoPill label="State" value={loc.region ?? "—"} />
								<InfoPill label="Country" value={loc.country ?? "—"} />
							</div>

							<div className="mt-4">
								<div className="text-sm font-semibold text-base-content/70 mb-2">Verification Types Allowed</div>

								<div className="flex flex-wrap gap-2">
									{baseVerifications.map((v) => {
										const isAllowed = allowedSet.has(v.type.toUpperCase());
										return (
											<span
												key={v.type}
												className={`badge px-3 py-3 text-xs rounded-lg ${
													isAllowed ? "badge-primary bg-orange-400" : "badge-primary badge-outline"
												}`}>
												{codeToLabel(v.type)}
											</span>
										);
									})}
								</div>
							</div>
						</div>

						{/* Right: VPN Ad / Notice */}
						<div className="relative overflow-hidden rounded-xl border border-base-200 bg-gradient-to-br from-base-200 to-base-100 p-4">
							<div className="flex items-start gap-4">
								<div className="shrink-0 text-primary">
									<Shield className="w-24 h-24 outline-blue fill-blue-400" />
								</div>
								<div className="flex-1">
									<h3 className="text-lg font-bold leading-tight">Using a VPN?</h3>
									<p className="text-sm text-base-content/70 mt-1">
										VPNs and proxies can change your detected location and may limit which verification options are
										available in your region.
									</p>
                                                      <h5 className="mt-2 font-bold">Why choosing a good VPN matters.</h5>
                                                      <p className="text-sm text-base-content/70 mt-1">
                                                            Shared VPN IPs often have poor reputation, which can lead to CAPTCHAs, extra identity prompts, or limited methods (e.g., card or ID only). Using a location that matches your real region keeps things fast and consistent.
                                                      </p>
									<div className="flex flex-wrap gap-2 mt-4">
										<button
											className="btn btn-primary btn-md bg-blue-400"
											onClick={() => {
												navigate("/members/vpns");
											}}>
											Recommended VPNs
										</button>
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
			<div className="bg-base-200 rounded-lg px-3 py-2 text-sm font-medium border border-base-300 text-orange-400">{value}</div>
		</div>
	);

	//||------------------------------------------------------------------------------------------------||
	//|| Required Banner
	//||------------------------------------------------------------------------------------------------||

	const VerifiedBanner: React.FC<{value: string}> = ({value}) => (
		<div className="w-full">
			<div className="bg-base-200 rounded-lg px-3 py-2 text-sm font-medium border border-base-300">Banner</div>
		</div>
	);

      const UnverifiedBanner: React.FC<{methods?: string}> = ({methods}) => {
            let minimumMethod = "IDEN";
            if (location?.types?.includes("FACE")) minimumMethod = "FACE";
            if (location?.types?.includes("CRCD")) minimumMethod = "CRCD";            
            return (
                  <div className="w-full">
                        <div className="bg-base-200 rounded-lg px-3 py-2 text-sm font-medium border border-base-300 mb-3">
                              <div className="flex items-center gap-3">
                                    <BadgeInfo className="w-5 h-5 text-secondary shrink-0" />
                                    <span className="text-secondary font-semibold">Action Required</span>
                                    <span className="text-base-content/80">Your age is not verified.</span>

                                    <span className="ml-auto text-xs text-base-content/60 whitespace-nowrap overflow-hidden text-ellipsis">
                                          <button className="btn bg-green-600" onClick={ () => {
                                                navigate(`/verification/init?type=${minimumMethod}`);
                                                } }
                                          >
                                                <CheckCircle className="w-4 h-4 mr-1" />
                                                Meet Minimum Verification Level Now!
                                          </button>
                                    </span>
                             </div>
                        </div>
                  </div>
            );
      }


      //||------------------------------------------------------------------------------------------------||
	//|| Default
	//||------------------------------------------------------------------------------------------------||

	return (
		<MembersLayout>
                  <div>
                        { (identity.verified && identity.verifiedAge && location && location.minAge && identity.verifiedAge > location.minAge) ? (
                              <VerifiedBanner value={String(identity.verifiedAge ?? "")} />
                        ) : (
                              <UnverifiedBanner methods={identity?.approved ?? []} />
                        ) }
                  </div>

			<div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
				<LocationCard loc={location ?? {allowed: baseVerifications.map((v) => v.type)}} />
			</div>

			<div className="mt-5 grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
				{verifications.map((v, idx) => (
					<div key={idx} className="flex flex-col justify-between p-4 bg-base-100 rounded-lg shadow border border-base-200">
						{/* Top row: icon + text */}
						<div className="flex items-start gap-3">
							{v.complete ? (
								<div className="p-2 rounded-full bg-success flex-grow-5">
									{React.cloneElement(v.icon as React.ReactElement, {size: 28})}
								</div>
							) : (
								<div className="p-2 rounded-full bg-white/20 flex-shrink-5ds">
									{React.cloneElement(v.icon as React.ReactElement, {size: 28})}
								</div>
							)}

							<div className="flex flex-col">
								<h3 className="text-base font-semibold">{v.label}</h3>
								<p className="text-sm text-base-content/70">{v.blurb}</p>
							</div>
						</div>

                                    { identity.verified && identity.verifiedAge && identity.verifiedAge > 18 && (
                                          <div className="mt-4">
                                                <div className="text-sm font-semibold text-base-content/70 mb-2">Age Verified</div>
                                                <div className="text-lg font-bold">{identity.verifiedAge}+</div>
                                          </div>
                                    ) }

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
