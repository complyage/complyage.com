//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useEffect, useState }                             from "react";
import { Ban, CheckCircle, CircleCheck, CircleX, Globe, User }    from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

import { VerificationTypes }                               from "../../interfaces/models/model.verify";
import { Identity, IdentityRecord }                        from "../../interfaces/verify/identity/identity";
import { AuthMe }                                          from "../../interfaces/auth/me";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import MembersLayout                                        from "../../layouts/MembersLayout";
import SpinnerCircle                                        from "../../components/base/SpinnerCircle";
import InlineAlert                                          from "../../components/base/InlineAlert";

//||------------------------------------------------------------------------------------------------||
//|| Data
//||------------------------------------------------------------------------------------------------||

import { getVerificationType, getVerificationIcon }       from "../../data/getVerificationData";


//||------------------------------------------------------------------------------------------------||
//|| Default Component
//||------------------------------------------------------------------------------------------------||

export default function MembersIdentity() {

      //||------------------------------------------------------------------------------------------------||
      //|| Use State
      //||------------------------------------------------------------------------------------------------||

      const [iden, setIden]                           = useState<Identity>({});
      const [error, setError]                         = useState<any>(null);
      const [loading, setLoading]                     = useState<boolean>(false);

      //||------------------------------------------------------------------------------------------------||
      //|| Load Data
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            if (loading) return;
            const loader = async () => {
                  setLoading(true);
                  const resp = await fetch("/auth/me", { credentials: "include" })
                  if (!resp.ok) {
                        setError(`Failed to load identity data: ${resp.status} ${resp.statusText}`);
                        setLoading(false);
                        return;
                  }
                  const payload = await resp.json();
                  if (!payload || !payload.success) {
                        setError(payload?.message || "Failed to load identity data");
                        setLoading(false);
                        return;
                  }                  
                  setIden(payload.data.identity || {});
                  console.log("Identity Data:", payload.data.identity);
                  setError(null);
                  setLoading(false);
            };
            loader();
      }, [] );      

      //||------------------------------------------------------------------------------------------------||
      //|| identityRow
      //||------------------------------------------------------------------------------------------------||

      function identityRow( verifyType : VerificationTypes, v : IdentityRecord ) : JSX.Element {
            //||------------------------------------------------------------------------------------------------||
            //|| identityRow
            //||------------------------------------------------------------------------------------------------||
            const isVerified  = !!v.verification && v.verification !== "";
            const rowBg       = isVerified ? "bg-black p-4 shadow-lg border-b border-gray-700" : "bg-gray-800 text-gray-400";
            const desc        = getVerificationType(verifyType) || verifyType;
            const Icon        = getVerificationIcon(verifyType) || User;
            //||------------------------------------------------------------------------------------------------||
            //|| identityRow
            //||------------------------------------------------------------------------------------------------||
            return (
                  <React.Fragment key={verifyType}>
                        <tr className={`${rowBg} transition-colors`}>
                              <td className="w-5"><Icon className="w-5 h-5" /></td>
                              <td>
                                    <span className="font-semibold">{desc}</span>
                              </td>
                              <td>
                                    <span className={(!isVerified) ? "opacity-50" : ""}>
                                          <pre className="whitespace-pre-wrap font-mono">
                                                {v.display || <span className="italic text-base-content/40">—</span>}
                                          </pre>
                                    </span>
                              </td>
                              <td className="text-center">
                                    {['FACE', 'IDEN'].includes(verifyType) ? (
                                          <span className="flex flex-col items-center gap-0.5 rounded px-3 py-1 text-sm font-semibold">
                                                <CircleCheck className="w-8 h-8 text-green-400 mb-1" />
                                                Yes
                                          </span>
                                    ) : verifyType == "CRCD" ? (
                                          <span className="flex flex-col items-center gap-0.5 rounded px-3 py-1 text-sm font-semibold">
                                                <Globe className="w-8 h-8 text-blue-400 mb-1" />
                                                Based on location
                                          </span>
                                    ) : (
                                          <span className="flex flex-col items-center gap-0.5 rounded px-3 py-1 text-sm font-semibold">
                                                <CircleX className="w-8 h-8 text-red-400 mb-1" />
                                                No
                                          </span>
                                    )}
                              </td>
                              <td className="text-right"> 
                                    {isVerified ? (
                                          <span className="inline-flex items-center gap-1 rounded px-5 py-2 bg-green-500/80 text-black"><CheckCircle className="w-4 h-4" />Verified</span>
                                    ) : ( 
                                          <button className="btn btn-secondary btn-md">Verify</button> 
                                    )}
                              </td>
                        </tr>
                  </React.Fragment>
            );
      }

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      return (
            <MembersLayout title="Your Verified Identity">
                  <div className="max-w-4xl mx-auto px-2 md:px-0">
                        <div className="mb-5 p-5 text-base leading-loose">
                              <b>We take your privacy seriously.</b> All sensitive data you provide—emails, phone numbers, and verification documents—is encrypted before it ever reaches our servers. <br />
                              Even we can't see your private data. Click <b>"See what we see"</b> to inspect your encrypted entries.
                        </div>
                        <div className="overflow-x-auto">

                              {/* Spinner (loading, no error) */}
                              {loading && !error && (
                                    <SpinnerCircle />
                              )}

                              {/* Error Alert (not loading, has error) */}
                              {!loading && error && (
                                    <InlineAlert isError message={error} />
                              )}

                              {/* Identity Verified/Not Verified & Table (not loading, no error) */}
                              {!loading && !error && (
                                    <>
                                          {
                                                iden.verified ? (
                                                      <div className="mb-5 p-5 bg-base-200 rounded-lg text-base leading-loose shadow text-center">
                                                            Your identity is  <span className="text-green-400">Verified</span> <span className="ml-4">Age: { iden.verifiedAge || "N/A" }</span>
                                                      </div>
                                                ) : (
                                                      <div className="mb-5 p-5 bg-base-200 rounded-lg leading-loose shadow text-center text-red-400 text-2xl">
                                                            <Ban className="inline w-8 h-8 mr-2" />
                                                            <span className="text-red-400">Your Age is Not Verified</span>
                                                            <button className="btn btn-secondary btn-md ml-6">Verify Now</button>
                                                      </div>
                                                )
                                          }
                                          <table className="table w-full bg-base-100 rounded-xl text-lg">
                                                <thead>
                                                      <tr className="bg-base-200 p-4 shadow-lg border-b border-gray-700">
                                                            <th colSpan={2}>Type</th>
                                                            <th>Details</th>
                                                            <th className="text-center">Verifies Age</th>
                                                            <th className="text-right">Status</th>
                                                      </tr>
                                                </thead>
                                                <tbody>
                                                      {iden.idCard          && identityRow("IDEN", iden.idCard)}
                                                      {iden.face            && identityRow("FACE", iden.face)}
                                                      {iden.creditCard      && identityRow("CRCD", iden.creditCard)}
                                                      {iden.email           && identityRow("MAIL", iden.email)}
                                                      {iden.phone           && identityRow("PHNE", iden.phone)}
                                                      {iden.address         && identityRow("ADDR", iden.address)}
                                                </tbody>
                                          </table>

                                    </>
                              )}
                        </div>
                  </div>
            </MembersLayout>
      );


}
