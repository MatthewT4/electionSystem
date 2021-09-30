package struction

import "time"

type Voter struct {
	Valid        bool   `bson:"valid"`
	Voted        bool   `bson:"voted"`
	Token        string `bson:"token"`
	NameElection string `bson:"name_election"`
}

type Election struct {
	Name               string         `bson:"name"`
	StartDate          time.Time      `bson:"start_date"`
	EndDate            time.Time      `bson:"end_date"`
	ElectionCandidates []string       `bson:"election_candidates"`
	ElectoralVotes     map[string]int `bson:"electoral_votes"`
	CoundVoting        int            `bson:"cound_voting"`
}
