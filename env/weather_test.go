package env

import "testing"
import "math/rand"
import "time"

func TestGetWeather(t *testing.T) {
	m := weather[Spring]
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(100)
	t.Log(i)
	for k, v := range m {
		if v.min <= i && i <= v.max {
			t.Log(k)
		}
	}
}
