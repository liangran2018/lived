package env

import "testing"

func TestGameTime_Add(t *testing.T) {
	tt := &gameTime{2018, 12, 31, 23, 59, 100}
	tt.Add(2, 0)
	t.Log(tt)
}

func TestGameTime_Month(t *testing.T) {
	i := NewTime().Month()
	t.Log(i)
}