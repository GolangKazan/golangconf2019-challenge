package challenge

import (
	"math/rand"
	"testing"
	"time"
)

func TestWinner(t *testing.T) {
	// filtered is a list of participating gophers.
	// Basically, a gophers list without ignored members.
	filtered := []gopher{}

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

		filtered = append(filtered, g)
	}

	if len(filtered) < 5 {
		rand.Seed(time.Now().UTC().UnixNano())
		winner := rand.Intn(len(filtered))
		t.Logf("The winner is: %s!", filtered[winner].key())
	} else {
		t.Log("Not enough participants")
	}
}
