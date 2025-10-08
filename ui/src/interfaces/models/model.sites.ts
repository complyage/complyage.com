//||------------------------------------------------------------------------------------------------||
//|| Interfaces : Site
//|| Model Site
//||------------------------------------------------------------------------------------------------||

//||------------------------------------------------------------------------------------------------||
//|| Site Scope
//||------------------------------------------------------------------------------------------------||

export interface SiteScope {
	code    : string;
	status  : string;
	enabled : boolean;
}

//||------------------------------------------------------------------------------------------------||
//|| Site Zone
//||------------------------------------------------------------------------------------------------||

export interface SiteZone {
	zone     : number;
	enforced : boolean;
}

//||------------------------------------------------------------------------------------------------||
//|| Model
//||------------------------------------------------------------------------------------------------||

export interface ModelSite {
            id                : number;
            fid_account       : string;
            name              : string;
            logo              : string;
            description       : string;
            url               : string;
            status            : string;
            enforcement       : "ALLZ" | "REGU" | "CSTM";
            zones             : SiteZone[];
            domains           : string;
            clientId          : string;
            private           : string;
            public            : string;
            redirect          : string;
            scopeAuto         : boolean;
            scopes            : SiteScope[];
            created           : string;
            gateSignup        : boolean;
            gateConfirm       : string;
            gateExit          : string;
            updated           : string;
            forceUpdate?      : string;
      }