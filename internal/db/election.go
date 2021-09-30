package db

import "go.mongodb.org/mongo-driver/mongo"

type ElectionRepo struct {
	collection *mongo.Collection
}

func NewElectionRepo(db *mongo.Database) *ElectionRepo {
	return &ElectionRepo{collection: db.Collection(NameElectionsCollection)}
}

func (e *ElectionRepo) GetCandidates() {

}