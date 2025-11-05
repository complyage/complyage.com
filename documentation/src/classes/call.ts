//||------------------------------------------------------------------------------------------------||
//|| classes/client/call.ts
//|| Call : Handles all calls to server-http
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| API
      //||------------------------------------------------------------------------------------------------||

      import apiURL                                                 from "../utils/apiURL";

      //||------------------------------------------------------------------------------------------------||
      //|| Types
      //||------------------------------------------------------------------------------------------------||

      export type CallStatuses = "PENDING" | "SUCCESS" | "ERROR" | "POLLING";

      //||------------------------------------------------------------------------------------------------||
      //|| API Response
      //||------------------------------------------------------------------------------------------------||
      
      type ApiResponse = {
            success     : boolean;
            data?       : Record<string, any>;
            message?    : string;
            redirect?   : string;
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Interfaces - Response Payload
      //||------------------------------------------------------------------------------------------------||

      export interface ResponseDataPayload {
            status       : number;
            message      : string;
            headers?     : Record<string, string>;
            redirect?    : string;
            cookies?     : string[];
            data         : Record<string, any>;
            route        : string;
            ttl?         : number;
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Mock Options
      //||------------------------------------------------------------------------------------------------||

      export interface MockOptions {
            message?     : string;
            headers?     : Record<string, string>;
            redirect?    : string;
            route?       : string;
            ttl?         : number;
            delay?       : number;
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Class
      //||------------------------------------------------------------------------------------------------||

      export default class Call {

            //||------------------------------------------------------------------------------------------------||
            //|| Static / Constants
            //||------------------------------------------------------------------------------------------------||

            public domain                         : string;
            public debug                          : boolean   = false;

            //||------------------------------------------------------------------------------------------------||
            //|| Request
            //||------------------------------------------------------------------------------------------------||

            public method                         : string    = "POST";
            public contentType                    : string    = "application/json";
            public secure                         : boolean   = true;
            public server                         : string    = "http://localhost:3000";
            public route                          : string    = "";
            public url                            : string    = "";
            public requestData                    : Record<string, object | boolean | string | number | null | undefined> = {};
            public headers                        : Headers   = new Headers();

            //||------------------------------------------------------------------------------------------------||
            //|| Response Payload
            //||------------------------------------------------------------------------------------------------||

            public responsePayload                : ResponseDataPayload = {
                  status       : 100,
                  message      : "PENDING",
                  headers      : {},
                  redirect     : "",
                  data         : {},
                  route        : "",
                  ttl          : 0,
            }

            //||------------------------------------------------------------------------------------------------||
            //|| State
            //||------------------------------------------------------------------------------------------------||

            public status                         : CallStatuses = "PENDING";
            public http                           : number = 100;

            //||------------------------------------------------------------------------------------------------||
            //|| Assignable Status Callback
            //||------------------------------------------------------------------------------------------------||

            public onState                        : (status: CallStatuses, response: ResponseDataPayload) => void = () => { return; };

            //||------------------------------------------------------------------------------------------------||
            //|| Mocking
            //||------------------------------------------------------------------------------------------------||

            public mockStatus                     : number    = 200;
            public mockMode                       : boolean   = false;
            public mockData                       : any       = {};
            public mockDelay                      : number    = 0;

            //||------------------------------------------------------------------------------------------------||
            //|| Constructor
            //||------------------------------------------------------------------------------------------------||

            constructor(route: string, data: Record<string, object | boolean | string | number | null | undefined> = {}) {
                  //||------------------------------------------------------------------------------------------------||
                  //|| Init Headers
                  //||------------------------------------------------------------------------------------------------||
                  if (!this.headers || !(this.headers instanceof Headers)) this.headers = new Headers();
                  //||------------------------------------------------------------------------------------------------||
                  //|| Setup the API Server
                  //||------------------------------------------------------------------------------------------------||
                  this.route                          = route.startsWith("/") ? route : `/${route}`;
                  this.requestData                    = data;
                  this.domain                         = this.hostname();
            }

            //||------------------------------------------------------------------------------------------------||
            //|| Set Payload
            //||------------------------------------------------------------------------------------------------||

            public setData(data: Record<string, object | boolean | string | number | null | undefined> = {}) {
                  this.requestData = data;
            }

            //||------------------------------------------------------------------------------------------------||
            //|| Mock Data
            //||------------------------------------------------------------------------------------------------||

            public mock(data: any, status = 200, options: MockOptions = {}) {
                  this.mockMode                      = true;
                  this.responsePayload               = {
                        status    : status,
                        message   : options.message || "MOCK MESSAGE",
                        headers   : options.headers || {},
                        redirect  : options.redirect || "",
                        data      : data,
                        route     : options.route || this.route,
                        ttl       : options.ttl || 0,
                  }
                  this.mockDelay                     = options.delay ?? 1;
            }

            //||------------------------------------------------------------------------------------------------||
            //|| Delay on Mock
            //||------------------------------------------------------------------------------------------------||

            public delay(delay: number) {
                  this.mockDelay = delay;
            }

            //||------------------------------------------------------------------------------------------------||
            //|| Execute
            //||------------------------------------------------------------------------------------------------||

            public async execute(): Promise<ResponseDataPayload> {

                  //||------------------------------------------------------------------------------------------------||
                  //|| Route
                  //||------------------------------------------------------------------------------------------------||

                  this.url = apiURL(this.route);

                  //||------------------------------------------------------------------------------------------------||
                  //|| Debug
                  //||------------------------------------------------------------------------------------------------||

                  if (this.debug) {
                        console.log("");
                        console.log("||================== CHIRP BEGIN ======================||");
                        console.log("URL   : ", this.url);
                  }

                  //||------------------------------------------------------------------------------------------------||
                  //|| Timing
                  //||------------------------------------------------------------------------------------------------||

                  const start = (typeof performance !== "undefined" ? performance.now() : Date.now());

                  //||------------------------------------------------------------------------------------------------||
                  //|| Content 
                  //||------------------------------------------------------------------------------------------------||

                  if (this.hasBody(this.method)) {
                        this.headers.set("Content-Type", this.contentType);
                  }

                  //||------------------------------------------------------------------------------------------------||
                  //|| Mock Mode
                  //||------------------------------------------------------------------------------------------------||

                  if (this.mockMode) {
                        this.updateStatus("POLLING");
                        await this.sleep(this.mockDelay);
                        this.http                          = this.responsePayload.status;
                        this.responsePayload.message       = `MOCK[${this.responsePayload.message}]`;
                        this.updateStatus((this.responsePayload.status === 200) ? "SUCCESS" : "ERROR");
                        this.responsePayload.ttl           = (typeof performance !== "undefined" ? performance.now() : Date.now()) - start;
                        this.debugPayload();
                        return this.responsePayload;
                  }

                  //||------------------------------------------------------------------------------------------------||
                  //|| Body
                  //||------------------------------------------------------------------------------------------------||

                  const encodedBody = this.encodeBody(this.requestData);

                  //||------------------------------------------------------------------------------------------------||
                  //|| Fetch
                  //||------------------------------------------------------------------------------------------------||

                  let response: Response | undefined;
                  try {
                        this.updateStatus("POLLING");
                        response = await fetch(this.url, {
                              method      : this.method,
                              headers     : this.headers,
                              credentials : "include",
                              body        : this.hasBody(this.method) ? encodedBody : undefined,
                        });
                  } catch (error) {
                        if (this.debug) {
                              console.error("CHIRP FETCH ERROR:", error);
                              console.log("||=================== CHIRP END =======================||");
                        }
                        this.responsePayload.status    = 501;
                        this.responsePayload.message   = "CHP_UNKNOWN";
                        this.responsePayload.headers   = {};
                        this.responsePayload.data      = {};
                        this.responsePayload.route     = this.route;
                        this.responsePayload.ttl       = (typeof performance !== "undefined" ? performance.now() : Date.now()) - start;
                        this.updateStatus("ERROR");
                        this.debugPayload();
                        return this.responsePayload;
                  }

                  //||------------------------------------------------------------------------------------------------||
                  //|| Parse JSON
                  //||------------------------------------------------------------------------------------------------||

                  let responseJSON: ApiResponse | undefined;
                  try {
                        responseJSON = await response!.json();
                  } catch {
                        try {
                              const txt = await response!.text();
                              this.responsePayload.message = txt || "CHP_BAD_JSON";
                        } catch {
                              this.responsePayload.message = "CHP_BAD_JSON";
                        }
                        this.responsePayload.status    = response!.status;
                        this.responsePayload.headers   = this.makeHeaders(response!.headers);
                        this.responsePayload.route     = this.route;
                        this.responsePayload.ttl       = (typeof performance !== "undefined" ? performance.now() : Date.now()) - start;
                        this.updateStatus("ERROR");
                        this.http                      = response!.status;
                        this.debugPayload();
                        return this.responsePayload;
                  }

                  //||------------------------------------------------------------------------------------------------||
                  //|| Map Response
                  //||------------------------------------------------------------------------------------------------||

                  if (responseJSON?.success) {
                        this.responsePayload = {
                              status   : response!.status,
                              message  : "",
                              headers  : this.makeHeaders(response!.headers),
                              redirect : "",
                              data     : (responseJSON.data || {}) as Record<string, any>,
                              route    : this.route,
                              ttl      : (typeof performance !== "undefined" ? performance.now() : Date.now()) - start,
                        };
                        this.updateStatus("SUCCESS");
                  } else {
                        this.responsePayload = {
                              status   : response!.status,
                              message  : responseJSON?.message || "Unknown error",
                              headers  : this.makeHeaders(response!.headers),
                              redirect : "",
                              data     : {},
                              route    : this.route,
                              ttl      : (typeof performance !== "undefined" ? performance.now() : Date.now()) - start,
                        };
                        this.updateStatus("ERROR");
                  }

                  //||------------------------------------------------------------------------------------------------||
                  //|| Finalize
                  //||------------------------------------------------------------------------------------------------||

                  this.http = response!.status;
                  this.debugPayload();
                  return this.responsePayload;
            }

            //||------------------------------------------------------------------------------------------------||
            //|| HTTP Verbs
            //||------------------------------------------------------------------------------------------------||

            public async get()    : Promise<ResponseDataPayload> { this.method = "GET";    return await this.execute(); }
            public async post()   : Promise<ResponseDataPayload> { this.method = "POST";   return await this.execute(); }
            public async put()    : Promise<ResponseDataPayload> { this.method = "PUT";    return await this.execute(); }
            public async delete() : Promise<ResponseDataPayload> { this.method = "DELETE"; return await this.execute(); }
            public async patch()  : Promise<ResponseDataPayload> { this.method = "PATCH";  return await this.execute(); }

            //||------------------------------------------------------------------------------------------------||
            //|| State Updates
            //||------------------------------------------------------------------------------------------------||

            public updateStatus(state: CallStatuses) {
                  this.status = state;
                  if (this.onState) this.onState(this.status, this.responsePayload);
            }

            //||------------------------------------------------------------------------------------------------||
            //|| Headers -> Record
            //||------------------------------------------------------------------------------------------------||

            public makeHeaders(headers: Headers): Record<string, string> {
                  const result: Record<string, string> = {};
                  headers.forEach((value, key) => { result[key] = value; });
                  return result;
            }

            //||------------------------------------------------------------------------------------------------||
            //|| Data Helpers
            //||------------------------------------------------------------------------------------------------||

            public hostname = (): string => {
                  if (typeof window === "undefined") return "";
                  return window.location.hostname.replace(/^www\./i, "").toLowerCase();
            }

            public ok(): boolean {
                  if (this.status !== "SUCCESS") return false;
                  return this.responsePayload.status >= 200 && this.responsePayload.status < 300;
            }

            public error(): string {
                  return this.responsePayload.message || "CHP_UNKNOWN";
            }

            public data<T = any>(key: string): T | undefined {
                  if (this.status !== "SUCCESS") return undefined;
                  if (this.responsePayload.data && (this.responsePayload.data as any)[key] !== undefined) {
                        return (this.responsePayload.data as any)[key] as T;
                  }
                  return undefined;
            }

            //||------------------------------------------------------------------------------------------------||
            //|| Debug Payload
            //||------------------------------------------------------------------------------------------------||

            public debugPayload() {
                  if (!this.debug) return;
                  console.groupCollapsed(`%c🐦 RESPONSE ${this.route} [${this.status}]`, "color: dodgerblue; font-weight: bold;");
                  console.log("%c→ URL     ", "color: gray;", this.url);
                  console.log("%c→ SERVER  ", "color: gray;", this.server);
                  console.log("%c→ HEADERS ", "color: gray;", Object.fromEntries(this.headers.entries()));
                  console.log("%c→ DATA    ", "color: gray;", this.requestData);
                  console.log("%c→ STATUS  ", "color: gray;", this.status);
                  console.log("%c→ PAYLOAD ", "color: gray;", this.responsePayload);
                  console.groupEnd();
                  console.log("||=================== CHIRP END =======================||");
                  console.log("");
            }

            //||------------------------------------------------------------------------------------------------||
            //|| Static Response Helper
            //||------------------------------------------------------------------------------------------------||

            static response(status: number, message: string, data?: any) {
                  return { status, message, data };
            }

            //||------------------------------------------------------------------------------------------------||
            //|| Private / Internals
            //||------------------------------------------------------------------------------------------------||

            private hasBody(method: string): boolean {
                  const m = method.toUpperCase();
                  return m === "POST" || m === "PUT" || m === "PATCH" || m === "DELETE";
            }

            private encodeBody(payload: Record<string, any>): string | undefined {
                  if (!this.hasBody(this.method)) return undefined;

                  if (this.contentType === "application/x-www-form-urlencoded") {
                        const formObj: Record<string, string> = {};
                        for (const k in payload) {
                              if (!Object.prototype.hasOwnProperty.call(payload, k)) continue;
                              const val = payload[k];
                              if (val !== null && val !== undefined) formObj[k] = String(val);
                        }
                        return new URLSearchParams(formObj).toString();
                  }

                  // default json
                  return JSON.stringify(payload ?? {});
            }

            private sleep(ms: number): Promise<void> {
                  const delay = Math.max(0, Number(ms) || 0);
                  return new Promise((resolve) => setTimeout(resolve, delay));
            }

            private normalizeServer(input: string): string {
                  const trimmed = input.trim();
                  if (/^https?:\/\//i.test(trimmed)) return trimmed.replace(/\/+$/,"");
                  return `https://${trimmed.replace(/\/+$/,"")}`;
            }

            private joinUrl(base: string, path: string): string {
                  const b = base.replace(/\/+$/,"");
                  const p = path.startsWith("/") ? path : `/${path}`;
                  return `${b}${p}`;
            }
      }
