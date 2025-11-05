//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React from "react";
import { CircleCheck, X } from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

interface PrivateCheckListProps {
      loggedIn    : boolean;
      level       : number;
      privateKey  : boolean;
}

//||------------------------------------------------------------------------------------------------||
//|| Default Component
//||------------------------------------------------------------------------------------------------||

export default function PrivateCheckList({ loggedIn, level, privateKey }: PrivateCheckListProps) {

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||
      return (
            <div className="mx-auto mt-4 w-full">
                  <div>

                        {/* single-row: three columns */}
                        <div className="flex gap-3 items-stretch">

                              {/* column 1 */}
                              <div className="flex-1 min-w-0 bg-black/40 rounded-lg p-4 flex flex-col items-center text-center">
                                    <div className="text-sm text-yellow-400 font-medium w-4/5 mb-2 border-dashed border-b pb-2 border-gray-500">Logged In</div>
                                    <div className="flex items-center gap-2">
                                          {loggedIn ? (
                                                <>
                                                      <span className="flex items-center justify-center w-5 h-5">
                                                            <CircleCheck className="w-5 h-5 text-green-300" />
                                                      </span>
                                                      <span className="text-sm font-semibold text-green-200 leading-none">Yes</span>
                                                </>
                                          ) : (
                                                <>
                                                      <span className="flex items-center justify-center w-5 h-5">
                                                            <X className="w-5 h-5 text-red-400" />
                                                      </span>
                                                      <span className="text-sm font-semibold text-red-300 leading-none">No</span>
                                                </>
                                          )}
                                    </div>
                              </div>

                              {/* column 2 */}
                              <div className="flex-1 min-w-0 bg-black/40 rounded-lg p-4 flex flex-col items-center text-center">
                                    <div className="text-sm text-yellow-400 font-medium w-4/5 mb-2 border-dashed border-b pb-2 border-gray-500">Managed By</div>
                                    <div className="flex items-center gap-2">
                                          {level < 2 ? (
                                                <>
                                                      <span className="flex items-center justify-center w-5 h-5">
                                                            <CircleCheck className="w-5 h-5 text-green-300" />
                                                      </span>
                                                      <span className="text-sm font-semibold text-green-200 leading-none">ComplyAge</span>
                                                </>
                                          ) : (
                                                <>
                                                      <span className="flex items-center justify-center w-5 h-5">
                                                            <X className="w-5 h-5 text-red-400" />
                                                      </span>
                                                      <span className="text-sm font-semibold text-red-300 leading-none">You</span>
                                                </>
                                          )}
                                    </div>
                              </div>

                              {/* column 3 */}
                              <div className="flex-1 min-w-0 bg-black/40 rounded-lg p-4 flex flex-col items-center text-center">
                                    <div className="text-sm text-yellow-400 font-medium w-4/5 mb-2 border-dashed border-b pb-2 border-gray-500">Private Key</div>
                                    <div className="flex items-center gap-2">
                                          {privateKey ? (
                                                <>
                                                      <span className="flex items-center justify-center w-5 h-5">
                                                            <CircleCheck className="w-5 h-5 text-green-300" />
                                                      </span>
                                                      <span className="text-sm font-semibold text-green-200 leading-none">{level > 1 ? "Temporarily Stored" : "Stored"}</span>
                                                </>
                                          ) : (
                                                <>
                                                      <span className="flex items-center justify-center w-5 h-5">
                                                            <X className="w-5 h-5 text-red-400" />
                                                      </span>
                                                      <span className="text-sm font-semibold text-red-300 leading-none">Requires Entry</span>
                                                </>
                                          )}
                                    </div>
                              </div>

                        </div>
                  </div>

            </div>
      );
}
