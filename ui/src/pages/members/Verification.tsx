import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import RequireMember, { User } from "../../components/dynamic/RequireMember";
import Sidebar from "../../components/nav/Sidebar";
import Topbar from "../../components/nav/Topbar";

interface Verification {
   id_verification: number;
   verification_type: string;
   verification_status: string;
   verification_meta: string;
}

export default function Verifications() {
   const navigate = useNavigate();
   const [verifications, setVerifications] = useState<Verification[]>([]);
   const [statusFilter, setStatusFilter] = useState<string>("ALL");

   const handleLogout = async () => {
      await fetch("/auth/logout", { credentials: "include" });
      navigate("/login");
   };

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

   useEffect(() => {
      fetchVerifications();
   }, []);

   const filteredVerifications =
      statusFilter === "ALL"
         ? verifications
         : verifications.filter(
              (v) => v.verification_status === statusFilter
           );

   return (
      <RequireMember>
         {(user: User) => (
            <div className="min-h-screen flex bg-base-200">
               <Sidebar onLogout={handleLogout} />
               <main className="flex-1 flex flex-col">
                  <Topbar email={user.email} userId={user.user_id} />

                  <div className="p-10 flex flex-col gap-6">
                     <h1 className="text-3xl font-bold mb-4">
                        Saved Verifications
                     </h1>

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
                              onClick={() => setStatusFilter(status)}
                           >
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
                              {filteredVerifications.length === 0 ? (
                                 <tr>
                                    <td colSpan={4} className="text-center p-4">
                                       No verifications found.
                                    </td>
                                 </tr>
                              ) : (
                                 filteredVerifications.map((v) => (
                                    <tr key={v.id_verification}>
                                       <td>{v.id_verification}</td>
                                       <td>{v.verification_type}</td>
                                       <td>{v.verification_status}</td>
                                       <td>
                                          <pre className="whitespace-pre-wrap text-xs">
                                             {v.verification_meta}
                                          </pre>
                                       </td>
                                    </tr>
                                 ))
                              )}
                           </tbody>
                        </table>
                     </div>
                  </div>
               </main>
            </div>
         )}
      </RequireMember>
   );
}
