export default function App() {
      return (
            <div class="font-sans h-screen w-screen text-[1.1em] flex flex-row">
                  
                  {/* Left Column (Background Image / Overlays) */}
                  <div
                        class="flex-1 flex justify-center items-center p-8 relative bg-black bg-cover bg-center"
                        style="background-image:url(/static/img/oauth.jpg)"
                  >
                        <aside id="initOverlay" class="w-full flex justify-center items-center">
                              <div class="flex flex-col items-center justify-center text-center w-4/5">
                                    <img
                                          class="mb-2"
                                          src="/static/img/complyage-w.webp"
                                          alt="ComplyAge Logo"
                                    />
                                    <h2 class="text-gray-400 text-2xl px-2 py-5 font-bold">
                                          Take Back Your Privacy. Keep Your Freedom.
                                    </h2>
                                    <p class="text-md leading-7 px-2 pt-1 pb-6 m-0 text-gray-200">
                                          You control your encrypted personal data, everything you submit
                                          is protected by end-to-end encryption. That means not even we can
                                          read, intercept, or sell your information. With ComplyAge, your
                                          private details stay exactly that.
                                    </p>

                                    <div class="showLoggedOut flex flex-col items-center justify-center text-center p-4 rounded-lg w-4/5">
                                          <div class="flex gap-4">
                                                <a
                                                      href="javascript:void(0)"
                                                      onClick={() =>
                                                            openOverlay('http://localhost:5173/overlay/signup')
                                                      }
                                                      class="inline-block text-2xl no-underline cursor-pointer bg-orange-500 text-white py-3 px-8 rounded-md font-bold transition hover:bg-orange-600 shadow"
                                                >
                                                      Sign Up
                                                </a>
                                                <a
                                                      href="javascript:void(0)"
                                                      onClick={() =>
                                                            openOverlay('http://localhost:5173/overlay/login')
                                                      }
                                                      class="inline-block text-2xl no-underline cursor-pointer bg-gray-600 text-white py-3 px-8 rounded-md font-bold transition hover:bg-gray-700 shadow"
                                                >
                                                      Log In
                                                </a>
                                          </div>
                                          <span class="my-4 text-2xl font-bold text-gray-500">
                                                Private and Secure.
                                          </span>
                                    </div>

                                    <div class="flex-col items-center justify-center text-center p-4 rounded-lg text-sm font-bold w-4/5 bg-black/50 text-white">
                                          <div>
                                                <span class="block text-md text-yellow-500">
                                                      Age Verification Enforced
                                                </span>
                                                <div class="bg-white/20 rounded-lg py-2 my-2">
                                                      <span>Sponsored VPN:</span>
                                                      <a
                                                            class="text-blue-400"
                                                            href="https://www.shadownetvpn.com"
                                                            alt="ShadowNet VPN"
                                                      >
                                                            ShadowNet VPN
                                                      </a>
                                                      <p class="font-normal text-xs text-gray-400">
                                                            Military-grade encryption with zero logs policy.
                                                      </p>
                                                </div>
                                                <div>
                                                      <span class="text-xs">68.110.15.226</span>
                                                      <span class="text-xs text-gray-400">Location: Arizona, US</span>
                                                </div>
                                          </div>
                                    </div>
                              </div>
                        </aside>
                  </div>

                  {/* Right Column (Main Panel) */}
                  <section class="flex-1 flex flex-col h-full w-full bg-[oklch(0.2326_0.014_253.1)] text-white">
                        <div class="flex flex-col h-full text-center min-h-0 pt-3">
                              <div class="flex-shrink-0 border-b border-black p-3 bg-black text-center">
                                    <div
                                          class="w-[60%] aspect-[5/1] bg-center bg-cover bg-no-repeat mb-3 border border-white/20 shadow-lg rounded-2xl mx-auto"
                                          style="background-image:url('http://localhost:12000/sites//static/img/complyage-w.webp');"
                                    ></div>
                                    <h1 class="text-gray-300 text-2xl font-bold m-0 p-1">My Website</h1>
                                    <b class="mb-2 text-yellow-400 block text-md">
                                          testurl.com <span class="text-gray-200">wants to access the following information</span>
                                    </b>
                              </div>

                              <div class="flex-grow overflow-y-auto bg-black text-left">
                                    <ul class="list-none m-0 p-4 space-y-3">
                                          <li class="flex items-center justify-between bg-white/5 shadow-sm rounded-lg border border-black p-4">
                                                <div class="flex items-center gap-4">
                                                      <div class="w-12 h-12 flex items-center justify-center rounded-full bg-black">
                                                            <img src="/static/img/icons/icon_IDEN.png" alt="IDEN" class="w-8 h-8" />
                                                      </div>
                                                      <div>
                                                            <h3 class="text-xl font-semibold text-gray-200">ID/Passport</h3>
                                                            <p class="text-lg text-red-400">ID/Passport is Not Verified</p>
                                                      </div>
                                                </div>
                                                <span class="bg-black/20 rounded-md text-white px-3 py-1 text-sm">Please log in</span>
                                          </li>

                                          <li class="flex items-center justify-between bg-white/5 shadow-sm rounded-lg border border-black p-4">
                                                <div class="flex items-center gap-4">
                                                      <div class="w-12 h-12 flex items-center justify-center rounded-full bg-black">
                                                            <img src="/static/img/icons/icon_CRCD.png" alt="CRCD" class="w-8 h-8" />
                                                      </div>
                                                      <div>
                                                            <h3 class="text-xl font-semibold text-gray-200">Credit Card</h3>
                                                            <p class="text-lg text-red-400">Credit Card is Not Verified</p>
                                                      </div>
                                                </div>
                                                <span class="bg-black/20 rounded-md text-white px-3 py-1 text-sm">Please log in</span>
                                          </li>
                                    </ul>
                              </div>

                              <div class="flex-shrink-0 bg-black/50 text-yellow-500 p-4 text-center">
                                    You must be logged in to access this feature.
                              </div>
                        </div>
                  </section>
            </div>
      )
}

// JS overlay helper functions
function bodyClass(newClass: string) {
      document.body.classList.remove('overlay-init', 'overlay-auth', 'overlay-preview')
      document.body.classList.add(newClass)
}

function openOverlay(url: string) {
      bodyClass('overlay-auth')
      const frame = document.getElementById('authFrame') as HTMLIFrameElement
      if (frame) frame.src = url
}

function closeOverlay() {
      bodyClass('overlay-init')
      const frame = document.getElementById('authFrame') as HTMLIFrameElement
      if (frame) frame.src = ''
}
