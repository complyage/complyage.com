//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, {useState}                        from "react";
import {useSearchParams, useNavigate}           from "react-router-dom";
import { useLocation }                          from "react-router-dom";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import NavMain                                  from "../../components/nav/NavMain";
import FooterMain                               from "../../components/footer/FooterMain";

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function TwoFactorVerify() {

	//||------------------------------------------------------------------------------------------------||
	//|| Var
	//||------------------------------------------------------------------------------------------------||

	const [searchParams]                      = useSearchParams();
	const navigate                            = useNavigate();
	const token                               = searchParams.get("token") || "";
	const [code, setCode]                     = useState("");
	const [statusMessage, setStatusMessage]   = useState("");
	const [loading, setLoading]               = useState(false);

      //||------------------------------------------------------------------------------------------------||
      //|| Extract oauth query param from current URL
      //||------------------------------------------------------------------------------------------------||

      const location      = useLocation();
      const params        = new URLSearchParams(location.search);
      const oauthParam    = params.get("oauth");
            
	//||------------------------------------------------------------------------------------------------||
	//|| Handle Submit
	//||------------------------------------------------------------------------------------------------||

	const handleSubmit = async (e: React.FormEvent) => {
            //||------------------------------------------------------------------------------------------------||
            //|| Generate the Payload
            //||------------------------------------------------------------------------------------------------||
            e.preventDefault();
		const payload = {
			token       : token,
                  code        : code.trim(),
		};            
            //||------------------------------------------------------------------------------------------------||
            //|| Handle the POST
            //||------------------------------------------------------------------------------------------------||
		try {
			setLoading(true);
			const res = await fetch("/auth/twofactor", {
				method      : "POST",
				headers     : { "Content-Type": "application/x-www-form-urlencoded"},
				body        : new URLSearchParams(payload).toString(),
			});
                  console.log("RES", res);
			const json = await res.json();
                  console.log("RESPONSE", json);
                  //||------------------------------------------------------------------------------------------------||
                  //|| Handle the Response
                  //||------------------------------------------------------------------------------------------------||
                  const redirectURL = (oauthParam) ? `/complete?oauth=${oauthParam}` : `/complete`;
			if (json.success) return navigate(redirectURL);
                  //||------------------------------------------------------------------------------------------------||
                  //|| Failed
                  //||------------------------------------------------------------------------------------------------||
                  setStatusMessage(`❌ ${json.error || "Verification failed."}`);
		} catch (err) {
			console.error(err);
			setStatusMessage("❌ Something went wrong. Please try again.");
		} finally {
			setLoading(false);
		}
	};

	//||------------------------------------------------------------------------------------------------||
	//|| JSX
	//||------------------------------------------------------------------------------------------------||

	return (
		<main className="min-h-screen flex flex-col">
			<NavMain />

			<div className="relative flex-1">
				<img
					src="https://picsum.photos/1920/1080"
					alt="Background"
					className="absolute inset-0 w-full h-full object-cover"
				/>
				<div className="absolute inset-0 bg-black/70"></div>

				<div className="relative z-10 flex flex-col items-center justify-center min-h-screen pt-24 pb-12 px-4">
					<div className="bg-base-100 bg-opacity-90 rounded-lg shadow-lg max-w-md w-full p-8 text-center">
						<h1 className="text-3xl font-bold mb-6 text-gray-300">
							Two-Factor Verification
						</h1>
						<p className="mb-4 text-base-content/70">Enter the 6-digit code we sent to your email.</p>

						<form
							onSubmit={handleSubmit}
							className="flex flex-col gap-4">
							<input
								type="text"
								placeholder="Code"
								value={code}
								onChange={(e) => setCode(e.target.value) }
								className="border border-white w-full text-center tracking-widest text-4xl p-5"
								required
							/>

							<button
								type="submit"
								disabled={
									code.trim().length ===
										0 || loading
								}
								className="btn btn-primary w-full">
								{loading
									? "Verifying..."
									: "Verify"}
							</button>
						</form>

						{statusMessage && (
							<p className="mt-4 text-sm">
								{statusMessage}
							</p>
						)}
					</div>
				</div>
			</div>

			<FooterMain />
		</main>
	);
}
