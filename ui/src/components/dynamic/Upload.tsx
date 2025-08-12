//||------------------------------------------------------------------------------------------------||
//|| Upload.tsx
//|| Drag-and-drop image upload with preview, onHandler callback, and reset feature
//||------------------------------------------------------------------------------------------------||

import React, { useState, useRef, DragEvent, ChangeEvent }        from "react";
import { XIcon }                                                  from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface UploadProps {
      onUpload: (file: File | Blob) => void;
      onReset?: () => void;
      onClose?: () => void;
}

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function Upload({ onUpload, onReset, onClose }: UploadProps) {

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
      //|| Handle Drag Events
      //||------------------------------------------------------------------------------------------------||

      const onDragOver = (e: DragEvent<HTMLDivElement>) => {
            e.preventDefault();
            e.stopPropagation();
      };

      const onDrop = (e: DragEvent<HTMLDivElement>) => {
            e.preventDefault();
            e.stopPropagation();
            const file = e.dataTransfer.files?.[0];
            if (file) handleFile(file);
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Reset Selection
      //||------------------------------------------------------------------------------------------------||
      const reset = () => {
            setPreview(null);
            if (inputRef.current) inputRef.current.value = "";
            if (onReset) onReset();
      };

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      return (
            <div className="flex flex-col items-center justify-center w-full max-w-md mx-auto">
                  <button onClick={onClose} className="absolute top-20 right-4 rounded-xl bg-black/90 hover:text-gray-700 transition p-2"><XIcon className="w-8 h-8 text-white" /></button>                  


                  {preview ? (
                        <div className="mt-4 flex flex-col items-center">
                              <img
                                    src={preview}
                                    alt="Preview"
                                    className="max-w-full max-h-64 rounded-lg shadow mb-3"
                              />
                              <button
                                    type="button"
                                    onClick={reset}
                                    className="px-4 py-2 bg-red-500 text-white rounded hover:bg-red-600 transition-colors mb-3"
                              >Reset</button>
                        </div>
                  ) : (
                        <div className="mt-4 flex flex-col items-center">
                              <img
                                    src="/img/id.png"
                                    alt="Preview"
                                    className="max-w-full max-h-64 rounded-lg shadow mb-3"
                              />
                        </div>                        
                  )}

                  <label
                        htmlFor="fileInput"
                        onDragOver={onDragOver}
                        onDrop={onDrop}
                        className="w-full border-2 border-dashed border-gray-400 rounded-lg p-6 text-center cursor-pointer hover:border-blue-500 transition-colors"
                  >
                        <p className="mb-2">Drag & Drop an image here or click to select</p>
                        <input
                              id="fileInput"
                              ref={inputRef}
                              type="file"
                              accept="image/*"
                              onChange={onChange}
                              className="hidden"
                        />
                  </label>

                  <button className="btn btn-secondary mt-5 text-2xl px-5" onClick={ () => { inputRef.current.click() } }>Upload a File</button>

            </div>
      );
}
