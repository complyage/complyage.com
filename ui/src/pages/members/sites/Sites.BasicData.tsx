//||------------------------------------------------------------------------------------------------||
//|| BasicDataSection
//|| Component for editing basic website info including logo upload + allowed domains
//||------------------------------------------------------------------------------------------------||

import React, { useState, useEffect, useRef }               from "react";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

import { ModelSite }                                        from "../../../interfaces/models/model.sites";
import SpinnerCircle                                        from "../../../components/base/SpinnerCircle";
import { X }                                                from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import LabelDescription                               from "../../../components/base/LabelDescription";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface BasicDataProps {
      data        : ModelSite;
      updateField : (field : string, value : string | boolean | number) => void;
      cacheBust   : string;
}

//||------------------------------------------------------------------------------------------------||
//|| Default Component
//||------------------------------------------------------------------------------------------------||

export default function BasicDataSection({ updateField, data, cacheBust }: BasicDataProps) {

      const [name, setName]                     = useState(data?.name || "");
      const [url, setUrl]                       = useState(data?.url || "");
      const [description, setDescription]       = useState(data?.description || "");
      const [logo, setLogo]                     = useState<File | null>(null);
      const [logoUploading, setLogoUploading]   = useState(false);
      const [status, setStatus]                 = useState<"Active" | "Pending">("Pending");
      const fileInputRef                        = useRef<HTMLInputElement | null>(null);
      const [cb, setCb]                         = useState<string>(cacheBust || new Date().toISOString());

      //||------------------------------------------------------------------------------------------------||
      //|| Allowed Domains
      //||------------------------------------------------------------------------------------------------||

      const [currentDomain, setCurrentDomain]   = useState("");
      const [domainList, setDomainList]         = useState<string[]>([]);

      //||------------------------------------------------------------------------------------------------||
      //|| Sync With Parent Data
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            if (!data) return;
            setName(data.name || "");
            setUrl(data.url || "");
            setDescription(data.description || "");
            setStatus((data.status as "Active" | "Pending") || "Pending");
            setLogo(null);
            setCb(cacheBust);
            if (data?.domains) {
                  setDomainList(data.domains.split(",").map(d => d.trim()).filter(Boolean));
            } else {
                  setDomainList([]);
            }
      }, [data, cacheBust]);

      //||------------------------------------------------------------------------------------------------||
      //|| Domain Management
      //||------------------------------------------------------------------------------------------------||

      const addDomain = () => {
            if (currentDomain && !domainList.includes(currentDomain)) {
                  const updatedDomains = [...domainList, currentDomain];
                  setDomainList(updatedDomains);
                  updateField("domains", updatedDomains.join(","));
                  setCurrentDomain("");
            }
      };

      const removeDomain = (domain: string) => {
            const updatedDomains = domainList.filter(d => d !== domain);
            setDomainList(updatedDomains);
            updateField("domains", updatedDomains.join(","));
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Handle File Upload (POST to API)
      //||------------------------------------------------------------------------------------------------||

      const uploadLogo = async (file: File) => {
            const formData = new FormData();
            formData.append("image", file);
            formData.append("siteId", String(data.id));

            setLogoUploading(true);
            try {
                  const res = await fetch("/v1/api/sites/upload", {
                        method: "POST",
                        body: formData,
                        credentials: "include",
                  });

                  if (!res.ok) {
                        const err = await res.json().catch(() => ({}));
                        alert("Logo upload failed: " + (err.message || res.statusText));
                        return null;
                  }

                  const json = await res.json();
                  return json.data?.object as string;
            } finally {
                  setLogoUploading(false);
            }
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Handle File Select + Upload
      //||------------------------------------------------------------------------------------------------||

      const handleUploadChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
            const file = e.target.files?.[0];
            if (!file) return;

            setLogo(file);

            const uploaded = await uploadLogo(file);
            if (uploaded) {
                  updateField("logo", uploaded);
                  setCb(Date.now().toString());
            }
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Handle Drop + Upload
      //||------------------------------------------------------------------------------------------------||
      
      const handleDrop = async (e: React.DragEvent<HTMLDivElement>) => {
            e.preventDefault();
            const file = e.dataTransfer.files?.[0];
            if (!file) return;

            setLogo(file);
            updateField("forceUpdate", new Date().toISOString());

            const uploaded = await uploadLogo(file);
            if (uploaded) {
                  updateField("logo", uploaded);
                  setCb(Date.now().toString());
            }
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Handle Delete Logo
      //||------------------------------------------------------------------------------------------------||

      const handleDeleteLogo = async () => {
            if (!data.id) return;

            const formData = new FormData();
            formData.append("siteId", String(data.id));  // no "image" -> handler will clear

            setLogoUploading(true);
            try {
                  const res = await fetch("/v1/api/sites/upload", {
                        method: "POST",
                        body: formData,
                        credentials: "include",
                  });

                  if (!res.ok) {
                        const err = await res.json().catch(() => ({}));
                        alert("Logo delete failed: " + (err.message || res.statusText));
                        return;
                  }

                  updateField("logo", "");
                  setLogo(null);
                  setCb(Date.now().toString());
            } finally {
                  setLogoUploading(false);
            }
      };


      //||------------------------------------------------------------------------------------------------||
      //|| Logo
      //||------------------------------------------------------------------------------------------------||

      const showLogo = () => {
            if (logoUploading) return (
                  <div className="flex flex-col items-center">
                        <SpinnerCircle />
                  </div>                  
            )

            if (data.logo) {
                  const isUrl = /^https?:\/\//i.test(data.logo);
                  const src   = isUrl
                        ? data.logo
                        : `${import.meta.env.VITE_COMPLYAGE_MINIO_URL}/sites/${data.logo}?cb=${cb}`;
                  
                  return (
                        <div className="flex flex-col items-center gap-3">
                              <img
                                    src={src}
                                    alt="Logo"
                                    className="max-w-xs h-24 w-24 border rounded-lg shadow"
                              />
                              <button
                                    type="button"
                                    onClick={handleDeleteLogo}
                                    className="bg-red-600 text-white px-4 py-2 rounded-md font-bold hover:bg-red-700 transition"
                              >
                                    Delete Logo
                              </button>
                        </div>
                  );
            }

            return (<div className="flex justify-center items-center text-gray-500 text-sm italic">No logo uploaded yet</div>);
      }


      //||------------------------------------------------------------------------------------------------||
      //|| No Site Loaded
      //||------------------------------------------------------------------------------------------------||

      if (!data) return null;

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      return (
            <>
                  <div className="w-full bg-base-100 shadow-lg rounded-lg p-8 mb-5">
                        <h2 className="text-2xl font-bold mb-6">Website Identity</h2>

                        <div className="grid grid-cols-1 gap-6">

                              <div>
                                    <LabelDescription
                                          id="sitename"
                                          label="Name"
                                          description="The name of your website or application. (i.e. ComplyAge Demo)"
                                    />                                    
                                    <input
                                          id="sitename"
                                          type="text"
                                          placeholder="Website Name"
                                          className="input input-bordered w-full"
                                          defaultValue={name}
                                          onChange={(e) => updateField("name", e.target.value)}
                                    />
                              </div>

                              <div>
                                    <LabelDescription
                                          id="siteurl"
                                          label="URL"
                                          description="Homepage of your website or application. (i.e. https://complyage.com)"
                                    />                                                                        
                                    <input
                                          id="siteurl"
                                          type="text"
                                          placeholder="Website URL"
                                          className="input w-full"
                                          defaultValue={url}
                                          onChange={(e) => updateField("url", e.target.value)}
                                    />
                              </div>

                              <div>
                                    <LabelDescription
                                          id="description"
                                          label="Description"
                                          description="A short description of your website or application."
                                    />                                                                        
                                    <textarea
                                          id="description"
                                          placeholder="Website Description"
                                          className="textarea textarea-bordered w-full"
                                          defaultValue={description}
                                          onChange={(e) => updateField("description", e.target.value)}
                                    ></textarea>
                              </div>

                              <div>
                                    <LabelDescription
                                          id="allowed"
                                          label="Allowed Domains"
                                          description="Which domains should be allowed to call Age Gate/OAuth/Agent? (e.g. complyage.com, mysite.org)"
                                    />                                    
                                    <div className="flex items-center">
                                          <input 
                                                id="allowed"
                                                type="text" 
                                                className="input input-bordered text-white flex-1"
                                                onChange={(e) => setCurrentDomain(e.target.value)} 
                                                onKeyDown={(e) => { if (e.key === 'Enter') { e.preventDefault(); addDomain(); }}}
                                                value={currentDomain} 
                                                placeholder="Add domain and press Enter"
                                          />
                                          <button
                                                onClick={addDomain}
                                                className="btn btn-sm btn-primary ml-2"
                                          >
                                                Add
                                          </button>
                                    </div>
                                    <div className="mt-2 flex flex-wrap gap-2">
                                          { domainList.length === 0 && (
                                                <p className="text-sm text-gray-500 text-center p-4 bg-black/20 border-lg mb-5 w-full">
                                                      No domains added. Please add allowed domains.
                                                </p>
                                          )}
                                          {domainList.map((domain, index) => (
                                                <div 
                                                      key={index}
                                                      className="flex items-center justify-between bg-gray-700 text-white text-sm px-3 py-1 rounded-md shadow-sm"
                                                >
                                                      <span>{domain}</span>
                                                      <button
                                                            onClick={() => removeDomain(domain)}
                                                            className="ml-2 flex items-center justify-center text-white hover:text-red-600 cursor-pointer"
                                                      >
                                                            <X size={16} />
                                                      </button>
                                                </div>
                                          ))}
                                    </div>
                              </div>

                              <LabelDescription
                                    id="allowed"
                                    label="Logo"
                                    description="Update a Logo for your website. (Max 2MB, JPG/PNG/GIF) This will be shown on the approval pages."
                              />     
                              {showLogo()}

                              <div
                                    onClick={() => fileInputRef.current?.click()}
                                    onDrop={handleDrop}
                                    onDragOver={(e) => e.preventDefault()}
                                    className="border border-dashed border-gray-400 rounded-lg p-6 text-center cursor-pointer hover:bg-gray-50/10"
                              >
                                    {logo ? (
                                          <div className="font-medium">{logo.name}</div>
                                    ) : (
                                          <div className="opacity-50">Drag & drop logo here or click to upload</div>
                                    )}
                                    <input
                                          ref={fileInputRef}
                                          type="file"
                                          accept="image/*"
                                          className="hidden"
                                          onChange={handleUploadChange}
                                    />
                              </div>
                        </div>
                  </div>
            </>
      );
}
