package models

import (
	"time"

	"github.com/rahul-sinha1908/go-mongoose/mongoose"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Accounts struct {
	ID     			primitive.ObjectID 	`json:"_id" bson:"_id,omitempty"`
	Name  			string             	`json:"accountName" bson:"accountName,omitempty"`
	Users 			primitive.A			`json:"-" bson:"users" mson:"collection=Users"`
	UserObjects		[]Users				`json:"users"`
	Agents 			primitive.A			`json:"-" bson:"agents" mson:"collection=Agents"`
	AgentObjects	[]Agents			`json:"agents"`
	CreationDate	time.Time			`json:"creationDate,omitempty" bson:"creationDate,omitempty"`
}

func (acct *Accounts) PopulateUsers() error {
	err := mongoose.PopulateObjectArray(acct,"Users",&acct.UserObjects);

	if(err != nil){
		return err;
	}
	return nil;
}

func (acct *Accounts) PopulateAgents() error {
	err := mongoose.PopulateObjectArray(acct,"Agents",&acct.AgentObjects);

	if(err != nil){
		return err;
	}
	return nil;
}