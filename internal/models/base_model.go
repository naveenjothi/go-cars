package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseModel struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IsActive  bool               `json:"isActive" bson:"isActive"`
	IsDeleted bool               `json:"isDeleted" bson:"isDeleted"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	DeletedAt *time.Time         `json:"deletedAt,omitempty" bson:"deletedAt,omitempty"`
}

func (bm *BaseModel) InitiliseDefaultValue() {
	now := time.Now().UTC()
	bm.IsDeleted = false
	bm.IsActive = true
	bm.CreatedAt = now
	bm.UpdatedAt = now
}

func (bm *BaseModel) UpdateDefaultValue() {
	now := time.Now().UTC()
	bm.UpdatedAt = now
}
