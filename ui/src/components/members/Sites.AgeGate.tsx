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

import type { Site }                                  from "../../interfaces/model.sites";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import PreviewGate                                    from "./Sites.GatePreview";

//||------------------------------------------------------------------------------------------------||
//|| Functions
//||------------------------------------------------------------------------------------------------||

import { isValidURL }                                 from "../../utils/validate";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface CustomAgeGateProps {
	data: Site;
	updateField: (field: string, value: any) => void;
}

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function CustomAgeGate({data, updateField}: CustomAgeGateProps) {

      //||------------------------------------------------------------------------------------------------||
	//|| State
	//||------------------------------------------------------------------------------------------------||

      const [showPreview, setShowPreview] = useState(false);
      const [validConfirm, setValidConfirm] = useState(false);
      const [validExit, setValidExit] = useState(false);


      useEffect(() => {
            // Validate URLs
            setValidConfirm(isValidURL(data.gateConfirm));
            setValidExit(isValidURL(data.gateExit));
      }, [data]);

      //||------------------------------------------------------------------------------------------------||
	//|| JSX
	//||------------------------------------------------------------------------------------------------||

      return (
		<div className="w-full bg-base-100 shadow-lg rounded-lg p-8 mb-5">
			<h2 className="text-2xl font-bold mb-6">Gate Custom Settings</h2>

			{/* Require Signup */}
			<div className="mb-4">
				<label htmlFor="requireSignup" className="block label pb-1">
					Show ComplyAge Signup 
				</label>
				<select 
                              id="requireSignup" 
                              value={String(data.gate_signup)} 
                              onChange={(e) => updateField("gate_signup", e.target.value === "1" ? "1" : "0")}
                              className="select w-full max-w-xs"
                        >
					<option value="1">True</option>
					<option value="0">False</option>
				</select>
			</div>

			{/* Custom Confirm URL */}
			<div>
				<label htmlFor="confirmURL" className="block label pb-1">
					Custom Confirm URL
				</label>
				<div className="flex items-center space-x-4">
                              <input
                                    type="text"
                                    value={data?.gateConfirm}
                                    onChange={(e) => updateField("gateConfirm", e.target.value)}
                                    placeholder="https://..."
                                    className="input input-bordered flex-1"
                              />
				</div>
                        {(!validConfirm && data.gateConfirm != "") && (
                              <p className="text-red-500 text-sm mt-1">
                                    Please enter a valid URL for the Confirm URL.
                              </p>
                        )}
			</div>

			{/* Custom Exit URL */}
			<div className="mt-6">
				<label htmlFor="exitURL" className="block label pb-1">
					Custom Exit URL
				</label>
				<div className="flex items-center space-x-4">
                              <input
                                    type="text"
                                    value={data?.gateExit}
                                    onChange={(e) => updateField("gateExit", e.target.value)}
                                    placeholder="https://..."
                                    className="input input-bordered flex-1"
                              />
				</div>
                        {(!validExit && data.gateExit != "") && (
                              <p className="text-red-500 text-sm mt-1">
                                    Please enter a valid URL for the Exit URL.
                              </p>
                        )}
			</div>
                  <>
                  <button onClick={() => setShowPreview(true)} className="btn btn-primary mt-6">Preview Gate</button>
                  { showPreview && (<PreviewGate data={data} onClose={() => setShowPreview(false)} />) }
                  </>
		</div>
	);
}
