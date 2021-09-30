package struction


type Voter struct {
	Valid bool 				`bson:"valid"`
	Voted bool				`bson:"voted"`
	Token string			`bson:"token"`
	NameElection string		`bson:"name_election"`
}
