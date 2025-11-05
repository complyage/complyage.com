//||------------------------------------------------------------------------------------------------||
//|| getEnv Utility
//|| src/data/getEnv.ts
//||------------------------------------------------------------------------------------------------||
           

      export function getEnv(key: string, defaultValue?: string): string {
            const val = import.meta.env?.[key] ?? (typeof process !== "undefined" ? process.env?.[key] : undefined);
            if (typeof val === "string" && val.length > 0) return val;
            if (typeof defaultValue !== "undefined") return defaultValue;
            throw new Error(`Missing environment variable: ${key}`);
      }