import React from "react";
import { authLogout } from "../../utils/auth.logout";

type TopbarProps = {
	email: string;
	userId: number;
};

export default function Topbar({email, userId}: TopbarProps) {
	return (
		<div className="w-full h-24 flex justify-between items-center p-4 border-b border-base-100 bg-base-200 border-r-1">
			<h1 className="text-2xl font-bold">Welcome, {email}</h1>
			<button className="btn btn-primary text-base-content" onClick={ () => authLogout() }>
				Logout
			</button>
		</div>
	);
}
