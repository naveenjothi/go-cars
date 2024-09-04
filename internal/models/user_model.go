package models

type Address struct{
	City string `json:"city" bson:"city"`
	Country string `json:"country" bson:"country"`
}

type UserModel struct{
	FirstName string `json:"firstName" bson:"firstName"`
	LastName string `json:"lastName" bson:"lastName"`
	Mobile string `json:"mobile" bson:"mobile"`
	Email string `json:"email" bson:"email"`
	ProfileName string `json:"profileName" bson:"profileName"`
	Address Address `json:"address" bson:"address"`
	BaseModel
}