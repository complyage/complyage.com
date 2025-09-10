//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { JSX, useEffect, useState, useRef }                      from "react";
import { CheckCircle, CircleCheck, CircleX, Globe, User, ArrowUpRight } from "lucide-react";
import { useNavigate }                                                  from "react-router-dom";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

import { VerificationTypes }                                            from "../../interfaces/models/model.verify";
import { IdentityRecord }                                               from "../../interfaces/verify/identity/identity";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import MembersLayout                                                    from "../../layouts/MembersLayout";
import SpinnerCircle                                                    from "../../components/base/SpinnerCircle";
import InlineAlert                                                      from "../../components/base/InlineAlert";

//||------------------------------------------------------------------------------------------------||
//|| Dashboard Data
//||------------------------------------------------------------------------------------------------||

import { DashboardData }                                                from "../../interfaces/members/dashboard";

//||------------------------------------------------------------------------------------------------||
//|| Call
//||------------------------------------------------------------------------------------------------||

import Call                                                             from "../../classes/call";

//||------------------------------------------------------------------------------------------------||
//|| Data
//||------------------------------------------------------------------------------------------------||

import { getVerificationType, getVerificationIcon }                    from "../../data/getVerificationData";


//||------------------------------------------------------------------------------------------------||
//|| Default Component
//||------------------------------------------------------------------------------------------------||

export default function MembersIdentity() {

      //||------------------------------------------------------------------------------------------------||
      //|| Navigate
      //||------------------------------------------------------------------------------------------------||

      const navigate                                                    = useNavigate();

      //||------------------------------------------------------------------------------------------------||
      //|| Default Data
      //||------------------------------------------------------------------------------------------------||

      const defaultData: DashboardData = {
            isVerified        : false,
            verifiedAge       : 0,
            minimumType       : "IDEN",
            ipAddress         : "---",
            location          : {
                  city        : "---",
                  region      : "---",
                  country     : "---",
                  latitude    : 0,
                  longitude   : 0
            },
            zone              : {
                  laws        : "",
                  requirements: ["IDEN", "FACE", "CRCD"],
                  effective   : "---",
                  minAge      : 0
            },
            identity          : {}
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Use State
      //||------------------------------------------------------------------------------------------------||

      const [ data, setData ]                                           = useState<DashboardData>(defaultData);
      const [ error, setError ]                                         = useState<any>(null);
      const [ loading, setLoading ]                                     = useState<boolean>(false);
      const didFetchRef                                                 = useRef(false);

      //||------------------------------------------------------------------------------------------------||
      //|| Get Statuses
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            if (didFetchRef.current) return;
            didFetchRef.current = true;
            setLoading(true);
            const fetchData = (async() => {
                  console.log("======> Fetching Dashboard Data");
                  const chirp = new Call("/user/dashboard", {});
                  chirp.debug  = true;
                  chirp.method = "GET";
                  await chirp.execute();
                  if (!chirp.ok()) {
                        setError(chirp.error());
                        setLoading(false);
                        return;
                  }
                  if (!chirp.responsePayload.data)  {
                        setError("No data received");
                        setLoading(false);
                        return;
                  }
                  console.log("LOADED DATA", chirp.responsePayload.data);
                  setData(chirp.responsePayload.data as DashboardData);
                  setLoading(false);
            });
            fetchData();
      }, []);

      //||------------------------------------------------------------------------------------------------||
      //|| identityRow
      //||------------------------------------------------------------------------------------------------||

      function identityRow(verifyType: VerificationTypes, v?: IdentityRecord) : JSX.Element {
            const isVerified  = !!(v && typeof v.verified !== "undefined" && v.verified === true);
            const rowBg       = isVerified ? "bg-black p-4 shadow-lg border-b border-gray-700" : "bg-gray-800 text-gray-400";
            const desc        = getVerificationType(verifyType) || verifyType;
            const Icon        = getVerificationIcon(verifyType) || User;

            return (
                  <React.Fragment key={verifyType}>
                        <tr className={`${rowBg} transition-colors`}>
                              <td className="w-5"><Icon className="w-5 h-5" /></td>
                              <td>
                                    <span className="font-semibold">{desc}</span>
                              </td>
                              <td>
                                    <span className={!isVerified ? "opacity-50" : ""}>
                                          <pre className="whitespace-pre-wrap font-mono">
                                                {v?.display || <span className="italic text-base-content/40">—</span>}
                                          </pre>
                                    </span>
                              </td>
                              <td className="text-center">
                                    {["FACE","IDEN"].includes(verifyType) ? (
                                          <span className="flex flex-col items-center gap-0.5 rounded px-3 py-1 text-sm font-semibold">
                                                <CircleCheck className="w-8 h-8 text-green-400 mb-1" />
                                          </span>
                                    ) : verifyType === "CRCD" ? (
                                          <span className="flex flex-col items-center gap-0.5 rounded px-3 py-1 text-sm font-semibold">
                                                <Globe className="w-8 h-8 text-blue-400 mb-1" />
                                          </span>
                                    ) : (
                                          <span className="flex flex-col items-center gap-0.5 rounded px-3 py-1 text-sm font-semibold">
                                                <CircleX className="w-8 h-8 text-red-400 mb-1" />
                                          </span>
                                    )}
                              </td>
                              <td className="text-center">
                                    {isVerified ? (
                                          <span className="inline-flex items-center gap-1 rounded-lg px-5 py-2 border-2 bg-green-400/20 border-green-400 text-green-400">
                                                <CheckCircle className="w-6 h-6" />
                                          </span>
                                    ) : (
                                          <button
                                                onClick={() => navigate(`/verification/init?type=${encodeURIComponent(verifyType)}`)}
                                                className="btn btn-secondary btn-md"
                                          >
                                                Verify <ArrowUpRight className="w-4 h-4" />
                                          </button>
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
                        <div className="overflow-x-auto rounded-lg">

                              {/* Spinner (loading, no error) */}
                              {loading && !error && (
                                    <div className="flex justify-center items-center h-64">
                                    <SpinnerCircle />
                                    </div>
                              )}

                              {/* Error Alert (not loading, has error) */}
                              {!loading && error && (
                                    <InlineAlert isError message={error} />
                              )}

                              {/* Identity Verified/Not Verified & Table (not loading, no error) */}
                              {!loading && !error && (
                                    <>
                                          { /* <textarea
                                                defaultValue={JSON.stringify(data, null, 2)}
                                                className="w-full h-64 bg-gray-800 text-gray-200 p-4 rounded-lg font-mono mb-5"
                                                readOnly
                                          /> */ }

                                          <table className="table w-full bg-base-100 rounded-xl text-lg">
                                                <thead>
                                                      <tr className="bg-base-200 p-4 shadow-lg border-b border-gray-700">
                                                            <th colSpan={2}>Type</th>
                                                            <th>Details</th>
                                                            <th className="text-center">Verifies Age</th>
                                                            <th className="text-center">Status</th>
                                                      </tr>
                                                </thead>
                                                <tbody>
                                                      {identityRow("IDEN", data.identity.IDEN)}
                                                      {identityRow("FACE", data.identity.FACE)}
                                                      {identityRow("CRCD", data.identity.CRCD)}
                                                      {identityRow("MAIL", data.identity.MAIL)}
                                                      {identityRow("PHNE", data.identity.PHNE)}
                                                      {identityRow("ADDR", data.identity.ADDR)}
                                                </tbody>
                                          </table>

                                    </>
                              )}
                        </div>
                  </div>
            </MembersLayout>
      );
}
