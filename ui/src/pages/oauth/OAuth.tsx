//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useState, useEffect }                       from "react";
import { useLocation }                                      from "react-router-dom";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import SpinnerCircle                                        from "../../components/base/SpinnerCircle";
import {ShieldCheck, Eye, Gift, User, Mail, CreditCard}     from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function OAuthPage() {
	const location                = useLocation();
	const [token, setToken]       = useState("");
	//||------------------------------------------------------------------------------------------------||
	//|| Var
	//||------------------------------------------------------------------------------------------------||
      const [siteData, setSiteData] = useState(null);
      const [error, setError]       = useState(null);
      //||------------------------------------------------------------------------------------------------||
	//|| Get the site
	//||------------------------------------------------------------------------------------------------||
	useEffect(() => {
		const query = new URLSearchParams(location.search);
		const tokenParam = query.get("token");
		if (!tokenParam) {
                  setError("Missing token parameter in URL");
                  return;
            }
            const response = fetch(`/v1/api/oauth?token=${tokenParam}`, { method: "GET", "credentials": "include" }).then((res) => {
                  if (!response || !response.ok) {
                        setError("Failed to fetch site data");
                        return;
                  }            
                  const json = res.json();
		});
	}, [location.search]);
	//||------------------------------------------------------------------------------------------------||
	//|| Scope
	//||------------------------------------------------------------------------------------------------||      
	const requestedScopes = [
		{
			icon: <Gift className="w-6 h-6" />,
			label: "Your Birthday",
			description: "Access your age and birthday.",
		},
		{
			icon: <Mail className="w-6 h-6" />,
			label: "Email Address",
			description: "Read your verified email address.",
		},
		{
			icon: <User className="w-6 h-6" />,
			label: "Your Goverment ID Information",
			description: "Access to your government ID information.",
		},
		{
			icon: <ShieldCheck className="w-6 h-6" />,
			label: "Security Settings",
			description: "Manage your security preferences.",
		},
		{
			icon: <CreditCard className="w-6 h-6" />,
			label: "Payment Details",
			description: "Use your stored payment methods.",
		},
	];
	//||------------------------------------------------------------------------------------------------||
	//|| Approve Changes
	//||------------------------------------------------------------------------------------------------||
	const handleApprove = () => {
		console.log("✅ User APPROVED access!");
		// Call your backend /oauth/authorize endpoint with a positive response
	};
	//||------------------------------------------------------------------------------------------------||
	//|| Deny Changes
	//||------------------------------------------------------------------------------------------------||
	const handleDeny = () => {
		console.log("❌ User DENIED access!");
		// Redirect or close window with an error
	};
	//||------------------------------------------------------------------------------------------------||
	//|| Approve Changes
	//||------------------------------------------------------------------------------------------------||
	return (
		<main className="fixed w-full h-full">
                  
                  <header className="top-0 left-0 right-0 bg-blue-300 font-bold justify-center flex p-3 w-full">
                        Unfortunately the location you are accessing from [Texas, USA] has age restriction laws in place. 
                  </header>


			{/* LEFT PANEL */}
                  <div className="w-full h-full flex flex-row">
                        <aside className="w-full bg-gray-100 md:bg-white md:w-1/2 p-2 md:p-8 flex flex-col justify-center items-center text-center">

                              {/*||------------------------------------------------------------------------------||
                              //|| Site Data
                              //||------------------------------------------------------------------------------||*/}

                              {(!error && siteData) && (                              
                                    <div className="flex w-full mx-auto rounded-lg justify-center">
                                          <div className="bg-black/10 w-full rounded-lg p-5 m-10">
                                                <div className="text-center p-4">
                                                      {(siteData && siteData.logo) && <img src={siteData.logo} alt={`${siteData.name} Logo`} className="h-auto w-48 mb-4 mx-auto" />}
                                                      {(siteData && siteData.name) && <span className="block text-sm font-bold mb-2 text-black">{siteData.name}</span>}
                                                      {(siteData && siteData.url) && <span className="block text-xs text-gray-500 mb-4">{siteData.url}</span>}
                                                      <hr className="border-gray-400 p-2 mt-4" />
                                                      <span className="block font-bold text-md text-red-600">Approving will give access to the requested information</span>
                                                      <span className="block text-black text-xs pt-2">
                                                            You can always revoke access later in your account settings, but we can not guarantee it will be removed from
                                                            their systems.
                                                      </span>
                                                </div>
                                          </div>
                                    </div>
                              )}

                              {/*||------------------------------------------------------------------------------||
                              //|| Error Handling
                              //||------------------------------------------------------------------------------||*/}

                              {(error && !siteData) && (                              
                                    <div className="flex w-full mx-auto rounded-lg justify-center">
                                          <div className="bg-red-100 w-full rounded-lg p-5 m-10">
                                                <div className="text-center p-4">
                                                      <span className="block text-sm font-bold mb-2 text-red-600">{error}</span>
                                                      <span className="block text-xs text-gray-500 mb-4">Please try again later.</span>
                                                </div>
                                          </div>
                                    </div>
                              )}                                              

                              {/*||------------------------------------------------------------------------------||
                              //|| Loading
                              //||------------------------------------------------------------------------------||*/}                                

                              {(!error && !siteData) && (                              
                                    <div className="flex w-full mx-auto rounded-lg justify-center">
                                          <SpinnerCircle />
                                    </div>
                              )}
                        </aside>
      
                        {/* RIGHT PANEL */}
                        <section className="w-full md:w-1/2 bg-base-200 p-10 flex flex-col justify-top">
                              <div className="flex flex-col items-center mb-8 text-center h-[10%]">
                                    <img src="/img/logow.png" alt={`${siteData?.name} Logo`} className="h-16 md:h-8 mb-4" />
                                    <h2 className="text-sm md:text-lg font-bold mb-2">
                                          <span className="text-yellow-400">{siteData?.name || window.location.hostname }</span> <span className="text-gray-500">wants to access this information</span>
                                    </h2>
                              </div>

                              <ul className="space-y-4 mb-4 h-[70%]">
                                    {requestedScopes.map((scope, idx) => (
                                          <li
                                                key={idx}
                                                className="flex items-start gap-4 p-2 bg-base-100 rounded-lg shadow border-l-2 border-orange-400 rounded-l-none">
                                                <div className="text-primary">{React.cloneElement(scope.icon as React.ReactElement, {className: "w-10 h-10"})}</div>
                                                <div>
                                                      <h3 className="font-bold text-lg md:text-lg">{scope.label}</h3>
                                                      <p className="text-base-content/70 text-base text-sm">{scope.description}</p>
                                                </div>
                                          </li>
                                    ))}
                              </ul>

                              <div className="flex flex-col md:flex-row gap-4 justify-end h-[10%]">
                                    <button
                                          onClick={handleApprove}
                                          className="btn btn-primary flex-1 text-xl md:text-2xl p-4 h-14 rounded-md
                              transform transition-transform duration-200 hover:scale-105">
                                          Approve
                                    </button>
                                    <button
                                          onClick={handleDeny}
                                          className="btn btn-neutral flex-1 text-xl md:text-2xl p-4 h-14 rounded-md
                              transform transition-transform duration-200 hover:scale-105">
                                          Deny
                                    </button>
                              </div>
                        </section>
                  </div>
		</main>
	);
}
