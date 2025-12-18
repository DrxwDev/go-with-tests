package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountDownOperations{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got: %q | want: %q", got, want)
		}
	})
	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPriter := &SpyCountDownOperations{}
		Countdown(spySleepPriter, spySleepPriter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPriter.Calls) {
			t.Errorf("wanted calls: %v | got: %v", want, spySleepPriter.Calls)
		}
	})
}

func TestConfurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.SetDurationSlept}
	sleeper.Sleep()
}
