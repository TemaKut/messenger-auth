package user

type ErrorReason = string

var (
	ErrorReasonUserEmailAlreadyExists ErrorReason = "auth.user-email-already-exist"
	ErrorInvalidCredentials           ErrorReason = "auth.invalid-user-credentials"
	ErrorReasonValidation             ErrorReason = "auth.validation"
)
