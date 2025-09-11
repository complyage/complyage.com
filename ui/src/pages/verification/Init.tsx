//||------------------------------------------------------------------------------------------------||
//|| Verification Init Page
//|| /src/pages/verification/Init.tsx
//||------------------------------------------------------------------------------------------------||

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useEffect, useState }                             from "react";
import { useSearchParams, useNavigate }                           from "react-router-dom";
import {  User }                                                  from "lucide-react";
import type { LucideIcon }                                        from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Data
//||------------------------------------------------------------------------------------------------||

import {getVerificationType, getVerificationIcon, getVerificationPageTitle}   from "../../data/getVerificationData";
import MembersLayout                                                          from "../../layouts/MembersLayout";
import SpinnerCircle                                                          from "../../components/base/SpinnerCircle";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import Timestamp                                                  from "../../components/base/Timestamp";
import { BadgeStatus }                                            from "../../components/badges/BadgeStatuses";

//||------------------------------------------------------------------------------------------------||
//|| Classes
//||------------------------------------------------------------------------------------------------||

import Call                                                       from "../../classes/call";

//||------------------------------------------------------------------------------------------------||
//|| Utils
//||------------------------------------------------------------------------------------------------||

import {timeAgo}                                                  from "../../utils/universalDate";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

import {ModelVerify, VerificationTypes, VerificationStatuses}     from "../../interfaces/models/model.verify";

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function VerificationInit() {

      //||------------------------------------------------------------------------------------------------||
      //|| Navigate
      //||------------------------------------------------------------------------------------------------||

      const navigate = useNavigate();

      //||------------------------------------------------------------------------------------------------||
      //|| Create New / Init Logic
      //||------------------------------------------------------------------------------------------------||

      const generateIdenUUID = async() => {
            try {
                  const res         = await fetch(`/v1/api/verify/id/init`);
                  const response    = await res.json();
                  console.log("Generate UUID :: ", response);
                  if (!res.ok || !response.success) {
                        throw new Error(response.message || "Verification init failed");
                  }
                  const identifier = response.data.identifier;
                  if (!identifier) throw new Error("Missing identifier in response");
                  navigate(`/verification/id/?identifier=${identifier}`);
            } catch (err: any) {
                  setError(err.message || "Unknown error");
            } finally {
                  setLoading(false);
            }                  
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Create New / Init Logic
      //||------------------------------------------------------------------------------------------------||

      const generateFaceUUID = async() => {
            try {
                  const res         = await fetch(`/v1/api/verify/face/init`);
                  const response    = await res.json();
                  console.log("Generate UUID :: ", response);
                  if (!res.ok || !response.success) {
                        throw new Error(response.message || "Verification init failed");
                  }
                  const identifier = response.data.identifier;
                  if (!identifier) throw new Error("Missing identifier in response");
                  navigate(`/verification/face/?identifier=${identifier}`);
            } catch (err: any) {
                  setError(err.message || "Unknown error");
            } finally {
                  setLoading(false);
            }                  
      }      

      //||------------------------------------------------------------------------------------------------||
      //|| Handle Init
      //||------------------------------------------------------------------------------------------------||

      const handleInit = () => {
            switch (typeVerify) {
                  case "MAIL": navigate(`/verification/email/`); return;
                  case "PHNE": navigate(`/verification/phone/`); return;
                  case "ADDR": navigate(`/verification/address/`); return;
                  case "CRCD": navigate(`/verification/card/`); return;
                  case "IDEN": generateIdenUUID(); return;
                  case "FACE": generateFaceUUID(); return;
                  default:     return;
            }
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Handle Verify
      //||------------------------------------------------------------------------------------------------||

      const handleVerify = (uuid: string) => {
            switch (typeVerify) {
                  case "MAIL":
                  case "PHNE":
                  case "IDEN":
                  case "ADDR":
                  case "CRCD":
                  case "FACE": navigate(`/verification/check?identifier=${uuid}`);  return;
                  default:     return;
            }
      };

      //||------------------------------------------------------------------------------------------------||
      //|| CheckStatus
      //||------------------------------------------------------------------------------------------------||

      const checkStatus = (uuid: string) => {
            navigate(`/verification/status?identifier=${uuid}`);
      };

	//||------------------------------------------------------------------------------------------------||
	//|| State
	//||------------------------------------------------------------------------------------------------||

	const [searchParams]                      = useSearchParams();
	const [error, setError]                   = useState<string | null>(null);
	const [loading, setLoading]               = useState(true);
	const [verifications, setVerifications]   = useState<ModelVerify[]>([]);
	const typeVerify                          = (searchParams.get("type") || "IDEN").toUpperCase();

	//||------------------------------------------------------------------------------------------------||
	//|| Fetch Verifications
	//||------------------------------------------------------------------------------------------------||

	useEffect(() => {
		(async () => {
			const chirp = new Call(`/v1/api/verify/list?type=${typeVerify}`, {});
			chirp.debug = true;
			chirp.method = "GET";
			await chirp.execute();
			if (!chirp.ok()) {
				setError(chirp.error());
				setLoading(false);
				return;
			}
			if (!chirp.responsePayload.data) {
				setError("No data received");
				setLoading(false);
				return;
			}
			console.log("LOADED DATA", chirp.responsePayload.data);
			setVerifications(chirp.responsePayload.data as ModelVerify[]);
			setLoading(false);
		})();
	}, []);

	//||------------------------------------------------------------------------------------------------||
	//|| Const
	//||------------------------------------------------------------------------------------------------||

	const typeTitle   = getVerificationPageTitle(typeVerify as VerificationTypes);
	const TypeIcon    = (getVerificationIcon(typeVerify as VerificationTypes) as LucideIcon) || User;

      //||------------------------------------------------------------------------------------------------||
	//|| Const
	//||------------------------------------------------------------------------------------------------||

	return (
		<MembersLayout title={typeTitle} icon={TypeIcon}>
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
								<h2 className="text-2xl font-bold mb-2 text-gray-300">
									No Previous{" "}
									<span className="text-gray-100">{getVerificationPageTitle(typeVerify as VerificationTypes)}</span>
								</h2>
								<p className="text-base text-gray-300 mb-6">
									When you create your first attempt, you’ll see your progress and status updates here.
									<br />
									<br />
									Verification is a quick and secure process that helps keep your account safe and trusted. If you have
									questions about verification, you can always contact our support team.
								</p>
							</div>
							<SpinnerCircle size={40} className="mb-4" />
							<button
								className="btn btn-primary text-lg font-semibold px-8 py-3 mt-2 shadow-lg bg-orange-400 hover:bg-orange-500 text-white"
								onClick={handleInit}>
								<TypeIcon className="inline-block mr-2" size={22} />
								Create New Verification
							</button>
						</div>
					</>
				) : (
					<div className="max-w-7xl">
                                    <div className="w-full text-center mr-4 flex mx-auto text-gray-400 mb-5">
                                          <div className="text-3xl font-semibold mb-2 bg-green-400/80 mx-auto rounded-full p-4">
                                                <TypeIcon className="block mx-auto text-white" size={128} />
                                          </div>
                                    </div>
                                    
                                    <div className="flex justify-center mb-6 pb-4">
                                          <button className="btn btn-xl btn-success" onClick={handleInit}>
                                                <span className="font-semibold">Get Verified Now!</span>
                                          </button>
                                    </div>         

						<div className="mb-4 text-left rounded-lg mt-4">
							<div className="overflow-x-auto rounded-lg shadow">
								<table className="table table-zebra table-lg bg-base-200 text-base-content">
									<thead>
										<tr className="bg-base-300">
											<th className="text-center">Type</th>
											<th>Display</th>
											<th className="text-left">Created</th>
											<th className="w-36 text-right">Action</th>
										</tr>
									</thead>
									<tbody>
										{verifications.length === 0 ? (
											<tr>
												<td colSpan={4} className="text-center py-8 text-lg opacity-60">
													No verifications yet.
												</td>
											</tr>
										) : (
											verifications.map((row, i) => {
												const Icon = getVerificationIcon(row.type as VerificationTypes);
												const typeLabel = getVerificationType(row.type as VerificationTypes);
												const created = timeAgo(new Date(row.created));

												return (
													<tr className="w-8 text-center" key={row.uuid || i}>
														<td className="w-6 text-center">
															{Icon && <Icon size={24} className="inline mr-2" />}
														</td>
														<td className="text-sm w-40 text-left">{row.display}</td>
														<td className="text-xs w-40 text-left">
															<Timestamp timestamp={row.created} type="timeago" />
														</td>
														<td className="text-xs text-left">
															<BadgeStatus
																uuid={row.uuid}
																type={row.type as VerificationTypes}
																status={row.status as VerificationStatuses}
                                                                                                handleInit={handleInit}
                                                                                                handleVerify={handleVerify}
                                                                                                checkStatus={checkStatus}                                                                                                
															/>
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
