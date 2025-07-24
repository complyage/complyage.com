//||------------------------------------------------------------------------------------------------||
//|| Types
//||------------------------------------------------------------------------------------------------||

export type VerificationStatusKeys =
      | "None"
      | "Pending"
      | "Verified"
      | "Failed";

export type VerificationStatusMap = Record<VerificationStatusKeys, string>;

//||------------------------------------------------------------------------------------------------||
//|| Constants
//||------------------------------------------------------------------------------------------------||

export const VerificationStatuses: VerificationStatusMap = {
      None:     "NONE",
      Pending:  "PEND",
      Verified: "VERF",
      Failed:   "FAIL",
};

//||------------------------------------------------------------------------------------------------||
//|| Utils
//||------------------------------------------------------------------------------------------------||

export function getAllVerificationStatuses(): string[] {
      return Object.values(VerificationStatuses);
}

export function getVerificationStatusMap(): Record<string, string> {
      return { ...VerificationStatuses };
}
