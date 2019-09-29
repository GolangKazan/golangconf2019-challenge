package challenge

import (
	"math/rand"
	"testing"
	"time"
)

func runWinnerTest(t *testing.T, gophers []gopher) {
	if len(gophers) >= 5 {
		rand.Seed(time.Now().UTC().UnixNano())
		winner := rand.Intn(len(gophers))
		t.Logf("The winner is: %s!", gophers[winner].key())
	} else {
		t.Log("Not enough participants")
	}
}

func TestWinner(t *testing.T) {
	kazanGophers := []gopher{}
	nizhnyGophers := []gopher{}

	// keys is a set of all participants keys.
	// Needed to detect collisions.
	keys := map[string]bool{}

	// Check whether all fields are set correctly.
	for i, g := range gophers {
		if g.name == "" {
			t.Errorf("gopher[%d] name is empty", i)
			continue
		}
		if g.post == "" {
			t.Errorf("gopher[%d] (%s) post is empty", i, g.name)
			continue
		}

		if keys[g.key()] {
			t.Errorf("gopher[%d] (%s) is a duplicate", i, g.name)
			continue
		}
		keys[g.key()] = true

		if g.tester {
			continue
		}

		if g.nizhnyNovgorod {
			nizhnyGophers = append(nizhnyGophers, g)
		} else {
			kazanGophers = append(kazanGophers, g)
		}
	}

	t.Run("Kazan", func(t *testing.T) {
		runWinnerTest(t, kazanGophers)
	})
	t.Run("NizhnyNovgorod", func(t *testing.T) {
		runWinnerTest(t, nizhnyGophers)
	})
}
