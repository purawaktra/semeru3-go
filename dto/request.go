package dto

type BodyProfile struct {
	AccountId   int    `json:"account-id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	City        int    `json:"city"`
	Province    int    `json:"province"`
	Zipcode     string `json:"zipcode"`
	PhoneNumber string `json:"phone_number"`
}

type Header struct {
	RequestId string `header:"request-id"`
	Host      string `header:"host"`
}
