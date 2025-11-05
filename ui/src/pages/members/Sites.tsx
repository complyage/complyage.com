//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useRef, useState }  from "react";
import { useNavigate }              from "react-router-dom";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

import { ModelSite }                from "../../interfaces/models/model.sites";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import MembersLayout                from "../../layouts/MembersLayout";
import { ChevronRightIcon }         from "lucide-react";
import apiURL                       from "../../utils/apiURL";

//||------------------------------------------------------------------------------------------------||
//|| Site Sections
//||------------------------------------------------------------------------------------------------||

import SiteSections                 from "../../components/nav/SiteSections";
import type { SectionTypes }        from "../../interfaces/models/model.sites";

//||------------------------------------------------------------------------------------------------||
//|| Status
//||------------------------------------------------------------------------------------------------||

import { getAccountStatus }                                 from "../../data/getAccountData";      
import { AccountStatusTypes }                               from "../../interfaces/models/model.account";

//||------------------------------------------------------------------------------------------------||
//|| Site Components
//||------------------------------------------------------------------------------------------------||

import BasicDataSection             from "./sites/Sites.BasicData";
import ZonesEnforcementSection      from "./sites/Sites.Zones";
import WebsiteManagerSection        from "./sites/Sites.Manager";
import CustomAgeGate                from "./sites/Sites.AgeGate";
import OAuthSettingsSection         from "./sites/Sites.OAuth";
import SiteAgentSection             from "./sites/Sites.Agent";
import SiteTestingSection           from "./sites/Sites.Testing";

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function Sites() {

      //||------------------------------------------------------------------------------------------------||
      //|| Var
      //||------------------------------------------------------------------------------------------------||
      const currentSite                   = useRef<number | null>(null);
      const [site, setSite]               = useState<ModelSite | null>(null);
      const [changed, setChanged]         = useState<boolean>(false);
      const [deleting, setDeleting]       = useState<boolean>(false);
      const [cacheBust, setCacheBust]     = useState<string>(Date.now().toString());
      const [siteSection, setSiteSection] = useState<SectionTypes>("basic");      
      const [testMode, setTestMode]       = useState(site?.testMode || false);

      //||------------------------------------------------------------------------------------------------||
      //|| Load Site
      //||------------------------------------------------------------------------------------------------||

      const loadSite = async (id : number | string | null) => {            
            if (changed && String(id) !== String(currentSite.current)) {                  
                  if (!confirm("Are you sure? You have unsaved changes that will be lost")) return;
            }
            const newSite = await fetch(`/v1/api/sites/load?id=${id}`, { method : "GET", credentials: "include" });
            if (!newSite.ok) {
                  alert("Failed to load site");
                  return;
            }
            const json = await newSite.json();
            const siteTemp = json.data.site;
            if (typeof siteTemp.zones  === "string") try {
                  siteTemp.zones = JSON.parse(siteTemp.zones);
            } catch (e) {
                  siteTemp.zones = {};
                  console.error("Failed to parse zones", e);
            }
            setSite(siteTemp);
            setChanged(false);
            setDeleting(false);
            currentSite.current = siteTemp.id;
            console.log("LOADED SITE", json.data.site);
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Create New 
      //||------------------------------------------------------------------------------------------------||

      const createNew = async () => {
            if (changed && !confirm("You have unsaved changes. Are you sure you want to create a new site?")) return;
            setSite(null);
            setChanged(false);
            setDeleting(false);
            currentSite.current = null;
            const response = await fetch(apiURL("/v1/api/sites/create"), { method: "POST", credentials: "include" });
            if (!response.ok) {
                  alert("Failed to create new site");
                  return;
            }
            const json = await response.json();
            console.log("CREATED NEW SITE", json.data.id);
            return json.data.id;
      }


      //||------------------------------------------------------------------------------------------------||
      //|| Copy Site
      //||------------------------------------------------------------------------------------------------||

      const copySite = async () => {
            if (site === null) return alert("No site selected to copy");
            if (changed) return alert("You must save your changes before copying");
            if (!confirm("Are you sure you want to copy this site?")) return;
            const response = await fetch("/v1/api/sites/copy?id=" + site.id, { method: "GET", credentials: "include" });
            if (!response.ok) {
                  alert("Failed to copy this site");
                  return;
            }     
            const json = await response.json();       
            setSite(json.data.id);
            setChanged(false);
            setDeleting(false);
            currentSite.current = json.data.id;

            return json.data.id;
      }


      //||------------------------------------------------------------------------------------------------||
      //|| Update Field
      //||------------------------------------------------------------------------------------------------||

      const updateField = (field: string, value: any) => {
            if (!site) return;
            setSite({ ...site, [field]: value, } as ModelSite); 
            if (field !== "logo") setChanged(true);
      };
          
      //||------------------------------------------------------------------------------------------------||
      //|| Update Field
      //||------------------------------------------------------------------------------------------------||

      const updateSite = async () => {
		if (!site) return;

		//||------------------------------------------------------------------------------------------------||
		//|| Update Site
		//||------------------------------------------------------------------------------------------------||

		const res = await fetch("/v1/api/sites/update", {
			method: "POST",
			headers: {"Content-Type": "application/json"},
			body: JSON.stringify(site),
			credentials: "include",
		});

		if (!res.ok) {
			const errorJson = await res.json();
			alert("Failed to update site: " + (errorJson.message || res.statusText));
			return;
		}

		//||------------------------------------------------------------------------------------------------||
		//|| Handle Response
		//||------------------------------------------------------------------------------------------------||

		const json = await res.json();
		console.log("UPDATED SITE", json.data.site);
		setChanged(false);
		setTimeout(() => {
			loadSite(site.id);
			setCacheBust(Date.now().toString());
		}, 100);
	};

      //||------------------------------------------------------------------------------------------------||
      //|| Delete
      //||------------------------------------------------------------------------------------------------||

      const deleteSite = async () => {
            if (!confirm("Are you sure you want to delete this site? \r\nThis action is permanent and cannot be undone.")) {
                  setDeleting(false);
                  return;                  
            }
            const response = await fetch(`/v1/api/sites/delete?id=${currentSite.current}`, {
                  method: "DELETE",
                  credentials: "include"
            });
            if (!response.ok) {
                  alert("Failed to delete site");
                  return;
            }
            setSite(null);
            setChanged(false);
            setDeleting(false);            
      }


      //||------------------------------------------------------------------------------------------------||
      //|| Toggle TEST MODE
      //||------------------------------------------------------------------------------------------------||

      const toggleTestMode = (toTestMode: boolean) => {
            updateField("testMode", toTestMode);
            setTestMode(toTestMode);
      }      

      //||------------------------------------------------------------------------------------------------||
      //|| JSX 
      //||------------------------------------------------------------------------------------------------||

      return (
            <MembersLayout title="Sites you're protecting">
                  <div className="flex flex-col items-center w-full">
                        <WebsiteManagerSection
                              data={ site }
                              onAddNew={ createNew }
                              onCopy={ copySite }
                              loadSite={loadSite}
                        />

                        { site && (
                              <>
                              <SiteSections value={ siteSection } setValue={ (value : SectionTypes) => { setSiteSection(value); } } />
                              <div className="bg-transparent w-full justify-center p-5 rounded-lg rounded-t-none">

                                    <div className="flex flex-wrap items-end gap-6 justify-top-center mb-4">

                                          <div className="flex-1 min-w-[200px] text-center">
                                                <label htmlFor="siteTestmode" className="block label pb-1">Test Mode</label>
                                                {testMode ? (
                                                      <button
                                                            onClick={() => toggleTestMode(false)}
                                                            className="btn btn-secondary cursor-pointer"
                                                      >
                                                            In Production
                                                      </button>
                                                ) : (
                                                      <button
                                                            onClick={() => toggleTestMode(true)}
                                                            className="btn btn-primary cursor-pointer"
                                                      >
                                                            In Test Mode
                                                      </button>
                                                )}
                                          </div>

                                          <div className="flex-1 min-w-[400px] text-center">
                                                <label htmlFor="clientId" className="block label pb-1">Client ID</label>
                                                <input id="clientId" className="input text-center inline-block text-gray-200" readOnly defaultValue={site?.clientId} />
                                          </div>                                    

                                          <div className="flex-1 min-w-[200px] text-center">
                                                <label htmlFor="siteStatus" className="block label pb-1">Site Status</label>
                                                <span id="siteStatus" className="inline-block bg-black/50 px-4 py-1 text-yellow-300">
                                                      {getAccountStatus(site?.status as AccountStatusTypes )}
                                                </span>
                                          </div>

                                    </div>


                              { siteSection === "basic" && (<BasicDataSection data={site} updateField={updateField} cacheBust={cacheBust} key={ `siteBasic-${site?.id || 0}` } />) }
                              { siteSection === "gate" && (
                                    <>
                                    <CustomAgeGate data={site} updateField={updateField} key={ `customGate-${site?.id || 0}` } />
                                    <ZonesEnforcementSection data={site} updateField={updateField} key={ `sitezZone-${site?.id || 0}` } /> 
                                    </>
                              )}
                              { siteSection === "agent" && (<SiteAgentSection data={site} updateField={updateField} key={ `siteOAuth-${site?.id || 0}` } /> ) }
                              { siteSection === "oauth" && (<OAuthSettingsSection data={site} updateField={updateField} key={ `siteOAuth-${site?.id || 0}` } /> ) }
                              { siteSection === "testing" && (<SiteTestingSection data={site} key={ `siteTesting-${site?.id || 0}` } /> ) }
            
                              {/* Delete */}
                              { siteSection === "basic" && site && (
                                    <div className="w-full bg-base-100 opacity-40 hover:opacity-80 bg-shadow-lg rounded-lg p-8 mb-10">
                                          { !deleting && ( 
                                                <>
                                                      <div className="flex flex-row justify-center items-center text-md font-normal">
                                                            <span className="mr-2">Remove this site and all its settings</span>
                                                            <a className="link" onClick={() => { setDeleting(true); }}>Delete Site</a>
                                                      </div>                                          
                                                </>
                                          )}

                                          { deleting && (
                                                <div className="flex flex-row justify-center items-center">
                                                      <div className="p-2">
                                                            Remove this site and all its settings. <b className="text-yellow-500">This action is permanent and cannot be undone.</b>
                                                      </div>
                                                      <div>
                                                            <button
                                                                  className="btn btn-primary mr-3 w-auto rounded-lg mx-auto"
                                                                  onClick={() => { setDeleting(false) }}
                                                            >Cancel</button>

                                                            <button
                                                                  className="btn btn-black w-auto rounded-lg mx-auto hover:bg-orange-400"
                                                                  onClick={() => { deleteSite(); }}
                                                            >
                                                                  Delete Site
                                                            </button>
                                                      </div>
                                                </div>
                                          )}
                                    </div>
                              )}
            
                              {/* Save */}
                              {changed && (
                                    <div className="fixed bottom-0 right-0 left-64 text-right p-3 pt-5 flex justify-end items-center bg-black/100">
                                          <span className="text-warning p-2 px-5 rounded-md mr-4 bg-white/10">
                                                You have unsaved changes <ChevronRightIcon className="inline" />
                                          </span>
                                          <button
                                                onClick={() => { updateSite() }}
                                                className="btn btn-secondary justify-end"
                                                disabled={!changed}
                                          >Save</button>
                                    </div>
                              )}
                        </div>
                        </>
                  )}                        
                  </div>
            </MembersLayout>
            
      );
      
      
}