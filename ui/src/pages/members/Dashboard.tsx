//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useEffect, useMemo, useState }           from "react";
import { useNavigate }                                   from "react-router-dom";

//||------------------------------------------------------------------------------------------------||
//|| Lucide
//||------------------------------------------------------------------------------------------------||

import { CheckCircle, Ban, MapPin, Shield, ArrowUpRight }   from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

import { VerificationTypes }                                from "../../interfaces/models/model.verify";
import { Identity }                                         from "../../interfaces/verify/identity/identity";

//||------------------------------------------------------------------------------------------------||
//|| Data
//||------------------------------------------------------------------------------------------------||

import { getAllVerificationTypes, getVerificationType, 
         getVerificationIcon, getVerificationDescription }  from "../../data/getVerificationData";
import { isVerified }                                       from "../../data/getIdentity";
import { BadgeVerified }                                    from "../../components/badges/BadgeStatuses";

//||------------------------------------------------------------------------------------------------||
//|| Call
//||------------------------------------------------------------------------------------------------||

import Call                                                 from "../../classes/call";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import MembersLayout                                        from "../../layouts/MembersLayout";
import SpinnerCircle                                        from "../../components/base/SpinnerCircle";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

import { DashboardData }                                    from "../../interfaces/members/dashboard";

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function Dashboard() {

      //||------------------------------------------------------------------------------------------------||
      //|| Default Data
      //||------------------------------------------------------------------------------------------------||

      const defaultData: DashboardData = {
            isVerified        : false,
            verifiedAge       : 0,
            minimumType       : "IDEN",
            ipAddress         : "---",
            location: {
                  city:      "---",
                  region:    "---",
                  country:   "---",
                  latitude:   0,
                  longitude:  0
            },
            zone: {
                  laws:          "",
                  requirements:  ["IDEN", "FACE", "CRCD"],
                  effective   : "---",
                  minAge:        0
            },
            identity: {}
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Const
      //||------------------------------------------------------------------------------------------------||

      const navigate                                          = useNavigate();
      const [ data, setData ]                                 = useState<DashboardData>(defaultData);
      const [ error, setError ]                               = useState<string | null>(null);
      const [ loading, setLoading ]                           = useState<boolean>(true);
      
      //||------------------------------------------------------------------------------------------------||
      //|| Base Verification
      //||------------------------------------------------------------------------------------------------||

      const baseVerifications = useMemo(() => getAllVerificationTypes(), []);
      
      //||------------------------------------------------------------------------------------------------||
      //|| Get Statuses
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {            
            setLoading(true);
            (async() => {
                  const chirp = new Call("/user/dashboard", {});
                  chirp.debug = true;
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
            })();
      }, []);


      //||------------------------------------------------------------------------------------------------||
      //|| UI: Info Pill
      //||------------------------------------------------------------------------------------------------||

      const InfoPill: React.FC<{ label: string; value?: string }> = ({ label, value }) => (
            <div className="w-full">
                  <div className="text-xs font-semibold uppercase tracking-wide text-base-content/60 mb-1">{label}</div>
                  <div className="bg-base-200 rounded-lg px-3 py-2 text-sm font-medium border border-base-300 text-white">{value ?? "—"}</div>
            </div>
      );

      //||------------------------------------------------------------------------------------------------||
      //|| Header
      //||------------------------------------------------------------------------------------------------||

      const VerifiedHeader: React.FC = () => {
            if (loading) return (
                  <div className="flex items-center gap-3 mb-4 border-b border-gray-600 pb-2">
                        <div className={`flex items-center gap-2 ${ data.isVerified ? "text-green-600" : "text-white" }`}>
                                    <>
                                          <SpinnerCircle className="w-6 h-6" />     
                                          <h2 className="text-xl text-gray-500 font-bold">Checking Location and Verifications</h2>
                                    </>
                        </div>
                  </div>                  
            );
            return (            
                  
                  <div className="flex items-center gap-3 mb-4 border-b border-gray-600 pb-2">
                        <div className={`flex items-center gap-2 ${ data.isVerified ? "text-green-600" : "text-white" }`}>
                              { data?.isVerified ? (
                                    <>
                                    <CheckCircle className="w-6 h-6" />
                                    <h2 className="text-xl font-bold">Congrats You are Age Verified in this Region!</h2>
                                    </>
                              ) : (
                                    <>
                                    <CheckCircle className="w-6 h-6" />
                                    <h2 className="text-xl font-bold">These methods are eligible for your region</h2>
                                    </>
                              ) }                        
                        </div>
                  </div>
            )
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Header
      //||------------------------------------------------------------------------------------------------||

      const AlertBanner: React.FC = () => {
            if (loading) return null;
            if (data.isVerified) return null;
            return (
                  <div className="flex items-center gap-3 p-2 bg-red-500/80 rounded-lg">
                        <div className="flex mx-auto text-center items-center gap-2 text-white">
                              <Ban className="w-6 h-6" />
                              <h2 className="text-xl font-bold">Warning : You are NOT Age Verified!</h2>
                        </div>
                  </div>
            )
      } 

      //||------------------------------------------------------------------------------------------------||
      //|| Generate the Required Verification Methods
      //||------------------------------------------------------------------------------------------------||

      const VerifyMethod: React.FC<{ type: VerificationTypes }> = ({ type }) => {
            if (!data || !data.zone || !data.zone.requirements) return null;
            //||------------------------------------------------------------------------------------------------||
            //|| Const
            //||------------------------------------------------------------------------------------------------||            
            const isAllowed         = data.zone.requirements.includes(type);
            const userVerified      = isVerified(type, data.identity);
            const Icon              = getVerificationIcon(type);
            //||------------------------------------------------------------------------------------------------||
            //|| Allowed
            //||------------------------------------------------------------------------------------------------||
            if (!isAllowed) return null;
            //||------------------------------------------------------------------------------------------------||
            //|| JSX
            //||------------------------------------------------------------------------------------------------||
            return (
                  <div
                        className={`w-full h-full p-4 rounded-lg border ${ userVerified ? " bg-green-500/10 text-green-500" : "border-base-300 bg-base-200" }`}
                  >
                        { loading ? (
                              <div className="animate-pulse space-y-4 opacity-10">
                                    {/* Header with icon and title */}
                                    <div className="flex items-center gap-3">
                                          <div className="bg-gray-300 rounded-full h-6 w-6" />
                                          <div className="h-4 bg-gray-300 rounded w-1/3" />
                                    </div>

                                    {/* Description lines */}
                                    <div className="space-y-2">
                                          <div className="h-3 bg-gray-300 rounded w-full" />
                                          <div className="h-3 bg-gray-300 rounded w-5/6" />
                                          <div className="h-3 bg-gray-300 rounded w-4/6" />
                                    </div>

                                    {/* Button placeholder */}
                                    <div className="mt-2 flex w-full justify-end">
                                          <div className="h-8 w-32 bg-gray-300 rounded" />
                                    </div>
                              </div>

                        ) : ( 
                              <>
                              <div className="flex items-center gap-3 mb-3">
                                    <div className="text-primary">
                                          {Icon ? <Icon className="text-white w-6 h-6" /> : null}
                                    </div>
                                    <h3 className="text-lg font-bold">{ getVerificationType(type) }</h3>
                                    {userVerified && <CheckCircle className="w-5 h-5 text-green-400 ml-auto" />}
                              </div>

                              <p className="text-sm text-base-content/70 mb-4">
                              {
                                    getVerificationDescription(type)
                              }
                              </p>
                              <div className="mt-2 flex w-full justify-end border-white">
                                    {
                                          !userVerified ? (
                                                <button
                                                      className="btn btn-secondary btn-md"
                                                      onClick={() => navigate(`/verification/init?type=${encodeURIComponent(type)}`)}>
                                                      Get Verified
                                                      <ArrowUpRight className="w-4 h-4 ml-2" />
                                                </button>
                                          ) : (
                                                <BadgeVerified />
                                          )
                                    }
                              </div>
                              </>
                        )}
                  </div>
                  
            );          
      };  

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      return (
		<MembersLayout>
			<div className="w-full">
				<AlertBanner />
				<div className="col-span-1 lg:col-span-3 mt-2">
					<div className="bg-base-100 rounded-2xl shadow-xl p-6 border border-base-200">
					      
                                    <VerifiedHeader />

                                    <div className="grid grid-cols-1 sm:grid-cols-3 md:grid-cols-3 lg:grid-cols-3 gap-4 mb-6 w-full">
                                          {baseVerifications.map((v) => (
                                                <VerifyMethod key={v} type={v} />
                                          ))}
                                    </div>

						{/* Heading inside left column */}
						<div className="flex items-center gap-3 mb-4 border-b border-gray-600 pb-2">
							<div className="text-primary">
								<MapPin className="w-6 h-6" />
							</div>
							<h2 className="text-xl font-bold">Your Location</h2>
							<div className="ml-auto text-center p-2 rounded-lg bg-black/20 shadow-2xl border border-gray-600">
								<span className="block text-xs text-gray-400">Min. Age</span>
								<span className="block font-bold text-white">{data?.zone?.minAge}</span>
							</div>
						</div>

						<div className="grid grid-cols-1 md:grid-cols-2 gap-6">
							{/* Left: Location + Allowed */}
							<div className="space-y-4">
								<div className="grid grid-cols-2 gap-3">
									<InfoPill label="IP Address" value={data?.ipAddress} />
									<InfoPill label="City" value={data?.location?.city} />
									<InfoPill label="State" value={data?.location?.region} />
									<InfoPill label="Country" value={data?.location?.country} />
								</div>

								<div className="mt-4">
									<div className="text-sm font-semibold text-base-content/70 mb-2">Age Verification Types Allowed</div>
									<div className="flex w-full gap-2">
										{baseVerifications.map((v) => {
                                                                  if (!data.zone || !data.zone.requirements) return null; 
                                                                  console.log(data?.zone?.requirements);
											const isAllowed = data?.zone?.requirements.includes(v);
											return (
												<span
													key={v}
													className={`flex-1 flex justify-center items-center badge px-3 py-5 text-xs rounded-lg ${
														isAllowed ? "badge-primary bg-orange-400" : "badge-primary badge-outline"
													}`}
													style={{minWidth: 0}}>
													{(() => {
														const Icon = getVerificationIcon(v);
														return Icon ? <Icon className="w-5 h-5" /> : null;
													})()}
												</span>
											);
										})}
									</div>
								</div>
							</div>

							{/* Right: VPN Ad / Notice */}
							<div className="relative overflow-hidden rounded-xl border border-base-200 bg-gradient-to-br from-base-200 to-base-100 p-4">
								<div className="flex items-start gap-4">
									<div className="shrink-0 text-primary">
										<Shield className="w-24 h-24" />
									</div>
									<div className="flex-1">
										<h3 className="text-lg font-bold leading-tight">Using a VPN?</h3>
										<p className="text-sm text-base-content/70 mt-1">
											VPNs and proxies can change your detected location and may limit which verification options
											are available in your region.
										</p>
										<p className="text-sm text-base-content/70 mt-1">
											<b className="text-white">Choose Wisely!</b> Low quality VPN IPs often have poor reputation,
											which can lead to CAPTCHAs or slow access.
										</p>
										<div className="flex flex-wrap gap-2 mt-4">
											<button
												className="btn btn-primary btn-md bg-blue-400"
												onClick={() => navigate("/members/vpns")}>
												Recommended VPNs
											</button>
										</div>
									</div>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		</MembersLayout>
	);
}
