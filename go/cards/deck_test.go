package cards

import (
	"testing"
)

func TestDeck36(t *testing.T) {
	testDeck(t, New36CardDeck(), Six, 1)
}

func TestDeck52(t *testing.T) {
	testDeck(t, New52CardDeck(), Two, 1)
}

func TestDeck54(t *testing.T) {
	testDeck(t, New54CardDeck(), Joker, 1)
}

func TestCustom(t *testing.T) {
	if _, err := NewCustomDeck(New54CardDeck(), New36CardDeck()); err == nil {
		t.Error("Custom deck cannot contain different deck types!")
	}
	cdeck, _ := NewCustomDeck(New54CardDeck())
	if _, err := NewCustomDeck(cdeck); err == nil {
		t.Error("Custom deck cannot contain custom decks!")
	}
	for n := 1; n < 10; n++ {
		if deck, err := NewCustomDeck(decks(n, New36CardDeck)...); err == nil {
			testDeck(t, deck, Six, n)
		}
		if deck, err := NewCustomDeck(decks(n, New52CardDeck)...); err == nil {
			testDeck(t, deck, Two, n)
		}
		if deck, err := NewCustomDeck(decks(n, New54CardDeck)...); err == nil {
			testDeck(t, deck, Joker, n)
		}
	}
}

func decks(n int, constructor DeckConstructor) []Deck {
	result := make([]Deck, n)
	for i := 0; i < n; i++ {
		result = append(result, constructor())
	}
	return result
}

func testDeck(t *testing.T, deck Deck, maxRank Rank, maxCount int) {
	joker := NewJoker()
	if !deck.HasCards() {
		t.Error("Deck should have cards!")
	}
	for suit := Hearts; suit < none_suit; suit++ {
		for rank := Ace; rank < none_rank; rank++ {
			if card, err := NewCard(suit, rank); err == nil {
				if deck.Put(card) {
					t.Error("Cannot add card to full deck!")
				}
				if (rank <= maxRank) {
					if !deck.Contains(card) {
						t.Error("Full deck should contain all cards!")
					}
				} else {
					if deck.Contains(card) {
						t.Error("Full deck shouldn't contain extra cards!")
					}
				}
			}
		}
	}
	if deck.Put(joker) {
		t.Error("Cannot add joker to full deck!")
	}
	if (joker.rank <= maxRank) {
		if !deck.Contains(joker) {
			t.Error("Full deck should contain all cards!")
		}
	} else {
		if deck.Contains(joker) {
			t.Error("Full deck shouldn't contain extra cards!")
		}
	}
	for {
		if card, err := deck.Take(); err == nil {
			if card == noneCard {
				t.Error("Taken card cannot be none!")
			}
			if card != joker && deck.Contains(card) && maxCount == 1 {
				t.Error("This card has be just taken!")
			}
		} else {
			if deck.HasCards() {
				t.Error("Empty deck should not have cards!")
			}
			break
		}
	}
	for suit := Hearts; suit < none_suit; suit++ {
		for rank := Ace; rank < maxRank; rank++ {
			if card, err := NewCard(suit, rank); err == nil {
				if !deck.Put(card) {
					t.Error("Cannot put card back to deck!")
				}
				for i := 1; i < maxCount; i++ {
					if !deck.Put(card) {
						t.Error("This card should be added several times!")
					}
				}
				if deck.Put(card) {
					t.Error("Cannot put same card back to deck twice!")
				}
				if !deck.HasCards() {
					t.Error("Some cards should be in deck!")
				}
			}
		}
	}
	if maxRank == joker.rank {
		for i := 0; i < maxCount; i++ {
			if !deck.Put(joker) {
				t.Error("Cannot put card back to deck!")
			}
			if !deck.Put(joker) {
				t.Error("Cannot put card back to deck!")
			}
		}
		if deck.Put(joker) {
			t.Error("Cannot put same card back to deck twice!")
		}
	}
}
