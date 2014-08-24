// euchredeck implements a deck of cards for Euchre.
// A euchre deck consistes of A-9 of each suit.
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/mohae/shuffler"
)

// Constants for suits.
const (
	_ = iota
	Club
	Diamond
	Heart
	Spade
)

// Constants for face cards and Ace.
const (
	_ = iota + 8
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

func init() {
	// set the seed
	rand.Seed(time.Now().UTC().UnixNano())
}

type decker interface {
	Shuffle() error
	Print()
}

// Deck is a deck of cards whose size is limited by Deck.size.
type Deck struct {
	// size of the deck
	size int
	// Cards are the actual cards in the deck.
	Cards []card
}

// card is a single card within a deck. Cards consist of a suit and a value.
type card struct {
	suit  int
	value int
}

// Deck.Shuffle implements the Fisher-Yates shuffle as designed by Durstenfield
// and popularized by Knuth
func (d *Deck) Shuffle() error {
	for i := 0; i < d.size; i++ {
		r := rand.Intn(i + 1)
		if i != r {
			d.Cards[r], d.Cards[i] = d.Cards[i], d.Cards[r]
		}
	}
	return nil
}

// Deck.Print prints the deck of cards.
func (d *Deck) Print() error {
	str := ""
	for i := 0; i < d.size; i++ {
		if i%5 == 0 {
			fmt.Println(str)
			str = ""
		} else {
			str += ", "
		}
		str += strconv.Itoa(d.Cards[i].suit) + ":" + strconv.Itoa(d.Cards[i].value)
	}
	if str != "" {
		fmt.Println(str)
	}
	return nil
}

// SetiSize sets the decksize based on the number of cards in the deck.
func (d *Deck) SetSize() {
	d.size = len(d.Cards)
}

// euchreDeck is a special form of deck.
type euchreDeck struct {
	Deck
}

// Hand is a set of cards that a player holds; a palyer's hand.
type Hand struct {
	Cards []card
}

// Deal's the hands for Euchre.
func (d *euchreDeck) Deal() []Hand {
	// Euchre has 4 hands. Each hand is dealt in 2 rounds, usually 2 and 3
	// cards in whatever order.
	idx := 0
	h := make([]Hand, 5)
	// Create each hand and randomly deal 2 or 3 cards in the first round of deal
	for i := 0; i < 4; i++ {
		h[i].Cards = make([]card, 2)
		h[i].Cards[0] = d.Deck.Cards[idx]
		h[i].Cards[1] = d.Deck.Cards[idx+1]
		idx += 2
		if rand.Intn(2) == 1 {
			h[i].Cards = append(h[i].Cards, d.Deck.Cards[idx])
			idx++
		}
	}

	// Deal the rest on the second round.
	for i := 0; i < 4; i++ {
		x := len(h[i].Cards)
		h[i].Cards = append(h[i].Cards, d.Deck.Cards[idx])
		h[i].Cards = append(h[i].Cards, d.Deck.Cards[idx+1])
		idx += 2
		if x == 2 {
			h[i].Cards = append(h[i].Cards, d.Deck.Cards[idx])
			idx++
		}
	}

	// The 5th 'hand' is the kitty. The top card, 20th, will be flipped up
	// to determine initial trump offer.
	h[4].Cards = []card{
		d.Deck.Cards[20],
		d.Deck.Cards[21],
		d.Deck.Cards[22],
		d.Deck.Cards[23],
	}

	return h
}

// convert a slice of cards to a slice of interfaces
func (d *euchreDeck) cardsToIface() []interface{} {
	ret := make([]interface{}, len(d.Cards))

	for i, v := range d.Cards {
		ret[i] = interface{}(v)
	}

	return ret
}

// convert a slice of interface to a slice of cards
func (d *euchreDeck) ifaceToCards(iface []interface{}) {
	for i, v := range iface {
		d.Cards[i] = v.(card)
	}
}
// newEuchreDeck initializes a euchre deck and returns it. This deck is not
// shuffled.
func newEuchreDeck() euchreDeck {
	return euchreDeck{
		Deck: Deck{
			size: 24,
			Cards: []card{
				{Club, Nine},
				{Club, Ten},
				{Club, Jack},
				{Club, Queen},
				{Club, King},
				{Club, Ace},
				{Diamond, Nine},
				{Diamond, Ten},
				{Diamond, Jack},
				{Diamond, Queen},
				{Diamond, King},
				{Diamond, Ace},
				{Heart, Nine},
				{Heart, Ten},
				{Heart, Jack},
				{Heart, Queen},
				{Heart, King},
				{Heart, Ace},
				{Spade, Nine},
				{Spade, Ten},
				{Spade, Jack},
				{Spade, Queen},
				{Spade, King},
				{Spade, Ace},
			},
		},
	}
}

// Main gets a new euchre deck, shuffles it, and deals the cards.
// The resulting hands will be printed out.
func main() {
	fmt.Println("Hello, here are your hands, after shuffling:")
	d := newEuchreDeck()
//	d.Shuffle()
	iCards := d.cardsToIface()
	shuffled := shuffler.FisherYates((iCards))
	d.ifaceToCards(shuffled)
	h := []Hand{}
	h = d.Deal()
	for i := 0; i < len(h); i++ {
		fmt.Printf("%v\n", h[i])
	}
}
