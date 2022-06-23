package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ContactForm struct {
	ID      primitive.ObjectID `json:"_id,onitempty" bson:"_id,onitempty"`
	Nom     string             `json:"nom,onitempty" bson:"nom,onitempty"`
	Prenom  string             `json:"prenom,onitempty" bson:"prenom,onitempty"`
	Email   string             `json:"email,onitempty,unique" bson:"email,onitempty,unique"`
	Message string             `json:"message,onitempty" bson:"message,onitempty"`
}

type Newsletter struct {
	Email string `json:"email,onitempty,unique" bson:"email,onitempty,unique"`
}
