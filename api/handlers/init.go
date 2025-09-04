package handlers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"api/handlers/public"
	"api/handlers/user"
	"api/handlers/utils"
	"api/handlers/v1/sites"
	"api/handlers/v1/verifier"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/auth/sentry"
)

//||------------------------------------------------------------------------------------------------||
//|| Auth
//||------------------------------------------------------------------------------------------------||

func InitRoutes() {
	//||------------------------------------------------------------------------------------------------||
	//|| Auth
	//||------------------------------------------------------------------------------------------------||
	app.HTTP["api"].Router.HandleFunc("/auth/signup", sentry.SignupHandler).Methods("POST")
	app.HTTP["api"].Router.HandleFunc("/auth/forgot", sentry.ForgotPasswordHandler).Methods("POST")
	app.HTTP["api"].Router.HandleFunc("/auth/twofactor", sentry.TwoFactorHandler).Methods("POST")
	app.HTTP["api"].Router.HandleFunc("/auth/me", sentry.AuthMeHandler).Methods("GET")
	app.HTTP["api"].Router.HandleFunc("/auth/complete", sentry.CompleteHandler).Methods("POST", "GET")
	app.HTTP["api"].Router.HandleFunc("/auth/login", sentry.LoginHandler).Methods("POST")
	app.HTTP["api"].Router.HandleFunc("/auth/logout", sentry.LogoutHandler).Methods("GET")
	app.HTTP["api"].Router.HandleFunc("/auth/generate", sentry.GenerateKeyPairHandler).Methods("GET")
	app.HTTP["api"].Router.HandleFunc("/auth/reset", sentry.ResetPasswordHandler).Methods("POST")
	app.HTTP["api"].Router.HandleFunc("/auth/quit", sentry.QuitHandler).Methods("POST")
	app.HTTP["api"].Router.HandleFunc("/auth/delete-account", sentry.DeleteAccountHandler).Methods("POST")
	//||------------------------------------------------------------------------------------------------||
	//|| User Routes
	//||------------------------------------------------------------------------------------------------||
	app.HTTP["api"].Router.HandleFunc("/user/dashboard", user.UserDashboard).Methods("GET")
	app.HTTP["api"].Router.HandleFunc("/user/verifications", user.UserVerifications).Methods("GET")
	app.HTTP["api"].Router.HandleFunc("/user/shared", user.UserSharedHandler).Methods("GET")
	app.HTTP["api"].Router.HandleFunc("/user/location", user.LocationUserDashboard).Methods("GET")
	app.HTTP["api"].Router.HandleFunc("/user/vpns/rate", user.VPNRatingUserHandler).Methods("GET")
	//||------------------------------------------------------------------------------------------------||
	//|| Get News/Zones Public
	//||------------------------------------------------------------------------------------------------||
	app.HTTP["api"].Router.HandleFunc("/v1/api/news", public.NewsHandler).Methods("GET")
	app.HTTP["api"].Router.HandleFunc("/v1/api/zones", public.ZoneHandler).Methods("GET")
	app.HTTP["api"].Router.HandleFunc("/v1/api/vpns", public.VPNHandler).Methods("GET")
	app.HTTP["api"].Router.HandleFunc("/v1/api/donate", public.DonateHandler).Methods("GET")
	//||------------------------------------------------------------------------------------------------||
	//|| Member Sites
	//||------------------------------------------------------------------------------------------------||
	app.HTTP["api"].Router.HandleFunc("/v1/api/sites/zones", sites.SitesZoneHandler).Methods("GET")
	app.HTTP["api"].Router.HandleFunc("/v1/api/sites/vtypes", sites.VerificationTypesListHandler).Methods("GET")
	app.HTTP["api"].Router.HandleFunc("/v1/api/sites/list", sites.SitesListHandler).Methods("GET")
	app.HTTP["api"].Router.HandleFunc("/v1/api/sites/load", sites.SitesLoadHandler).Methods("GET")
	app.HTTP["api"].Router.HandleFunc("/v1/api/sites/upload", sites.UploadHandler).Methods("POST")
	app.HTTP["api"].Router.HandleFunc("/v1/api/sites/create", sites.SitesNewHandler).Methods("POST")
	app.HTTP["api"].Router.HandleFunc("/v1/api/sites/copy", sites.SitesCopyHandler).Methods("GET")
	app.HTTP["api"].Router.HandleFunc("/v1/api/sites/update", sites.SitesUpdateHandler).Methods("POST")
	app.HTTP["api"].Router.HandleFunc("/v1/api/sites/delete", sites.SitesDeleteHandler).Methods("DELETE")
	//||------------------------------------------------------------------------------------------------||
	//|| Verify
	//||------------------------------------------------------------------------------------------------||
	app.HTTP["api"].Router.HandleFunc("/v1/api/verify/list", verifier.GetVerificationsList).Methods("GET")
	app.HTTP["api"].Router.HandleFunc("/v1/api/verify/load", verifier.VerificationCodeLoad).Methods("GET")
	app.HTTP["api"].Router.HandleFunc("/v1/api/verify/check", verifier.VerificationCodeCheck).Methods("POST")
	app.HTTP["api"].Router.HandleFunc("/v1/api/verify/status", verifier.VerifyStatusHandler).Methods("GET")
	app.HTTP["api"].Router.HandleFunc("/v1/api/verify/qr/generate", verifier.QRCodeGenerate).Methods("GET")
	//||------------------------------------------------------------------------------------------------||
	//|| Verify - ID
	//||------------------------------------------------------------------------------------------------||
	app.HTTP["api"].Router.HandleFunc("/v1/api/verify/id/init", verifier.IdentifierVerifyInitHandler).Methods("GET")
	app.HTTP["api"].Router.HandleFunc("/v1/api/verify/id/media/upload", verifier.VerifyIDMediaUpload).Methods("POST")
	app.HTTP["api"].Router.HandleFunc("/v1/api/verify/id/media/fetch", verifier.VerifyIDMediaFetch).Methods("GET")
	app.HTTP["api"].Router.HandleFunc("/v1/api/verify/id/success", verifier.VerifyIDSuccessHandler).Methods("POST")
	//||------------------------------------------------------------------------------------------------||
	//|| Verify - Face
	//||------------------------------------------------------------------------------------------------||
	app.HTTP["api"].Router.HandleFunc("/v1/api/verify/face/init", verifier.FaceVerifyInitHandler).Methods("GET")
	app.HTTP["api"].Router.HandleFunc("/v1/api/verify/face/media/upload", verifier.VerifyFaceMediaUpload).Methods("POST")
	app.HTTP["api"].Router.HandleFunc("/v1/api/verify/face/media/fetch", verifier.VerifyFaceMediaFetch).Methods("GET")
	app.HTTP["api"].Router.HandleFunc("/v1/api/verify/face/success", verifier.VerifyFaceSuccessHandler).Methods("POST")
	//||------------------------------------------------------------------------------------------------||
	//|| Verify - Phone
	//||------------------------------------------------------------------------------------------------||
	app.HTTP["api"].Router.HandleFunc("/v1/api/verify/phone", verifier.PhoneVerifyInitHandler).Methods("POST")
	//||------------------------------------------------------------------------------------------------||
	//|| Verify - Card
	//||------------------------------------------------------------------------------------------------||
	app.HTTP["api"].Router.HandleFunc("/v1/api/verify/card/init", verifier.CCVerifyInitHandler).Methods("POST")
	app.HTTP["api"].Router.HandleFunc("/v1/api/verify/card/success", verifier.CCVerifySuccessHandler).Methods("POST")
	//||------------------------------------------------------------------------------------------------||
	//|| Verify - Address
	//||------------------------------------------------------------------------------------------------||
	app.HTTP["api"].Router.HandleFunc("/v1/api/verify/address/init", verifier.AddressVerifyInitHandler).Methods("POST")
	app.HTTP["api"].Router.HandleFunc("/v1/api/verify/address/success", verifier.AddressVerifySuccessHandler).Methods("POST")
	//||------------------------------------------------------------------------------------------------||
	//|| Tools
	//||------------------------------------------------------------------------------------------------||
	app.HTTP["api"].Router.HandleFunc("/v1/api/currency", utils.ConvertUSDHandler).Methods("GET")
	app.HTTP["api"].Router.HandleFunc("/health", utils.HealthHandler).Methods("GET")
	//||------------------------------------------------------------------------------------------------||
	//|| Admin
	//||------------------------------------------------------------------------------------------------||
	app.HTTP["api"].Router.HandleFunc("/v1/api/verify/json", verifier.TestingJSONVerification).Methods("GET")
	app.HTTP["api"].Router.HandleFunc("/v1/api/verify/id/reset", verifier.TestingResetVerification).Methods("GET")
}
