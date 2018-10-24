package hangman

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	// Active describes an in progress game
	Active Status = iota
	// Lost the game
	Lost
	// Won the game
	Won
	// AlreadyGuessed a letter
	AlreadyGuessed

	// MaxGuesses tracks the maximum number of guesses
	MaxGuesses = 7
)

type (
	// Status of the game
	Status int

	// Tally tracks a game progress
	Tally struct {
		TurnsLeft int    `json:"turns_left"`
		Letters   []rune `json:"letters"`
		Status    Status `json:"status"`
	}
)

var promTally = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "hangman_tally_total",
	Help: "The total number of game won or lost",
	ConstLabels: map[string]string{
		"app": "hangman",
	},
})

// NewTally initializes a tally
func NewTally(word []rune) *Tally {
	return &Tally{TurnsLeft: MaxGuesses, Letters: updateLetters(word, []rune{})}
}

// Update the tally with a new guess
func (t *Tally) Update(word, guesses []rune) {
	t.Letters = updateLetters(word, guesses)
	if !t.guessesLeft() {
		promTally.Inc()
		t.Status = Won
	}
	if t.TurnsLeft == 0 {
		promTally.Dec()
		t.Status = Lost
	}
}

func (t *Tally) guessesLeft() bool {
	for _, l := range t.Letters {
		if l == '_' {
			return true
		}
	}
	return false
}

func updateLetters(word, guesses []rune) []rune {
	ll := make([]rune, len(word))
	for i, l := range word {
		if inGuesses(guesses, l) {
			ll[i] = l
		} else {
			ll[i] = '_'
		}
	}
	return ll
}

func inGuesses(guesses []rune, g rune) bool {
	for _, l := range guesses {
		if g == l {
			return true
		}
	}
	return false
}
