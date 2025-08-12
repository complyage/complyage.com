//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, {useEffect, useState}                   from "react";
import {useNavigate}                                  from "react-router-dom";
import SpinnerCircle                                  from "../base/SpinnerCircle";

//||------------------------------------------------------------------------------------------------||
//|| User Type (Matches /auth/me API Response)
//||------------------------------------------------------------------------------------------------||

export type User = {
	id                : number;
	email?            : string;
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
	const navigate = useNavigate();
	const [loading, setLoading] = useState(true);
	const [user, setUser] = useState<User | null>(null);

	useEffect(() => {
		const checkAuth = async () => {
			try {
				const res = await fetch("/auth/me", {method: "GET", credentials: "include"});
				if (!res.ok) {
					navigate("/login");
					return;
				}

				const json = await res.json();
				if (json.success && json.data) {
					setUser(json.data as User);
				} else {
					navigate("/login");
				}
			} catch (err) {
				console.error("Auth check failed", err);
				navigate("/login");
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
                  <main className="min-h-screen flex items-center justify-center">
                        <div className="flex justify-center items-center w-[300px]">
                              <SpinnerCircle className="inline-block mx-auto" />
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
