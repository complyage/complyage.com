//||------------------------------------------------------------------------------------------------||
//|| BasicDataSection
//|| Component for editing basic website info including logo upload
//||------------------------------------------------------------------------------------------------||

import React, { useState, useEffect, useRef }               from "react";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

import { Site }                                             from "../../interfaces/model.sites";
import { getAccountStatus }                                 from "../../interfaces/status.account";      

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface BasicDataProps {
      data        : Site;
      updateField : (field : string, value : string) => void;
      cacheBust   : string;
}

//||------------------------------------------------------------------------------------------------||
//|| Default Component
//||------------------------------------------------------------------------------------------------||

export default function BasicDataSection({ updateField, data, cacheBust }: BasicDataProps) {

      const [name, setName]               = useState(data?.name || "");
      const [url, setUrl]                 = useState(data?.url || "");
      const [description, setDescription] = useState(data?.description || "");
      const [logo, setLogo]               = useState<File | null>(null);
      const [status, setStatus]           = useState<"Active" | "Pending">("Pending");
      const fileInputRef                  = useRef<HTMLInputElement | null>(null);
      const [logoExists, setLogoExists]   = useState(false);
      const [cb, setCb]                   = useState<string>(cacheBust || new Date().toISOString());
      const [logoPreview, setLogoPreview] = useState<string | null>(null);
      const [logoHash, setLogoHash]       = useState<string>(data?.logoHash || "");
      const [logoMissing, setLogoMissing] = useState<boolean>(data?.logoMissing || true);


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
            setLogoHash(data.logoHash || "");
            setLogoMissing(data.logoMissing);
            setCb(cacheBust);
      }, [data, cacheBust]);

      //||------------------------------------------------------------------------------------------------||
      //|| Upload File
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            if (logoHash != "") {
               const url = `/media/sites/logos/${logoHash}.webp?cachebust=${cb}`;
               fetch(url, { method: "HEAD", cache: "no-store" })
                  .then((res) => setLogoExists(res.ok))
                  .catch(() => setLogoExists(false));
            }
      }, [data, cb]);

      //||------------------------------------------------------------------------------------------------||
      //|| Handle File Select
      //||------------------------------------------------------------------------------------------------||

      const handleUploadChange = (e: React.ChangeEvent<HTMLInputElement>) => {
            const file = e.target.files?.[0];
            if (file) {
                  setLogo(file);
                  updateField("forceUpdate", new Date().toISOString());
                  updateField("logo", file);
      
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
                  updateField("forceUpdate", new Date().toISOString());
      
                  const reader = new FileReader();
                  reader.onloadend = () => {
                        setLogoPreview(reader.result as string);
                  };
                  reader.readAsDataURL(file);
            }
      };

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
                        <h2 className="text-2xl font-bold mb-6">Basic Website Data</h2>

                        <div className="grid grid-cols-1 gap-6">

                              <div>
                                    <label htmlFor="siteStatus" className="block label pb-1">Site Status</label>
                                    <span id="siteStatus" className="inline-block bg-black/50 px-4 py-1 text-yellow-300">{getAccountStatus(data?.status)}</span>
                              </div>


                              <div>
                                    <label htmlFor="sitename" className="label pb-1">Name</label>
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
                                    <label htmlFor="siteurl" className="label pb-1">URL</label>
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
                                    <label htmlFor="description" className="label pb-1">Description</label>
                                    <textarea
                                          id="description"
                                          placeholder="Website Description"
                                          className="textarea textarea-bordered w-full"
                                          defaultValue={description}
                                          onChange={(e) => updateField("description", e.target.value)}
                                    ></textarea>
                              </div>

                              {(logoPreview || !logoMissing) && (
                                    <div className="flex flex-col items-center">
                                          <img
                                                key={logoPreview || cb}
                                                src={ logoPreview ? logoPreview : `${import.meta.env.VITE_COMPLYAGE_MINIO_URL}/${import.meta.env.VITE_MINIO_BUCKET}/sites/logos/${logoHash}.webp?cachebust=${cb}` }
                                                alt="Logo preview"
                                                className="max-w-xs h-24 w-24 border rounded-lg shadow"
                                          />
                                    </div>
                              )}

                              {!(logoPreview || logoMissing) && (
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
