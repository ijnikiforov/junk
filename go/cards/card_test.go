package cards

import "testing"

func TestNewJoker(t *testing.T) {
	joker := NewJoker()
	if joker.rank != Joker {
		t.Error("Not a Joker!")
	}
	if joker.suit != none_suit {
		t.Error("Joker should not have suit!")
	}
	if joker.Rank() != Joker {
		t.Error("Not an Joker!")
	}
	if _, err := joker.Suit(); err == nil {
		t.Error("Joker should not have suit!")
	}
	if joker.Color() != None {
		t.Error("Jokes has no color!")
	}
}

func TestNewCard(t *testing.T) {
	if _, err := NewCard(17, Two); err == nil {
		t.Error("Invalid suit!")
	}
	if _, err := NewCard(Diamonds, 77); err == nil {
		t.Error("Invalid rank!")
	}
	if _, err := NewCard(Clubs, Joker); err == nil {
		t.Error("Should not create joker!")
	}
	for suit := Hearts; suit < none_suit; suit++ {
		for rank := Ace; rank < none_rank; rank++ {
			if card, err := NewCard(suit, rank); err == nil {
				if card.rank != rank {
					t.Error("Invalid rank!")
				}
				if card.suit != suit {
					t.Error("Invalid suit!")
				}
				if card.Rank() != rank {
					t.Error("Invalid rank!")
				}
				if s, err := card.Suit(); err == nil {
					if s != suit {
						t.Error("Invalid suit!")
					}
				} else {
					t.Error("Cannot get suit!")
				}
				if card.Color() == None {
					t.Error("Color cannot be None!")
				}
			} else {
				t.Error("Cannot create valid card!")
			}
		}
	}
}