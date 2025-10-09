//||------------------------------------------------------------------------------------------------||
//|| GateCustomSection
//|| Component for configuring site gate custom settings
//||------------------------------------------------------------------------------------------------||

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, {useState, useEffect}                   from "react";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

import type { ModelSite }                             from "../../../interfaces/models/model.sites";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import LabelDescription                               from "../../../components/base/LabelDescription";

//||------------------------------------------------------------------------------------------------||
//|| Functions
//||------------------------------------------------------------------------------------------------||

import { isValidURL }                                 from "../../../utils/validate";
import { integrationCode }                            from "../../../utils/integration.code";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface CustomAgeGateProps {
	data                    : ModelSite;
	updateField             : (field: string, value: any) => void;
}

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function CustomAgeGate({data, updateField}: CustomAgeGateProps) {

      //||------------------------------------------------------------------------------------------------||
	//|| State
	//||------------------------------------------------------------------------------------------------||

      const [showPreview, setShowPreview]             = useState(false);
      const [validConfirm, setValidConfirm]           = useState(false);
      const [validExit, setValidExit]                 = useState(false);
      const [codePreview, setCodePreview]             = useState(integrationCode(data?.clientId || "123"));

      //||------------------------------------------------------------------------------------------------||
	//|| JSX
	//||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            setValidConfirm(isValidURL(data.webhook));
      }, [data]);

      //||------------------------------------------------------------------------------------------------||
	//|| JSX
	//||------------------------------------------------------------------------------------------------||

      return (
		<div className="w-full bg-base-100 shadow-lg rounded-lg p-8 mb-5">
			<h2 className="text-2xl font-bold mb-6">Age Gate Setup</h2>

			{/* Custom Confirm URL */}
			<div className="mb-4">
				<LabelDescription id="confirmURL" label="Webhook URL" description="This is the URL where age verification results will be sent." />
				<div className="flex items-center space-x-4">
					<input
						type="text"
						value={data?.webhook}
						onChange={(e) => updateField("webhook", e.target.value)}
						placeholder="https://..."
						className="input input-bordered flex-1"
					/>
				</div>
				{!validConfirm && data.webhook != "" && <p className="text-red-500 text-sm mt-1">Please enter a valid URL for the Webhook</p>}
			</div>

			{/* Check Key */}
			<div className="mb-4">
				<LabelDescription
					id="checkKey"
					label="Validation Key"
					description="This secret key will be passed to the webhook to ensure data is from ComplyAge."
				/>

				<div className="flex items-center space-x-4">
					<input
						type="text"
						value={data?.checkKey}
						onChange={(e) => updateField("gateExit", e.target.value)}
						className="input input-bordered flex-1"
					/>
				</div>
			</div>

			<div>
				<LabelDescription
					id="integrationCode"
					label="Integration Code"
					description="Javascript to integrate the age gate on your website. Place this code before the closing </body> tag."
				/>
				<input
					id="integrationCode"
					readOnly
					className="w-full font-mono text-xs bg-base-200 p-3 text-gray-200"
					value={codePreview}
                        />
			</div>
		</div>
	);
}
