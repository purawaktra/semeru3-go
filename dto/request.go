package dto

type RequestBody struct {
	RequestId string
	Data      any
}

type RequestDataAccount struct {
	AccountId    uint   `json:"account-id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	EmailAddress string `json:"email_address"`
}
