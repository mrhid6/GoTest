package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AgentConfigSFVersion struct{
	Installed 			int		`json:"installed" bson:"installed"`
	Available 			int		`json:"available" bson:"available"`

}

type AgentConfig struct{
	Version 			string					`json:"version" bson:"version"`
	WorkerThreads 		int						`json:"workerThreads" bson:"workerThreads"`
	Branch 				string					`json:"sfBranch" bson:"sfBranch"`
	MaxPlayers 			int						`json:"maxPlayers" bson:"maxPlayers"`
	SFVersion			AgentConfigSFVersion 	`json:"sfVersion" bson:"sfVersions"`
}

type Agents struct {
	ID     			primitive.ObjectID 	`json:"_id" bson:"_id,omitempty"`
	Name			string				`json:"name" bson:"agentName"`
	APIKey			string				`json:"-" bson:"apiKey"`
	Online			bool				`json:"online" bson:"online"`
	Running			bool				`json:"running" bson:"running"`
	Installed		bool				`json:"installed" bson:"installed"`
	NeedsUpdate		bool				`json:"needsUpdate" bson:"needsUpdate"`
	SFPortNum		int					`json:"portNumber" bson:"sfPortNum"`
	Config			AgentConfig			`json:"config" bson:"config"`
	LastComm		time.Time			`json:"lastComm,omitempty" bson:"lastCommDate,omitempty"`
	CreationDate	time.Time			`json:"creationDate,omitempty" bson:"creationDate,omitempty"`
}