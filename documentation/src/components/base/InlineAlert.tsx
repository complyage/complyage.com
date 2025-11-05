//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React from "react";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface InlineAlertProps {
      message      : string;
      isError?     : boolean;
}

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function InlineAlert({ message, isError = true }: InlineAlertProps) {

      if (!message) return null;

      return (
            <div 
                  className={`w-full p-1 mb-4 rounded-lg shadow 
                              ${isError ? "bg-red-500" : "bg-green-500"} 
                              text-white flex items-center gap-1`}>
                  
                  {/* Icon */}
                  <svg xmlns="http://www.w3.org/2000/svg" 
                        className="h-8 w-8 shrink-0 stroke-current" 
                        fill="none" 
                        viewBox="0 0 24 24" 
                        stroke="currentColor">
                        {isError ? (
                              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" 
                                    d="M12 9v2m0 4h.01M12 5a7 7 0 110 14a7 7 0 010-14z" />
                        ) : (
                              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" 
                                    d="M5 13l4 4L19 7" />
                        )}
                  </svg>

                  {/* Message */}
                  <span className="text-sm font-medium">{message}</span>
            </div>
      );
}
