package blogic

import (
	"context"
	"electionSystem/internal/db"
	"electionSystem/internal/struction"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"math/rand"
	"strings"
	"time"
)

type IBVoting interface {
	Vote(token string, VotingCandidate string) (int, string)
	Login(token string) (bool, string)
	GetVoteInElection(nameElection string) (int, string)
	GetCandidates(nameElection string) (map[string]string, error)
}

type BVoting struct {
	DBVoit db.IVoting
	DBElec db.IElections
}

func CreateBVoting(dbs *mongo.Database) *BVoting {
	return &BVoting{DBVoit: db.NewVotRepo(dbs), DBElec: db.NewElectionRepo(dbs)}
}
func (v *BVoting) GetCandidates(nameElection string) (map[string]string, error) {
	elect, err := v.DBElec.GetElection(context.TODO(), nameElection)
	return elect.ElectionCandidates, err
}
func (v *BVoting) Login(token string) (bool, string) {
	voter, err := v.DBVoit.GetInfoInToken(context.TODO(), token)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return false, "Токен не найден"
		}
		return false, err.Error()
	}
	if voter.Valid != true {
		return false, "Токен не валиден"
		//return false, "token is not valid"
	}
	if voter.Voted == true {
		//return false, "token is voted already"
		return false, "Токен уже был использован"
	}
	return true, voter.NameElection
}

func (v *BVoting) Vote(token string, VotingCandidate string) (int, string) {
	//получить нф о токене
	voter, err := v.DBVoit.GetInfoInToken(context.TODO(), token)
	if err != nil {
		return 404, ""
	}
	//проверить на валидность и отсуствие голоса
	if voter.Valid != true {
		return 403, "Токен не валиден"
	}
	if voter.Voted != false {
		return 423, "Токен уже был использован"
	}
	//получить данные голосования
	election, er := v.DBElec.GetElection(context.TODO(), voter.NameElection)
	if er != nil {
		return 500, ""
	}

	//проверить голосоване на даты
	if election.StartDate.Unix() > time.Now().Unix() {
		return 404, "Голосование ещё не началось"
	}
	if election.EndDate.Unix() < time.Now().Unix() {
		return 404, "Голосование уже закончилось"
	}

	//поставить на токен наличие голоса, в запросе указать проверку на валидность и голос
	countUpd, errr := v.DBVoit.VotedToken(context.TODO(), token)
	if errr != nil {
		return 500, ""
	}
	if countUpd < 1 {
		return 400, ""
	}
	//приплюсовать голос кандидату
	cUpd, e := v.DBElec.VotingIncrement(context.TODO(), voter.NameElection, VotingCandidate)
	if e != nil {
		return 500, ""
	}
	fmt.Printf("token: %v candidate: %v Return Election Update: %v", token, VotingCandidate, cUpd)
	return 200, "OK"
}

func (b *BVoting) BAddVoted(token, nameElection string) {
	var data = struction.Voter{
		Token:        token,
		NameElection: nameElection,
		Voted:        false,
		Valid:        true,
	}

	_, err := b.DBVoit.AddVoter(context.TODO(), data)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func (b *BVoting) GetVoteInElection(nameElection string) (int, string) {
	elect, err := b.DBElec.GetElection(context.TODO(), nameElection)
	if err != nil {
		return 400, err.Error()
	}
	jss, er := json.Marshal(elect.ElectoralVotes)
	if er != nil {
		return 400, er.Error()
	}
	return 200, string(jss)
}

func GenerateToken(passwordLength int) string {
	lowerCharSet := "abcdedfghijklmnopqrst"
	upperCharSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//specialCharSet := "!@#$%&*"
	numberSet := "0123456789"
	allCharSet := lowerCharSet + upperCharSet /*+ specialCharSet*/ + numberSet
	rand.Seed(time.Now().Unix())
	minSpecialChar := 1
	minNum := 1
	minUpperCase := 1

	var password strings.Builder

	/*
		//Set special character
		for i := 0; i < minSpecialChar; i++ {
			random := rand.Intn(len(specialCharSet))
			password.WriteString(string(specialCharSet[random]))
		}*/

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}
