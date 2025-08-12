//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useEffect } from "react";
import { useNavigate } from "react-router-dom";

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function Logout() {
    const navigate = useNavigate();

    //||------------------------------------------------------------------------------------------------||
    //|| Logout on Mount
    //||------------------------------------------------------------------------------------------------||

    useEffect(() => {
        const performLogout = async () => {
            try {
                const res = await fetch("/auth/logout", {
                    method: "GET",
                    credentials: "include",
                });

                const json = await res.json();
                if (json.success) {
                    // Clear any client-readable cookies if present
                    document.cookie = "userSession=; path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT";

                    // Redirect to login or provided URL
                    navigate(json.next || "/login");
                } else {
                    navigate("/login");
                }
            } catch (err) {
                console.error("❌ Logout failed", err);
                navigate("/login");
            }
        };

        performLogout();
    }, [navigate]);

    //||------------------------------------------------------------------------------------------------||
    //|| UI While Logging Out
    //||------------------------------------------------------------------------------------------------||

    return (
        <main className="min-h-screen flex flex-col items-center justify-center text-center">
            <h1 className="text-2xl font-bold mb-4">Logging you out...</h1>
            <p className="opacity-70">Please wait while we end your session.</p>
        </main>
    );
}
