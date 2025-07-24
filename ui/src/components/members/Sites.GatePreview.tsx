//||------------------------------------------------------------------------------------------------||
//|| SiteGatePreview
//|| Component that renders a full‑screen mockup overlay based on site and gate_custom settings
//||------------------------------------------------------------------------------------------------||

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useState, useEffect }           from "react";
import { X }                                    from "lucide-react";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||
import type {Site} from "../../interfaces/model.sites";

interface SiteGatePreviewProps {
	data: Site;
	onClose: () => void;
}

//||------------------------------------------------------------------------------------------------||
//|| Component
//||------------------------------------------------------------------------------------------------||

export default function SiteGatePreview({data, onClose}: SiteGatePreviewProps) {	

      //||------------------------------------------------------------------------------------------------||
      //|| Upload File
      //||------------------------------------------------------------------------------------------------||
      
      const [logoHash, setLogoHash]       = useState<string>(data?.logoHash || "");
      const [logoMissing, setLogoMissing] = useState<boolean>(data?.logoMissing || true);
      const [logoExists, setLogoExists]   = useState(false);

      //||------------------------------------------------------------------------------------------------||
      //|| Sync With Parent Data
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            if (!data) return;
            setLogoHash(data.logoHash || "");
            setLogoMissing(data?.logoMissing);
      }, [data]);

      //||------------------------------------------------------------------------------------------------||
      //|| Upload File
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            if (logoHash != "") {
               const url = `/media/sites/logos/${logoHash}.webp`;
               fetch(url, { method: "HEAD", cache: "no-store" })
                  .then((res) => setLogoExists(res.ok))
                  .catch(() => setLogoExists(false));
            }
      }, [data]);

      //||------------------------------------------------------------------------------------------------||
      //|| Const
      //||------------------------------------------------------------------------------------------------||

      console.log("GATE", data.gate);
      const showSignup = data.gate?.signup === true && data.gate?.customConfirm !== null && data.gate?.customExit !== null;
      const signupCSS = showSignup ? "fixed left-1/5 right-1/5 top-1/4 bottom-1/4 bg-gray-200 rounded-lg text-center" : "fixed left-1/3 right-1/3 top-1/4 bottom-1/4 bg-gray-200 rounded-lg text-center border-2 border-white";


      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      return (
		<div className="fixed inset-0 bg-black bg-opacity-80 z-[9997] font-sans" onClick={onClose}>
                  <a className="fixed top-10 right-10 bg-white/40 text-black rounded-lg p-2 hover:bg-white cursor-pointer"><X size="36"/></a>
			<div className={signupCSS} onClick={(e) => e.stopPropagation()}>
				<div className="flex flex-row justify-center items-center h-full w-full z-[9999]">
					{/* Left Panel */}
                              { showSignup && (
                                    <div className="flex-1 p-8 w-1/2 flex flex-col items-center justify-center bg-white h-full text-center shadow-[8px_0_8px_-4px_rgba(0,0,0,0.3)] text-black">
                                          <img
                                                src={ logoMissing ? "/static/media/alert.svg" : `${import.meta.env.VITE_COMPLYAGE_MINIO_URL}/${import.meta.env.VITE_MINIO_BUCKET}/sites/logos/${logoHash}.webp?cachebust=${ new Date() }` }
                                                alt="Logo preview"
                                                className="max-w-xs h-24 w-24 border rounded-lg shadow"
                                          />

                                          <img
                                                className="agegate-logo mb-4 h-16 w-auto"
                                                src={`${import.meta?.env.VITE_COMPLYAGE_CLIENT_URL}/static/media/complyage.webp`}
                                                alt="Logo"
                                                onError={(e) => {
                                                      (e.currentTarget as HTMLImageElement).src = "/static/media/alert.svg";
                                                }}
                                          />
                                          <h2 className="text-gray-600 text-base leading-[1.4] pl-2 mb-4">
                                                Take Back Your Privacy.
                                                <br />
                                                Keep Your Freedom.
                                          </h2>
                                          <p className="text-xs text-black py-4">
                                                You control your encrypted personal data, everything you submit is protected by end-to-end encryption. That means not
                                                even we can read, intercept, or sell your information. With ComplyAge, your private details stay exactly that.
                                                <b> Private and Secure.</b>
                                          </p>
                                          <a
                                                href={data.gate?.customConfirm!}
                                                target="_parent"
                                                className="inline-block bg-black text-white px-6 py-4 rounded text-xl font-bold mb-2 hover:bg-gray-800">
                                                Signup. Securely Today
                                          </a>
                                          <span className="text-xs font-bold p-2">Take Back Your Privacy.</span>
                                    </div>
                              )}
					{/* Right Panel */}
                              <div className="flex-1 flex flex-col justify-start items-center h-full z-[9998] text-black">
						<div className="flex-col flex rounded-t-lg bg-black bg-opacity-75 w-full text-white py-2">
							<div className="flex items-center justify-center gap-2">
								<span className="font-bold text-2xl leading-none">Location, NA</span>
								<span className="text-xs bg-yellow-300 text-black rounded px-2 py-0.5">2.19.144.0</span>
							</div>
							<span className="text-lg text-gray-300">
								Effective Date:&nbsp;<span className="text-yellow-300">January 2000</span>
							</span>
						</div>
						<h1 className="mt-0 pt-5 text-2xl">{data.name}</h1>
						<b className="text-blue-500 mb-2">{data.url}</b>
						<img
							src={data.logoHash || "http://localhost:8089/static/media/alert.svg"}
							alt="Site Logo"
							className="w-auto mb-4"
						/>
						<p className="px-5 pt-5 pb-5 text-sm">Please verify your age to continue accessing this content.</p>
						<div className="pt-1">
							<a
								href={data.gate?.customConfirm || "https://complyage.com/oauth?apiKey=" + data.public}
								target="_blank"
								className="text-2xl no-underline mr-2 bg-blue-600 px-4 py-2 text-white rounded hover:bg-blue-700">
								I can confirm my age.
							</a>
							<br />
							<a
								href={data.gate?.customExit || "https://complyage.com/exit"}
								target="_parent"
								className="text-base no-underline block text-gray-800 font-bold py-4">
								I would like to exit
							</a>
						</div>
					</div>
				</div>
			</div>
		</div>
	);
          
}
