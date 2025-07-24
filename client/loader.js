window._complyage = {
	apiKey: "",
	apiUrl: "localhost:8888/v1/api",
	debug: false,
	language: (navigator.language || navigator.userLanguage || defaultLang).split("-")[0].toLowerCase(),
	salt: "",
	sessionId: decodeURIComponent((document.cookie.match(/(?:^|; )complyage_session=([^;]*)/) || [])[1] || ""),
};
(async () => {
	const cfg = window._complyage;
	const {apiUrl, sessionId} = cfg;
	const headers = {"Content-Type": "application/json"};

	// 1. If we have a session token, verify it's still valid
	if (sessionId) {
		try {
			const res = await fetch(`${apiUrl}/verify`, {
				method: "POST",
				headers: {
					...headers,
					Authorization: `Bearer ${sessionId}`,
				},
			});
			if (res.ok) {
				// token is valid → nothing to do
				return;
			}
		} catch (e) {
			console.warn("ComplyAge verify failed:", e);
		}
	}

	// 2. Otherwise, ask backend if this session needs age verification
	let needsGate = false;
	try {
		const res = await fetch(`${apiUrl}/sessions/${sessionId || ""}/status`, {
			method: "GET",
			headers,
		});
		if (res.ok) {
			const {needsVerification, enforce} = await res.json();
			needsGate = needsVerification && enforce;
		}
	} catch (e) {
		console.error("ComplyAge status check failed:", e);
	}

	// 3. If we do need to show the gate, fetch your modal HTML and inject it
	if (needsGate) {
		try {
			const modalHtml = await fetch(`${apiUrl}/modal`, {
				method: "GET",
				headers,
			}).then((r) => r.text());
			document.body.insertAdjacentHTML("beforeend", modalHtml);
		} catch (e) {
			console.error("ComplyAge modal fetch failed:", e);
		}
	}
})();
