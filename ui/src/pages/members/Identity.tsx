//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useEffect, useState } from "react";
import MembersLayout                  from "../../layouts/MembersLayout";
import { getVerificationType }        from "../../data/getVerificationData";
import { Mail, Phone, CreditCard, MapPin, User, IdCard, Smile, CheckCircle, XCircle, ChevronDown, ChevronRight } from "lucide-react";

// Types for new Identity structure
interface IdentityRecord {
      data?: string;
      display?: string;
      [key: string]: any;
}
interface Identity {
      id?: number;
      email?: IdentityRecord;
      age?: IdentityRecord;
      phone?: IdentityRecord;
      address?: IdentityRecord;
      creditCard?: IdentityRecord;
      face?: IdentityRecord;
      idCard?: IdentityRecord;
      usernames?: { [key: string]: any };
      approved?: string[];
      verified?: boolean;
      verifiedAge?: number;
}

interface ModelVerification {
      id: number;
      type: string;
      encrypt: string;
      data: string;
      meta: string;
      status: string;
      complete: boolean;
}

// Icon mapping for each verification type
const TYPE_ICON: Record<string, JSX.Element> = {
      "IDEN":  <IdCard className="w-10 h-10 text-blue-400" />, // ID card
      "CRCD":  <CreditCard className="w-8 h-8 text-pink-500" />,
      "FACE":  <Smile className="w-8 h-8 text-yellow-400" />,
      "MAIL":  <Mail className="w-8 h-8 text-purple-500" />,
      "PHNE":  <Phone className="w-8 h-8 text-emerald-400" />,
      "ADDR":  <MapPin className="w-8 h-8 text-red-400" />,
      "AGE":   <User className="w-8 h-8 text-cyan-400" />,
      "PROF":  <User className="w-8 h-8 text-gray-400" />,
};

const TYPE_DESC: Record<string, string> = {
      "IDEN": "Government-issued ID",
      "CRCD": "Credit Card",
      "FACE": "Selfie / Facial",
      "MAIL": "Email Address",
      "PHNE": "Phone Number",
      "ADDR": "Mailing Address",
      "AGE":  "Date of Birth",
      "PROF": "Profile / Usernames",
};

//||------------------------------------------------------------------------------------------------||
//|| Default Component
//||------------------------------------------------------------------------------------------------||

