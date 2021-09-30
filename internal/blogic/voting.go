package blogic

import (
	"context"
	"electionSystem/internal/db"
	"electionSystem/internal/struction"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"math/rand"
	"strings"
	"time"
)

type BVoting struct {
	DB db.IVoting
}

func CreateBVoting(dbs *mongo.Database) *BVoting {
	return &BVoting{DB: db.NewVotRepo(dbs)}
}

func (b *BVoting) BAddVoted() {
	var data = struction.Voter{
		Token: "ffff",
		NameElection: "test",
		Voted: true,
		Valid: true,
	}

	_, err := b.DB.AddVoter(context.TODO(), data)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func GenerateToken(passwordLength int) string {
	lowerCharSet := "abcdedfghijklmnopqrst"
	upperCharSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//specialCharSet := "!@#$%&*"
	numberSet := "0123456789"
	allCharSet := lowerCharSet + upperCharSet /*+ specialCharSet*/ + numberSet
	time.Sleep(1000*time.Millisecond)
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