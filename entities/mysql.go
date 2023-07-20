package entities

type Accounts struct {
	AccountId   uint   `gorm:"primary_key"`
	FirstName   string `gorm:"column:first_name"`
	LastName    string `gorm:"column:last_name"`
	Address     string `gorm:"column:address"`
	City        uint   `gorm:"column:city"`
	Province    uint   `gorm:"column:province"`
	Zipcode     string `gorm:"column:zipcode"`
	PhoneNumber uint   `gorm:"column:phone_number"`
}
