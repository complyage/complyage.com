//||------------------------------------------------------------------------------------------------||
//|| Dashboard Data
//||------------------------------------------------------------------------------------------------||

export interface DashboardData {
      isVerified              : boolean;
      verifiedAge             : number;
      minimumType             : VerificationTypes;
      ipAddress               : string;
      location    : {
            city              : string;
            region            : string;
            country           : string;
            latitude?         : number;
            longitude?        : number;
      },
      zone : {
            laws              : string;
            requirements      : VerificationTypes[];
            effective         : string;
            minAge            : number;
      },
      identity                : Identity;
}
