//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React from "react";

//||------------------------------------------------------------------------------------------------||
//|| SectionBar Component
//||------------------------------------------------------------------------------------------------||

import type { SectionTypes } from "../../interfaces/types.sitesections";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface SiteSectionsProps {
      value    : SectionTypes;
      setValue : (value: SectionTypes) => void;
}

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function SiteSections({ value, setValue }: SiteSectionsProps) {
      const sections: Record<SectionTypes, string> = {
            basic       : "Basic",
            zones       : "Zones",
            integration : "Integration",
            gate        : "Age Gate",
            oauth       : "OAuth",
      };

      return (
            <nav className="flex space-x-4 bg-black/20 my-5 rounded-lg">
                  {Object.entries(sections).map(([key, label]) => (
                        <button
                              key={key}
                              onClick={() => setValue(key as SectionTypes)}
                              className={`py-2 px-4 -mb-px border-b-2 cursor-pointer ${
                                    value === key
                                          ? "border-yellow-500 font-semibold text-yellow-500 px-2 mx-2"
                                          : "border-transparent text-gray-300 hover:text-gray-400 mx-2"
                              }`}
                        >
                              {label}
                        </button>
                  ))}
            </nav>
      );
}
