//||------------------------------------------------------------------------------------------------||
//|| ADvanded Choose
//||------------------------------------------------------------------------------------------------||

      //||------------------------------------------------------------------------------------------------||
      //|| Import
      //||------------------------------------------------------------------------------------------------||

      import React, {useState, useEffect}                        from "react";
      import {useNavigate}                                       from "react-router-dom";
      import { useLocation }                                     from "react-router-dom";

      //||------------------------------------------------------------------------------------------------||
      //|| Utils
      //||------------------------------------------------------------------------------------------------||

      import {makeBIPList}                                       from "../../utils/bip39";

      //||------------------------------------------------------------------------------------------------||
      //|| Components
      //||------------------------------------------------------------------------------------------------||

      import NavMain                                             from "../../components/nav/NavMain";
      import FooterMain                                          from "../../components/footer/FooterMain";
      import SpinnerCircle                                       from "../../components/base/SpinnerCircle";
      import BIP39                                               from "../../components/base/BIP39";
      import KeyPairGenerator                                    from "../../components/base/KeyPairGenerator";
      import InlineAlert                                         from "../../components/base/InlineAlert";

      //||------------------------------------------------------------------------------------------------||
      //|| Interfaces
      //||------------------------------------------------------------------------------------------------||

      import { BIPList }                                         from "../../interfaces/auth/bip.list";
      import { SecurityLevel }                                   from "../../interfaces/auth/security.levels";

      //||------------------------------------------------------------------------------------------------||
      //|| Default
      //||------------------------------------------------------------------------------------------------||

      export default function SignupComplete() {
            //||------------------------------------------------------------------------------------------------||
            //|| Var
            //||------------------------------------------------------------------------------------------------||
            const navigate                                  = useNavigate();
            const [checkStatus, setCheckStatus]             = useState("CHECK");
            const [wordList, setWordList]                   = useState<BIPList>(makeBIPList());
            const [password, setPassword]                   = useState("");
            const [encryptionLevel, setEncryptionLevel]     = useState<number>(1);
            const [statusMessage, setStatusMessage]         = useState("");
            const [privateKey, setPrivateKey]               = useState<string>("");
            const [publicKey, setPublicKey]                 = useState<string>("");
            const [loading, setLoading]                     = useState(true);
            const levels: SecurityLevel[] = [
                  {
                        id: 1,
                        name: "Standard [Recommended]",
                        description: "We securely manage your keys with enterprise-grade encryption, so you never have to worry about losing access. This is recommended for most users."
                  },
                  {
                        id: 2,
                        name: "Advanced",
                        description: "Protect your account with a 6-word BIP39 passphrase – easy to remember, strong enough for advanced users. Ideal for those who want a balance of security and convenience."
                  },
                  {
                        id: 3,
                        name: "Expert",
                        description: "Take full control by managing your own private security key. Maximum responsibility, maximum freedom. This is for advanced users who want complete control over their security."
                  }
            ];

            //||------------------------------------------------------------------------------------------------||
            //|| Extract oauth query param from current URL
            //||------------------------------------------------------------------------------------------------||

            const location      = useLocation();
            const params        = new URLSearchParams(location.search);
            const oauthParam    = params.get("oauth");
            
            //||------------------------------------------------------------------------------------------------||
            //|| Check Session
            //||------------------------------------------------------------------------------------------------||
            useEffect(() => {
                  const checkMe = async () => {
                        try {
                              //||------------------------------------------------------------------------------------------------||
                              //|| Fetch
                              //||------------------------------------------------------------------------------------------------||
                              console.log("AUTH/me")
                              const res = await fetch("/auth/me", {
                                    method: "GET",
                                    credentials: "include",
                              });
                              //||------------------------------------------------------------------------------------------------||
                              //|| Check JSON
                              //||------------------------------------------------------------------------------------------------||
                              const json = await res.json();
                              console.log(json);
                              if (!json.success) {
                                    setCheckStatus("FAIL");
                              }
                              //||------------------------------------------------------------------------------------------------||
                              //|| Check JSON
                              //||------------------------------------------------------------------------------------------------||
                              if (json.data.status == "ACTV") {
                                    if (oauthParam) {
                                          window.location.href = `${import.meta.env.VITE_COMPLYAGE_OAUTH_URL}/v1/return?oauth=${oauthParam}`;
                                    } else {
                                          navigate("/members");
                                    }
                              }
                              //||------------------------------------------------------------------------------------------------||
                              //|| Verified
                              //||------------------------------------------------------------------------------------------------||
                              if (json.data.status == "VERF") {
                                    setCheckStatus("VERF");
                              }
                              //||------------------------------------------------------------------------------------------------||
                              //|| Other
                              //||------------------------------------------------------------------------------------------------||
                              if (json.data.status == "ACTV") {
                                    setCheckStatus(json.data.status);
                              }
                        } catch (err) {
                              setCheckStatus("FAIL");
                        }
                  };
                  checkMe();
            }, []);

            //||------------------------------------------------------------------------------------------------||
            //|| Handle Submit
            //||------------------------------------------------------------------------------------------------||

            const handleSubmit = async (e: React.FormEvent) => {
                  //||------------------------------------------------------------------------------------------------||
                  //|| Generate the Payload
                  //||------------------------------------------------------------------------------------------------||
                  e.preventDefault();
                  const payload = new URLSearchParams({
                        password          : password,
                        wordList          : JSON.stringify(wordList),
                        encryptionLevel   : encryptionLevel.toString(),
                        privateKey        : privateKey || "",
                        publicKey         : publicKey || ""
                  }).toString();
                  //||------------------------------------------------------------------------------------------------||
                  //|| Handle the POST
                  //||------------------------------------------------------------------------------------------------||
                  try {
                        setLoading(true);
                        console.log("AUTH/Complete")
                        const res = await fetch("/auth/complete", {
                              method: "POST",
                              headers: {"Content-Type": "application/x-www-form-urlencoded"},
                              body: payload
                        });
                        const json = await res.json();
                        console.log("RESPONSE", json);
                        //||------------------------------------------------------------------------------------------------||
                        //|| Handle the Response
                        //||------------------------------------------------------------------------------------------------||
                        const redirectURL = (oauthParam) ? `/members/?oauth=${oauthParam}` : `/members/`;
                        if (json.success) return navigate(redirectURL);
                        //||------------------------------------------------------------------------------------------------||
                        //|| Failed
                        //||------------------------------------------------------------------------------------------------||
                        setStatusMessage(`${json.error || "Save failed."}`);
                  } catch (err) {
                        console.error(err);
                        setStatusMessage("Something went wrong. Please try again.");
                  } finally {
                        setLoading(false);
                  }
            };

            //||------------------------------------------------------------------------------------------------||
            //|| Check Status Failed
            //||------------------------------------------------------------------------------------------------||
            if (checkStatus === "CHECK") {
                  return (
                        <main className="flex flex-col h-[70vw]">
                              <NavMain />
                              <div className="flex-1 flex items-center justify-center bg-gray-500">
                                    <div className="flex flex-col items-center justify-center">
                                          <SpinnerCircle />
                                          <div className="mt-4 text-center">Checking account status...</div>
                                    </div>
                              </div>
                              <FooterMain />
                        </main>
                  );
            }
            //||------------------------------------------------------------------------------------------------||
            //|| Check Status Failed
            //||------------------------------------------------------------------------------------------------||
            if (checkStatus === "FAIL") {
                  return (
                        <main className="flex flex-col min-h-screen">
                              <NavMain />
                              <div className="flex-1 flex flex-col items-center justify-center bg-white text-black px-4">
                                    <h1 className="text-2xl font-bold mb-4 text-center">Could not find your session</h1>
                                    <p className="mb-8 text-center text-gray-600">Please log in or sign up to continue.</p>
                                    <div className="flex space-x-4">
                                          <a href="/login" className="px-6 py-3 bg-black text-white rounded-md hover:bg-gray-800 transition">Log In</a>
                                          <a href="/signup" className="px-6 py-3 bg-white text-black border border-black rounded-md hover:bg-gray-100 transition">Sign Up</a>
                                    </div>
                              </div>
                              <FooterMain />
                        </main>
                  );
            }
            //||------------------------------------------------------------------------------------------------||
            //|| JSX / FORM
            //||------------------------------------------------------------------------------------------------||
            
            return (
                  <main className="min-h-screen flex flex-col">
                        <NavMain />

                        <div className="relative flex-1">
                              <img src="https://picsum.photos/1920/1080" alt="Background" className="absolute inset-0 w-full h-full object-cover" />
                              <div className="absolute inset-0 bg-black/70"></div>

                              <div className="relative z-10 flex items-center justify-center min-h-screen px-4 pt-24 pb-12">
                                    <div className="bg-base-100 bg-opacity-90 rounded-lg shadow-lg w-full max-w-[80vw] p-8 text-center mx-auto">

                                          <h1 className="text-3xl font-bold mb-6 text-gray-200">Complete Your Signup</h1>

                                          <p className="mb-6 text-base-content/70">Set your password and encryption level.</p>
                                          <InlineAlert message={statusMessage} isError={true} />

                                          <form onSubmit={handleSubmit} className="flex flex-col gap-4 text-left">
                                                <label className="font-semibold">Password</label>
                                                <input
                                                      type="password"
                                                      placeholder="Enter your password"
                                                      value={password}
                                                      onChange={(e) => setPassword(e.target.value)}
                                                      className="input input-bordered w-full"
                                                      required
                                                />

                                                <label className="font-semibold mt-4">Encryption Level</label>
                                                <div className="flex flex-col gap-6">
                                                      {/* Three Columns */}
                                                      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
                                                            {levels.map(level => (
                                                                  <div
                                                                        key={level.id}
                                                                        onClick={() => setEncryptionLevel(level.id)}
                                                                        className={`cursor-pointer p-6 rounded-lg text-center border-2 transition ${
                                                                              encryptionLevel === level.id
                                                                                    ? "border-primary bg-gray-400/50 shadow-lg"
                                                                                    : "border-gray-300 hover:border-primary"
                                                                        }`}
                                                                  >
                                                                        <h2 className="text-xl font-bold mb-2">{level.name}</h2>
                                                                        <p className="text-sm opacity-70">
                                                                              {level.id === 1 && "Best for simplicity and convenience."}
                                                                              {level.id === 2 && "Balanced security with easy memorization."}
                                                                              {level.id === 3 && "Maximum control for advanced users."}
                                                                        </p>
                                                                  </div>
                                                            ))}
                                                      </div>

                                                      {/* Dynamic Description Box */}
                                                      <div className="p-4 rounded-lg bg-black/40 text-center">
                                                            <h3 className="font-bold text-yellow-500 text-lg">{levels[encryptionLevel - 1].name}</h3>                                                            
                                                            <p className="text-base">{levels[encryptionLevel - 1].description}</p>
                                                      </div>

                                                      {encryptionLevel === 2 && (
                                                            <BIP39 initialWords={wordList} setValue={(words) => {
                                                                  setWordList(words);     
                                                            }} />
                                                      ) }

                                                      {encryptionLevel === 3 && (
                                                            <KeyPairGenerator defaultPrivate={privateKey} defaultPublic={publicKey} setValue={({ privateKey, publicKey }) => {
                                                                  setPrivateKey(privateKey);
                                                                  setPublicKey(publicKey);
                                                                  console.log("Generated Key Pair:", { privateKey, publicKey });
                                                            }} />
                                                      )}
                                                </div>

                                                <button className="btn btn-secondary text-xl p-5 py-5 mt-6" onClick={ handleSubmit }>
                                                      {"Complete Signup"}
                                                </button>
                                          </form>
                                    </div>
                              </div>
                        </div>

                        <FooterMain />
                  </main>
            );
      }
