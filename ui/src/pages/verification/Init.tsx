//||------------------------------------------------------------------------------------------------||
//|| Verification Init Page
//|| /src/pages/verification/Init.tsx
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||

      import React, { useEffect, useState }          from "react";
      import { useSearchParams, useNavigate }        from "react-router-dom";
      import { XCircle, Timer, CheckCircle }            from "lucide-react";

      //||------------------------------------------------------------------------------------------------||
      //|| Data
      //||------------------------------------------------------------------------------------------------||

      import type { VerificationTypes, VerificationStatuses }     from "../../interfaces/models/model.verification";
      import { 
            getVerificationType, 
            getVerificationStatus,
            getVerificationIcon
      }                                                           from "../../data/getVerificationData";
      import MembersLayout                                        from "../../layouts/MembersLayout";
      import SpinnerCircle                                        from "../../components/base/SpinnerCircle";

      //||------------------------------------------------------------------------------------------------||
      //|| Types
      //||------------------------------------------------------------------------------------------------||

      interface VerificationRow {
            uuid      : string;
            type      : string;
            display   : string;
            meta      : string;
            status    : string;
            createdAt: string;
            updatedAt: string;
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Component
      //||------------------------------------------------------------------------------------------------||

      export default function VerificationInit() {

            //||------------------------------------------------------------------------------------------------||
            //|| State
            //||------------------------------------------------------------------------------------------------||

            const [searchParams]                     = useSearchParams();
            const navigate                           = useNavigate();
            const [error, setError]                  = useState<string | null>(null);
            const [loading, setLoading]              = useState(true);
            const [verifications, setVerifications]  = useState<VerificationRow[]>([]);
            const type                               = (searchParams.get("type") || "IDEN").toUpperCase();

            //||------------------------------------------------------------------------------------------------||
            //|| Fetch Verifications
            //||------------------------------------------------------------------------------------------------||

            const fetchList = async () => {
                  setLoading(true);
                  setError(null);
                  try {
                        const res = await fetch(`/v1/api/verify/list?type=${type}`);
                        const body = await res.json();
                        if (!res.ok || !body?.success) {
                              throw new Error(body?.message || "Failed to fetch verifications");
                        }
                        const rows = Array.isArray(body?.data) ? body.data : [];
                        console.log("ROWS", rows);
                        if (!rows.length) {
                              // No records: Init process
                              await handleInit();
                              return;
                        }
                        setVerifications(rows);
                  } catch (err: any) {
                        setError(err.message || "Unknown error");
                  } finally {
                        setLoading(false);
                  }
            };

            //||------------------------------------------------------------------------------------------------||
            //|| Create New / Init Logic
            //||------------------------------------------------------------------------------------------------||

            const handleInit = async () => {
                  setLoading(true);
                  setError(null);
                  switch(type) {
                        case "MAIL": window.location.href = `/verification/email/`; return;
                        case "PHNE": window.location.href = `/verification/phone/`; return;
                        case "ADDR": window.location.href = `/verification/address/`; return;
                        case "CRCD": window.location.href = `/verification/card/`; return;
                        default: setError("Invalid verification type"); return;
                  }
                  try {
                        const res = await fetch(`/v1/api/verify/init?type=${type}`);
                        const body = await res.json();
                        if (!res.ok || !body?.success) {
                              throw new Error(body?.message || "Verification init failed");
                        }
                        const identifier = body?.data?.identifier;
                        if (!identifier) throw new Error("Missing identifier in response");
                        navigate(`/verification/id/?identifier=${identifier}`);
                  } catch (err: any) {
                        setError(err.message || "Unknown error");
                  } finally {
                        setLoading(false);
                  }
            };

            //||------------------------------------------------------------------------------------------------||
            //|| handleVerify
            //||------------------------------------------------------------------------------------------------||

            const handleVerify = (type : VerificationTypes, uuid : string ) => {
                  switch(type) {
                        case "MAIL": window.location.href = `/verification/mail/?identifier=${uuid}`; break;
                        case "PHNE": window.location.href = `/verification/phone/?identifier=${uuid}`; break;
                        case "UAGE": window.location.href = `/verification/age/?identifier=${uuid}`; break;
                        case "ADDR": window.location.href = `/verification/address/verify?identifier=${uuid}`; break;   
                        case "CRCD": window.location.href = `/verification/card/verify?identifier=${uuid}`; break;
                        default: return;
                  }
            };

            //||------------------------------------------------------------------------------------------------||
            //|| Button Action
            //||------------------------------------------------------------------------------------------------||

            const buttonAction = (uuid: string, type: VerificationTypes, status: VerificationStatuses) => {
                  switch (status) {
                        case "PEND": return ( <button className="btn btn-error btn-md rounded-md px-4 flex items-center gap-1" title="Try again" onClick={handleInit}><XCircle size={18} className="inline-block" /></button>);
                        case "PEVF": return ( <button className="btn btn-success btn-md rounded-md px-4" title="Enter Verification" onClick={() => handleVerify(type, uuid)} >Enter Verification</button> );
                        case "APPR": return ( <span className="badge badge-outline badge-info flex items-center gap-1 text-xs"><Timer size={16} className="inline-block" /> { getVerificationStatus(status) }</span>);
                        case "VERF": return ( <span className="badge badge-outline badge-info flex items-center gap-1"><CheckCircle size={16} className="inline-block" /></span>);
                        case "RJCT": return ( <button className="btn btn-error btn-md rounded-md px-4 flex items-center gap-1" title="Try again" onClick={handleInit}><XCircle size={18} className="inline-block" /></button>);
                        case "ESCL": return ( <span className="badge badge-warning badge-outline"><Timer size={16} className="inline-block" /> { getVerificationStatus(status) }</span>);
                        case "EXPD": return ( <span className="badge badge-outline badge-info flex items-center gap-1">{ getVerificationStatus(status) }</span>);
                        case "CNCL": return ( <span className="badge badge-outline badge-info flex items-center gap-1">{ getVerificationStatus(status) }</span>);                        
                        case "MISS": return ( <span className="badge badge-outline badge-secondary flex items-center gap-1">{ getVerificationStatus(status) }</span>);
                        default: return (<span className="badge badge-outline badge-error flex items-center gap-1">Error - {status }</span>);                        
                  }
            };

            //||------------------------------------------------------------------------------------------------||
            //|| On Mount
            //||------------------------------------------------------------------------------------------------||

            useEffect(() => { fetchList(); /* eslint-disable-next-line */ }, [type]);

            //||------------------------------------------------------------------------------------------------||
            //|| UI
            //||------------------------------------------------------------------------------------------------||

            return (
			<MembersLayout title="Your Previous Verification Attempts">
				<div className="max-w-4xl mx-auto py-8">
					{loading ? (
						<SpinnerCircle />
					) : error ? (
						<div className="text-red-500 bg-red-900/60 border border-red-400 p-4 rounded-xl mb-6 text-center">
							<strong>Error:</strong> {error}
						</div>
					) : (
						<div className="max-w-7xl">
							<div className="mb-2 text-left">
								<h3 className="text-2xl font-semibold mb-2 text-left text-gray-400">
									Below is a list of your verification attempts.
								</h3>
								<p>You can continue or re-verify as needed.</p>
							</div>
							<div className="mb-4 text-left bg-black/20 p-4 rounded-lg mt-4">
								<div className="flex justify-end mb-6 border-b border-gray-400 pb-4">
									<button className="btn btn-secondary btn-md" onClick={handleInit}>
										<span className="font-semibold">Create New Verification</span>
									</button>
								</div>

								<div className="overflow-x-auto rounded-lg shadow">
									<table className="table table-zebra table-lg bg-base-200 text-base-content">
										<thead>
											<tr className="bg-base-300">
												<th></th>
                                                                        <th>Type</th>
                                                                        <th>Description</th>
												<th>Created</th>
												<th>Updated</th>
												<th className="w-36 text-center">Action</th>
											</tr>
										</thead>
										<tbody>
											{verifications.length === 0 ? (
												<tr>
													<td colSpan={5} className="text-center py-8 text-lg opacity-60">
														No verifications yet.
													</td>
												</tr>
											) : (
												verifications.map((row, i) => {
													const Icon = getVerificationIcon(row.type as VerificationTypes);
													const typeLabel = getVerificationType(row.type as VerificationTypes);
													const created = new Date(row.createdAt).toLocaleString();
													const updated = new Date(row.updatedAt).toLocaleString();

													return (
														<tr className="w-8 text-center" key={row.uuid || i} className="hover">
															<td className="w-6">
																{Icon && <Icon size={20} className="inline mr-2" />}																
															</td>
															<td className="text-sm w-40">{typeLabel}</td>
                                                                                          <td className="text-sm w-40">{row.display}</td>
                                                                                          <td className="text-xs w-40">{created}</td>
															<td className="text-xs w-40">{updated}</td>
															<td className="text-xs text-center">
																{ buttonAction(row.uuid, row.type as VerificationTypes, row.status as VerificationStatuses) }
															</td>
														</tr>
													);
												})
											)}
										</tbody>
									</table>
								</div>
							</div>
						</div>
					)}
				</div>
			</MembersLayout>
		);
      }
