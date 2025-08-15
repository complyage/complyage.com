//||------------------------------------------------------------------------------------------------||
//|| Interfaces : Site
//|| Model Shared
//||------------------------------------------------------------------------------------------------||

      export interface ModelShared {
            id_shared                  : number;
            fid_site                   : number;    
            fid_verification           : number;
            shared_timestamp           : string;
            site_name                  : string;
            site_url                   : string;
            verification_type          : string;
            verification_status        : string;
      }