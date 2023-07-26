package main

import (
	"context"
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"time"
)

//const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // +"abcdefghijklmnopqrstuvwxyz"
const (
	charset      = "AEFHLMNPRWY34679"
	charsetLen   = len(charset)
	appealIDSize = 10
	maxAttempt   = 1000000
)

//var seededRand *rand.Rand
type appeal struct {
	AppealID    string
	TccAppealID string
}

func generateID(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		//b[i] = charset[seededRand.Intn(charsetLen)]
		b[i] = charset[rand.Intn(charsetLen)]
	}
	return string(b)
}

func init() {
	// basic
	// rand.Seed(time.Now().UnixNano())
	bSize := 8
	b := make([]byte, bSize)
	_, err := crand.Read(b)
	var seed int64
	if err != nil {
		fmt.Println("Cannot seed generator without random number")
		seed = time.Now().UnixNano()
	} else {
		seed = int64(binary.LittleEndian.Uint64(b))
	}
	rand.Seed(seed)
	//seededRand = rand.New(rand.NewSource(seed))
}

func main() {
	aID := generateID(appealIDSize, charset)
	// fmt.Println(aID)

	ap := []appeal{}
	for i := 0; i < 10; i++ {
		aID = generateID(appealIDSize, charset)
		ap = append(ap, appeal{AppealID: aID})
	}

	// fmt.Println(ap)

	apMap := make(map[string]bool)
	for _, v := range ap {
		apMap[v.AppealID] = true
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	var cnt int
	var err error
loop:
	for {
		cnt++
		select {
		case <-ctx.Done():
			aID = ""
			err = ctx.Err()
			err = fmt.Errorf("unable to get new appeal ID: %v", err)
			break loop
		default:
			aID = generateID(appealIDSize, charset)
			_, ok := apMap[aID]
			if !ok {
				err = nil
				break loop
			}
		}
		if cnt > maxAttempt {
			aID = ""
			err = fmt.Errorf("Hit Attempt max: %v", maxAttempt)
			//break loop
			cancel()
		}

	}
	if err != nil {
		log.Println(err)
	}
	fmt.Println("New AppealID:", aID)
	fmt.Println("attemps:", cnt)

	apMap = make(map[string]bool)
	var dup int
	n := 2000000
	cnt = 0
	var i int
	for i = 0; i < n; i++ {
		aID = generateID(appealIDSize, charset)
		_, ok := apMap[aID]
		if ok {
			dup++
			break
		}
		apMap[aID] = true
		cnt++
	}
	fmt.Printf("After %v Id found dupes:%v non-dups:%v\n", i, dup, cnt)
}
