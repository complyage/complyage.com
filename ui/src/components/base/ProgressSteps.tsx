/*||------------------------------------------------------------------------------------------------||
//|| Progress Steps Component
//|| ProgressSteps
//||------------------------------------------------------------------------------------------------||*/

/*||------------------------------------------------------------------------------------------------||
//|| React
//||------------------------------------------------------------------------------------------------||*/

import React                  from "react";
import {Check}                from "lucide-react";

/*||------------------------------------------------------------------------------------------------||
//|| Step
//||------------------------------------------------------------------------------------------------||*/

export type ProgressStep = {      
	label             : string;
	description?      : string;
};

/*||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||*/

interface ProgressStepsProps {
	steps             : ProgessStep[];
	currentStep       : number;
	className?        : string;
}

/*||------------------------------------------------------------------------------------------------||
//|| Helpers
//||------------------------------------------------------------------------------------------------||*/

const clamp = (n: number, min: number, max: number) => Math.max(min, Math.min(max, n));

/*||------------------------------------------------------------------------------------------------||
//|| Progress Steps
//||------------------------------------------------------------------------------------------------||*/

const ProgressSteps: React.FC<ProgressStepsProps> = ({steps, currentStep, className}) => {
      const maxSteps = steps.length || 1;
      const allComplete = currentStep > maxSteps;
      const stepNow = clamp(currentStep, 1, maxSteps);
      const activeStep = allComplete ? undefined : steps[stepNow - 1];

      return (
            <div className={`w-full ${className ?? ""}`}>
                  {/* Current step description (only once) */}
                  {!allComplete && activeStep?.description ? (
                        <h2 className="text-2xl leading-snug text-gray-200 font-bold border-b border-b-gray-300 py-2 mb-4">
                              {activeStep.description}
                        </h2>
                  ) : null}

                  {/* Stepper */}
                  <div className="flex items-center w-full" role="list" aria-label="Progress">
                        {steps.map((s, i) => {
                              const stepNum = i + 1;
                              const isCompleted = allComplete || stepNum < stepNow;
                              const isActive = !allComplete && stepNum === stepNow;

                              return (
                                    <React.Fragment key={stepNum}>
                                          {/* Step Column */}
                                          <div className="flex-1 basis-0 min-w-0 flex flex-col items-center" role="listitem">
                                                {/* Step Circle */}
                                                <div
                                                      className={[
                                                            "flex items-center justify-center w-8 h-8 rounded-full border-2 transition-all duration-200",
                                                            isCompleted
                                                                  ? "bg-green-500 border-green-500 text-white"
                                                                  : isActive
                                                                  ? "border-blue-500 text-blue-500"
                                                                  : "border-gray-300 text-gray-400",
                                                      ].join(" ")}
                                                      aria-current={isActive ? "step" : undefined}
                                                      aria-label={`${stepNum}. ${s.label}${s.description ? ` — ${s.description}` : ""}`}>
                                                      {isCompleted ? <Check className="w-5 h-5" /> : stepNum}
                                                </div>

                                                {/* Step Text */}
                                                <div className="mt-2 text-center px-2">
                                                      <div 
                                                            className={[
                                                                  "text-xs font-medium",
                                                                  isCompleted ? "text-green-400" : isActive ? "text-blue-500" : "text-gray-200"
                                                            ].join(" ")}                                                      
                                                      >{s.label}</div>
                                                </div>
                                          </div>

                                          {/* Connecting Line */}
                                          {stepNum < maxSteps && (
                                                <div className={["flex-none w-12 h-0.5 mx-2", isCompleted ? "bg-green-500" : isActive ? "bg-blue-500" : "bg-gray-300"].join(" ")} aria-hidden="true" />
                                          )}
                                    </React.Fragment>
                              );
                        })}
                  </div>
            </div>
      );
};

/*||------------------------------------------------------------------------------------------------||
//|| Export
//||------------------------------------------------------------------------------------------------||*/

export default ProgressSteps;
