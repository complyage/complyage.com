//||------------------------------------------------------------------------------------------------||
//|| LabelDescription
//|| General-purpose labeled field description component
//||------------------------------------------------------------------------------------------------||

import React from "react";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface LabelDescriptionProps {
      id?            : string;
      label          : string;
      description?   : string;
      required?      : boolean;
      className?     : string;
}

//||------------------------------------------------------------------------------------------------||
//|| Default Component
//||------------------------------------------------------------------------------------------------||

export default function LabelDescription({ id, label, description, required = false, className = "" }: LabelDescriptionProps) {

      return (
            <div className={`mb-2 ${className}`}>
                  <label 
                        htmlFor={id}
                        className="block label font-semibold text-sm text-gray-300"
                  >
                        {label}
                        {required && <span className="text-red-500 ml-1">*</span>}
                  </label>
                  {description && (
                        <span className="block text-xs text-gray-400 leading-snug">
                              {description}
                        </span>
                  )}
            </div>
      );
}
