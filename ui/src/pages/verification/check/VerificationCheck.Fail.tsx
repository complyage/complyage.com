//||------------------------------------------------------------------------------------------------||
//|| CCVerification (Container)
//|| src/pages/verification/CCVerification.tsx
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||

      import React, {useRef, useState, useEffect}           from "react";
      import {useNavigate, useSearchParams}                 from "react-router-dom";
      import {CheckCircle, Home}                            from "lucide-react";

      //||------------------------------------------------------------------------------------------------||
      //|| Data
      //||------------------------------------------------------------------------------------------------||

      import type { VerificationTypes } from "../../../interfaces/models/model.verify";
      import { getVerificationIcon, getVerificationType } from "../../../data/getVerificationData";

      //||------------------------------------------------------------------------------------------------||
      //|| Components
      //||------------------------------------------------------------------------------------------------||

      import MembersLayout                                  from "../../../layouts/MembersLayout";
      import ProgressSteps, { ProgessStep }                 from "../../../components/base/ProgressSteps";
      import InlineAlert                                    from "../../../components/base/InlineAlert";
      import SpinnerCircle                                  from "../../../components/base/SpinnerCircle";

      //||------------------------------------------------------------------------------------------------||
      //|| Page
      //||------------------------------------------------------------------------------------------------||

      export default function VerificationCheckFail({ error, verificationId, onBack } : { error?: string, verificationId?: any, onBack: () => void }) {
            //||------------------------------------------------------------------------------------------------||
            //|| Navigate
            //||------------------------------------------------------------------------------------------------||

            const navigate                     = useNavigate();

            //||------------------------------------------------------------------------------------------------||
            //|| Verification Data
            //||------------------------------------------------------------------------------------------------||

            return (
                  <div className="w-full max-w-2xl mx-auto p-6">                        
                        <div className="text-center bg-black/20 p-5 text-gray-200 mt-4">
                              <h2 className="text-2xl text-white border-b border-gray-400 p-2 mb-5">Verification Error</h2>
                              <InlineAlert message={error || "Unknown Error"} isError />                              
                              <p className="text-gray-200 font-bold p-5">                                    
                                    <span className="text-gray-500 p-2">{verificationId}</span>
                              </p>
                              <button
                                    className="btn btn-secondary mt-4"
                                    onClick={ onBack }
                              >
                                    Try Again
                              </button>
                        </div>
                  </div>
            );
      }

