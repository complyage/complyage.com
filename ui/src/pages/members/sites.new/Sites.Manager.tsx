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

import { ModelSite }                            from "../../../interfaces/models/model.sites";


//||------------------------------------------------------------------------------------------------||
//|| Lucite Icons
//||------------------------------------------------------------------------------------------------||

import { RefreshCcw }                           from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface SiteManagerProps {
      fetchWebsites   : () => void;
      websites          : ModelSite[];
      site              : ModelSite | null;
      onCopy            : (site: ModelSite | null) => void;
      onAddNew          : () => void;
      loadSite          : (id: number) => void;
}

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function SiteManager({ websites, site, onCopy, onAddNew, loadSite }: SiteManagerProps) {

	//||------------------------------------------------------------------------------------------------||
	//|| State
	//||------------------------------------------------------------------------------------------------||

	const [selected, setSelected] = useState<number | null>(null);
	const [loading, setLoading]   = useState(true);      

	//||------------------------------------------------------------------------------------------------||
	//|| Create a New Website
	//||------------------------------------------------------------------------------------------------||

      const handleOnNew = async () => {
            const id = await onAddNew();
            if (typeof id === "number") {
                  loadSite(id);
                  setSelected(id); 
            }
      }

	//||------------------------------------------------------------------------------------------------||
	//|| Create a New Website
	//||------------------------------------------------------------------------------------------------||

      const handleOnCopy = async () => {
            const id = await onCopy(site);
            if (typeof id === "number") {
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
                                          value={site?.id}
                                          onChange={(e) => {
                                                const id = Number(e.target.value);
                                                setSelected(id);
                                                loadSite(id);
                                          }}
                                    >
                                          <option disabled value="">Select a website</option>
                                          {websites.map((site : ModelSite) => (
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
                                    <button className="btn btn-primary"  onClick={handleOnCopy}>Copy</button>
					      <button className="btn btn-secondary" onClick={handleOnNew}>Add New</button>
					</div>
				</div>
			)}
		</div>
	);
}
