//||------------------------------------------------------------------------------------------------||
//|| components/base/PrivateKey.tsx
//|| Private key input with live PKCS#8 RSA validation (uses validateRSAPrivateKeyPEM)
//||------------------------------------------------------------------------------------------------||

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||
import React, { useEffect, useRef, useState }        from "react";
import { CircleCheck, X }                             from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Utils
//||------------------------------------------------------------------------------------------------||
import validateRSAPrivateKeyPEM, { ValidateResult }   from "../../utils/validateRSAPrivateKey";

//||------------------------------------------------------------------------------------------------||
//|| Props
//||------------------------------------------------------------------------------------------------||
interface PrivateKeyProps {
      value?          : string;
      onChange?       : (val: string) => void;
      showType?       : boolean;
      className?      : string;
}

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||
export default function PrivateKey({ value = "", onChange, showType = true, className = "" }: PrivateKeyProps) {

      //||------------------------------------------------------------------------------------------------||
      //|| State + refs
      //||------------------------------------------------------------------------------------------------||
      const [text, setText]         = useState<string>(value || "");
      const [ok, setOk]             = useState<boolean>(false);
      const [reason, setReason]     = useState<string>("");
      const timerRef                = useRef<number | null>(null);
      const validatingRef           = useRef<boolean>(false);

      // keep controlled value in sync if parent updates
      useEffect(() => {
            if (value !== text) setText(value || "");
            // eslint-disable-next-line react-hooks/exhaustive-deps
      }, [value]);

      // debounce validate on text change
      useEffect(() => {
            // clear prior timer
            if (timerRef.current) window.clearTimeout(timerRef.current);

            // small debounce so typing doesn't hammer subtle.importKey
            timerRef.current = window.setTimeout(() => {
                  validatingRef.current = true;
                  (async () => {
                        try {
                              const res: ValidateResult = await validateRSAPrivateKeyPEM(text);
                              setOk(!!res.ok);
                              setReason(res.reason ?? "");
                        } catch (err) {
                              setOk(false);
                              setReason("exception");
                        } finally {
                              validatingRef.current = false;
                        }
                  })();
            }, 250);

            return () => {
                  if (timerRef.current) {
                        window.clearTimeout(timerRef.current);
                        timerRef.current = null;
                  }
            };
      }, [text]);

      // proxy change to parent if provided
      const handleChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
            const v = e.target.value;
            setText(v);
            if (onChange) onChange(v);
      };

      // small helper to render badge text
      const badgeText = () => {
            if (!text || text.trim() === "") return "empty";
            if (validatingRef.current) return "checking";
            if (ok) return "rsa-pkcs8";
            return reason || "invalid";
      };

      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||
      return (
            <div className={`relative w-full ${className}`}>
                  {/* status badge (top-right) */}
                  <div className="absolute right-2 top-2 z-20 flex items-center gap-2">
                        {ok ? (
                              <div className="flex items-center gap-2 bg-black/50 border border-green-700 text-green-200 rounded-full px-2 py-1">
                                    <CircleCheck className="w-4 h-4 text-green-300" />
                                    {showType && <span className="text-xs font-mono text-green-200 leading-none">{badgeText()}</span>}
                              </div>
                        ) : (
                              <div className="flex items-center gap-2 bg-black/50 border border-red-700 text-red-200 rounded-full px-2 py-1">
                                    <X className="w-4 h-4 text-red-300" />
                                    {showType && <span className="text-xs font-mono text-red-200 leading-none">{badgeText()}</span>}
                              </div>
                        )}
                  </div>

                  <textarea
                        aria-label="Private key"
                        spellCheck={false}
                        wrap="soft"
                        value={text}
                        onChange={handleChange}
                        className="h-64 p-3 text-sm font-mono bg-gray-900 text-yellow-300 border border-gray-700 rounded-md w-full resize-y overflow-auto placeholder-gray-500 caret-yellow-400 break-words whitespace-pre-wrap"
                        placeholder="Paste your private key"
                  />
            </div>
      );
}
