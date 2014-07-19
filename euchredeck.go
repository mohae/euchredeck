package main

import (
	"fmt"	
	"math/rand"		
	"strconv"
	"time"
)

func init () {
	// set the seed
	rand.Seed(time.Now().UTC().UnixNano())
}

type decker interface {
	Shuffle() error
	Print()
}

type Deck struct {
	// size of the deck
	size int 
	// Cards are the actual cards in the deck.
	Cards []card
}

type card struct {
	suit int
	value int
}

// Implements Fisher-Yates shuffle as designed by Durstenfield 
// and popularized by Knuth
func (d *Deck) Shuffle() error {
	for i := 0; i < d.size; i++ {
		r := rand.Intn(i+1)
		if i != r {
			d.Cards[r], d.Cards[i] = d.Cards[i], d.Cards[r]
		}
	}	
	return nil
}
	
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

// Set's the decksize based on the number of cards in the deck.
func (d *Deck) SetSize() {
	d.size = len(d.Cards)
}

type euchreDeck struct {
	Deck
}

type Hand struct {
	Cards []card
}

// Deal's the hands for Euchre.
func (d *euchreDeck) Deal() []Hand {
	// Euchre has 4 hands. Each hand is dealt in 2 rounds, usually 2 and 3
	// cards in whatever order.
	idx := 0
	h := make([]Hand,4)
	// Create each hand and randomly deal 2 or 3 cards in the first round of deal
	for i := 0; i < 4; i++ {
		h[i].Cards = make([]card,2)
		h[i].Cards[0] = d.Deck.Cards[idx];
		h[i].Cards[1] = d.Deck.Cards[idx + 1];
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

	return h
}

func newEuchreDeck() euchreDeck {
	return euchreDeck{Deck: Deck{Cards: []card{
		{1, 9},
		{1, 10},
		{1, 11},
		{1, 12},
		{1, 13},
		{1, 14},				
		{2, 9},
		{2, 10},
		{2, 11},
		{2, 12},
		{2, 13},
		{2, 14},				
		{3, 9},
		{3, 10},
		{3, 11},
		{3, 12},
		{3, 13},
		{3, 14},				
		{4, 9},
		{4, 10},
		{4, 11},
		{4, 12},
		{4, 13},
		{4, 14},				
	}}}
}
func main() {
	fmt.Println("Hello, here are your hands, after shuffling:")
	d := newEuchreDeck()
	d.SetSize()
	d.Shuffle()
	h := []Hand{}
	h = d.Deal()
	fmt.Printf("%v", h)
	d.Print()
	
}
