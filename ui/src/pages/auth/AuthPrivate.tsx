//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import React, { useState, useEffect }                       from "react";
import { useOverlayNavigate }                               from "../../hooks/useOverlay";

//||------------------------------------------------------------------------------------------------||
//|| Utils
//||------------------------------------------------------------------------------------------------||

import Call                                                  from "../../classes/call";

//||------------------------------------------------------------------------------------------------||
//|| API
//||------------------------------------------------------------------------------------------------||

import apiURL                                                 from "../../utils/apiURL";

//||------------------------------------------------------------------------------------------------||
//|| Components
//||------------------------------------------------------------------------------------------------||

import SpinnerCircle                                         from "../../components/base/SpinnerCircle";
import BIP39                                                 from "../../components/base/BIP39";
import InlineAlert                                           from "../../components/base/InlineAlert";
import PrivateCheckList                                      from "../../components/dynamic/PrivateCheckList";
import PrivateKey                                            from "../../components/base/PrivateKey";

//||------------------------------------------------------------------------------------------------||
//|| Interfaces
//||------------------------------------------------------------------------------------------------||

import { BIPList }                                           from "../../interfaces/auth/bip.list";

//||------------------------------------------------------------------------------------------------||
//|| Default Component
//||------------------------------------------------------------------------------------------------||

export default function AuthPrivate() {

      const navigate                                  = useOverlayNavigate();
      const [wordList, setWordList]                   = useState<BIPList>(["","","","","",""]);
      const [userPrivate, setUserPrivate]             = useState<string>("");
      const [level, setLevel]                         = useState<number>(0);
      const [loggedIn, setLoggedIn]                   = useState<boolean>(false);
      const [privateKey, setPrivateKey]               = useState<boolean>(false);
      const [statusMessage, setStatusMessage]         = useState("");
      const [loading, setLoading]                     = useState(true);
      const [minutes, setMinutes]                     = useState<number>(5);

      //||------------------------------------------------------------------------------------------------||
      //|| State
      //||------------------------------------------------------------------------------------------------||

      useEffect(() => {
            const fetchData = async () => {
                  setLoading(true);
                  setStatusMessage("");
                  const chirp = new Call(apiURL("/user/private/fetch"), {});
                  chirp.method = "GET";
                  chirp.debug = true;
                  await chirp.execute();
                  setLoading(false);
                  if (chirp.ok()) {
                        setLoggedIn(chirp.responsePayload.data.loggedIn as boolean);
                        setLevel(chirp.responsePayload.data.level as number);
                        setPrivateKey(chirp.responsePayload.data.private as boolean);
                  } else setStatusMessage(chirp.error());
            };
            fetchData();
      }, []);

      //||------------------------------------------------------------------------------------------------||
      //|| Handle Submit
      //||------------------------------------------------------------------------------------------------||

      const handleSubmit = async (e: React.FormEvent) => {
            e.preventDefault();            
      };

      //||------------------------------------------------------------------------------------------------||
      //|| Level
      //||------------------------------------------------------------------------------------------------||

      const levelToWordCount = (level: number) : 6 | 12 | 18 | 24 => {
            switch(level) {
                  case 2 : return 6;
                  case 3 : return 12;
                  case 4 : return 18;
                  case 5 : return 24;
            }
            return 6;
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Not Loaded
      //||------------------------------------------------------------------------------------------------||

      if (loading) {
            return (
                  <div className="relative z-10 flex items-center justify-center min-h-screen px-4 pt-24 pb-12">
                        <SpinnerCircle />
                  </div>
            );
      }

      //||------------------------------------------------------------------------------------------------||
      //|| Not Logged In
      //||------------------------------------------------------------------------------------------------||

      if (!loggedIn) return (
            <div className="relative z-10 flex items-center justify-center min-h-screen px-4 pt-24 pb-12">
                  <div className="bg-base-100 bg-opacity-90 rounded-lg shadow-lg w-full max-w-3xl p-8 text-center mx-auto">
                        <h1 className="text-3xl font-bold mb-6 text-gray-200">Not Logged In</h1>
                        <InlineAlert message="You must be logged in to access this page." isError={true} />
                        <button className="btn btn-secondary btn-lg p-5 py-5 mt-6" onClick={() => navigate("/auth/login")}>Login</button>
                  </div>
            </div>
      );


      //||------------------------------------------------------------------------------------------------||
      //|| JSX
      //||------------------------------------------------------------------------------------------------||

      return (
		<div className="relative z-10 flex items-center justify-center min-h-screen px-4">
			<div className="bg-base-100 bg-opacity-90 rounded-lg shadow-lg w-full max-w-3xl py-3 px-8 text-center mx-auto">
				<h1 className="text-3xl font-bold mt-5 text-gray-200 border-b border-gray-600 pb-5">Enter your Private Key</h1>
				<InlineAlert message={statusMessage} isError={true} />
				<PrivateCheckList loggedIn={loggedIn} level={level} privateKey={privateKey} />
				<form onSubmit={handleSubmit} className="flex flex-col gap-4 text-left">
					<div className="p-4 bg-gray-800 rounded-lg mt-4">
						<div className="flex flex-col gap-6">
							{level > 1 && level < 6 && <BIP39 mode="VERIFY" wordCount={levelToWordCount(level)} setValue={setWordList} />}
							{level == 6 && <PrivateKey onChange={(val) => setUserPrivate(val)} />}
						</div>
					</div>

					<div className="flex flex-col md:flex-row items-center gap-3 text-sm ">
						<label htmlFor="store-minutes" className="w-full md:w-1/2 text-gray-200 text-center">
							Store your private key for
						</label>

						<div className="relative w-full md:w-1/2">
							<select
								id="store-minutes"
								aria-label="Store private key duration"
								value={minutes}
								onChange={(e) => setMinutes(parseInt(e.target.value))}
								className="select select-auto w-64">
								<option value={5}>5 minutes</option>
								<option value={15}>15 minutes</option>
								<option value={60}>1 hour</option>
								<option value={240}>4 hours</option>
								<option value={1440}>24 hours</option>
								<option value={10080}>7 days</option>
								<option value={43200}>30 days</option>
								<option value={525600}>1 year</option>
								<option value={0}>Indefinitely</option>
							</select>

						</div>
					</div>

					<button className="btn btn-secondary btn-lg p-5 py-5 mt-3 mb-6" disabled={loading}>
						{loading ? <SpinnerCircle /> : "Store Private Key"}
					</button>
				</form>
			</div>
		</div>
	);
}
