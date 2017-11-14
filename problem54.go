package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Card struct {
	val  int
	suit string
}

func NewCard(card string) *Card {
	split := strings.Split(card, "")
	val := 0
	if split[0] == "A" {
		val = 14
	} else if split[0] == "K" {
		val = 13
	} else if split[0] == "Q" {
		val = 12
	} else if split[0] == "J" {
		val = 11
	} else if split[0] == "T" {
		val = 10
	} else {
		val, _ = strconv.Atoi(strings.Join(split[:len(split)-1], ""))
	}
	return &Card{val, split[len(split)-1]}
}

type Hand struct {
	cards Cards
}

func NewHand(hand []string) *Hand {
	cards := make([]Card, 5)
	for i := 0; i < 5; i++ {
		cards[i] = *NewCard(hand[i])
	}
	h := &Hand{cards}
	return h
}

type Cards []Card

func (c Cards) Len() int {
	return len(c)
}
func (c Cards) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c Cards) Less(i, j int) bool {
	return c[i].val > c[j].val
}

func (h Cards) Sort() Cards {
	sort.Sort(Cards(h))
	return h
}

const ( // iota is reset to 0
	r_hc            = iota // 0
	r_pair          = iota // 1
	r_two_pairs     = iota // 2
	r_threeofakind  = iota // 3
	r_straight      = iota // 4
	r_flush         = iota // 5
	r_fullhouse     = iota // 6
	r_four          = iota // 7
	r_straightflush = iota // 8
	r_royalflush    = iota // 9
)

func IsStraight(cards Cards) bool {
	cards = cards.Sort()
	prev_c := cards[0].val
	for i := 1; i < len(cards); i++ {
		if prev_c != cards[i].val+1 {
			return false
		}
		prev_c = cards[i].val
	}
	return true
}

func (h Hand) GetRank() ([]Card, int) {
	cards_by_val := make(map[int]Cards, 0)
	// Not num_pairs, it's basically , 1 = one pair, 2 = two pairs or threeofakind.

	suit := ""
	h.cards = h.cards.Sort()
	straight := IsStraight(h.cards)
	for _, c := range h.cards {
		cards_by_val[c.val] = append(cards_by_val[c.val], c)
		if suit != "N" {
			if suit == "" {
				suit = c.suit
			}
			if suit != c.suit {
				suit = "N"
			}
		}
	}

	if straight {
		if suit != "N" {
			return h.cards, r_straightflush
		}
		return h.cards, r_straight
	}
	if suit != "N" {
		return h.cards, r_flush
	}

	num_pairs, num_three := 0, 0
	pairs := make(Cards, 0)
	threes := make(Cards, 0)
	for _, c := range cards_by_val {
		if len(c) == 2 {
			num_pairs++
			pairs = append(pairs, c[0])
		} else if len(c) == 3 {
			num_three++
			threes = c
		} else if len(c) == 4 {
			// Four of a kind return
			return c, r_four
		}
	}
	if num_pairs == 1 && num_three == 1 {
		// Only return the threes here, since they will always differ when comparing.
		return threes, r_fullhouse
	}
	if num_three == 1 {
		return threes, r_threeofakind
	}
	if num_pairs == 2 {
		return pairs, r_two_pairs
	}
	if num_pairs == 1 {
		return pairs, r_pair
	}
	return h.cards, r_hc
}

func P1Win(p1, p2 Hand) bool {
	p1c, p1r := p1.GetRank()
	p2c, p2r := p2.GetRank()
	if p1r > p2r {
		return true
	} else if p1r == p2r {
		// check highest cards in p1c p2c, if they are equal, check highest cards.
		for i, _ := range p1c {
			if p1c[i].val > p2c[i].val {
				return true
			}
			if p1c[i].val < p2c[i].val {
				return false
			}
		}
		for i, _ := range p1.cards {
			if p1.cards[i].val > p2.cards[i].val {
				return true
			}
			if p1.cards[i].val < p2.cards[i].val {
				return false
			}
		}
	}
	return false
}

func main() {
	f, err := ioutil.ReadFile("p054_poker.txt")
	if err != nil {
		panic(err)
	}

	games := strings.Split(string(f), "\n")
	p1w := 0
	for _, game := range games {
		if game != "" {
			hands := strings.Split(game, " ")
			p1 := NewHand(hands[:5])
			p2 := NewHand(hands[5:])
			p1win := P1Win(*p1, *p2)
			if p1win {
				p1w++
			}
		}
	}
	fmt.Println(p1w)

}
