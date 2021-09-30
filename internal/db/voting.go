package db

import (
	"context"
	"electionSystem/internal/struction"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type IVoting interface {
	GetInfoInToken(ctx context.Context, token string) (struction.Voter, error)
	AddVoter(ctx context.Context, voter struction.Voter) (interface{}, error)
	VotedToken(ctx context.Context, token string) (int, error)
}

type Voting struct {
	Valid        bool   `bson:"valid"`
	Voted        bool   `bson:"voted"`
	Token        string `bson:"token"`
	NameElection string `bson:"name_election"`
}

type VotRepo struct {
	collection *mongo.Collection
}

func NewVotRepo(db *mongo.Database) *VotRepo {
	return &VotRepo{collection: db.Collection(NameVoitingCollection)}
}

func (v *VotRepo) GetInfoInToken(ctx context.Context, token string) (struction.Voter, error) {
	filter := bson.M{"token": token}
	var voter struction.Voter
	err := v.collection.FindOne(ctx, filter).Decode(&voter)
	return voter, err
}

func (v *VotRepo) AddVoter(ctx context.Context, voter struction.Voter) (interface{}, error) {
	bs, er := bson.Marshal(voter)
	if er != nil {
		log.Fatal(er.Error())
		return 0, er
	}
	result, err := v.collection.InsertOne(ctx, bs)
	if err != nil {
		return 0, err
	}
	return result.InsertedID, err
}

func (v *VotRepo) VotedToken(ctx context.Context, token string) (int, error) {
	filter := bson.M{"token": token, "voted": false, "valid": true}
	update := bson.D{
		{"$set", bson.D{
			{"voted", true},
		}},
	}
	updResult, err := v.collection.UpdateOne(ctx, filter, update)
	fmt.Println(updResult)
	return int(updResult.MatchedCount), err
}
