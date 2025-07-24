//||------------------------------------------------------------------------------------------------||
//|| ZonesEnforcementSection
//|| Component for enforcing zones based on selected mode
//||------------------------------------------------------------------------------------------------||

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useEffect, useState }              from "react";

//||------------------------------------------------------------------------------------------------||
//|| Hooks
//||------------------------------------------------------------------------------------------------||

import { useEffectOnce }                          from "../../hooks/useEffectOnce";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

import { Site }                                   from "../../interfaces/model.sites";
import { Zone }                                   from "../../interfaces/zones";
import { ZoneRequirementPlain }                   from "../../interfaces/zoneRequirements";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface ZonesEnforcementSectionProps {
      data                    : Site;
      updateField             : <K extends keyof Site>(field: K, value: Site[K]) => void;
}

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function ZonesEnforcementSection({ data, updateField }: ZonesEnforcementSectionProps) {

      //||------------------------------------------------------------------------------------------------||
      //|| State
      //||------------------------------------------------------------------------------------------------||

      const [mode, setMode]                               = useState<string>(data?.enforcement || "ALLZ");
      const [zones, setZones]                             = useState<Zone[]>([]);
      const [customZones, setCustomZones]                 = useState<Record<number, number>>({});
      const [loading, setLoading]                         = useState<boolean>(true);

      //||------------------------------------------------------------------------------------------------||
      //|| Zones
      //||------------------------------------------------------------------------------------------------||

      console.log("ZONES DATA", data.zones);

      //||------------------------------------------------------------------------------------------------||
      //|| Load Zones
      //||------------------------------------------------------------------------------------------------||

      useEffectOnce(() => {
            const fetchZones = async () => {
                  try {
                        const res  = await fetch("/v1/api/sites/zones", { method : "GET", credentials: "include" });
                        const json = await res.json();
                        console.log("ZONES JSON", json.data);
                        if (Array.isArray(json.data)) {
                              const newZones = json.data;
                              newZones.push({
                                    id          : "9999",
                                    state       : "All",
                                    country     : "All",
                                    requirements: "none"
                              } as Zone);
                              setZones(newZones);
                        }
                  } catch (err) {
                        console.error("Failed to load zones:", err);
                  } finally {
                        setLoading(false);
                  }
            };

            fetchZones();
      });

      //||------------------------------------------------------------------------------------------------||
      //|| Sync With Parent Data
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            setMode(data?.enforcement || "");
            if (data?.zones === undefined || data.zones === null) {
                  setCustomZones({});
                  return;
            }
            if (typeof data?.zones !== "object" && !Array.isArray(data?.zones)) {
                  try { 
                        setCustomZones(JSON.parse(data?.zones) as Record<number, number>);
                  } catch(e) { 
                        setCustomZones({});
                  }
            } else {
                  try {
                        setCustomZones(data?.zones as Record<number, number>);
                  } catch (e) {
                        setCustomZones({});
                  }
            }
      }, [data]);

      //||------------------------------------------kora------------------------------------------------------||
      //|| Handle Mode Change
      //||------------------------------------------------------------------------------------------------||

      const handleModeChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
            const newMode = e.target.value;
            setMode(newMode);
            updateField("enforcement", newMode);
      };

      //||------------------------------------------------------------------------------------------------||
      //|| UpdateZone
      //||------------------------------------------------------------------------------------------------||

      const updateZone = (zoneId: string, action: 'ENFORCE' | 'IGNORE') => {
            if (!data) return;
            const allZones = { ...(data.zones || {}) } as Record<number, number>;
            const num = parseInt(zoneId, 10);
            if (isNaN(num)) {
              console.error('Invalid zone ID:', zoneId);
              return;
            }
            allZones[num] = action === 'ENFORCE' ? 1 : 0;
            updateField('zones', allZones);
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Hide if no data
      //||------------------------------------------------------------------------------------------------||

      if (!data)  return null;

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      console.log("CUSTOM ZONES", customZones);

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      return (
            <div className="w-full bg-base-100 shadow-lg rounded-lg p-8 mb-5">
                  <h2 className="text-2xl font-bold mb-6">Zones Enforcement</h2>

                  <div className="flex flex-col sm:flex-row sm:items-center gap-4 mb-6">
                        <select
                              className="select select-bordered w-full sm:max-w-xs"
                              defaultValue={data.enforcement}
                              onChange={handleModeChange}
                        >
                              <option value="ALLZ">Force All traffic</option>
                              <option value="REGU">Enforce Regulated Zones Only</option>
                              <option value="CSTM">Custom</option>
                        </select>
                  </div>

                  {mode === "ALLZ" && (
                        <div className="mb-4 text-center bg-white/5 p-4">
                              ALL traffic will be enforced. 
                        </div>
                  )}

                  {(mode === "CSTM" || mode === "REGU") && (
                        loading ? (
                              <div className="text-sm opacity-60">Loading zones...</div>
                        ) : (
                              <div className="overflow-x-auto">
                                    <table className="table table-auto w-full">
                                          <thead >
                                                <tr className="border-b-[1px] border-gray-500">
                                                      <th>State</th>
                                                      <th>Country</th>
                                                      <th>Requirements</th>
                                                      <th>Action</th>
                                                </tr>
                                          </thead>
                                          <tbody>
                                                {zones.map((zone : Zone) => {
                                                      const requirementsList = !zone.requirements ? (<li>N/A</li>): zone.requirements.split(",").map((req, idx) => (
                                                            <li className="p-1 text-gray-300 text-xs" key={idx}>{ZoneRequirementPlain(req.trim())}</li>
                                                      ));
                                                      if (zone.id === "9999") {
                                                            return (
                                                                  <>
                                                                        <tr key={`unknown${zone.id}`} className="bg-black/50">
                                                                              <td className="align-center text-center text-yellow-400 text-xs p-5" colSpan="4">
                                                                                    Some IP Addresses may be un-resolvable to a certain state or country. How would you like to handle these?
                                                                              </td>
                                                                        </tr>                                                                  
                                                                        <tr key={`location${zone.id}`} className="border-t-0 border-transparent bg-black/50">
                                                                              <td className="align-top"><span className="text-yellow-200 text-xs bg-black p-1">Unknown</span></td>
                                                                              <td className="align-top"><span className="text-yellow-200 text-xs bg-black p-1">Location</span></td>
                                                                              <td className="align-top "><ul><li className="p-1 text-gray-300 text-xs">3rd Party Validation</li></ul></td>
                                                                              <td className="align-top">
                                                                                    { (mode === "REGU") ? (<span className="font-bold bg-black text-yellow-500 p-2 text-xs">Enforced</span>) : (
                                                                                    <select
                                                                                          className="select select-sm select-bordered"
                                                                                          defaultValue={ ( String(customZones[zone.id]) === "0" ) ? "IGNORE" : "ENFORCE" } 
                                                                                          onChange={(e) => { updateZone(zone.id, e.target.value as 'ENFORCE' | 'IGNORE') } }
                                                                                    >
                                                                                          <option value="ENFORCE">✅ ENFORCE</option>
                                                                                          <option value="IGNORE">❌ IGNORE </option>
                                                                                    </select>
                                                                              )}
                                                                              </td>
                                                                        </tr>
                                                                  </>
                                                            );
                                                      }                                                                 
                                                      return (<tr key={zone.id} className="border-t-[1px] border-gray-600">
                                                            <td className="align-top">{zone.state === null ? (<span className="text-yellow-200 text-xs bg-black p-1">{`Countrywide`}</span>) : zone.state}</td>
                                                            <td className="align-top">{zone.country}</td>
                                                            <td className="align-top"><ul>{requirementsList}</ul></td>
                                                            <td className="align-top">
                                                                  { (mode === "REGU") ? (<span className="font-bold bg-black text-yellow-500 p-2 text-xs">Enforced</span>) : (
                                                                  <select
                                                                        className="select select-sm select-bordered"
                                                                        defaultValue={ ( String(customZones[zone.id]) === "0" ) ? "IGNORE" : "ENFORCE" } 
                                                                        onChange={(e) => { updateZone(zone.id, e.target.value as 'ENFORCE' | 'IGNORE') } }
                                                                  >
                                                                        <option value="ENFORCE">✅ ENFORCE</option>
                                                                        <option value="IGNORE">❌ IGNORE </option>
                                                                  </select>
                                                            )}
                                                            </td>
                                                      </tr>
                                                )}
                                          )}
                                          </tbody>
                                    </table>
                              </div>
                        )
                  )}
            </div>
      );
}