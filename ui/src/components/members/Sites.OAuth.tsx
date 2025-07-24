//||------------------------------------------------------------------------------------------------||
//|| OAuthSettingsSection
//|| Component for configuring OAuth redirect URL, private key, and verification permissions
//||------------------------------------------------------------------------------------------------||

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, {useEffect, useState}       from "react";

//||------------------------------------------------------------------------------------------------||
//|| Intefaces
//||------------------------------------------------------------------------------------------------||

import { Site }                           from "../../interfaces/model.sites";

//||------------------------------------------------------------------------------------------------||
//|| Hooks
//||------------------------------------------------------------------------------------------------||

import {useEffectOnce}                    from "../../hooks/useEffectOnce";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

interface VerificationType {
	id                : string;
	code              : string;
	description       : string;
	level             : number;
}

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface OAuthSettingsSectionProps {
	data              : Site;
	updateField       : (field: string, value: any) => void;
}

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function OAuthSettingsSection({data, updateField}: OAuthSettingsSectionProps) {
      //||------------------------------------------------------------------------------------------------||
      //|| Var
      //||------------------------------------------------------------------------------------------------||
      const [types, setTypes]             = useState<VerificationType[]>([]);
	const [loading, setLoading]         = useState(true);
	const [selected, setSelected]       = useState<Set<string>>(new Set(data.permissions ? data.permissions.split(",") : []));
      //||------------------------------------------------------------------------------------------------||
      //|| Load Verification Types
      //||------------------------------------------------------------------------------------------------||
	useEffectOnce(() => {
		(async () => {
			try {
				const res = await fetch("/v1/api/sites/vtypes", {
					method: "GET",
					credentials: "include",
				});
				const json = await res.json();
                        console.log(json);
				if (Array.isArray(json.data.verification_types)) {
					setTypes(json.data.verification_types);
				}
			} catch (err) {
				console.error(err);
			} finally {
				setLoading(false);
			}
		})();
	});
      //||------------------------------------------------------------------------------------------------||
      //|| Permissions
      //||------------------------------------------------------------------------------------------------||
	useEffect(() => {
		setSelected(new Set(data.permissions?.split(",") || []));
	}, [data.permissions]);

      //||------------------------------------------------------------------------------------------------||
      //|| On Chnage
      //||------------------------------------------------------------------------------------------------||
	const onChangeField = (field: "redirect" | "private") => (e: React.ChangeEvent<HTMLInputElement>) => {
		updateField(field, e.target.value);
	};
      //||------------------------------------------------------------------------------------------------||
      //|| Handle Permissions
      //||------------------------------------------------------------------------------------------------||
	const toggleCode = (code: string) => {
		const next = new Set(selected);
		if (next.has(code)) next.delete(code);
		else next.add(code);
		setSelected(next);
		const arr = Array.from(next).sort();
		updateField("permissions", arr.join(","));
	};
      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||
	return (
		<div className="w-full bg-base-100 shadow-lg rounded-lg p-8 mb-5">
			<h2 className="text-2xl font-bold mb-6">OAuth Settings</h2>

			{/* Redirect URL */}
			<div className="mb-4">
				<label className="block text-sm font-medium mb-1">Redirect URL</label>
				<input type="text" className="input input-bordered w-full" value={data.redirect} onChange={onChangeField("redirect")} />
			</div>

			{/* Client ID */}
			<div className="mb-6">
				<label className="block text-sm font-medium mb-1">Client ID</label>
				<input type="text" className="input input-bordered w-full" value={data.public} />
			</div>

			{/* Private Key */}
			<div className="mb-6">
				<label className="block text-sm font-medium mb-1">Private Key</label>
				<input type="text" className="input input-bordered w-full" value={data.private} />
			</div>

			{/* Permissions Table */}
			<h3 className="text-xl font-semibold mb-4">Verification Permissions</h3>
			{loading ? (
				<div className="text-sm opacity-60">Loading…</div>
			) : (
				<div className="overflow-x-auto">
					<table className="table table-auto w-full">
						<thead>
							<tr className="border-b-[1px] border-gray-500">
								<th className="p-2 w-12" className="text-center">Enable</th>
								<th className="p-2">Code</th>
								<th className="p-2">Description</th>
                                                <th className="p-2 text-left">Require Approval</th>
							</tr>
						</thead>
						<tbody>
							{types.map((t) => (
								<tr key={t.id} className="border-t-[1px] border-gray-600">
									<td className="p-2 w-12 text-center">
										<input type="checkbox" checked={selected.has(t.code)} onChange={() => toggleCode(t.code)} />
									</td>
									<td className="p-2"><span className="bg-gray-800 border-dashed border-gray-200 rounded-lg font-mono code p-2">{t.code}</span></td>
									<td className="p-2 text-xs">{t.description}</td>
                                                      <td className="p-2">{t.level > 0 ? (<span className="flex justify-center w-12 text-center bg-black text-yellow-500">Yes</span>) : (<span className="flex justify-center w-12 text-center bg-gray-800 text-gray">No</span>)}</td>
								</tr>
							))}
						</tbody>
					</table>
				</div>
			)}
		</div>
	);
}
