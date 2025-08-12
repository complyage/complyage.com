package constants

type VerificationType string

const (
	VerificationEmail        VerificationType = "MAIL"
	VerificationPhone        VerificationType = "PHNE"
	VerificationAge          VerificationType = "UAGE"
	VerificationAddress      VerificationType = "ADDR"
	VerificationCreditCard   VerificationType = "CRCD"
	VerificationProfilePhoto VerificationType = "PROF"
	VerificationUsername     VerificationType = "UNAM"
)

var AllVerificationTypes = []VerificationType{
	VerificationEmail,
	VerificationPhone,
	VerificationAge,
	VerificationAddress,
	VerificationCreditCard,
	VerificationProfilePhoto,
	VerificationUsername,
}
