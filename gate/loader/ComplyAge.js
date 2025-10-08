
//||------------------------------------------------------------------------------------------------||
//|| ComplyAge Class
//||------------------------------------------------------------------------------------------------||

      export class ComplyAge {


            //||------------------------------------------------------------------------------------------------||
            //|| Setup
            //||------------------------------------------------------------------------------------------------||

            SETUP = {
                  clientId    : null,
                  publicKey   : null,
                  debug       : false,
                  version     : null,
                  versionPath : null,
                  urlGate     : null,
                  urlOAuth    : null,
            };

            //||------------------------------------------------------------------------------------------------||
            //|| JWT
            //||------------------------------------------------------------------------------------------------||

            DYNAMIC = {
                  registered : false,
                  sessionId : null,
                  created   : null,
                  expires   : null,
                  ipAddress : null,
                  verified  : false,
                  enforced  : false,
                  age       : null,
                  status    : "INIT",
            };

            //||------------------------------------------------------------------------------------------------||
            //|| Construct
            //||------------------------------------------------------------------------------------------------||

            constructor(setup) {
                  console.log("\x1b[1m\x1b[32m[ComplyAge] Starting...\x1b[0m");                             
                  console.log(`\x1b[1m\x1b[31mComplyAge v${setup.version} \x1b[0m`);
                  console.log(`\x1b[1m\x1b[31m[ComplyAge] Client ID: ${setup.clientId}\x1b[0m`);
                  this.SETUP.clientId      = setup.clientId;
                  this.SETUP.debug         = setup.debug === true;
                  this.SETUP.version       = setup.version || 1.0;
                  this.SETUP.versionPath   = "v" + Math.floor(setup.version);
                  this.SETUP.urlGate       = setup.urlGate;
                  this.SETUP.urlOAuth      = setup.urlOAuth;
                  this.reset();
                  if (this.SETUP.debug === true) console.log("\x1b[1m\x1b[33m[ComplyAge] Debug mode enabled\x1b[0m");
            };

            //||------------------------------------------------------------------------------------------------||
            //|| Register
            //||------------------------------------------------------------------------------------------------||

            register() {
                  this.defunc("register");
                  if (this.DYNAMIC.registered === true) return;
                  this.DYNAMIC.registered = true;
                  window.addEventListener("message", (event) => {
                        if (event.data?.action === "complyAgeReload") {
                              window.location.reload();
                        }
                  });               
            };

            //||------------------------------------------------------------------------------------------------||
            //|| Call
            //||------------------------------------------------------------------------------------------------||

            async authorize() {
                  this.defunc("authorize");
                  //||------------------------------------------------------------------------------------------------||
                  //|| Register the Listener
                  //||------------------------------------------------------------------------------------------------||
                  this.register()
                  //||------------------------------------------------------------------------------------------------||
                  //|| Load Session
                  //||------------------------------------------------------------------------------------------------||
                  await this.call("check");
                  this.logger("API Call", "Load complete");
                  //||------------------------------------------------------------------------------------------------||
                  //|| Verified
                  //||------------------------------------------------------------------------------------------------||
                  if (this.DYNAMIC.verified === true) return true;
                  //||------------------------------------------------------------------------------------------------||
                  //|| Enforced
                  //||------------------------------------------------------------------------------------------------||
                  if (this.DYNAMIC.enforced === false) return true;
                  //||------------------------------------------------------------------------------------------------||
                  //|| Not Verified && Is Enforced
                  //||------------------------------------------------------------------------------------------------||
                  return this.ageGate();
            }

            //||------------------------------------------------------------------------------------------------||
            //|| Call
            //||------------------------------------------------------------------------------------------------||

            async call(handler, args) {
                  this.defunc("apiCall", handler, args);
                  const version     = this.SETUP.version;
                  const returnURL   = btoa(window.location.href);
                  const url         = this.SETUP.urlGate + "/" + this.SETUP.versionPath + "/" + handler + "?client_id=" + this.SETUP.clientId + "&return_url=" + returnURL;
                  const start = performance.now();
                  let response;
                  try {
                        response = await fetch(url, {
                              method      : args ? "POST" : "GET",
                              headers     : {"Content-Type": "application/json"},
                              credentials : "include",
                              body        : args ? JSON.stringify(args) : undefined,
                        }); 
                  } catch (err) {
                        this.logger("[ComplyAge] fetch failed", err);
                        return this.errorGate("ComplyAge API error: Network error");
                  }      
                  let json;
                  try {
                        json = await response.json();
                  } catch (err) {
                        this.logger("[ComplyAge] response json", json);
                        return this.errorGate("ComplyAge API error: " + (json.message || "Unknown error"));
                  }
                  const elapsed = performance.now() - start;
                  this.logger(`[ComplyAge] API ${handler} took ${elapsed.toFixed(2)} ms`);                  
                  //||------------------------------------------------------------------------------------------------||
                  //|| Set the Dynamic
                  //||------------------------------------------------------------------------------------------------||
                  this.DYNAMIC.sessionId = json.keeperId || null,
                  this.DYNAMIC.ipAddress = json.ipAddress || null,
                  this.DYNAMIC.verified  = json.verified || false;
                  this.DYNAMIC.enforced  = json.enforced || false;
                  this.DYNAMIC.age       = json.age || 0;
                  this.DYNAMIC.status    = json.status || "ERROR";
                  this.logger("Check complete", this.DYNAMIC);
            }

            //||------------------------------------------------------------------------------------------------||
            //|| Reset
            //||------------------------------------------------------------------------------------------------||

            reset() {
                  this.defunc("reset");
                  this.DYNAMIC.sessionId = null;
                  this.DYNAMIC.created   = null;
                  this.DYNAMIC.expires   = null;
                  this.DYNAMIC.ipAddress = null;
                  this.DYNAMIC.verified  = false;
                  this.DYNAMIC.enforced  = true;
                  this.DYNAMIC.age       = 0;
                  this.DYNAMIC.status    = "INIT";
                  return null;
            }

            //||------------------------------------------------------------------------------------------------||
            //|| Overlay
            //||------------------------------------------------------------------------------------------------||

            ageGate() {
                  this.defunc("ageGate");
                  const src = `${this.SETUP.urlOAuth}/${this.SETUP.versionPath}/age?client_id=${this.SETUP.clientId}`;
                  const spinner = document.createElement("div");
                  Object.assign(spinner.style, {
                        position    : "fixed",
                        top         : "0",
                        left        : "0",
                        width       : "100%",
                        height      : "100%",
                        display     : "flex",
                        alignItems  : "center",
                        justifyContent: "center",
                        background  : "rgba(0,0,0,0.6)",
                        zIndex      : "9998"
                  });
                  spinner.innerHTML = `
                        <div style="
                              width:48px; height:48px;
                              border: 6px solid #f3f3f3;
                              border-top: 6px solid #3498db;
                              border-radius: 50%;
                              animation: spin 1s linear infinite;
                        "></div>
                  `;

                  const style = document.createElement("style");
                  style.textContent = `
                        @keyframes spin {
                              0% { transform: rotate(0deg); }
                              100% { transform: rotate(360deg); }
                        }
                  `;
                  document.head.appendChild(style);
                  document.body.appendChild(spinner);
                  const iframe = document.createElement("iframe");
                  iframe.src = src;
                  iframe.classList.add("no-blur");
                  Object.assign(iframe.style, {
                        position    : "fixed",
                        top         : "0",
                        left        : "0",
                        width       : "100%",
                        height      : "100%",
                        border      : "none",
                        margin      : "0",
                        padding     : "0",
                        zIndex      : "9999",
                        background  : "transparent",
                  });
                  iframe.onload = () => {
                        if (spinner && spinner.parentNode) {
                              spinner.parentNode.removeChild(spinner);
                        }
                  };
                  document.body.appendChild(iframe);
            }


            //||------------------------------------------------------------------------------------------------||
            //|| Error
            //||------------------------------------------------------------------------------------------------||

            errorGate(message) {
                  this.defunc("ageGate");
                  const container = document.createElement("div");
                  container.classList.add("complyage-error");
                  Object.assign(container.style, {
                        position     : "fixed",
                        bottom       : "20px",
                        left         : "50%",
                        transform    : "translateX(-50%)",
                        background   : "#f44336",
                        color        : "#fff",
                        padding      : "12px 20px",
                        borderRadius : "8px",
                        boxShadow    : "0 4px 8px rgba(0,0,0,0.2)",
                        fontFamily   : "system-ui, sans-serif",
                        fontSize     : "14px",
                        maxWidth     : "90%",
                        zIndex       : "10000",
                        display      : "flex",
                        alignItems   : "center",
                        gap          : "12px",
                  });
                  const text = document.createElement("span");
                  text.textContent = message;
                  text.style.flexGrow = "1";
                  const button = document.createElement("button");
                  button.textContent = "×";
                  Object.assign(button.style, {
                        background   : "transparent",
                        border       : "none",
                        color        : "#fff",
                        fontSize     : "18px",
                        fontWeight   : "bold",
                        cursor       : "pointer",
                        lineHeight   : "1",
                  });
                  button.addEventListener("click", () => container.remove());
                  container.appendChild(text);
                  container.appendChild(button);
                  document.body.appendChild(container);
                  return null;
            }

            //||------------------------------------------------------------------------------------------------||
            //|| Debug
            //||------------------------------------------------------------------------------------------------||

            defunc(functionName, ...args){
                  if (!this.SETUP.debug) return;
                  console.log(`\x1b[34m[COMPLYAGE]->[${functionName}]\x1b[0m`, ...args);
            }

            logger(...args){
                  if (!this.SETUP.debug) return;
                  console.log("[COMPLYAGE]", ...args);
            }            
            //||------------------------------------------------------------------------------------------------||
            //|| EOF
            //||------------------------------------------------------------------------------------------------||

      }
