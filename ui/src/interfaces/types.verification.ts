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

export type VerificationTypeValues = typeof VerificationTypes[VerificationTypeKeys];


//||------------------------------------------------------------------------------------------------||
//|| Utils
//||------------------------------------------------------------------------------------------------||

export function getVerificationKeyFromCode(code: string): VerificationTypeKeys | undefined {
      const entry = Object.entries(VerificationTypes).find(([_, value]) => value === code);
      return entry ? (entry[0] as VerificationTypeKeys) : undefined;
}

export function getAllVerificationTypes(): string[] {
      return Object.values(VerificationTypes);
}

export function getVerificationTypeMap(): Record<string, string> {
      return { ...VerificationTypes };
}
