//||------------------------------------------------------------------------------------------------||
//|| WidgetPanel
//|| components/widgets/WidgetPanel.tsx
//||------------------------------------------------------------------------------------------------||

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useEffect, useState }                       from "react";

//||------------------------------------------------------------------------------------------------||
//|| Classes
//||------------------------------------------------------------------------------------------------||

import Call                                                 from "../../classes/call";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface WidgetPanelProps {
      title             : string;
      route             : string;
      selectOptions?    : string[]; // optional filter buttons
}

//||------------------------------------------------------------------------------------------------||
//|| Widget Panel
//||------------------------------------------------------------------------------------------------||

export default function WidgetPanel({ title, route, selectOptions }: WidgetPanelProps) {

      //||------------------------------------------------------------------------------------------------||
      //|| State
      //||------------------------------------------------------------------------------------------------||

      const [data, setData]                   = useState<any>(null);
      const [selected, setSelected]           = useState<string>(selectOptions?.[0] || "");
      const [loading, setLoading]             = useState<boolean>(false);
      const [error, setError]                 = useState<string | null>(null);

      //||------------------------------------------------------------------------------------------------||
      //|| Fetch Data
      //||------------------------------------------------------------------------------------------------||

      const fetchData = async () => {
            setLoading(true);
            setError(null);

            let finalRoute = route;
            if (selected) {
                  const param = encodeURIComponent(selected);
                  finalRoute += route.includes("?") ? `&filter=${param}` : `?filter=${param}`;
            }

            const call = new Call(finalRoute);
            call.method = "GET";
            const res = await call.get();

            if (call.ok()) {
                  const value = res.data?.count ?? "0";
                  setData(value);
            } else {
                  setError(call.error());
            }

            setLoading(false);
      };

      //||------------------------------------------------------------------------------------------------||
      //|| useEffect
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            fetchData();
      }, [route, selected]);

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      return (
            <div className="bg-black/40 rounded-lg shadow p-5 flex flex-col justify-between min-h-[150px]">
                  {/* Title */}
                  <h3 className="text-lg font-semibold text-gray-300 mb-2">{title}</h3>

                  {/* Value */}
                  <div className="flex-grow flex items-center justify-start mt-2 mb-4">
                        {loading ? (
                              <p className="text-2xl font-bold text-gray-400">Loading...</p>
                        ) : error ? (
                              <p className="text-sm text-red-400">{error}</p>
                        ) : (
                              <p className="text-3xl font-bold text-white">{data}</p>
                        )}
                  </div>

                  {/* Options */}
                  {selectOptions && (
                        <div className="flex flex-wrap gap-2 mt-auto">
                              {selectOptions.map(opt => (
                                    <button
                                          key={opt}
                                          onClick={() => setSelected(opt)}
                                          className={`text-xs px-3 py-1 rounded font-medium ${
                                                selected === opt
                                                      ? "bg-green-400 text-gray-900"
                                                      : "bg-gray-600 text-gray-300 hover:bg-gray-500"
                                          }`}
                                    >
                                          {opt.toUpperCase()}
                                    </button>
                              ))}
                        </div>
                  )}
            </div>
      );
}