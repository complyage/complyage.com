//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, {useEffect, useState} from "react";
import { useNavigate }              from "react-router-dom";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

import  Verification                from "../../interfaces/Verification";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import MembersLayout                from "../../layouts/MembersLayout";
   
//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function Encrypted() {

      //||------------------------------------------------------------------------------------------------||
      //|| Var
      //||------------------------------------------------------------------------------------------------||

      const [verifications, setVerifications] = useState<Verification[]>([]);
      const [statusFilter, setStatusFilter] = useState<string>("ALL");

      //||------------------------------------------------------------------------------------------------||
      //|| Fetch
      //||------------------------------------------------------------------------------------------------||            

      const fetchVerifications = async () => {
            try {
                  const res = await fetch("/members/verifications", {
                        credentials: "include",
                  });
                  const json = await res.json();
                  if (json.success) {
                        setVerifications(json.data);
                  }
            } catch (err) {
                  console.error("Failed to load verifications:", err);
            }
      };

      //||------------------------------------------------------------------------------------------------||
      //|| OnLoad
      //||------------------------------------------------------------------------------------------------||            

      useEffect(() => {
            fetchVerifications();
      }, []);

      const filteredVerifications = statusFilter === "ALL" ? verifications : verifications.filter( (v) => v.verification_status === statusFilter );

      //||------------------------------------------------------------------------------------------------||
      //|| Return
      //||------------------------------------------------------------------------------------------------||            

      return (
            <MembersLayout title="Your encrypted data">
            <>
                  {/* Filter Toolbar */}
                  <div className="flex gap-4 mb-4">
                        {["ALL", "ACTV", "QUEU", "DENY"].map((status) => (
                              <button
                                    key={status}
                                    className={`btn btn-sm ${
                                          statusFilter === status
                                                ? "btn-primary"
                                                : "btn-ghost"
                                    }`}
                                    onClick={() =>
                                          setStatusFilter(status)
                                    }>
                                    {status === "ALL"
                                          ? "All"
                                          : status === "ACTV"
                                          ? "Active"
                                          : status === "QUEU"
                                          ? "Queued"
                                          : "Denied"}
                              </button>
                        ))}
                  </div>

                  {/* Table */}
                  <div className="overflow-x-auto">
                        <table className="table w-full bg-base-100 shadow rounded-lg">
                              <thead>
                                    <tr>
                                          <th>ID</th>
                                          <th>Type</th>
                                          <th>Status</th>
                                          <th>Meta</th>
                                    </tr>
                              </thead>
                              <tbody>
                                    {filteredVerifications.length ===
                                    0 ? (
                                          <tr>
                                                <td
                                                      colSpan={4}
                                                      className="text-center p-4">
                                                      No stored data.
                                                </td>
                                          </tr>
                                    ) : (
                                          filteredVerifications.map(
                                                (v) => (
                                                      <tr
                                                            key={
                                                                  v.id_verification
                                                            }>
                                                            <td>
                                                                  {
                                                                        v.id_verification
                                                                  }
                                                            </td>
                                                            <td>
                                                                  {
                                                                        v.verification_type
                                                                  }
                                                            </td>
                                                            <td>
                                                                  {
                                                                        v.verification_status
                                                                  }
                                                            </td>
                                                            <td>
                                                                  <pre className="whitespace-pre-wrap text-xs">
                                                                        {
                                                                              v.verification_meta
                                                                        }
                                                                  </pre>
                                                            </td>
                                                      </tr>
                                                )
                                          )
                                    )}
                              </tbody>
                        </table>
                  </div>
		</>
      </MembersLayout>
      );
}
