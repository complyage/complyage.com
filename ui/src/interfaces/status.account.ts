//||------------------------------------------------------------------------------------------------||
//|| Types
//||------------------------------------------------------------------------------------------------||

export type AccountStatusKeys =
      | "None"
      | "Created"
      | "Pending"
      | "Verified"
      | "Active"
      | "Banned"
      | "Closed";

export type AccountStatusMap = Record<AccountStatusKeys, string>;

//||------------------------------------------------------------------------------------------------||
//|| Constants
//||------------------------------------------------------------------------------------------------||

export const AccountStatuses: AccountStatusMap = {
      None:      "NONE",
      Created:   "PNEW",
      Pending:   "PEND",
      Verified:  "VERF",
      Active:    "ACTV",
      Banned:    "BNND",
      Closed:    "RMVD",
};

//||------------------------------------------------------------------------------------------------||
//|| Utils
//||------------------------------------------------------------------------------------------------||

export function getAccountStatus(code : string): string[] {
      return Object.entries(AccountStatuses)
            .filter(([_, value]) => value === code)
            .map(([key]) => key);
}

export function getAllAccountStatuses(): string[] {
      return Object.values(AccountStatuses);
}

export function getAccountStatusMap(): Record<string, string> {
      return { ...AccountStatuses };
}
