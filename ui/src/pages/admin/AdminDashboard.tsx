//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useEffect, useState }                      from "react";
import { CheckCircle, XCircle, Loader2 }                   from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import AdminLayout                                         from "../../layouts/AdminLayout";
import WidgetPanel                                         from "../../components/admin/WidgetPanel";

//||------------------------------------------------------------------------------------------------||
//|| Types
//||------------------------------------------------------------------------------------------------||

interface Verification {
      id: string;
      name: string;
      type: string;
      status: "approved" | "declined" | "pending";
      date: string;
}

interface ServiceHealth {
      name: string;
      status: "healthy" | "degraded" | "offline";
}

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function AdminDashboard() {

      //||------------------------------------------------------------------------------------------------||
      //|| State
      //||------------------------------------------------------------------------------------------------||

      const [verifications, setVerifications]                     = useState<Verification[]>([]);
      const [services, setServices]                               = useState<ServiceHealth[]>([]);

      //||------------------------------------------------------------------------------------------------||
      //|| useEffect
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            // Simulate data fetch
            setVerifications([
                  { id: "v001", name: "John Doe", type: "IDEN", status: "approved", date: "2025-09-09" },
                  { id: "v002", name: "Jane Smith", type: "ADDR", status: "declined", date: "2025-09-09" },
                  { id: "v003", name: "Emily Rose", type: "FACE", status: "approved", date: "2025-09-08" },
                  { id: "v004", name: "Chris Lee", type: "UAGE", status: "pending", date: "2025-09-08" },
            ]);

            setServices([
                  { name: "oauth",  status: "healthy" },
                  { name: "client", status: "healthy" },
                  { name: "ui",     status: "degraded" },
                  { name: "api",    status: "offline" },
                  { name: "agent",  status: "healthy" },
            ]);
      }, []);

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      return (
            <AdminLayout>
                  <div className="max-w-6xl mx-auto px-6 py-10">
                        <h1 className="text-3xl font-bold mb-6">Admin Dashboard</h1>

                        {/* Widgets */}
                        <div className="grid grid-cols-1 sm:grid-cols-3 gap-6 mb-10">
                              <WidgetPanel title="Total Sales" route="/admin/widget/verificaiton" />
                              <WidgetPanel title="Total Bills" route="/api/admin/stats/bills" />
                              <WidgetPanel
                                    title="Verifications"
                                    route="/admin/widget/verifications"
                                    selectOptions={["VERF", "REJT", "PEND", "PEVF"]}
                              />
                        </div>

                        {/* Recent Verifications Table */}
                        <div className="bg-black/20 shadow rounded-lg p-6 mb-10">
                              <h2 className="text-xl font-semibold mb-4">Recent Verifications</h2>
                              <div className="overflow-x-auto">
                                    <table className="min-w-full shadow">
                                          <thead className="bg-black/40">
                                                <tr>
                                                      <th className="px-4 py-2 text-left text-sm font-semibold">Name</th>
                                                      <th className="px-4 py-2 text-left text-sm font-semibold">Type</th>
                                                      <th className="px-4 py-2 text-left text-sm font-semibold">Status</th>
                                                      <th className="px-4 py-2 text-left text-sm font-semibold">Date</th>
                                                </tr>
                                          </thead>
                                          <tbody>
                                                {verifications.map(v => (
                                                      <tr key={v.id} className="border-b border-gray-600">
                                                            <td className="px-4 py-2">{v.name}</td>
                                                            <td className="px-4 py-2">{v.type}</td>
                                                            <td className="px-4 py-2">
                                                                  {v.status === 'approved' && <span className="text-green-600">Approved</span>}
                                                                  {v.status === 'declined' && <span className="text-red-600">Declined</span>}
                                                                  {v.status === 'pending'  && <span className="text-yellow-600">Pending</span>}
                                                            </td>
                                                            <td className="px-4 py-2">{v.date}</td>
                                                      </tr>
                                                ))}
                                          </tbody>
                                    </table>
                              </div>
                        </div>

                        {/* Service Health List */}
                        <div className="bg-black/20 shadow rounded-lg p-4">
                              <h2 className="text-xl font-semibold mb-4 border-b border-gray-400 pb-2">Service Health</h2>
                              <ul className="space-y-3">
                                    {services.map(service => (
                                          <li key={service.name} className="flex items-center justify-between">
                                                <span className="font-medium capitalize">{service.name}</span>
                                                <span className={`inline-flex items-center px-3 py-1 rounded-full text-sm font-medium 
                                                      ${service.status === "healthy"  ? "bg-green-100 text-green-700" : ""}
                                                      ${service.status === "degraded" ? "bg-yellow-100 text-yellow-700" : ""}
                                                      ${service.status === "offline"  ? "bg-red-100 text-red-700" : ""}
                                                `}>
                                                      {service.status}
                                                      {service.status === "healthy"  && <CheckCircle className="w-4 h-4 ml-2" />}
                                                      {service.status === "degraded" && <Loader2 className="w-4 h-4 ml-2 animate-spin" />}
                                                      {service.status === "offline"  && <XCircle className="w-4 h-4 ml-2" />}
                                                </span>
                                          </li>
                                    ))}
                              </ul>
                        </div>
                  </div>
            </AdminLayout>
      );
}