export default function Identity() {

      const [verifications, setVerifications]         = useState<ModelVerification[]>([]);
      const [page, setPage]                           = useState<number>(1);
      const [pageSize]                                = useState<number>(10);
      const [total, setTotal]                         = useState<number>(0);
      const [loading, setLoading]                     = useState<boolean>(true);
      const [showEncrypted, setShowEncrypted]         = useState<string[]>([]);

      // Animated caret for encrypted toggle
      function Caret({ open }: { open: boolean }) {
            return open
                  ? <ChevronDown className="inline w-4 h-4 transition-transform" />
                  : <ChevronRight className="inline w-4 h-4 transition-transform" />;
      }

      const toggleEncrypted = (type: string) => {
            setShowEncrypted(prev =>
                  prev.includes(type) ? prev.filter(t => t !== type) : [...prev, type]
            );
      };

      // Fetch Verifications
      const fetchVerifications = async () => {
            setLoading(true);
            try {
                  const res = await fetch(`/auth/me`, { credentials: "include" });
                  const json = await res.json();

                  if (json.success) {
                        let identity: Identity = json.data.identity;
                        if (typeof identity === "string") {
                              try { identity = JSON.parse(identity); } catch { identity = {}; }
                        }
                        let approved: string[] = [];
                        const rawApproved = identity?.approved;
                        if (Array.isArray(rawApproved)) approved = rawApproved;
                        else if (typeof rawApproved === "string") {
                              try { approved = JSON.parse(rawApproved); } catch { approved = []; }
                        }

                        const baseVerifications: ModelVerification[] = [
                              { id: 1, type: "IDEN", encrypt: identity?.idCard?.data || "",      data: identity?.idCard?.display || "",     meta: "", status: "", complete: false },
                              { id: 2, type: "CRCD", encrypt: identity?.creditCard?.data || "",  data: identity?.creditCard?.display || "", meta: "", status: "", complete: false },
                              { id: 3, type: "FACE", encrypt: identity?.face?.data || "",        data: identity?.face?.display || "",       meta: "", status: "", complete: false },
                              { id: 4, type: "MAIL", encrypt: identity?.email?.data || "",       data: identity?.email?.display || "",      meta: "", status: "", complete: false },
                              { id: 5, type: "PHNE", encrypt: identity?.phone?.data || "",       data: identity?.phone?.display || "",      meta: "", status: "", complete: false },
                              { id: 6, type: "ADDR", encrypt: identity?.address?.data || "",     data: identity?.address?.display || "",    meta: "", status: "", complete: false },
                              { id: 7, type: "AGE",  encrypt: identity?.age?.data || "",         data: identity?.age?.display || "",        meta: "", status: "", complete: false },
                              { id: 8, type: "PROF", encrypt: "",                                data: identity?.usernames ? Object.keys(identity.usernames).length + " usernames" : "", meta: "", status: "PEND", complete: false }
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

      useEffect(() => { fetchVerifications(); }, []);

      const totalPages = Math.ceil(total / pageSize);

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      return (
            <MembersLayout title="Your Verified Identity">
                  <div className="max-w-4xl mx-auto px-2 md:px-0">
                        <div className="mb-5 p-5 bg-base-200 rounded-lg text-base leading-loose shadow">
                              <b>We take your privacy seriously.</b> All sensitive data you provide—emails, phone numbers, and verification documents—is encrypted before it ever reaches our servers. <br />
                              Even we can't see your private data. Click <b>"See what we see"</b> to inspect your encrypted entries.
                        </div>
                        <div className="overflow-x-auto">
                              <table className="table w-full bg-base-100 rounded-xl text-lg">
                                    <tbody>
                                          {loading ? (
                                                <tr>
                                                      <td colSpan={5} className="text-center p-6 text-base-content/60">
                                                            Loading identity records…
                                                      </td>
                                                </tr>
                                          ) : verifications.length === 0 ? (
                                                <tr>
                                                      <td colSpan={5} className="text-center p-6 text-base-content/60">
                                                            No stored identity data.
                                                      </td>
                                                </tr>
                                          ) : (
                                                verifications.map((v, i) => {
                                                      const rowBg = v.complete ? "bg-black p-4 shadow-lg border-b border-gray-700" : "opacity-30";
                                                      const desc = TYPE_DESC[v.type] || v.type;
                                                      const icon = TYPE_ICON[v.type] || <User className="w-5 h-5" />;
                                                      const statusChip = v.complete
                                                            ? <span className="inline-flex items-center gap-1 rounded bg-green-400 px-5 py-2 text-black font-semibold"><CheckCircle className="w-4 h-4" /> Verified</span>
                                                            : <></>;
                                                      return (
                                                            <React.Fragment key={v.id}>
                                                                  <tr className={`${rowBg} transition-colors`}>
                                                                        <td>{icon}</td>
                                                                        <td>
                                                                              <span className="font-semibold">{desc}</span>
                                                                              {v.type === "PROF" && !!v.data && (
                                                                                    <span className="block  text-base-content/70">{v.data}</span>
                                                                              )}
                                                                        </td>
                                                                        <td>
                                                                              <span className={(!v.complete) ? "opacity-50" : ""}>
                                                                                    <pre className="whitespace-pre-wrap  font-mono">{v.data || <span className="italic text-base-content/40">—</span>}</pre>
                                                                              </span>
                                                                        </td>
                                                                        <td>{statusChip}</td>
                                                                        <td className="text-right">
                                                                              {v.complete ? (
                                                                                    <button
                                                                                          className="btn btn-primary btn-md"
                                                                                          onClick={() => { toggleEncrypted(v.type); }}
                                                                                    >
                                                                                          <Caret open={showEncrypted.includes(v.type)} /> See what we see
                                                                                    </button>
                                                                              ) : (
                                                                                    <button className="btn btn-secondary btn-md" disabled>Verify</button>
                                                                              )}
                                                                        </td>
                                                                  </tr>
                                                                  {v.encrypt && showEncrypted.includes(v.type) && (
                                                                        <tr>
                                                                              <td colSpan={5} className="bg-black/80">
                                                                                    <textarea
                                                                                          readOnly={true}
                                                                                          className="w-full h-28 rounded font-mono bg-black/70 text-yellow-400 p-3"
                                                                                    >{v.encrypt}</textarea>
                                                                              </td>
                                                                        </tr>
                                                                  )}
                                                            </React.Fragment>
                                                      )
                                                })
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
                  </div>
            </MembersLayout>
      );
}
