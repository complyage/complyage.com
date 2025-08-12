
//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useEffect, useState }                             from "react";
import { useNavigate }                                            from "react-router-dom";

import { 
            CheckCircle, ArrowRight, CreditCard, 
            MapPin, Mail, Phone, IdCard, 
            Image as ImageIcon 
       }                                                          from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

import { VerificationStatus }                                     from "../../interfaces/verificationStatus";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import MembersLayout                                              from "../../layouts/MembersLayout";

//||------------------------------------------------------------------------------------------------||
//|| API Verification Interface
//||------------------------------------------------------------------------------------------------||

interface ApiVerification {
      vType: string;
      vStatus: string;
}

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function Dashboard() {

      //||------------------------------------------------------------------------------------------------||
      //|| Const
      //||------------------------------------------------------------------------------------------------||

      const [verifications, setVerifications]  = useState<VerificationStatus[]>([]);
      const navigate                           = useNavigate();

      //||------------------------------------------------------------------------------------------------||
      //|| Base Verification
      //||------------------------------------------------------------------------------------------------||

      const baseVerifications: VerificationStatus[] = [
            { type: "MAIL", label: "Email", blurb: "Confirm your email to receive account updates.", complete: false, icon: <Mail /> },
            { type: "IDEN", label: "ID / Age", blurb: "Verify your age with a government ID.", complete: false, icon: <IdCard /> },
            { type: "PHNE", label: "Phone", blurb: "Add and confirm your phone number.", complete: false, icon: <Phone /> },
            { type: "ADDR", label: "Address", blurb: "Verify your billing or home address.", complete: false, icon: <MapPin /> },
            { type: "CRCD", label: "Credit Card", blurb: "Secure your account with a valid card on file.", complete: false, icon: <CreditCard /> },
            { type: "PROF", label: "Profile Photo", blurb: "Upload a clear profile picture.", complete: false, icon: <ImageIcon /> },
      ];

      //||------------------------------------------------------------------------------------------------||
      //|| Get Statuses
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            //||------------------------------------------------------------------------------------------------||
            //|| Get Statuses
            //||------------------------------------------------------------------------------------------------||
            const fetchStatuses = async () => {
                  try {
                        const res = await fetch("/auth/me", { credentials: "include" });
                        const json = await res.json();
                        //||------------------------------------------------------------------------------------------------||
                        //|| Success
                        //||------------------------------------------------------------------------------------------------||
                        if (json.success) {
                              //||------------------------------------------------------------------------------------------------||
                              //|| Identity 
                              //||------------------------------------------------------------------------------------------------||
                              let identity = json.data.identity;
                              if (typeof identity === "string") {
                                    try {
                                          identity = JSON.parse(identity);
                                    } catch (e) {
                                          console.error("Failed to parse identity JSON:", identity);
                                          identity = {};
                                    }
                              }
                              //||------------------------------------------------------------------------------------------------||
                              //|| Approved
                              //||------------------------------------------------------------------------------------------------||
                              let approved = [];
                              if (Array.isArray(identity.approved)) {
                                    approved = identity.approved;
                              } else if (typeof identity.approved === "string" && identity.approved.trim() !== "" && identity.approved !== "undefined") {
                                    try {
                                          approved = JSON.parse(identity.approved);
                                    } catch (e) {
                                          console.warn("Failed to parse approved JSON:", identity.approved);
                                    }
                              }
                              //||------------------------------------------------------------------------------------------------||
                              //|| Set the New Verifications
                              //||------------------------------------------------------------------------------------------------||
                              const newVerifications = baseVerifications.map((v) => {
                                    if (approved.includes(v.type)) v.complete = true;
                                    return v;
                              });
                              setVerifications(newVerifications);
                        } else {
                              console.warn("Unexpected API response for verifications", json);
                              setVerifications(baseVerifications);
                        }
                  } catch (err) {
                        console.error("Failed to fetch verification statuses:", err);
                        setVerifications(baseVerifications);
                  }
            };

            fetchStatuses();
      }, []);

      //||------------------------------------------------------------------------------------------------||
      //|| Default
      //||------------------------------------------------------------------------------------------------||

      return (
            <MembersLayout title="Your Verification Checklist">
                  <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
                        {verifications.map((v, idx) => (
                              <div key={idx} className="flex flex-col justify-between text-left p-6 bg-base-100 rounded-xl shadow gap-3 h-full">
                                    <div className="flex flex-col gap-2 flex-grow">
                                          <div className="text-primary mb-2">{React.cloneElement(v.icon as React.ReactElement, { size: 48 })}</div>
                                          <h3 className="text-xl font-bold">{v.label}</h3>
                                          <p className="text-base-content/70 text-sm">{v.blurb}</p>
                                    </div>
                                    <div className="flex justify-end mt-4">
                                          {v.complete ? (
                                                <button className="btn btn-success btn-sm">
                                                      <CheckCircle className="w-4 h-4 mr-1" /> Verified
                                                </button>
                                          ) : (
                                                <button className="btn btn-primary btn-sm" onClick={ () => navigate(`/verification/init/?type=${v.type}`) }>
                                                      Complete <ArrowRight className="w-4 h-4 ml-1" />
                                                </button>
                                          )}
                                    </div>
                              </div>
                        ))}
                  </div>
            </MembersLayout>
      );
}
