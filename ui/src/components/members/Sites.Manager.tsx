//||------------------------------------------------------------------------------------------------||
//|| WebsiteManagerSection
//|| Component for managing websites with dropdown + action buttons
//||------------------------------------------------------------------------------------------------||

//||------------------------------------------------------------------------------------------------||
//|| React
//||------------------------------------------------------------------------------------------------||

import React, { useEffect, useState }           from "react";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

import { Site }                                 from "../../interfaces/model.sites";

//||------------------------------------------------------------------------------------------------||
//|| Hooks
//||------------------------------------------------------------------------------------------------||

import { useEffectOnce }                        from "../../hooks/useEffectOnce";

//||------------------------------------------------------------------------------------------------||
//|| Lucite Icons
//||------------------------------------------------------------------------------------------------||

import { RefreshCcw }                           from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface WebsiteManagerSectionProps {
      data              : Site | null;
      onCopy            : (site: Website | null) => void;
      onAddNew          : () => void;
      loadSite          : (id: number) => void;
}

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function WebsiteManagerSection({ data, onCopy, onAddNew, loadSite }: WebsiteManagerSectionProps) {
	//||------------------------------------------------------------------------------------------------||
	//|| State
	//||------------------------------------------------------------------------------------------------||

	const [websites, setWebsites] = useState<Website[]>([]);
	const [selected, setSelected] = useState<number | null>(null);
	const [loading, setLoading]   = useState(true);      

      //||------------------------------------------------------------------------------------------------||
	//|| useEffect
	//||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            console.log("WebsiteManagerSection data changed: REFRESHING", data);
            fetchWebsites();
            if (data === null || data.id === undefined) {
                  setSelected(null);
            }
      }, [data]);

      //||------------------------------------------------------------------------------------------------||
	//|| Load Websites
	//||------------------------------------------------------------------------------------------------||

      const fetchWebsites = async () => {
            try {
                  const res = await fetch("/v1/api/sites/list", {credentials: "include"});
                  console.log(res);
                  const json = await res.json();
                  console.log(json);
                  if (Array.isArray(json.data.sites)) {
                        setWebsites(json.data.sites);
                  }
            } catch (err) {
                  console.error("Failed to load websites:", err);
            } finally {
                  setLoading(false);
            }
      };

      //||------------------------------------------------------------------------------------------------||
	//|| Load Websites
	//||------------------------------------------------------------------------------------------------||

	useEffectOnce(() => {
		fetchWebsites();
	});

	//||------------------------------------------------------------------------------------------------||
	//|| Create a New Website
	//||------------------------------------------------------------------------------------------------||

      const handleOnNew = async () => {
            const id = await onAddNew();
            if (typeof id === "number") {
                  await fetchWebsites();
                  loadSite(id);
                  setSelected(id); 
            }
      }

	//||------------------------------------------------------------------------------------------------||
	//|| Create a New Website
	//||------------------------------------------------------------------------------------------------||

      const handleOnCopy = async () => {
            const id = await onCopy();
            if (typeof id === "number") {
                  await fetchWebsites();
                  loadSite(id);
                  setSelected(id); 
            }
      }      

	//||------------------------------------------------------------------------------------------------||
	//|| JSX
	//||------------------------------------------------------------------------------------------------||

	return (
		<div className="w-full bg-base-100 shadow-lg rounded-lg p-5 ">
			{loading ? (
				<div className="text-sm opacity-60">Loading websites...</div>
			) : (                        
				<div className="flex flex-col sm:flex-row sm:items-center gap-4">
                              <div className="justicy-start w-[60%]">
                                    <select
                                          className="select select-bordered w-full"
                                          value={selected ?? ""}
                                          onChange={(e) => {
                                                const id = Number(e.target.value);
                                                setSelected(id);
                                                loadSite(id);
                                          }}
                                    >
                                          <option disabled value="">Select a website</option>
                                          {websites.map((site : Site) => (
                                                <option key={site.id} value={site.id}>
                                                      {site.url + `-` + site.name }
                                                </option>
                                          ))}
                                    </select>
                              </div>
                              <div className="flex gap-2 w-[40%] justify-start">
                                    <button className="btn btn-neutral" onClick={() => { fetchWebsites() } }><RefreshCcw /></button>                                    
                              </div>
					<div className="flex gap-2 w-[40%] justify-end">
                                    <button className="btn btn-neutral"  onClick={handleOnCopy}>Copy</button>
					      <button className="btn btn-primary" onClick={handleOnNew}>Add New</button>
					</div>
				</div>
			)}
		</div>
	);
}
