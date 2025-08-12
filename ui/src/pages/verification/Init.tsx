//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useEffect, useState }          from "react";
import { useSearchParams, useNavigate }        from "react-router-dom";
import MembersLayout                           from "../../layouts/MembersLayout";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import SpinnerCircle                           from "../../components/base/SpinnerCircle";

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

export default function VerificationInit() {
      
      //||------------------------------------------------------------------------------------------------||
      //|| Var
      //||------------------------------------------------------------------------------------------------||

      const [searchParams]                     = useSearchParams();
      const navigate                           = useNavigate();
      const [error, setError]                  = useState<string | null>(null);
      const setType                            = searchParams.get("type") || "IDEN";

      //||------------------------------------------------------------------------------------------------||
      //|| Create the Verification Record
      //||------------------------------------------------------------------------------------------------||
      
      useEffect(() => {
            const run = async () => {
                  try {

                        //||------------------------------------------------------------------------------------------------||
                        //|| Call
                        //||------------------------------------------------------------------------------------------------||

                        const res = await fetch(`/v1/api/verify/init?type=${setType}`, {
                              method      : "GET",
                              headers     : { "Content-Type": "application/json" },
                        });

                        //||------------------------------------------------------------------------------------------------||
                        //|| Parse
                        //||------------------------------------------------------------------------------------------------||

                        const body = await res.json();
                        console.log("API Response Text:", body);

                        if (!res.ok || !body?.success) {
                              throw new Error(body?.message || "Verification init failed");
                        }

                        const identifier = body?.data?.identifier;
                        if (!identifier) throw new Error("Missing identifier in response");
                        
                        switch(setType) {
                              case "IDEN":
                                    navigate(`/verification/id/?identifier=${identifier}`);
                                    break;
                              default:
                                    throw new Error("Unknown verification type");
                        }
                        

                  } catch (err: any) {
                        setError(err.message || "Unknown error occurred");
                  }
            };

            run();
      }, [setType, navigate]);

      //||------------------------------------------------------------------------------------------------||
      //|| Return the Error Message if needed
      //||------------------------------------------------------------------------------------------------||

      return (
            <MembersLayout title="Initializing Verification...">                  
                  {error ? (
                        <div className="text-red-600 bg-red-100 border border-red-300 p-4 rounded-md mt-6 max-w-xl mx-auto text-center">
                              <strong>Error:</strong> {error}
                        </div>
                  ) : (
                        <SpinnerCircle />
                  )}
            </MembersLayout>
      );
}
