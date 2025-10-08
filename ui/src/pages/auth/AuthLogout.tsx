//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useState }                from "react";
import { useNavigate, useLocation }       from "react-router-dom";
import SpinnerCircle                      from "../../components/base/SpinnerCircle";

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function AuthLogout() {

      const navigate        = useNavigate();
      const location        = useLocation();
      const inOverlay       = location.pathname.startsWith("/overlay");
      const prefix          = inOverlay ? "/overlay" : "";
      const [isLoading, setIsLoading] = useState(false);

      //||------------------------------------------------------------------------------------------------||
      //|| Perform Logout
      //||------------------------------------------------------------------------------------------------||

      const performLogout = async () => {
            try {
                  setIsLoading(true);

                  const serv = "";
                  const res = await fetch(serv+"/auth/logout", {
                        method: "GET",
                        credentials: "include",
                  });

                  const json = await res.json();
                  if (json.success) {
                        // Clear any client-readable cookies if present
                        document.cookie = "userSession=; path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT";

                        if (inOverlay) {
                              // Post a message to parent window: logout success, close + refresh
                              window.parent.postMessage({ action: "refreshParent" }, "*");
                              return;
                        } else {
                              // Redirect to next or default login
                              navigate(json.next ? `${prefix}${json.next}` : `${prefix}/login`);
                        }
                  } else {
                        navigate(`${prefix}/login`);
                  }
            } catch (err) {
                  console.error("❌ Logout failed", err);
                  navigate(`${prefix}/login`);
            }
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Cancel Logout
      //||------------------------------------------------------------------------------------------------||

      const cancelLogout = () => {
            if (inOverlay) {
                  // Post a message to parent window to just close overlay
                  window.parent.postMessage({ action: "closeOverlay" }, "*");
            } else {
                  navigate(-1);
            }
      };

      //||------------------------------------------------------------------------------------------------||
      //|| UI: Prompt to Logout
      //||------------------------------------------------------------------------------------------------||

      if (isLoading) {
            return (
                  <div className="min-h-screen flex flex-col items-center justify-center text-center">
                        <SpinnerCircle />
                        <p className="mt-4 temt-sm text-gray-400">Logging you out...</p>
                  </div>
            );
      }

      return (
            <div className="min-h-screen flex flex-col items-center justify-center text-center">
                  <h1 className="text-2xl font-bold mb-6">Are you sure you want to log out?</h1>
                  <div className="flex gap-4">
                        <button
                              onClick={performLogout}
                              className="btn btn-lg btn-black hover:bg-red-700 transition border-gray-400"
                        >
                              Yes, log me out
                        </button>
                        <button
                              onClick={cancelLogout}
                              className="btn btn-lg btn-tertiary hover:bg-gray-500 transition border-gray-400"
                        >
                              Cancel
                        </button>
                  </div>
            </div>
      );
}
