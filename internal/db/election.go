package db

import (
	"context"
	"electionSystem/internal/struction"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IElections interface {
	GetElection(ctx context.Context, nameElection string) (struction.Election, error)
	VotingIncrement(ctx context.Context, nameElection, VotingCandidate string) (int, error)
}
type ElectionRepo struct {
	collection *mongo.Collection
}

func NewElectionRepo(db *mongo.Database) *ElectionRepo {
	return &ElectionRepo{collection: db.Collection(NameElectionsCollection)}
}

func (e *ElectionRepo) GetElection(ctx context.Context, nameElection string) (struction.Election, error) {
	filter := bson.M{"name": nameElection}
	var election struction.Election
	err := e.collection.FindOne(ctx, filter).Decode(&election)
	return election, err
}

func (e *ElectionRepo) VotingIncrement(ctx context.Context, nameElection, VotingCandidate string) (int, error) {
	filter := bson.M{"name": nameElection}
	update := bson.D{
		{"$inc", bson.D{
			{"electoral_votes." + VotingCandidate, 1},
		}},
	}
	updResult, err := e.collection.UpdateOne(ctx, filter, update)
	fmt.Println(updResult)
	return int(updResult.MatchedCount), err
}
