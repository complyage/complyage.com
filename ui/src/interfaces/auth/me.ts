//||------------------------------------------------------------------------------------------------||
//|| Identity
//||------------------------------------------------------------------------------------------------||

import { Identity }             from "../verify/identity/identity";

//||------------------------------------------------------------------------------------------------||
//|| Auth/Me
//||------------------------------------------------------------------------------------------------||

export interface AuthMe {
      id             : number;
      status         : string;
      type           : string;
      email          : string;
      username       : string;
      level          : number;
      security       : number;
      created        : number;
      expires        : number;
      identity       : Identity;
}