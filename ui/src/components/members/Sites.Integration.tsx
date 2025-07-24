//||------------------------------------------------------------------------------------------------||
//|| IntegrationSection
//|| Component for managing API integration keys and settings
//||------------------------------------------------------------------------------------------------||

import React, { useState, useEffect } from "react";

//||------------------------------------------------------------------------------------------------||
//|| Interface
//||------------------------------------------------------------------------------------------------||

import { Site }                     from "../../interfaces/model.sites";
import { integrationCode }          from "../../utils/integration.code";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface IntegrationSectionProps {
      data           : Site;
      updateField    : (field: string, value: string) => void;
}

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function IntegrationSection({ data, updateField }: IntegrationSectionProps) {

      //||------------------------------------------------------------------------------------------------||
      //|| State
      //||------------------------------------------------------------------------------------------------||

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
      }, [data]);

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
                              <label htmlFor="publicKey" className="label pb-1">Public Key</label>
                              <input
                                    id="publicKey"
                                    type="text"
                                    placeholder="Public Key"
                                    className="input input-bordered w-full text-white"
                                    readOnly
                                    value={publicKey}
                              />
                        </div>
                        <div>
                              <label htmlFor="allowed" className="label pb-1">Allowed Domains</label>
                              <textarea
                                    id="allowed"
                                    placeholder="Allowed Domains (comma-separated)"
                                    className="textarea textarea-bordered w-full"
                                    defaultValue={domains}
                                    onChange={(e) => updateField("domains", e.target.value)}
                              ></textarea>
                        </div>
                        <div>
                              <label htmlFor="testMode" className="label pb-1">Test Mode</label>
                              <select
                                    id="testMode"
                                    placeholder="Test Mode"
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
