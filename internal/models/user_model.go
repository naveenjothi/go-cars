package models

type Address struct {
	City      string `json:"city" bson:"city"`
	CityId    string `json:"cityId" bson:"cityId"`
	State     string `json:"state" bson:"state"`
	CountryId string `json:"countryId" bson:"countryId"`
	Country   string `json:"country" bson:"country"`
}

type UserModel struct {
	FirstName        string  `json:"firstName" bson:"firstName"`
	LastName         string  `json:"lastName" bson:"lastName"`
	Mobile           string  `json:"mobile" bson:"mobile"`
	PhotoUrl         string  `json:"photoUrl" bson:"photoUrl"`
	Email            string  `json:"email" bson:"email"`
	ProfileName      string  `json:"profileName" bson:"profileName"`
	Address          Address `json:"address" bson:"address"`
	IsMobileVerified bool    `json:"isMobileVerified" bson:"isMobileVerified"`
	IsKYCVerified    bool    `json:"isKYCVerified" bson:"isKYCVerified"`
	BaseModel        `bson:",inline"`
}
