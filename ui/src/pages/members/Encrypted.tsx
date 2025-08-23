//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useEffect, useState }                 from "react";
import MembersLayout                                  from "../../layouts/MembersLayout";
import { ModelVerification }                          from "../../interfaces/models/model.verify";
import { getVerificationType }                        from "../../data/getVerificationData";

//||------------------------------------------------------------------------------------------------||
//|| Default Component
//||------------------------------------------------------------------------------------------------||

export default function Encrypted() {

      //||------------------------------------------------------------------------------------------------||
      //|| Const
      //||------------------------------------------------------------------------------------------------||

      const [verifications, setVerifications]         = useState<ModelVerification[]>([]);
      const [statusFilter, setStatusFilter]           = useState<string>("ALL");
      const [page, setPage]                           = useState<number>(1);
      const [pageSize]                                = useState<number>(10);
      const [total, setTotal]                         = useState<number>(0);
      const [loading, setLoading]                     = useState<boolean>(true);
      const [showEncrypted, setShowEncrypted]         = useState<string[]>([]);

      //||------------------------------------------------------------------------------------------------||
      //|| Show Encrypted
      //||------------------------------------------------------------------------------------------------||

      const toggleEncrypted = (type: string) => {
            setShowEncrypted(prev => 
                  prev.includes(type) ? prev.filter(t => t !== type) : [...prev, type]
            );
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Fetch Verifications (Now Uses Identity)
      //||------------------------------------------------------------------------------------------------||

      const fetchVerifications = async () => {
            setLoading(true);
            try {
                  const res = await fetch(`/auth/me`, { credentials: "include" });
                  const json = await res.json();
                  console.log("JSON", json);

                  if (json.success) {
                        let identity = json.data.identity;

                        // ✅ If identity is a string, parse it
                        if (typeof identity === "string") {
                              try {
                                    identity = JSON.parse(identity);
                              } catch (e) {
                                    console.error("Failed to parse identity JSON:", identity);
                                    identity = {};
                              }
                        }
                        console.log("IDENTITY:", identity);
                        // ✅ Extract approved list
                        let approved: string[] = [];
                        const rawApproved = identity?.approved;
                        if (Array.isArray(rawApproved)) {
                              approved = rawApproved;
                        } else if (typeof rawApproved === "string") {
                              try {
                                    approved = JSON.parse(rawApproved);
                              } catch {
                                    approved = [];
                              }
                        }

                        console.log("APPROVED CODES:", approved);

                        // ✅ Mark verification types as complete
                        const baseVerifications : AnalyserNode = [
                              { id: 1, type: "MAIL", encrypt: identity?.email?.data || "", data: identity?.email?.display || "", meta: "", status: "PEND", complete: false },
                              { id: 2, type: "UAGE", encrypt: identity?.age?.data || "",data: identity?.age?.display || "", meta: "", status: "PEND", complete: false },
                              { id: 3, type: "PHNE", encrypt: identity?.phone?.data || "",data: identity?.phone?.display || "", meta: "", status: "PEND", complete: false },
                              { id: 4, type: "ADDR", encrypt: identity?.address?.data || "",data: identity?.address?.display || "", meta: "", status: "PEND", complete: false },
                              { id: 5, type: "CRCD", encrypt: identity?.creditCard?.data || "",data: identity?.creditCard?.display || "", meta: "", status: "PEND", complete: false },
                              { id: 6, type: "PROF", encrypt: identity?.usernames?.data || "",data: identity?.usernames ? Object.keys(identity.usernames).length + " usernames" : "", meta: "", status: "PEND", complete: false }
                        ];

                        const updatedVerifications = baseVerifications.map(v => ({
                              ...v,
                              complete: approved.includes(v.type)
                        }));

                        setVerifications(updatedVerifications);
                        setTotal(updatedVerifications.length);
                  }
            } catch (err) {
                  console.error("❌ Failed to load identity:", err);
            } finally {
                  setLoading(false);
            }
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Effects
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            fetchVerifications();
      }, [page, statusFilter]);

      const totalPages = Math.ceil(total / pageSize);

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      return (
            <MembersLayout title="Your Verified Identity">
                  <>


                        {/* Table */}
                        <div className="overflow-x-auto">
                              <p className="mb-5 p-5 bg-black/10 text-md text-center leading-loose"><b>We take your privacy seriously</b>, which is why all sensitive information you provide, such as emails, phone numbers, and verification documents is encrypted before it ever reaches our servers.
                              This means that even we cannot see the contents of your sensitive data. However, you can choose to view the encrypted data we store for you by clicking the "See what we see" button below each entry.
                              </p>
                              <table className="table w-full bg-base-100 shadow rounded-lg">
                                    <thead>
                                          <tr>
                                                <th>Type</th>
                                                <th>Display</th>
                                                <th>Status</th>
                                                <th></th>
                                          </tr>
                                    </thead>
                                    <tbody>
                                          {loading ? (
                                                <tr>
                                                      <td colSpan={3} className="text-center p-4">
                                                            Loading...
                                                      </td>
                                                </tr>
                                          ) : verifications.length === 0 ? (
                                                <tr>
                                                      <td colSpan={3} className="text-center p-4">
                                                            No stored identity data.
                                                      </td>
                                                </tr>
                                          ) : (
                                                verifications.map((v) => {
                                                      return (                                                            
                                                            <>
                                                                  <tr key={v.id} className="text-xs">
                                                                        <td className={(!v.complete) ? "opacity-20" : ""}>{ getVerificationType(v.type) }</td>
                                                                        <td className={(!v.complete) ? "opacity-20" : ""}><pre className="whitespace-pre-wrap text-xs">{v.data}</pre></td>
                                                                        <td className={(!v.complete) ? "opacity-20" : ""}>{v.complete ? "✅ Verified" : "❌ Pending"}</td>
                                                                        <td className="text-right">
                                                                        { v.encrypt ?  ( 
                                                                              <button className="btn btn-primary cursor-pointer" onClick={ () => { toggleEncrypted(v.type) } } >See what we see</button>
                                                                        ) : (
                                                                              <button className="btn btn-secondary">Verify</button>
                                                                        )}
                                                                        </td>
                                                                  </tr>
                                                                  { v.encrypt && showEncrypted.includes(v.type) && (
                                                                        <tr key={`${v.id}-encrypted`}>
                                                                              <td colSpan={4}>
                                                                                    <textarea readOnly={true} className="bg-black h-24 w-full whitespace-pre-wrap text-yellow-500 p-3 text-xs">{v.encrypt}</textarea>
                                                                              </td>
                                                                        </tr>
                                                                  )}
                                                            </>
                                                )})
                                          )}
                                    </tbody>
                              </table>
                        </div>

                        {/* Paging Controls */}
                        {totalPages > 1 && (
                              <div className="flex justify-center mt-4 gap-2">
                                    <button disabled={page === 1} className="btn btn-sm" onClick={() => setPage(page - 1)}>Prev</button>
                                    <span className="px-2">Page {page} of {totalPages}</span>
                                    <button disabled={page === totalPages} className="btn btn-sm" onClick={() => setPage(page + 1)}>Next</button>
                              </div>
                        )}
                  </>
            </MembersLayout>
      );
}
