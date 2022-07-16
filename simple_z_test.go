package main

// Call: go run test -v *.go

import (
	"time"
	"testing"
	"math/rand"
)

const (
	// Min/Max length of test string
	MinLength = 10
	MaxLength = 100
)


func init () {
	rand.Seed(time.Now().UnixNano())
}


func randomN (low, hi int) int {
	return rand.Intn(hi - low) + low
}


func createRandomString () string {
	length := randomN(MinLength, MaxLength)
	i := 0
	str := ""

	for i < length {
		// Use any printable unicode character except '$'
		str = str + string(rune(randomN(48,127)))
		i++
	}

	return str
}


func cutPattern (str string) string {
	length := len(str)

	start := randomN(0, length - 2)
	end := randomN(start + 1, length)

	return str[start: end]
}


func Test_Simple_Z (t *testing.T) {
	fiveSecTimer := time.NewTimer(time.Duration(time.Duration(5) * time.Second))
	count := 0

	for {
		select{
		case <- fiveSecTimer.C:
			t.Logf("%d tests successfully run.", count)
			return
		default:
			count++
			str := createRandomString()
			pat := cutPattern(str)
			
			z, err := Simple_Z(pat, str)
			if err != nil {
				t.Error(err)
			}
			
			// Proof modified to only print for errors.
			proof := pat == str[z.Start: z.End]
			if !proof {
				t.Error("simple_z is not returning accurate results.")
			}
		}
	}
}	