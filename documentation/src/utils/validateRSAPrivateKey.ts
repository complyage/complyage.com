//||------------------------------------------------------------------------------------------------||
//|| utils/validateRSAPrivateKey.ts
//|| Validate RSA PKCS#8 private key PEM (the format produced by your Go GenerateKeyPair)
//|| - Checks PEM delimiters, decodes base64 body, then attempts subtle.importKey('pkcs8', ...)
//|| - Returns { ok: boolean, reason?: string }
//||------------------------------------------------------------------------------------------------||

export interface ValidateResult {
      ok: boolean;
      reason?: string;
}

const rePKCS8 = /-----BEGIN PRIVATE KEY-----([\s\S]+?)-----END PRIVATE KEY-----/;

function b64ToArrayBuffer(b64: string): Uint8Array {
      // remove whitespace/newlines
      const clean = b64.replace(/\s+/g, "");
      const binary = atob(clean);
      const len = binary.length;
      const bytes = new Uint8Array(len);
      for (let i = 0; i < len; i++) bytes[i] = binary.charCodeAt(i);
      return bytes;
}

/**
 * Attempt to import the PKCS#8 ArrayBuffer with a couple RSA algorithms.
 * If import succeeds, assume it's a valid RSA PKCS#8 private key.
 */
async function tryImportPKCS8(pkcs8: ArrayBuffer): Promise<boolean> {
      if (!("crypto" in window) || !("subtle" in (window as any).crypto)) {
            // WebCrypto not available (SSR or very old browser) — can't import, fallback to header-only check
            return false;
      }

      const subtle = (window as any).crypto.subtle as SubtleCrypto;

      // Try RSASSA-PKCS1-v1_5 (signing) and RSA-OAEP (encryption). If either imports it's likely a valid RSA PKCS#8 key.
      const algs: any[] = [
            { name: "RSASSA-PKCS1-v1_5", hash: { name: "SHA-256" } },
            { name: "RSA-OAEP", hash: { name: "SHA-256" } },
      ];

      for (const alg of algs) {
            try {
                  // subtle.importKey will throw on invalid format
                  // set extractable false and empty usages (we only check importability)
                  // usages vary per algorithm; 'sign' for RSASSA, 'decrypt' for OAEP (but import will still validate)
                  const usages = alg.name === "RSASSA-PKCS1-v1_5" ? ["sign"] : ["decrypt"];
                  // @ts-ignore - subtle.importKey signature
                  const key = await subtle.importKey("pkcs8", pkcs8, alg, false, usages);
                  if (key) return true;
            } catch (err) {
                  // ignore and try next algorithm
            }
      }

      return false;
}

/**
 * Public function: validate RSA PKCS#8 PEM private key.
 * Returns { ok, reason }.
 */
export default async function validateRSAPrivateKeyPEM(pem: string): Promise<ValidateResult> {
      if (!pem || typeof pem !== "string" || pem.trim().length === 0) {
            return { ok: false, reason: "empty" };
      }

      const m = pem.match(rePKCS8);
      if (!m) {
            return { ok: false, reason: "missing-pkcs8-delimiters" };
      }

      const b64 = m[1];
      let bytes: Uint8Array;
      try {
            bytes = b64ToArrayBuffer(b64);
      } catch (err) {
            return { ok: false, reason: "invalid-base64" };
      }

      // If subtle.importKey is available, try to import
      if ("crypto" in window && "subtle" in (window as any).crypto) {
            try {
                  const ok = await tryImportPKCS8(bytes.buffer);
                  if (ok) return { ok: true };
                  return { ok: false, reason: "subtle-import-failed" };
            } catch (err) {
                  return { ok: false, reason: "import-error" };
            }
      }

      // Fallback: we have delimiters + decodable body -> accept as "structurally valid" but not fully verified.
      return { ok: true, reason: "no-subtle-fallback-accepted" };
}
