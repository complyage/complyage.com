//||------------------------------------------------------------------------------------------------||
//|| Upload.tsx
//|| Drag-and-drop image upload with preview, onHandler callback, and reset feature
//||------------------------------------------------------------------------------------------------||

import React, { useState, useRef, DragEvent, ChangeEvent }        from "react";
import { XIcon, IdCard, CreditCard, Smile, Image }                from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

import { Media }                                                  from "../../../interfaces/base/media";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface UploadProps {
      which             : "front" | "back" | "selfie" | "other";
      onUpload          : (file: File | Blob) => void;
      getUpload?        : (which: string) => Media;
      onClose?          : () => void;
}

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function Upload({ which, onUpload, getUpload, onClose }: UploadProps) {

      //||------------------------------------------------------------------------------------------------||
      //|| Var
      //||------------------------------------------------------------------------------------------------||

      const [preview, setPreview]         = useState<string | null>(null);
      const inputRef                      = useRef<HTMLInputElement>(null);

      //||------------------------------------------------------------------------------------------------||
      //|| Handle File Selection
      //||------------------------------------------------------------------------------------------------||

      const handleFile = (file: File) => {
            if (file && file.type.startsWith("image/")) {
                  const reader = new FileReader();
                  reader.onloadend = () => {
                        const result = reader.result as string;
                        setPreview(result);
                        if (onUpload) onUpload(file);
                  };
                  reader.readAsDataURL(file);
            } else {
                  alert("Only image files are allowed.");
            }
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Handle File Input Change
      //||------------------------------------------------------------------------------------------------||

      const onChange = (e: ChangeEvent<HTMLInputElement>) => {
            const file = e.target.files?.[0];
            if (file) handleFile(file);
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Handle Drag 
      //||------------------------------------------------------------------------------------------------||

      const onDragOver = (e: DragEvent<HTMLLabelElement>) => {
            e.preventDefault();
            e.stopPropagation();
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Handle Drop
      //||------------------------------------------------------------------------------------------------||

      const onDrop = (e: DragEvent<HTMLLabelElement>) => {
            e.preventDefault();
            e.stopPropagation();
            const file = e.dataTransfer.files?.[0];
            if (file) handleFile(file);
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Reset Selection
      //||------------------------------------------------------------------------------------------------||
      
      const handleReset = (e) => {
            e.preventDefault();
            e.stopPropagation();
            setPreview(null);
            if (inputRef.current) inputRef.current.value = "";
            if (onUpload) onUpload(null);

      };

      //||------------------------------------------------------------------------------------------------||
      //|| Upload Icon
      //||------------------------------------------------------------------------------------------------||

      const uploadIcon = (which: string) => {
            switch (which) {
                  case "front":
                        return <IdCard className="max-w-full h-96 w-256 group-hover:text-blue-600" />
                  case "back":
                        return <CreditCard className="max-w-full h-96 w-256 group-hover:text-blue-600" />
                  case "selfie":
                        return <Smile className="max-w-full h-96 w-256 group-hover:text-blue-600" />
                  default:
                        return <Image className="max-w-full h-96 w-256 group-hover:text-blue-600" />
            }
      }

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      return (
            <div className="flex flex-col w-full relative items-center justify-center mx-auto">

                  <button onClick={onClose} className="absolute top-2 right-2 rounded-xl bg-black/90 hover:text-gray-700 transition p-2 cursor-pointer">
                        <XIcon className="w-8 h-8 text-white" />
                  </button>

                  <label
                        htmlFor="fileInput"
                        onDragOver={onDragOver}
                        onDrop={onDrop}
                        className={`group w-full border-4 border-dashed border-gray-400 rounded-2xl pt-3 pb-3 px-6 font-bold text-center cursor-pointer flex flex-col items-center justify-center transition-colors
                              ${preview ? "bg-black/20" : "hover:border-blue-500"}
                        `}
                        style={{ minHeight: 300 }}
                  >
                        {/* Image preview or default image */}
                        <div className="w-full flex flex-col items-center mb-2">
                              {preview ? (
                                    <>
                                          <img
                                                src={preview}
                                                alt="Preview"
                                                className="max-w-full h-96 w-256 group-hover:text-blue-600"
                                          />

                                          <button
                                                type="button"
                                                onClick={(e) => { handleReset(e) }}
                                                className="mt-5 px-4 py-2 bg-red-500 text-white rounded hover:bg-red-600 transition-colors"
                                          >
                                                Reset
                                          </button>
                                    </>
                              ) : (
                                    <>
                                          { uploadIcon(which) }
                                          <p className="mb-2 text-gray-400 group-hover:text-blue-400">
                                                Drag & Drop an image here or click to select
                                          </p>
                                    </>
                                    
                              )}
                        </div>

                        <input
                              id="fileInput"
                              ref={inputRef}
                              type="file"
                              accept="image/*"
                              onChange={onChange}
                              className="hidden"
                        />
                  </label>
            </div>
      );


}
