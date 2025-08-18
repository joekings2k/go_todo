package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomInt(min,max int64)int64{
	return min + seededRand.Int63n(max-min+1) 
}


func RandomString(n int)string{
	var sb strings.Builder
	k := len(alphabet)
	for i:=0;i<n;i++{
		c := alphabet[seededRand.Intn(k)] // returns a random int n, 0 <= n < k
		sb.WriteByte(c)

	}
	return sb.String()
}

func RandomUsername()string{
	return  RandomString(6)
}

func RandomEmail()string{
	return RandomString(6) + "@gmail.com"
}

func RandomPassword()string{
	return RandomString(10)
}

func RandomTodoTitle()string{
	return RandomString(10)
}

func RandomTodoStatus()string{
	status := []string{"pending","in_progress","completed"}
	n := len(status)
	return status[seededRand.Intn(n)]
}