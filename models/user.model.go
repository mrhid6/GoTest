package models

import (
	"time"

	"github.com/rahul-sinha1908/go-mongoose/mongoose"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID				primitive.ObjectID		`json:"_id,omitempty" bson:"_id,omitempty"`
	Email			string					`json:"email,omitempty" bson:"email,omitempty"`
	Password		string					`json:"-" bson:"password,omitempty"`
	IsAccountAdmin	bool					`json:"isAccountAdmin" bson:"isAccountAdmin"`
	TwoFASecret		string					`json:"twoFASecret" bson:"twoFASecret"`
	TwoFASetup		bool					`json:"twoFASetup" bson:"twoFASetup"`
	IsActive		bool					`json:"active" bson:"active"`
	Role			primitive.ObjectID		`json:"-" bson:"role" mson:"collection=UserRoles"`
	RoleObject		UserRoles				`json:"role"`
	CreationDate	time.Time				`json:"creationDate,omitempty" bson:"creationDate,omitempty"`
	LastActive		time.Time				`json:"lastActive,omitempty" bson:"lastActiveDate,omitempty"`
}

func (user *Users) PopulateUserRole() {
	mongoose.PopulateObject(user,"Role",&user.RoleObject);
}

type UserRoles struct {
	ID				primitive.ObjectID		`json:"_id,omitempty" bson:"_id,omitempty"`
	Name			string					`json:"name,omitempty" bson:"roleName,omitempty"`
	Editable        bool					`json:"editable" bson:"canEdit"`
}