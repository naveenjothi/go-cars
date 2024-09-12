package models

type DriverModel struct {
	FirstName           string  `json:"firstName" bson:"firstName"`
	LastName            string  `json:"lastName" bson:"lastName"`
	Mobile              string  `json:"mobile" bson:"mobile"`
	PhotoUrl            string  `json:"photoUrl" bson:"photoUrl"`
	Email               string  `json:"email" bson:"email"`
	Address             Address `json:"address" bson:"address"`
	IsMobileVerified    bool    `json:"isMobileVerified" bson:"isMobileVerified"`
	IsKYCVerified       bool    `json:"isKYCVerified" bson:"isKYCVerified"`
	Earnings            float64 `json:"earnings" bson:"earnings"`
	DriverLicenseNumber string  `json:"driverLicenseNumber" bson:"driverLicenseNumber"`
	ExperienceYears     int     `json:"experienceYears" bson:"experienceYears"`
	Availability        bool    `json:"availability" bson:"availability"`
	BaseModel           `bson:",inline"`
}
