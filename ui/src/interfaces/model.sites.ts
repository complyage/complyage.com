export interface Site {
      id                : number;
      fid_account       : string;
      name              : string;
      logo              : string;
      description       : string;
      url               : string;
      status            : string;
      enforcement       : "ALLZ" | "REGU" | "CSTM";
      zones             : Record<number, number> | null;
      domains           : string;
      private           : string;
      public            : string;
      redirect          : string;
      permissions       : string;
      created           : string;
      gateSignup        : boolean;
      gateConfirm       : string;
      gateExit          : string;
      updated           : string;
      forceUpdate?      : string;
}