//||------------------------------------------------------------------------------------------------||
//|| Verification Init Page
//|| /src/pages/verification/Init.tsx
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||

      import React, { useEffect, useState }          from "react";
      import { useSearchParams, useNavigate }        from "react-router-dom";
      import { XCircle, IdCard, CreditCard, 
            Phone, User, Timer, Home,
            CheckCircle }            from "lucide-react";
      import type { LucideIcon }                            from "lucide-react";

      //||------------------------------------------------------------------------------------------------||
      //|| Data
      //||------------------------------------------------------------------------------------------------||

      import { 
            getVerificationType, 
            getVerificationStatus,
            getVerificationIcon
      }                                                           from "../../data/getVerificationData";
      import MembersLayout                                        from "../../layouts/MembersLayout";
      import SpinnerCircle                                        from "../../components/base/SpinnerCircle";

      //||------------------------------------------------------------------------------------------------||
      //|| Utils
      //||------------------------------------------------------------------------------------------------||

      import { timeAgo }                                            from "../../utils/universalDate";

      //||------------------------------------------------------------------------------------------------||
      //|| Interfaces
      //||------------------------------------------------------------------------------------------------||
      
      import { ModelVerify, VerificationTypes, VerificationStatuses }     from "../../interfaces/models/model.verify";

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
            const [verifications, setVerifications]  = useState<ModelVerify[]>([]);
            const typeVerify                         = (searchParams.get("type") || "IDEN").toUpperCase();

            //||------------------------------------------------------------------------------------------------||
            //|| Fetch Verifications
            //||------------------------------------------------------------------------------------------------||

            const fetchList = async () => {
                  setLoading(true);
                  setError(null);
                  try {
                        const res = await fetch(`/v1/api/verify/list?type=${typeVerify}`);
                        const body = await res.json();
                        if (!res.ok || !body?.success) {
                              throw new Error(body?.message || "Failed to fetch verifications");
                        }
                        const rows = Array.isArray(body?.data) ? body.data : [];
                        console.log("Fetched verifications:", rows);
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
                  switch(typeVerify) {
                        case "MAIL": navigate(`/verification/email/`); return;
                        case "PHNE": navigate(`/verification/phone/`); return;
                        case "ADDR": navigate(`/verification/address/`); return;
                        case "CRCD": navigate(`/verification/card/`); return;
                        case "IDEN": generateIdenUUID(); return;
                        default: setError("Invalid verification type"); return;
                  }
            };

            //||------------------------------------------------------------------------------------------------||
            //|| Create New / Init Logic
            //||------------------------------------------------------------------------------------------------||

            const generateIdenUUID = async() => {
                  try {
                        const res  = await fetch(`/v1/api/verify/id/init`);
                        const body = await res.json();
                        if (!res.ok || !body?.success) {
                              throw new Error(body?.message || "Verification init failed");
                        }
                        const identifier = body?.data?.uuid;
                        if (!identifier) throw new Error("Missing identifier in response");
                        navigate(`/verification/id/?identifier=${identifier}`);
                  } catch (err: any) {
                        setError(err.message || "Unknown error");
                  } finally {
                        setLoading(false);
                  }                  
            }

            //||------------------------------------------------------------------------------------------------||
            //|| handleVerify
            //||------------------------------------------------------------------------------------------------||

            const handleVerify = (uuid : string ) => {
                  switch(typeVerify) {
                        case "MAIL": navigate(`/verification/check/?identifier=${uuid}`); break;
                        case "PHNE": navigate(`/verification/check/?identifier=${uuid}`); break;
                        case "UAGE": navigate(`/verification/check/?identifier=${uuid}`); break;
                        case "ADDR": navigate(`/verification/check?identifier=${uuid}`); break;   
                        case "CRCD": navigate(`/verification/check?identifier=${uuid}`); break;
                        default: return;
                  }
            };

            //||------------------------------------------------------------------------------------------------||
            //|| Button Action
            //||------------------------------------------------------------------------------------------------||

            const buttonAction = (uuid: string, status: VerificationStatuses) => {
                  switch (status) {
                        case "PEND": return ( <button className="btn btn-error btn-md rounded-md px-4 flex items-center gap-1" title="Try again" onClick={handleInit}><XCircle size={18} className="inline-block" /></button>);
                        case "PEVF": return ( <button className="btn btn-success btn-md rounded-md px-4" title="Enter Verification" onClick={() => handleVerify(uuid)} >Enter Verification</button> );
                        case "INPR": return ( <span className="badge badge-outline badge-info flex items-center p-5 text-xs"><Timer size={24} className="inline-block" /><span className="w-18 block font-bold">In Progress</span></span>);
                        case "VERF": return ( <span className="badge badge-outline badge-info flex items-center gap-1"><CheckCircle size={24} className="inline-block" /><span className="w-18 block font-bold">Verified</span></span>);
                        case "RJCT": return ( <button className="btn btn-error btn-md rounded-md px-4 flex items-center gap-1" title="Try again" onClick={handleInit}><XCircle size={18} className="inline-block" /></button>);
                        case "ESCL": return ( <span className="badge badge-warning badge-outline"><Timer size={16} className="inline-block" /> { getVerificationStatus(status) }</span>);
                        case "EXPD": return ( <span className="badge badge-outline badge-info flex items-center gap-1">{ getVerificationStatus(status) }</span>);
                        default: return (<span className="badge badge-outline badge-error flex items-center gap-1">Error - {status }</span>);                        
                  }
            };

            //||------------------------------------------------------------------------------------------------||
            //|| Button Action
            //||------------------------------------------------------------------------------------------------||

            const pageTitle = () : {icon : LucideIcon, title : string} => {
                  switch (typeVerify) {
                        case "ADDR": return { icon: Home, title: "Address Verification" };
                        case "MAIL": return { icon: CheckCircle, title: "Email Verification" };
                        case "IDEN": return { icon: IdCard, title: "Identification Verification" };
                        case "PHNE": return { icon: Phone, title: "Phone Verification" };
                        case "CRCD": return { icon: CreditCard, title: "Address Verification" };
                        case "USER": return { icon: User, title: "Address Verification" };
                        default: return { icon: CheckCircle, title: "Verification" };
                  }
            };


            //||------------------------------------------------------------------------------------------------||
            //|| On Mount
            //||------------------------------------------------------------------------------------------------||

            useEffect(() => { 
                  fetchList(); 
            }, [typeVerify]);

            //||------------------------------------------------------------------------------------------------||
            //|| UI
            //||------------------------------------------------------------------------------------------------||

            const { icon: Icon, title } = pageTitle();

            return (
			<MembersLayout title={pageTitle().title} icon={pageTitle().icon}>
				<div className="max-w-4xl mx-auto py-8">
					{loading ? (
						<SpinnerCircle />
					) : error ? (
						<div className="text-red-500 bg-red-900/60 border border-red-400 p-4 rounded-xl mb-6 text-center">
							<strong>Error:</strong> {error}
						</div>
					) : verifications.length === 0 ? (
						<>
							<div className="flex flex-col items-center justify-center min-h-[300px] p-8 bg-black/20 rounded-lg shadow mb-10">
								<div className="max-w-lg text-center">
									<h2 className="text-2xl font-bold mb-2 text-gray-300">No Previous <span className="text-gray-100">{ pageTitle().title }</span></h2>
									<p className="text-base text-gray-300 mb-6">
										When you create your first attempt, you’ll see
										your progress and status updates here.
										<br />
										<br />
										Verification is a quick and secure process that helps keep your account safe and trusted. If you
										have questions about verification, you can always contact our support team.
									</p>
								</div>
								<button
									className="btn btn-primary text-lg font-semibold px-8 py-3 mt-2 shadow-lg bg-orange-400 hover:bg-orange-500 text-white"
									onClick={handleInit}>
									   <Icon className="inline-block mr-2" size={22} />Create New Verification
								</button>
							</div>
						</>
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
													<td colSpan={6} className="text-center py-8 text-lg opacity-60">
														No verifications yet.
													</td>
												</tr>
											) : (
												verifications.map((row, i) => {
													const Icon = getVerificationIcon(row.type as VerificationTypes);
													const typeLabel = getVerificationType(row.type as VerificationTypes);
													const created = timeAgo(row.created);
													const updated = timeAgo(row.updated);

													return (
														<tr className="w-8 text-center" key={row.uuid || i} className="hover">
															<td className="w-6">
																{Icon && <Icon size={20} className="inline mr-2" />}
															</td>
															<td className="text-sm w-40">{typeLabel}</td>
															<td className="text-sm w-40">{row.display}</td>
															<td className="text-xs w-40">{row.created}</td>
															<td className="text-xs w-40">{row.updated}</td>
															<td className="text-xs text-center">
																{buttonAction(
																	row.uuid,
																	row.status as VerificationStatuses
																)}
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
