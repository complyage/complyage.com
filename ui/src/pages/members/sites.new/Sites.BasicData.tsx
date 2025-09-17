//||------------------------------------------------------------------------------------------------||
//|| BasicDataSection
//|| Component for editing basic website info including logo upload
//||------------------------------------------------------------------------------------------------||

import React, { useState, useEffect, useRef }               from "react";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

import { ModelSite }                                        from "../../../interfaces/models/model.sites";
import { getAccountStatus }                                 from "../../../data/getAccountData";      
import type { AccountStatusTypes }                          from "../../../interfaces/models/model.account";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface BasicDataProps {
      site        : ModelSite;
      updateSite  : (field : string, value : string | number) => void;
}

//||------------------------------------------------------------------------------------------------||
//|| Default Component
//||------------------------------------------------------------------------------------------------||

export default function BasicDataSection({ updateSite, site }: BasicDataProps) {

      //||------------------------------------------------------------------------------------------------||
      //|| Input
      //||------------------------------------------------------------------------------------------------||

      const fileInputRef                  = useRef<HTMLInputElement | null>(null);
      const [logo, setLogo]               = useState<File | null>(null);
      const [logoPreview, setLogoPreview] = useState<string | null>(null);

      //||------------------------------------------------------------------------------------------------||
      //|| Handle File Select
      //||------------------------------------------------------------------------------------------------||

      const handleUploadChange = (e: React.ChangeEvent<HTMLInputElement>) => {
            const file = e.target.files?.[0];
            if (file) {
                  setLogo(file);
                  const reader = new FileReader();
                  reader.onloadend = () => {
                        setLogoPreview(reader.result as string);
                  };
                  reader.readAsDataURL(file);
            }
      };
         

      //||------------------------------------------------------------------------------------------------||
      //|| Handle Drop
      //||------------------------------------------------------------------------------------------------||

      const handleDrop = (e: React.DragEvent<HTMLDivElement>) => {
            e.preventDefault();
            const file = e.dataTransfer.files?.[0];
            if (file) {
                  setLogo(file);
                  const reader = new FileReader();
                  reader.onloadend = () => {
                        setLogoPreview(reader.result as string);
                  };
                  reader.readAsDataURL(file);
            }
      };

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      return (
            <>
                  <div className="w-full bg-base-100 shadow-lg rounded-lg p-8 mb-5">
                        <h2 className="text-2xl font-bold mb-6">Basic Website Data</h2>

                        <div className="grid grid-cols-1 gap-6">

                              <div>
                                    <label htmlFor="siteStatus" className="block label pb-1">Site Status</label>
                                    <span id="siteStatus" className="inline-block bg-black/50 px-4 py-1 text-yellow-300">{getAccountStatus(site?.status as AccountStatusTypes)}</span>
                              </div>


                              <div>
                                    <label htmlFor="sitename" className="label pb-1">Name</label>
                                    <input
                                          id="sitename"
                                          type="text"
                                          placeholder="Website Name"
                                          className="input input-bordered w-full"
                                          defaultValue={site.name}
                                          onChange={(e) => updateSite("name", e.target.value)}
                                    />
                              </div>

                              <div>
                                    <label htmlFor="siteurl" className="label pb-1">URL</label>
                                    <input
                                          id="siteurl"
                                          type="text"
                                          placeholder="Website URL"
                                          className="input w-full"
                                          defaultValue={site.url}
                                          onChange={(e) => updateSite("url", e.target.value)}
                                    />
                              </div>
                              <div>
                                    <label htmlFor="description" className="label pb-1">Description</label>
                                    <textarea
                                          id="description"
                                          placeholder="Website Description"
                                          className="textarea textarea-bordered w-full"
                                          defaultValue={site.description}
                                          onChange={(e) => updateSite("description", e.target.value)}
                                    ></textarea>
                              </div>

                              {(site.logo) && (
                                    <div className="flex flex-col items-center">
                                          <img
                                                src={ site.logo }
                                                alt="Logo preview"
                                                className="max-w-xs h-24 w-24 border rounded-lg shadow"
                                          />
                                    </div>
                              )}

                              { (!site.logo || site.logo === "") && (
                                    <div className="flex justify-center items-center text-gray-500 text-sm italic">
                                          No logo uploaded yet
                                    </div>
                              )}                              

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
