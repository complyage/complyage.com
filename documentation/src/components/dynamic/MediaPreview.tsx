//||------------------------------------------------------------------------------------------------||
//|| Media Preview
//|| Displays a preview of uploaded media with options to reset or close
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| React
      //||------------------------------------------------------------------------------------------------||
      
      import React                                                from "react";
      import { IdCard, XIcon }                                    from "lucide-react";

      //||------------------------------------------------------------------------------------------------||
      //|| Props
      //||------------------------------------------------------------------------------------------------||

      import { VerificationMedia }                                from "../../interfaces/verify/id/process";

      //||------------------------------------------------------------------------------------------------||
      //|| Props
      //||------------------------------------------------------------------------------------------------||

      interface MediaPreviewProps {
            media    : VerificationMedia;
            onEdit?  : () => void;
      }

      //||------------------------------------------------------------------------------------------------||
      //|| React
      //||------------------------------------------------------------------------------------------------||

      const MediaPreview: React.FC<MediaPreviewProps> = ({ media, onEdit }) => (
            <div className="flex flex-col w-full relative items-center justify-center mx-auto g">

                  <div className={`group w-full border-4 border-dashed border-gray-400 rounded-2xl pt-3 pb-3 px-6 font-bold text-center flex flex-col items-center justify-center transition-colors bg-black/20`} style={{ minHeight: 300 }}>
                        <div className="w-full min-h-96 flex flex-col items-center mb-2">
                              {media?.blob ? (
                                    <>
                                          <img
                                                src={`data:${media.mime || "image/jpeg"};base64,${media.blob}`}
                                                alt="Preview"
                                                className="max-w-full w-256 group-hover:text-blue-600"
                                          />
                                          {onEdit && (
                                                <button
                                                      type="button"
                                                      onClick={onEdit}
                                                      className="mt-5 px-4 py-2 bg-red-500 text-white rounded hover:bg-red-600 transition-colors"
                                                >
                                                      Reset
                                                </button>
                                          )}
                                    </>
                              ) : (
                                    <>
                                          <IdCard className="max-w-full h-96 w-256 group-hover:text-blue-600" />
                                          <p className="mb-2 text-gray-400 group-hover:text-blue-400">
                                                No file uploaded yet
                                          </p>
                                    </>
                              )}
                        </div>
                  </div>
            </div>
      );

      //||------------------------------------------------------------------------------------------------||
      //|| Export
      //||------------------------------------------------------------------------------------------------||

      export default MediaPreview;
