//||------------------------------------------------------------------------------------------------||
//|| IntegrationSection
//|| Component for managing API integration keys and settings
//||------------------------------------------------------------------------------------------------||

import React, { useState, useEffect } from "react";

//||------------------------------------------------------------------------------------------------||
//|| Interface
//||------------------------------------------------------------------------------------------------||

import { ModelSite }                     from "../../../interfaces/models/model.sites";
import { integrationCode }               from "../../../utils/integration.code";
import { X } from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface IntegrationSectionProps {
      data           : ModelSite;
      updateField    : (field: string, value: string) => void;
}

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function IntegrationSection({ data, updateField }: IntegrationSectionProps) {

      //||------------------------------------------------------------------------------------------------||
      //|| State
      //||------------------------------------------------------------------------------------------------||

      const [currentDomain, setCurrentDomain]           = useState("");
      const [domainList, setDomainList]                 = useState<string[]>([]);
      const [clientId, setClientId]                     = useState(data?.clientId || "");
      const [privateKey, setPrivateKey]                 = useState(data?.private || "");
      const [publicKey, setPublicKey]                   = useState(data?.public || "");
      const [domains, setDomains]                       = useState(data?.domains || "");
      const [codePreview, setCodePreview]               = useState(integrationCode(data?.public || "123"));

      //||------------------------------------------------------------------------------------------------||
      //|| Sync With Parent Data
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            setPrivateKey(data?.private || "");
            setPublicKey(data?.public || "");
            setDomains(data?.domains || "");
            setCodePreview(integrationCode(data?.public || "123"));
            if (data?.domains) {
                  setDomainList(data.domains.split(",").map(d => d.trim()).filter(Boolean));
            } else {
                  setDomainList([]);
            }            
      }, [data]);

      //||------------------------------------------------------------------------------------------------||
      //|| Data
      //||------------------------------------------------------------------------------------------------||

      const addDomain = () => {
            if (currentDomain && !domainList.includes(currentDomain)) {
                  const updatedDomains = [...domainList, currentDomain];
                  setDomainList(updatedDomains);
                  setDomains(updatedDomains.join(","));
                  updateField("domains", updatedDomains.join(","));
                  setCurrentDomain("");
            }
      };

      const removeDomain = (domain: string) => {
            const updatedDomains = domainList.filter(d => d !== domain);
            setDomainList(updatedDomains);
            setDomains(updatedDomains.join(","));
            updateField("domains", updatedDomains.join(","));
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Data
      //||------------------------------------------------------------------------------------------------||

      if (!data) return null;

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      return (
            <div className="w-full bg-base-100 shadow-lg rounded-lg p-8 mb-4">
                  <h2 className="text-2xl font-bold mb-6">Integration Settings</h2>

                  <div className="grid grid-cols-1 gap-6">
                        <div>
                              <label htmlFor="clientId" className="label pb-1">ClientId</label>
                              <input
                                    id="clientId"
                                    type="text"
                                    placeholder="Public Key"
                                    className="input input-bordered w-full text-white"
                                    readOnly
                                    value={clientId}
                              />
                        </div>
                        <div>
                              <label htmlFor="allowed" className="label pb-1">Allowed Domains</label>
                              <div className="flex items-center">
                                    <input 
                                    id="allowed"
                                    type="text" 
                                    className="input input-bordered text-white flex-1"
                                    onChange={(e) => setCurrentDomain(e.target.value)} 
                                    onKeyDown={(e) => { if (e.key === 'Enter') { e.preventDefault(); addDomain(); }}}
                                    value={currentDomain} 
                                    placeholder="Add domain and press Enter"
                                    />
                                    <button
                                    onClick={addDomain}
                                    className="btn btn-sm btn-primary ml-2"
                                    >
                                    Add
                                    </button>
                              </div>
                              <div className="mt-2">
                                    { domainList.length === 0 && (<p className="text-sm text-gray-500 text-center p-4 bg-black/20 border-lg mb-5">No domains added. Please add allowed domains.</p>) }
                                    {domainList.map((domain, index) => (
                                    <div 
                                          key={index}
                                          className="flex items-center justify-between w-fit bg-gray-700 text-white text-sm px-3 py-1 rounded-md mr-2 mb-2 shadow-sm"
                                    >
                                          <span>{domain}</span>
                                          <button
                                                onClick={() => removeDomain(domain)}
                                                className="ml-2 flex items-center justify-center text-white hover:text-red-600 cursor-pointer"
                                          >
                                          <X size={24} />
                                          </button>
                                    </div>
                                    ))}
                              </div>
                        </div>
                        <div>
                              <label htmlFor="testMode" className="label pb-1">Test Mode</label>
                              <select
                                    id="testMode"
                                    className="select select-bordered w-full"
                                    defaultValue={domains}
                                    onChange={(e) => updateField("testMode", e.target.value)}
                              >
                                    <option value="1">Development Mode - Test</option>
                                    <option value="0">Production Mode - Live</option>
                              </select>
                        </div>                        
                        <div>
                              <label htmlFor="integrationCode" className="label pb-1">Integration Code</label>
                              <textarea
                                    id="integrationCode"
                                    readOnly
                                    className="w-full font-mono text-xs bg-base-200 p-3 text-gray-200"
                                    rows={10}
                                    value={ codePreview }
                              ></textarea>
                        </div>
                  </div>
            </div>
      );
}
