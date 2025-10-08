//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, {useEffect, useState}                   from "react";
import {useNavigate}                                  from "react-router-dom";
import SpinnerCircle                                  from "../base/SpinnerCircle";
import Call                                           from "../../classes/call";

//||------------------------------------------------------------------------------------------------||
//|| User Type (Matches /auth/me API Response)
//||------------------------------------------------------------------------------------------------||

export type User = {
	id                : number;
	email?            : string;
      username?         : string;
	status            : string;
	type              : string;
	level             : number;
	security          : number;
	verifications     : any[];
};

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||

type RequireMemberProps = {
	children: (user: User) => React.ReactNode;
};

//||------------------------------------------------------------------------------------------------||
//|| RequireMember Component
//||------------------------------------------------------------------------------------------------||

export default function RequireMember({children}: RequireMemberProps) {

      //||------------------------------------------------------------------------------------------------||
      //|| Const
      //||------------------------------------------------------------------------------------------------||

      const navigate                      = useNavigate();
      const [loading, setLoading]         = useState(true);
	const [user, setUser]               = useState<User | null>(null);
      const [loggedOut, setLoggedOut]     = useState(false);

      //||------------------------------------------------------------------------------------------------||
      //|| Effect
      //||------------------------------------------------------------------------------------------------||

	useEffect(() => {
		const checkAuth = async () => {
                  try {
                        const chirp = new Call("/auth/me", {});
                        chirp.method = "GET";
                        chirp.debug  = true;
                        await chirp.execute();
                        if (!chirp.ok()) {
                              setLoggedOut(true);
                        }
                        setUser(chirp.responsePayload.data as User);
                        setLoading(false);
			} catch (err) {
				console.error("Auth check failed", err);
                        setLoggedOut(true);
			} finally {
				setLoading(false);
			}
		};

		checkAuth();
	}, [navigate]);

	//||------------------------------------------------------------------------------------------------||
	//|| Render States
	//||------------------------------------------------------------------------------------------------||

	if (loading) {
		return (
                  <main className="fixed top-0 left-0 right-0 bottom-0 w-full h-full z-99 justify-center items-center bg-black/60">
                              <SpinnerCircle className="inline-block mx-auto" />
                              {children({ user: null } as any)}
                  </main>
		);
	}

      if (loggedOut) {
		return (
                  <main className="fixed inset-0 z-[99] flex justify-center items-center bg-black/20 text-center">
                        <div className="p-5">
                              <img src="/complyage.svg" alt="Logo" className="w-20 h-20 mx-auto mb-4" />
                              <h1 className="font-bold text-2xl">You have been logged out. Please log in again.</h1>
                              <button onClick={() => window.location.href="/login/"} className="mt-3 px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">
                                    Login
                              </button>
                        </div>
                  </main>
		);
      }

	if (!user) {
		return null; // In case of unexpected state
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Render Protected Content
	//||------------------------------------------------------------------------------------------------||
	return <>{children(user)}</>;
}
