package cards

import "errors"

type Suit uint8

const (
	Hearts Suit = iota
	Diamonds
	Clubs
	Spades
	none_suit
)

type Rank uint8

const (
	Ace Rank = iota
	King
	Queen
	Jack
	Ten
	Nine
	Eight
	Seven
	Six
	Five
	Four
	Three
	Two
	none_rank
)

const Joker Rank = 255

type Color uint8

const (
	Red Color = iota
	Black
	None
)

type Card struct {
	suit Suit
	rank Rank
}

var noneCard = Card{ none_suit, none_rank}

func NewJoker() Card {
	return Card {none_suit, Joker}
}

func NewCard(suit Suit, rank Rank) (Card, error) {
	if suit >= none_suit {
		return Card {none_suit, none_rank}, errors.New("Invalid suit!")
	}
	if rank >= none_rank {
		return Card {none_suit, none_rank}, errors.New("Invalid rank!")
	}
	if rank == Joker {
		return Card {none_suit, none_rank}, errors.New("Use NewJoker() to create ace!")
	}
	return Card {suit, rank}, nil
}

func (card Card) Suit() (Suit, error) {
	if card.suit == none_suit {
		return none_suit, errors.New("Jockers do not have suit!")
	}
	return card.suit, nil
}

func (card Card) Rank() Rank {
	return card.rank
}

func (card Card) Color() Color {
	if card.rank == Joker {
		return None
	}
	if card.suit == Diamonds || card.suit == Hearts {
		return Red
	}
	return Black
}



