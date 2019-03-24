package blueutil

import (
	"bytes"
	"math/rand"
	"time"
)

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

// RandStringN generate length n string
func RandStringN(n int) string {
	var result bytes.Buffer
	var temp string
	for i := 0; i < n; {
		if string(randInt(65, 90)) != temp {
			temp = string(randInt(65, 90))
			result.WriteString(temp)
			i++
		}
	}
	return result.String()
}
