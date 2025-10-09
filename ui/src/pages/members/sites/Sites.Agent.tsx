//||------------------------------------------------------------------------------------------------||
//|| SiteAgentSettings
//|| Component describing the upcoming "Agent" (AI User Verification API)
//||------------------------------------------------------------------------------------------------||

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React from "react";

//||------------------------------------------------------------------------------------------------||
//|| Model
//||------------------------------------------------------------------------------------------------||

import { ModelSite } from "../../../interfaces/models/model.sites";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface SiteAgentSettingsProps {
      data              : ModelSite;
      updateField       : (field: string, value: any) => void;
}

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function SiteAgentSettings({ data, updateField }: SiteAgentSettingsProps) {

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      return (
            <div className="w-full bg-base-100 shadow-lg rounded-lg p-8 mb-5">
                  <h2 className="text-2xl font-bold mb-4">Agent (Coming Soon)</h2>
                  <p className="text-sm text-gray-300 leading-6">
                        The <strong>Agent</strong> is an upcoming AI-powered user verification API designed to 
                        automatically verify identity, age, and compliance data with minimal developer setup. 
                        It will integrate directly into your site’s compliance flow, allowing automated validation 
                        of government IDs, facial recognition, address confirmation, and behavioral signals — all 
                        securely processed through the ComplyAge ecosystem.
                  </p>
                  <p className="text-sm text-gray-400 mt-4 leading-6">
                        Once released, this section will allow you to configure Agent integration keys, API endpoints, 
                        and verification options, enabling frictionless, privacy-focused compliance checks powered by 
                        on-device AI and real-time evaluation models.
                  </p>
                  <p className="text-xs text-gray-500 italic mt-6">
                        Feature currently disabled — this integration will become available in an upcoming release.
                  </p>
            </div>
      );
}
