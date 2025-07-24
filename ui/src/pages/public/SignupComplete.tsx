//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, {useState, useEffect}                         from "react";
import NavMain                                              from "../../components/nav/NavMain";
import FooterMain                                           from "../../components/footer/FooterMain";
import SpinnerCircle                                        from "../../components/base/SpinnerCircle";
import {useNavigate}                                        from "react-router-dom";

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function SignupComplete() {
	//||------------------------------------------------------------------------------------------------||
	//|| Var
	//||------------------------------------------------------------------------------------------------||
	const navigate                                  = useNavigate();
	const [checkStatus, setCheckStatus]             = useState("CHECK");
	const [password, setPassword]                   = useState("");
	const [encryptionLevel, setEncryptionLevel]     = useState<"secure" | "advanced">("secure");
	const [statusMessage, setStatusMessage]         = useState("");
	const [loading, setLoading]                     = useState(true);
	//||------------------------------------------------------------------------------------------------||
	//|| Check Session
	//||------------------------------------------------------------------------------------------------||
	useEffect(() => {
		const checkMe = async () => {
			try {
				//||------------------------------------------------------------------------------------------------||
				//|| Fetch
				//||------------------------------------------------------------------------------------------------||
				const res = await fetch("/auth/me", {
					method: "GET",
					credentials: "include",
				});
				//||------------------------------------------------------------------------------------------------||
				//|| Check JSON
				//||------------------------------------------------------------------------------------------------||
				const json = await res.json();
				console.log(json);
                        if (!json.success) {
                              setCheckStatus("FAIL");
                        }
				//||------------------------------------------------------------------------------------------------||
				//|| Check JSON
				//||------------------------------------------------------------------------------------------------||
                        if (json.data.status == "ACTV") {
                              navigate("/members");
                        }
				//||------------------------------------------------------------------------------------------------||
				//|| Verified
				//||------------------------------------------------------------------------------------------------||
                        if (json.data.status == "VERF") {
                              setCheckStatus("VERF");
                        }
				//||------------------------------------------------------------------------------------------------||
				//|| Other
				//||------------------------------------------------------------------------------------------------||
                        if (json.data.status == "ACTV") {
                              setCheckStatus(json.data.status);
                        }

                  } catch (err) {
				setCheckStatus("FAIL");
			}
		};
		checkMe();
	}, []);


	//||------------------------------------------------------------------------------------------------||
	//|| Handle Submit
	//||------------------------------------------------------------------------------------------------||

	const handleSubmit = async (e: React.FormEvent) => {
            //||------------------------------------------------------------------------------------------------||
            //|| Generate the Payload
            //||------------------------------------------------------------------------------------------------||
            e.preventDefault();
		const payload = {
			password    : password,
                  advanced    : encryptionLevel === "advanced" ? "advanced" : "standard",
		};            
            //||------------------------------------------------------------------------------------------------||
            //|| Handle the POST
            //||------------------------------------------------------------------------------------------------||
		try {
			setLoading(true);
			const res = await fetch("/auth/complete", {
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
                  const next = json.data?.next || "/complete";
			if (json.success) return navigate(next);
                  //||------------------------------------------------------------------------------------------------||
                  //|| Failed
                  //||------------------------------------------------------------------------------------------------||
                  setStatusMessage(`❌ ${json.error || "Save failed."}`);
		} catch (err) {
			console.error(err);
			setStatusMessage("❌ Something went wrong. Please try again.");
		} finally {
			setLoading(false);
		}
	};

      //||------------------------------------------------------------------------------------------------||
      //|| Check Status Failed
      //||------------------------------------------------------------------------------------------------||                              
	if (checkStatus === "CHECK") {
		return (
                  <main className="flex flex-col h-[70vw]">
                        <NavMain />
                              <div className="flex-1 flex items-center justify-center bg-gray-500">
                                    <div className="flex flex-col items-center justify-center">
                                          <SpinnerCircle />
                                          <div className="mt-4 text-center">Checking account status...</div>
                                    </div>
                              </div>
                        <FooterMain />
                  </main>
		);
	}
      //||------------------------------------------------------------------------------------------------||
      //|| Check Status Failed
      //||------------------------------------------------------------------------------------------------||                              
	if (checkStatus === "FAIL") {
            return (
                  <main className="flex flex-col min-h-screen">
                     <NavMain />
                     <div className="flex-1 flex flex-col items-center justify-center bg-white text-black px-4">
                        <h1 className="text-2xl font-bold mb-4 text-center">
                           Could not find your session
                        </h1>
                        <p className="mb-8 text-center text-gray-600">
                           Please log in or sign up to continue.
                        </p>
                        <div className="flex space-x-4">
                           <a
                              href="/login"
                              className="px-6 py-3 bg-black text-white rounded-md hover:bg-gray-800 transition"
                           >
                              Log In
                           </a>
                           <a
                              href="/signup"
                              className="px-6 py-3 bg-white text-black border border-black rounded-md hover:bg-gray-100 transition"
                           >
                              Sign Up
                           </a>
                        </div>
                     </div>
                     <FooterMain />
                  </main>
               );
	}      
      //||------------------------------------------------------------------------------------------------||
      //|| JSX / FORM
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
						<h1 className="text-3xl font-bold mb-6 text-primary">
							Complete Your Signup
						</h1>

						<p className="mb-6 text-base-content/70">
							Set your password and encryption
							level.
						</p>

						<form
							onSubmit={handleSubmit}
							className="flex flex-col gap-4 text-left">
							<label className="font-semibold">
								Password
							</label>
							<input
								type="password"
								placeholder="Enter your password"
								value={password}
								onChange={(e) =>
									setPassword(
										e.target.value
									)
								}
								className="input input-bordered w-full"
								required
							/>

							<label className="font-semibold mt-4">
								Encryption Level
							</label>
							<label className="flex items-center gap-2">
								<input
									type="radio"
									value="secure"
									checked={
										encryptionLevel ===
										"secure"
									}
									onChange={() =>
										setEncryptionLevel(
											"secure"
										)
									}
									className="radio radio-primary"
								/>
								Secure (we manage your keys)
							</label>
							<label className="flex items-center gap-2">
								<input
									type="radio"
									value="advanced"
									checked={
										encryptionLevel ===
										"advanced"
									}
									onChange={() =>
										setEncryptionLevel(
											"advanced"
										)
									}
									className="radio radio-secondary"
								/>
								Advanced (you keep your private
								key)
							</label>

							<button
								type="submit"
								className="btn btn-primary w-full mt-4">
								Complete Signup
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
