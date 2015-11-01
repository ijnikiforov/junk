package cards

import (
	"errors"
	"math/rand"
)

type Deck interface {
	Take() (Card, error)
	Put(card Card) bool
	Contains(card Card) bool
	HasCards() bool
}

type simpleDeck struct {
	size uint8
	cards map[Card]int
	maxRank Rank
}

type DeckConstructor func() Deck

func New36CardDeck() Deck {
	deck := &simpleDeck{36, make(map[Card]int), Six}
	for suit := Hearts; suit < none_suit; suit++ {
		for rank := Ace; rank < Five; rank++ {
			deck.cards[Card{suit, rank}] = 1
		}
	}
	return deck
}

func New52CardDeck() Deck {
	deck := &simpleDeck{52, make(map[Card]int), Two}
	for suit := Hearts; suit < none_suit; suit++ {
		for rank := Ace; rank < none_rank; rank++ {
			deck.cards[Card{suit, rank}] = 1
		}
	}
	return deck
}

func New54CardDeck() Deck {
	deck := &simpleDeck{54, make(map[Card]int), Joker}
	for suit := Hearts; suit < none_suit; suit++ {
		for rank := Ace; rank < none_rank; rank++ {
			deck.cards[Card{suit, rank}] = 1
		}
	}
	deck.cards[NewJoker()] = 2
	return deck
}

func (deck *simpleDeck) Take() (Card, error) {
	if !deck.HasCards() {
		return noneCard, errors.New("The deck is empty!")
	}
	result := noneCard
	for result == noneCard {
		for key, num := range deck.cards {
			if rand.Intn(2) == 0 {
				result = key
				if num == 2 {
					deck.cards[key] = 1
				} else {
					delete(deck.cards, key)
				}
				break
			}
		}
	}
	return result, nil
}

func (deck *simpleDeck) Put(card Card) bool {
	if (card.rank > deck.maxRank) {
		return false
	}
	_, exists := deck.cards[card]
	if exists {
		if card.rank == Joker && deck.cards[card] == 1 {
			deck.cards[card] = 2
			return true
		}
		return false
	} else {
		deck.cards[card] = 1
		return true
	}
}

func (deck *simpleDeck) Contains(card Card) bool {
	_, result := deck.cards[card]
	return result
}

func (deck *simpleDeck) HasCards() bool {
	return len(deck.cards) > 0
}

type compositeDeck struct {
	decks []simpleDeck
}

func NewCustomDeck(decks... Deck) (Deck, error) {
	size := uint8(0)
	max := none_rank
	result := &compositeDeck{}
	result.decks = make([]simpleDeck, 0, len(decks))
	for _, deck := range decks {
		if sdeck, simple := deck.(*simpleDeck); simple {
			if size == 0 {
				size = sdeck.size
			}
			if max == none_rank {
				max = sdeck.maxRank
			}
			if size != sdeck.size || max != sdeck.maxRank {
				return &compositeDeck{}, errors.New("Custom deck should contain from equal decks!")
			}
			result.decks = append(result.decks, *sdeck)
		}
		if _, composite := deck.(*compositeDeck); composite {
			return &compositeDeck{}, errors.New("Composite decks are not allowed here!")
		}
	}
	return result, nil
}

func (deck *compositeDeck) Take() (Card, error) {
	if !deck.HasCards() {
		return noneCard, errors.New("The deck is empty!")
	}
	result := noneCard
	for result == noneCard {
		for _, d := range deck.decks {
			if rand.Intn(2) == 0 {
				if card, err := d.Take(); err == nil {
					result = card
					break
				}
			}
		}
	}
	return result, nil
}

func (deck *compositeDeck) Put(card Card) bool {
	for _, d := range deck.decks {
		if d.Put(card) {
			return true
		}
	}
	return false
}

func (deck *compositeDeck) Contains(card Card) bool {
	for _, d := range deck.decks {
		if d.Contains(card) {
			return true
		}
	}
	return false
}

func (deck *compositeDeck) HasCards() bool {
	for _, d := range deck.decks {
		if d.HasCards() {
			return true
		}
	}
	return false
}