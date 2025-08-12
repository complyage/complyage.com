//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, {useEffect, useState}                   from "react";
import {useNavigate}                                  from "react-router-dom";
import MembersLayout                                  from "../../layouts/MembersLayout";

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function Quit() {

      //||------------------------------------------------------------------------------------------------||
      //|| Var
      //||------------------------------------------------------------------------------------------------||

      const [quitToken, setQuitToken]           = useState("");
	const [loading, setLoading]               = useState(true);
	const [error, setError]                   = useState("");
	const navigate                            = useNavigate();

      //||------------------------------------------------------------------------------------------------||
      //|| Use Effect
      //||------------------------------------------------------------------------------------------------||

	useEffect(() => {
		const fetchToken = async () => {
			try {
				const res = await fetch("/user/quit", {method: "POST", credentials: "include"},);
				const json = await res.json();
				if (json.success) setQuitToken(json.data.quitToken);
				else setError("Failed to load security token");
			} catch {
				setError("Error fetching quit token");
			} finally {
				setLoading(false);
			}
		};
		fetchToken();
	}, []);

      //||------------------------------------------------------------------------------------------------||
      //|| Handle Confirm
      //||------------------------------------------------------------------------------------------------||

	const handleConfirmDelete = async () => {
		try {
			const res = await fetch("/user/delete-account", {
				method: "POST",
				headers: {"Content-Type": "application/x-www-form-urlencoded"},
				credentials: "include",
				body: new URLSearchParams({quitToken}).toString(),
			});
			const json = await res.json();
			if (json.success) navigate("/signup");
			else setError(json.message || "Failed to delete account");
		} catch {
			setError("Error deleting account");
		}
	};

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

	return (
		<MembersLayout title="Confirm Account Deletion">
			{loading ? (
				<p>Loading security check…</p>
			) : error ? (
				<p className="text-red-500">{error}</p>
			) : (
				<div className="p-6 bg-base-100 rounded-lg shadow max-w-xl">
					<h2 className="text-xl font-bold mb-4">Final Confirmation</h2>
					<p className="mb-4 text-sm text-base-content/70">
						This action is <strong>irreversible</strong>. Your account will be permanently deleted.
					</p>
					<button onClick={handleConfirmDelete} className="btn btn-error">
						Confirm Delete
					</button>
				</div>
			)}
		</MembersLayout>
	);
}
