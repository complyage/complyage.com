//||------------------------------------------------------------------------------------------------||
//|| Function :: authLogout
//||------------------------------------------------------------------------------------------------||

export async function authLogout( navigate: ReturnType<typeof useNavigate> ): Promise<void> {
	try {
		//||------------------------------------------------------------------------------------------------||
		//|| API Call
		//||------------------------------------------------------------------------------------------------||
		const response = await fetch("/auth/logout", {
			method: "GET", // Changed back to POST for security
			headers: {
				"Content-Type": "application/json",
			},
		});

		const data = await response.json();

		//||------------------------------------------------------------------------------------------------||
		//|| Handle Response
		//||------------------------------------------------------------------------------------------------||
		if (response.ok && data.success) {
			window.location.href = "/login";
		} else {
			console.error(
				"Logout API error:",
				data.error || data.message || "Logout failed."
			);
			throw new Error(data.message || "Logout failed."); // Propagate error for calling component to handle
		}
	} catch (err: any) {
		//||------------------------------------------------------------------------------------------------||
		//|| Handle Network Error
		//||------------------------------------------------------------------------------------------------||
		console.error("Logout network error:", err);
		throw new Error(
			err.message ||
				"Network error during logout. Please check your connection."
		);
	}
}
