//||---------------------------------------------------------------------------------------------||
//|| Donation API Types
//||---------------------------------------------------------------------------------------------||

export interface DonateCryptoOption {
      name    : string;
      symbol  : string;
      address : string;
      prefix  : string;
      qr      : string;
      color?  : string;
}

export interface DonateAddress {
      name      : string;
      address1  : string;
      address2?: string;
      city      : string;
      state     : string;
      postal    : string;
      country   : string;
}

export interface DonateApiResponse {
      crypto      : DonateCryptoOption[];
      address     : DonateAddress;
}
