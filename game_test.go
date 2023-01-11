package poker_test

import "testing"

func TestGame_Start(t *testing.T) {
	t.Run("schedules alerts on game start for 5 players", func(t *testing.T) {
		BlindAlerter := &poker.SpyBlindAlerter{}
	})
}
