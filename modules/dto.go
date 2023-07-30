package modules

type Profile struct {
	AccountId   int    `json:"account-id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	City        int    `json:"city"`
	Province    int    `json:"province"`
	Zipcode     string `json:"zipcode"`
	PhoneNumber string `json:"phone_number"`
}
