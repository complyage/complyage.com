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

import { ModelSite }                      from "../../../interfaces/models/model.sites";

//||------------------------------------------------------------------------------------------------||
//|| Hooks
//||------------------------------------------------------------------------------------------------||

import {useEffectOnce}                    from "../../../hooks/useEffectOnce";

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
	data              : ModelSite;
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
      const [mode, setMode]               = useState<"auto" | "manual">("auto");
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
      //|| Change Mode
      //||------------------------------------------------------------------------------------------------||
      const changeMode = (newMode: "auto" | "manual") => {
            setMode(newMode);
            updateField("permission_mode", newMode);
      }
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
			<div className="mb-4 flex gap-4 mx-auto max-w-lg justify-center p-3">
				<button
					type="button"
					className={`btn transition ${
						mode === "auto" ? "btn-secondary" : "btn-primary text-gray-400"
					}`}
					onClick={() => changeMode("auto")}>
					Automatic
				</button>

				<button
					type="button"
					className={`btn transition ${
						mode === "manual" ? "btn-secondary" : "btn-primary text-gray-400"
					}`}
					onClick={() => changeMode("manual")}>
					Manual
				</button>
			</div>

			{mode === "auto" ? (
				<div className="mb-4 bg-black/40 p-4 py-10 rounded-xl text-center">
					<p className="text-sm text-yellow-400">Permissions are automatically managed based on the end-users's location and applicable laws.</p>
				</div>
			) : (
				<div className="overflow-x-auto">
					<table className="table table-auto w-full">
						<thead>
							<tr className="border-b-[1px] border-gray-500">
								<th className="p-2 w-12 text-center">Enable</th>
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
									<td className="p-2">
										<span className="bg-gray-800 border-dashed border-gray-200 rounded-lg font-mono code p-2">
											{t.code}
										</span>
									</td>
									<td className="p-2 text-md">{t.description}</td>
									<td className="p-2">
										{t.level > 0 ? (
											<span className="flex rounded-lg justify-center w-12 text-center bg-black text-yellow-500 py-1">Yes</span>
										) : (
											<span className="flex rounded-lg justify-center w-12 text-center bg-gray-800 text-gray-400 py-1">No</span>
										)}
									</td>
								</tr>
							))}
						</tbody>
					</table>
				</div>
			)}
		</div>
	);
}
