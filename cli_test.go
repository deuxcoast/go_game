package poker_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	poker "github.com/duexcoast/go_game"
)

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

type GameSpy struct {
	StartCalled     bool
	StartCalledWith int
	BlindAlert      []byte

	FinshCalled      bool
	FinishCalledWith string
}

func (g *GameSpy) Start(numberOfPlayers int, out io.Writer) {
	g.StartCalled = true
	g.StartCalledWith = numberOfPlayers
	out.Write(g.BlindAlert)
}

func (g *GameSpy) Finish(winner string) {
	g.FinshCalled = true
	g.FinishCalledWith = winner
}

func TestCLI(t *testing.T) {
	t.Run("start game with 6 players and finish game with Chris as winner", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userSends("6", "Chris wins")

		game := &GameSpy{}
		cli := poker.NewCLI(in, stdout, game)

		cli.PlayPoker()
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 6)
		assertFinishCalledWith(t, game, "Chris")

	})

	t.Run("starts game with 8 players, and record Cleo win", func(t *testing.T) {
		in := userSends("8", "Cleo wins")
		game := &GameSpy{}
		cli := poker.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		assertGameStartedWith(t, game, 8)
		assertFinishCalledWith(t, game, "Cleo")
	})

	t.Run("it prints an error when a non numeric value is entered and does not start game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userSends("Stupid")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		if game.StartCalled {
			t.Errorf("game should not have been started")
		}

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})

	t.Run("it prints an error when user input does not indicate winner, and does not finish game", func(t *testing.T) {
		in := userSends("8", "Lloyd is a killer")
		stdout := &bytes.Buffer{}

		game := &GameSpy{}
		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		if game.FinshCalled {
			t.Errorf("game should not have been finished")
		}

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadWinnerInputMessage)
	})
}

func userSends(messages ...string) io.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}

func assertScheduledAlert(t testing.TB, got, want poker.ScheduledAlert) {
	t.Helper()

	if got.Amount != want.Amount {
		t.Errorf("got amount %d want %d", got.Amount, want.Amount)
	}

	if got.At != want.At {
		t.Errorf("got schedule time %v want %v", got.At, want.At)
	}
}

func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()

	want := strings.Join(messages, "")
	got := stdout.String()

	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}

func assertFinishCalledWith(t testing.TB, game *GameSpy, winner string) {
	t.Helper()

	if game.FinishCalledWith != winner {
		t.Errorf("expected finish called with %q, but got %q", winner, game.FinishCalledWith)
	}
}

func assertGameStartedWith(t testing.TB, game *GameSpy, numberOfPlayers int) {
	if game.StartCalledWith != numberOfPlayers {
		t.Errorf("wanted Start called with %d but got %d", numberOfPlayers, game.StartCalledWith)
	}
}
