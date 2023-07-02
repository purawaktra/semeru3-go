package dto

type RequestBody struct {
	RequestId string
	Timestamp string `json:"timestamp" validate:"required,len=13"`
	Data      interface{}
}
type RequestDataAccount struct {
	AccountId    uint   `json:"account-id" validate:"numeric"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	EmailAddress string `json:"email_address"`
}
