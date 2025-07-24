(() => {
  // src/utils.ts
  var Utils = class _Utils {
    //||------------------------------------------------------------------------------------------------||
    //|| GSet
    //||------------------------------------------------------------------------------------------------||            
    static gset(name, value) {
      if (typeof window.__complyage__ === "undefined") window.__complyage__ = {};
      if (value !== void 0) window.__complyage__[name] = value;
      return window.__complyage__[name] || void 0;
    }
    //||------------------------------------------------------------------------------------------------||
    //|| getCookie
    //||------------------------------------------------------------------------------------------------||            
    static getCookie(name) {
      const nameEQ = `${encodeURIComponent(name)}=`;
      const parts = document.cookie.split(";");
      for (let part of parts) {
        part = part.trim();
        if (part.startsWith(nameEQ)) return decodeURIComponent(part.substring(nameEQ.length));
      }
      return null;
    }
    //||------------------------------------------------------------------------------------------------||
    //|| CallAPI
    //||------------------------------------------------------------------------------------------------||            
    static async callAPI(handler) {
      const baseUrl = String(_Utils.gset("baseUrl") || "");
      const version = String(_Utils.gset("version") || "");
      const url = `${baseUrl}/${version}/${handler}`;
      alert(`${url}`);
      const response = await fetch(`${url}`, { credentials: "include" });
      if (!response.ok) return null;
      const data = await response.json();
      return data;
    }
    //||------------------------------------------------------------------------------------------------||
    //|| Check Session
    //||------------------------------------------------------------------------------------------------||            
    static async checkSession() {
      const token = this.getCookie("x_complyage");
      if (!token || token === "") return null;
      const data = await this.callAPI(`session?apiKey=${this.gset("apiKey")}&token=${token || ""}`);
      return data;
    }
    //||------------------------------------------------------------------------------------------------||
    //|| Age Gate
    //||------------------------------------------------------------------------------------------------||            
    static ageGate() {
      console.log("WINDOW COMPLYAGE", window.__complyage__);
      const iframe = document.createElement("iframe");
      iframe.src = `http://localhost:8089/v1/gate.html?apiKey=${_Utils.gset("apiKey")}`;
      iframe.classList.add("no-blur");
      Object.assign(iframe.style, {
        position: "fixed",
        top: "0",
        left: "0",
        width: "100%",
        height: "100%",
        border: "none",
        margin: "0",
        padding: "0",
        zIndex: "9999",
        background: "transparent"
      });
      document.body.appendChild(iframe);
    }
    //||------------------------------------------------------------------------------------------------||
    //|| Check Location
    //||------------------------------------------------------------------------------------------------||            
    static async checkLocation() {
      const data = await this.callAPI(`enforce?apiKey=${this.gset("apiKey") || ""}`);
      return data;
    }
    //||------------------------------------------------------------------------------------------------||
    //|| Get Lang
    //||------------------------------------------------------------------------------------------------||            
    static getLang() {
      const locale = navigator.language || navigator.languages && navigator.languages[0] || navigator.userLanguage || "en";
      return locale.split(/[-_]/)[0].toLowerCase();
    }
    //||------------------------------------------------------------------------------------------------||
    //|| EOC
    //||------------------------------------------------------------------------------------------------||            
  };

  // src/index.ts
  var start = async () => {
    console.log(`ComplyAge Client Loader - ${Utils.gset("version")}`);
    console.log(`API Key: ${Utils.gset("apiKey")}`);
    try {
      const verified = await Utils.checkSession();
      if (verified === null) {
        console.log("Session is invalid. Redirecting to age gate.");
        Utils.ageGate();
      } else {
        console.log("Session is valid.");
      }
    } catch (error) {
      console.error("An error occurred:", error);
    }
  };
  (async () => {
    await start();
  })();
})();
