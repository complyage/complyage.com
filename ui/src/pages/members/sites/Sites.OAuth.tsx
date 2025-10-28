//||------------------------------------------------------------------------------------------------||
//|| OAuthSettingsSection
//|| Component for configuring OAuth redirect URL, private key, and verification permissions
//||------------------------------------------------------------------------------------------------||

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, {useEffect, useState, useRef}       from "react";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import LabelDescription                               from "../../../components/base/LabelDescription";

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
      const [mode, setMode]               = useState<"manual">("manual");
	const [selected, setSelected]       = useState<Set<string>>(new Set([]));
      const initializedRef                = useRef(false);
      const hasMounted                    = useRef(false);
      const lastScopesRef                 = useRef<string>("");      
      //||------------------------------------------------------------------------------------------------||
      //|| Load Verification Types
      //||------------------------------------------------------------------------------------------------||
      useEffect(() => {
            // prevent multiple initializations (strict mode safe)
            if (initializedRef.current) return;
            initializedRef.current = true;

            (async () => {
                  try {
                        const res = await fetch("/v1/api/sites/scopes", {
                              method      : "GET",
                              credentials : "include",
                        });
                        const json = await res.json();
                        console.log("JSON SCOPES", json);
                        setTypes(json.data || []);

                        //||------------------------------------------------------------------------------------------------||
                        //|| Initialize only if scopes are missing
                        //||------------------------------------------------------------------------------------------------||
                        if ((!data.scopes || data.scopes.length === 0) && Array.isArray(json.data)) {
                              const initialScopes = json.data.map((t: VerificationType) => ({
                                    code        : t.code,
                                    status      : "VERF",
                                    enabled     : false,
                              }));

                              console.log("[INIT] Setting default scopes locally (no parent update)");

                              // just set local default values without calling updateField
                              setSelected(new Set());
                              lastScopesRef.current = JSON.stringify(initialScopes);
                        }

                  } catch (err) {
                        console.error(err);
                  } finally {
                        setLoading(false);
                  }
            })();
      }, []);
      //||------------------------------------------------------------------------------------------------||
      //|| Sync Scopes With Selected
      //||------------------------------------------------------------------------------------------------||
      useEffect(() => {
            console.log("[EFFECT selected]", Array.from(selected));

            // Don't sync until data is loaded
            if (loading || types.length === 0) return;

            // Skip first render entirely
            if (!hasMounted.current) {
                  hasMounted.current = true;
                  return;
            }

            // Build updated scopes directly from types (not data.scopes)
            const updatedScopes = types.map(t => ({
                  code    : t.code,
                  status  : "VERF",
                  enabled : selected.has(t.code),
            }));

            const serialized = JSON.stringify(updatedScopes);
            if (serialized !== lastScopesRef.current) {
                  lastScopesRef.current = serialized;
                  console.log("[SYNC] updateField scopes (user interaction)");
                  updateField("scopes", updatedScopes);
            }
      }, [selected, loading, types]);

      //||------------------------------------------------------------------------------------------------||
      //|| On Change
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
            console.log("[TOGGLE]", code, "->", Array.from(next));
            setSelected(next);
      };
      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||
	return (
		<div className="w-full bg-base-100 shadow-lg rounded-lg p-8 mb-5">
			<h2 className="text-2xl font-bold mb-6">OAuth Settings</h2>

			{/* Redirect URL */}
			<div className="mb-4">
                        <LabelDescription
                              id="sitename"
                              label="Redirect URL"
                              description="After OAuth confirmation, users will be redirected here with a 1 time access code."
                        />                           
				<input type="text" className="input input-bordered w-full" value={data.redirect} onChange={onChangeField("redirect")} />
			</div>

			{/* Permissions Table */}
                  <LabelDescription
                        id="scope"
                        label="Requested Scope"
                        description="Which user data your application will request access to. (e.g. age verification, phone verification, etc.)"
                  />                           

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
		</div>
	);
}
