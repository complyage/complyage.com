(() => {
  // src/env.ts
  var Env = class {
    //||------------------------------------------------------------------------------------------------||
    //|| Version
    //||------------------------------------------------------------------------------------------------||            
    static version = "0.1.0";
    static clientVersion = "v1";
    static clientURL = `//localhost:8089/${this.clientVersion}/client/`;
    //||------------------------------------------------------------------------------------------------||
    //|| EOC
    //||------------------------------------------------------------------------------------------------||            
  };

  // src/utils.ts
  var Utils = class {
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
      const response = await fetch(Env.clientURL + handler, { credentials: "include" });
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
    try {
      console.log(`ComplyAge Client Loader - ${Env.clientVersion}`);
      console.log(`API Key: ${Utils.gset("apiKey")}`);
      const data = await Utils.checkSession();
      if (data === null) {
        console.log("No Session Cookie Found. Checking Zone/Public Key");
      }
      const location = await Utils.checkLocation();
      console.log("AGE GATE", location.ageGate);
      if (location.enforce) {
        const blurAmount = "8px";
        const style = document.createElement("style");
        style.textContent = `
                body > *:not(.no-blur) {
                  filter: blur(${blurAmount});
                  transition: filter 0.3s ease;
                }
              `;
        document.head.appendChild(style);
        const iframe = document.createElement("iframe");
        iframe.src = `http://localhost:8089/verify?apiKey=${encodeURIComponent(apiKey)}`;
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
    } catch (error) {
      console.error("An error occurred:", error);
    }
  };
  (async () => {
    await start();
  })();
})();
