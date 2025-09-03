import React, {useEffect, useMemo, useState} from "react";
import {useNavigate} from "react-router-dom";
import {Shield, Globe, ExternalLink, Lock, Zap, Star} from "lucide-react";
import MembersLayout from "../../layouts/MembersLayout";
import Stars from "../../components/dynamic/Stars";
//||------------------------------------------------------------------------------------------------||
//|| Types: Backend VPN Interface
//||------------------------------------------------------------------------------------------------||

export interface VPN {
	id_vpn: number;
	vpn_name: string | null;
	vpn_url: string | null;
	vpn_blurb: string | null;
	vpn_highlights: string | null;
	vpn_region: string | null;
	vpn_price: string | null;
	vpn_rating: number | null;
}

//||------------------------------------------------------------------------------------------------||
//|| Page
//||------------------------------------------------------------------------------------------------||

export default function VPNs() {
	const [query, setQuery] = useState("");
	const [sort, setSort] = useState<"top" | "price" | "alpha">("top");
	const [vpns, setVpns] = useState<VPN[] | null>(null);
	const [loading, setLoading] = useState(true);
	const [userRatings, setUserRatings] = useState<{[vpnId: number]: number}>({});
	const [ratingSubmitting, setRatingSubmitting] = useState<number | null>(null);

	useEffect(() => {
		setLoading(true);
		fetch("/v1/api/vpns")
			.then((res) => {
				if (!res.ok) throw new Error("Failed to fetch VPNs");
				return res.json();
			})
			.then((data) => {
				setVpns(Array.isArray(data.data) ? data.data : []);
			})
			.catch(() => setVpns([]))
			.finally(() => setLoading(false));
	}, []);

	// Send rating
	function handleRate(vpn: VPN, stars: number) {
		setRatingSubmitting(vpn.id_vpn);
		fetch(`/user/vpns/rate?vpn=${vpn.id_vpn}&rating=${stars}`, {
			method: "GET",
			credentials: "include",
		})
			.then((res) => res.json())
			.then((result) => {
				setUserRatings((old) => ({
					...old,
					[vpn.id_vpn]: stars,
				}));
				if (result && typeof result.vpn_rating_percent === "number" && vpns) {
					setVpns( (vpns) => vpns?.map((v) => v.id_vpn === vpn.id_vpn ? {...v,vpn_rating: result.vpn_rating_percent,}: v ) ?? [] );
				}
			})
			.finally(() => setRatingSubmitting(null));
	}

	const filtered = useMemo(() => {
		if (!vpns) return [];
		const q = query.trim().toLowerCase();
		let list = vpns.filter(
			(v) =>
				(v.vpn_name ?? "").toLowerCase().includes(q) ||
				(v.vpn_blurb ?? "").toLowerCase().includes(q) ||
				(v.vpn_highlights ?? "").toLowerCase().includes(q)
		);
		if (sort === "alpha") {
			list = list.slice().sort((a, b) => (a.vpn_name ?? "").localeCompare(b.vpn_name ?? ""));
		} else if (sort === "price") {
			list = list.slice().sort((a, b) => (a.vpn_price ?? "").length - (b.vpn_price ?? "").length);
		} else {
			list = list.slice().sort((a, b) => (b.vpn_rating ?? 0) - (a.vpn_rating ?? 0) || (a.vpn_name ?? "").localeCompare(b.vpn_name ?? ""));
		}
		return list;
	}, [vpns, query, sort]);

	return (
		<MembersLayout title="VPN Providers" icon={Shield}>
			<div className="mb-6 flex flex-col sm:flex-row gap-3 sm:items-center sm:justify-between">
				<div className="flex items-center gap-2">
					<h2 className="text-xl font-bold text-gray-300">Our Top Virtual Private Network Picks</h2>
				</div>
				<div className="flex items-center gap-2">
					<input
						type="text"
						value={query}
						onChange={(e) => setQuery(e.target.value)}
						placeholder="Search VPNs..."
						className="input input-bordered input-md w-56"
					/>
					<select value={sort} onChange={(e) => setSort(e.target.value as any)} className="select select-bordered select-md">
						<option value="top">Top Rated</option>
						<option value="alpha">A → Z</option>
						<option value="price">Budget-ish</option>
					</select>
				</div>
			</div>
			{loading ? (
				<div className="py-12 text-center text-base-content/60 text-sm">Loading VPNs…</div>
			) : (
				<div className="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-5">
					{filtered.length === 0 ? (
						<div className="col-span-full py-20 text-center text-base-content/50 text-lg">No VPNs found.</div>
					) : (
						filtered.map((vpn) => (
							<div
								key={vpn.id_vpn}
								className="group relative overflow-hidden rounded-2xl border border-base-300 bg-base-100 shadow hover:shadow-xl transition-shadow">
								<div className="absolute inset-x-0 -top-10 h-24 bg-gradient-to-r from-primary/30 via-secondary/30 to-accent/30 blur-2xl opacity-50 group-hover:opacity-80 transition-opacity" />
								<div className="relative p-5 flex flex-col gap-4">
									<div className="flex items-center gap-4">
										<div className="w-12 h-12 rounded-xl bg-base-200 border border-base-300 flex items-center justify-center overflow-hidden">
											<Zap className="w-8 h-8 text-primary/60" />
										</div>
										<div className="flex-1">
											<h3 className="text-lg font-bold tracking-tight">{vpn.vpn_name ?? "Unnamed VPN"}</h3>
											<div className="flex items-center justify-between gap-3">
												<Stars
													n={userRatings[vpn.id_vpn] ?? Math.round((vpn.vpn_rating ?? 0) / 20)}
													avg={typeof vpn.vpn_rating === "number" ? vpn.vpn_rating / 20 : undefined}
													interactive
													userRating={userRatings[vpn.id_vpn]}
													disabled={ratingSubmitting === vpn.id_vpn}
													onRate={(val) => handleRate(vpn, val)}
												/>
											</div>
											<div className="text-xs text-base-content/60 flex items-center gap-2 mt-0.5">
												<Globe className="w-3.5 h-3.5" />
												<span>{vpn.vpn_region ?? "Unknown"}</span>
											</div>
										</div>
									</div>
									<p className="text-sm text-base-content/80 leading-snug">{vpn.vpn_blurb}</p>
									{vpn.vpn_highlights && (
										<div className="flex text-center gap-1">
											{vpn.vpn_highlights.split(",").map((h) => (
												<span
													key={h.trim()}
													className="px-2.5 py-1 rounded-lg border border-base-300 bg-base-200 text-xs font-medium">
													{h.trim()}
												</span>
											))}
										</div>
									)}
									<div className="mt-1 flex items-center justify-between">
										<span className="text-sm text-yellow-400 flex items-center gap-2">
                                                                        <Lock className="w-3.5 h-3.5" />
												<span className="font-bold">{vpn.vpn_price ?? "Unknown"}</span>
                                                            </span>
										<span className="btn btn-primary btn-sm">
											Visit <ExternalLink className="w-4 h-4 ml-1" />
										</span>
									</div>
								</div>
								<div className="pointer-events-none absolute -right-10 -bottom-10 w-40 h-40 rounded-full bg-primary/10 blur-2xl group-hover:bg-primary/20 transition-colors" />
							</div>
						))
					)}
				</div>
			)}
			<div className="mt-6 text-center text-xs text-base-content/60">
				Speed and availability can vary by region and ISP. For best results, try multiple protocols (e.g., WireGuard) and nearest locations.
			</div>
		</MembersLayout>
	);
}
