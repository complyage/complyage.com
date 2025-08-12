//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useEffect, useState }                 from "react";
import { useNavigate }                                from "react-router-dom";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import MembersLayout                                  from "../../layouts/MembersLayout";

//||------------------------------------------------------------------------------------------------||
//|| Default
//||------------------------------------------------------------------------------------------------||

export default function MembersSettings() { 

      //||------------------------------------------------------------------------------------------------||
      //|| Const
      //||------------------------------------------------------------------------------------------------||

      const navigate = useNavigate();

      //||------------------------------------------------------------------------------------------------||
      //|| Const
      //||------------------------------------------------------------------------------------------------||

      const [password, setPassword]                   = useState("");
	const [confirmPassword, setConfirmPassword]     = useState("");
      const [statusMessage, setStatusMessage]         = useState<React.ReactNode>("");

      //||------------------------------------------------------------------------------------------------||
      //|| Alert
      //||------------------------------------------------------------------------------------------------||

      const showAlert = (message: string, type: "success" | "error" = "success") => {
            const alertBox = (
                  <div className={`w-full max-w-2xl text-sm p-4 ${type === "success" ? "bg-green-500" : "bg-red-500"} text-white mb-4 rounded-lg shadow`}>
                        {message}
                  </div>
            );
            setStatusMessage(alertBox);
            setTimeout(() => setStatusMessage(""), 5000);
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Update Password
      //||------------------------------------------------------------------------------------------------||
      
	const handleLogout = async () => {
		try {
			const res = await fetch("/auth/logout", {
				method: "POST",
				headers: {"Content-Type": "application/x-www-form-urlencoded"},
				credentials: "include",
			});

			const json = await res.json();
			if (json.success) {
                        navigate("/auth/login");
                        
			} else showAlert("Logout failed", "error");

		} catch (err) {
                  showAlert("Something went wrong. Please try again.", "error");
		}
	};

      //||------------------------------------------------------------------------------------------------||
      //|| Update Password
      //||------------------------------------------------------------------------------------------------||
      
      const handlePasswordUpdate = async () => {
            if (!password || !confirmPassword) {
                  showAlert("Please fill out both password fields.", "error");
                  return;
            }
            if (password !== confirmPassword) {
                  showAlert("Passwords do not match.", "error");
                  return;
            }

            try {
                  const res = await fetch("/user/reset", {
                        method: "POST",
                        headers: { "Content-Type": "application/x-www-form-urlencoded" },
                        credentials: "include",
                        body: new URLSearchParams({
                              password,
                              confirmPassword,
                        }).toString(),
                  });

                  const json = await res.json();
                  if (json.success) {
                        showAlert("Password updated successfully.");
                        setPassword("");
                        setConfirmPassword("");
                  } else {
                        showAlert(json.message || "Failed to update password", "error");
                  }
            } catch {
                  showAlert("Something went wrong. Please try again.", "error");
            }
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Handle Delete Account
      //||------------------------------------------------------------------------------------------------||

      const handleDeleteAccount = async () => {
            if (!window.confirm("This will permanently delete your account. Continue?")) return;
            try {
                  navigate("/members/quit"); // Ensure user is logged out before deleting account
            } catch {
                  showAlert("❌ Error deleting account.", "error");
            }
      };

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

	return (
		<MembersLayout title="Settings">
                  <>  
                        {/* Change Password */}
                        <div className="w-full max-w-2xl bg-base-100 shadow-lg rounded-lg p-8 mb-10">
                              <h2 className="text-2xl font-bold mb-4">Change Password</h2>
                              {statusMessage}
                              <input
                                    type="password"
                                    placeholder="Current password"
                                    className="input input-bordered w-full mb-4"
                                    value={password}
                                    onChange={(e) => setPassword(e.target.value)}
                              />
                              <input
                                    type="password"
                                    placeholder="New password"
                                    className="input input-bordered w-full mb-4"
                                    value={confirmPassword}
                                    onChange={(e) => setConfirmPassword(e.target.value)}
                              />
                              <button onClick={handlePasswordUpdate} className="btn btn-primary">
                                    Update Password
                              </button>
                        </div>

                        <h1 className="text-2xl font-bold bg-black mb-10 p-4">Warning! The following actions are highly destructive and can not be reversed!</h1>


                        {/* Delete Account */}
                        <div className="w-full max-w-2xl bg-base-100 shadow-lg rounded-lg p-8">
                              <h2 className="text-2xl font-bold mb-4">Delete Account</h2>
                              <p className="mb-4 text-sm text-base-content/70">
                                    This action is irreversible. All your data will be deleted permanently.
                              </p>
                              <button onClick={handleDeleteAccount} className="btn btn-error">
                                    Delete Account
                              </button>
                        </div>
                  </>
		</MembersLayout>
	);
}
