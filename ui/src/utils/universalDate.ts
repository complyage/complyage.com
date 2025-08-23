//||------------------------------------------------------------------------------------------------||
//|| Universal Date Utilities
//|| src/utils/universalDate.ts
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Layout
      //||------------------------------------------------------------------------------------------------||

      const UNIVERSAL_DATE_LAYOUT = "yyyy-MM-dd HH:mm:ss";

      //||------------------------------------------------------------------------------------------------||
      //|| Format: Date to Universal String (UTC)
      //||------------------------------------------------------------------------------------------------||

      export function toUniversal(date: Date): string {
            const pad = (n: number) => n.toString().padStart(2, '0');
            return (
                  date.getUTCFullYear() + '-' +
                  pad(date.getUTCMonth() + 1) + '-' +
                  pad(date.getUTCDate()) + ' ' +
                  pad(date.getUTCHours()) + ':' +
                  pad(date.getUTCMinutes()) + ':' +
                  pad(date.getUTCSeconds())
            );
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Parse: Universal String to Date (UTC)
      //||------------------------------------------------------------------------------------------------||

      export function fromUniversal(str: string): Date | null {
            if (!str) return new Date('1970-01-01T00:00:00Z');
            const m = str.match(/^(\d{4})-(\d{2})-(\d{2}) (\d{2}):(\d{2}):(\d{2})$/);
            if (!m) return new Date('1970-01-01T00:00:00Z');
            const [ , year, month, day, hour, minute, second ] = m.map(Number);
            //|| Month is zero-based in JS
            return new Date(Date.UTC(year, month - 1, day, hour, minute, second));
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Universal Now: Returns current UTC date
      //||------------------------------------------------------------------------------------------------||

      export function universalNow(): Date {
            return new Date(Date.now());
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Try Format: Validates & formats a string or returns original
      //||------------------------------------------------------------------------------------------------||

      export function tryFormatUniversal(str: string): string {
            const date = fromUniversal(str);
            return date ? toUniversal(date) : str;
      }

      //||------------------------------------------------------------------------------------------------||
      //|| TimeAgo: Returns a human-readable "time ago" string from universal date
      //||------------------------------------------------------------------------------------------------||

      export function timeAgo(dateInput: string | Date): string {
            let date: Date | null;

            if (typeof dateInput === "string") {
                  // Try using fromUniversal, but fallback to new Date if needed
                  date = fromUniversal(dateInput);
                  if (!date || isNaN(date.getTime())) {
                        // Fallback: try native Date parsing (handles Go/ISO8601)
                        date = new Date(dateInput);
                  }
            } else {
                  date = dateInput;
            }

            if (!date || isNaN(date.getTime())) return "";

            const now     = universalNow ? universalNow() : new Date();
            const diffSec = Math.floor((now.getTime() - date.getTime()) / 1000);
            if (diffSec < 5)  return "just now";
            if (diffSec < 60) return `${diffSec} seconds ago`;
            if (diffSec < 90) return "a minute ago";
            if (diffSec < 3600) {
                  const min = Math.floor(diffSec / 60);
                  return min === 1 ? "1 minute ago" : `${min} minutes ago`;
            }
            if (diffSec < 5400) return "an hour ago";
            if (diffSec < 86400) {
                  const hours = Math.floor(diffSec / 3600);
                  return hours === 1 ? "1 hour ago" : `${hours} hours ago`;
            }
            const days = Math.floor(diffSec / 86400);
            if (days === 1) return "yesterday";
            if (days < 7)   return `${days} days ago`;
            if (days < 31)  return `${Math.floor(days/7)} week${Math.floor(days/7) > 1 ? "s" : ""} ago`;
            if (days < 365) return `${Math.floor(days/30)} month${Math.floor(days/30) > 1 ? "s" : ""} ago`;
            const years = Math.floor(days/365);
            return years === 1 ? "a year ago" : `${years} years ago`;
      }
