//||------------------------------------------------------------------------------------------------||
//|| Types
//||------------------------------------------------------------------------------------------------||

export type VerificationTypeKeys =
      | "Email"
      | "Phone"
      | "Age"
      | "Address"
      | "CreditCard"
      | "ProfilePhoto"
      | "Username";

export type VerificationTypeMap = Record<VerificationTypeKeys, string>;

//||------------------------------------------------------------------------------------------------||
//|| Constants
//||------------------------------------------------------------------------------------------------||

export const VerificationTypes: VerificationTypeMap = {
      Email:        "MAIL",
      Phone:        "PHNE",
      Age:          "UAGE",
      Address:      "ADDR",
      CreditCard:   "CRCD",
      ProfilePhoto: "PROF",
      Username:     "UNAM",
};

//||------------------------------------------------------------------------------------------------||
//|| Utils
//||------------------------------------------------------------------------------------------------||

export function getAllVerificationTypes(): string[] {
      return Object.values(VerificationTypes);
}

export function getVerificationTypeMap(): Record<string, string> {
      return { ...VerificationTypes };
}
