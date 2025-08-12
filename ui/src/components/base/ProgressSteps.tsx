// ProgressSteps.tsx
import React from "react";
import { Check } from "lucide-react";

interface ProgressStepsProps {
   maxSteps: number;
   currentStep: number;
}

const ProgressSteps: React.FC<ProgressStepsProps> = ({ maxSteps, currentStep }) => {
   return (
      <div className="flex items-center w-full">
         {Array.from({ length: maxSteps }, (_, i) => {
            const step = i + 1;
            const isCompleted = step < currentStep;
            const isActive = step === currentStep;

            return (
               <React.Fragment key={step}>
                  {/* Step Circle */}
                  <div
                     className={`flex items-center justify-center w-8 h-8 rounded-full border-2 transition-all duration-200 ${
                        isCompleted
                           ? "bg-green-500 border-green-500 text-white"
                           : isActive
                           ? "border-blue-500 text-blue-500"
                           : "border-gray-300 text-gray-400"
                     }`}
                  >
                     {isCompleted ? <Check className="w-5 h-5" /> : step}
                  </div>

                  {/* Connecting Line */}
                  {step < maxSteps && (
                     <div
                        className={`flex-1 h-0.5 mx-2 ${
                           isCompleted
                              ? "bg-green-500"
                              : isActive
                              ? "bg-blue-500"
                              : "bg-gray-300"
                        }`}
                     />
                  )}
               </React.Fragment>
            );
         })}
      </div>
   );
};

export default ProgressSteps;
