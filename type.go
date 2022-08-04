package esms

type StatusCode string

// each sms type has different phone number and cost
type SmsType string
type SandboxMode string

const (
	CodeSuccess             StatusCode = "100"
	CodeUnknown             StatusCode = "99"
	CodeUnauthorized        StatusCode = "101"
	CodeAccountBlocked      StatusCode = "102"
	CodeInsufficientBalance StatusCode = "103"
	CodeInvalidBrandName    StatusCode = "104"
	CodeInvalidSmsType      StatusCode = "118"
	// advertise brand name message requires at least 20 recipients
	CodeBrandNameMessageMinRecipients StatusCode = "119"
	CodeDuplicatedRequestID           StatusCode = "124"
	// advertise brand name message content should have at least 422 characters
	CodeBrandNameMessageMinContent StatusCode = "131"
	// the user don't have permission to send message from phone number 8755
	CodePhoneNumber8755NotAllowed StatusCode = "132"
	CodeNonRegisteredBrandName    StatusCode = "177"
	// the RequestID should not have more than 120 characters
	CodeInvalidRequestID                         StatusCode = "159"
	CodeInvalidSocialNetworkTemplate             StatusCode = "145"
	CodeInvalidBrandNameCustomerSupportBrandName StatusCode = "146"

	SmsTypeBrandName       SmsType = "2"
	SmsTypeStatic          SmsType = "8"
	SmsTypeZaloPrioritized SmsType = "24"
	SmsTypeZalo            SmsType = "25"

	NoSandbox SandboxMode = "0"
	Sandbox   SandboxMode = "1"
)

type Credential struct {
	ApiKey    string `json:"ApiKey"`
	SecretKey string `json:"SecretKey"`
}
