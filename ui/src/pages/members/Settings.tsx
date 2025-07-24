//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useEffect, useState }                 from "react";

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

      const [currentPassword, setCurrentPassword]                 = useState("");
	const [newPassword, setNewPassword]                         = useState("");
	const [statusMessage, setStatusMessage]                     = useState("");

      //||------------------------------------------------------------------------------------------------||
      //|| Update Password
      //||------------------------------------------------------------------------------------------------||
      
	const handlePasswordUpdate = async () => {
		try {
			const res = await fetch("/members/password", {
				method: "POST",
				headers: {"Content-Type": "application/x-www-form-urlencoded"},
				credentials: "include",
				body: new URLSearchParams({
					currentPassword,
					newPassword,
				}).toString(),
			});

			const json = await res.json();
			if (json.success) {
				setStatusMessage("✅ Password updated successfully.");
				setCurrentPassword("");
				setNewPassword("");
			} else {
				setStatusMessage(`❌ ${json.error}`);
			}
		} catch (err) {
			setStatusMessage("❌ Something went wrong. Please try again.");
		}
	};

      //||------------------------------------------------------------------------------------------------||
      //|| Handle Delete Private Keys
      //||------------------------------------------------------------------------------------------------||

	const handleDeletePrivateKeys = async () => {
		if (!window.confirm("Are you sure you want to delete your private keys?")) return;

		try {
			const res = await fetch("/members/delete-private-keys", {
				method: "POST",
				credentials: "include",
			});
			const json = await res.json();
			if (json.success) {
				setStatusMessage("✅ Private keys deleted.");
			} else {
				setStatusMessage(`❌ ${json.error}`);
			}
		} catch (err) {
			setStatusMessage("❌ Error deleting private keys.");
		}
	};

      //||------------------------------------------------------------------------------------------------||
      //|| Handle Delete Account
      //||------------------------------------------------------------------------------------------------||

	const handleDeleteAccount = async () => {
		if (!window.confirm("This will permanently delete your account. Continue?")) return;

		try {
			const res = await fetch("/members/delete-account", {
				method: "POST",
				credentials: "include",
			});
			const json = await res.json();
			if (json.success) {
				await handleLogout();
			} else {
				setStatusMessage(`❌ ${json.error}`);
			}
		} catch (err) {
			setStatusMessage("❌ Error deleting account.");
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
                              <input
                                    type="password"
                                    placeholder="Current password"
                                    className="input input-bordered w-full mb-4"
                                    value={currentPassword}
                                    onChange={(e) => setCurrentPassword(e.target.value)}
                              />
                              <input
                                    type="password"
                                    placeholder="New password"
                                    className="input input-bordered w-full mb-4"
                                    value={newPassword}
                                    onChange={(e) => setNewPassword(e.target.value)}
                              />
                              <button onClick={handlePasswordUpdate} className="btn btn-primary">
                                    Update Password
                              </button>
                        </div>

                        <h1 className="text-2xl font-bold bg-black mb-10 p-4">Warning! The following actions are highly destructive and can not be reversed!</h1>

                        {/* Delete Private Keys */}
                        <div className="w-full max-w-2xl bg-base-100 shadow-lg rounded-lg p-8 mb-10">
                              <h2 className="text-2xl font-bold mb-4">Delete Private Keys</h2>
                              <p className="mb-4 text-sm text-base-content/70">
                                    This action will permanently delete any stored private keys.
                              </p>
                              <button onClick={handleDeletePrivateKeys} className="btn btn-warning">
                                    Delete Private Keys
                              </button>
                        </div>

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

                        {statusMessage && <div className="w-full max-w-2xl text-sm p-4 bg-base-100 rounded-lg shadow">{statusMessage}</div>}
                  </>
		</MembersLayout>
	);
}
