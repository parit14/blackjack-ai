package blackjack

import (
	"deck"
	"fmt"
)

type State int8

const (
	statePlayerTurn State = iota
	stateDealerTurn
	stateHandOver
)

type Game struct {
	// unexported fields
	deck   []deck.Card
	state  State
	player []deck.Card
	dealer []deck.Card
}

func Deal(gs Game) Game {
	gs.player = make([]deck.Card, 0, 5)
	gs.dealer = make([]deck.Card, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, gs.deck = draw(gs.deck)
		gs.player = append(gs.player, card)
		card, gs.deck = draw(gs.deck)
		gs.dealer = append(gs.dealer, card)
	}
	gs.state = statePlayerTurn
	return gs
}

func (g *Game) Play(ai AI) {
	g.deck = deck.New(deck.Deck(3), deck.Shuffle)
	for i := 0; i < 10; i++ {
		gs = Deal(gs)
		var input string
		for gs.State == statePlayerTurn {
			fmt.Println("Player:", gs.Player)
			fmt.Println("Dealer:", gs.Dealer.DealerString())
			fmt.Println("What will you do? (h)it, (s)tand")
			fmt.Scanf("%s\n", &input)
			switch input {
			case "h":
				gs = Hit(gs)
				// fmt.Println(gs.Player)
			case "s":
				gs = Stand(gs)
			default:
				fmt.Println("Invalid Option:", input)
			}
		}
	}
	//If dealer score <= 16, we hit
	//If dealer has a soft 17, then we hit.
	for gs.State == stateDealerTurn {
		if gs.Dealer.Score() <= 16 || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() != 17) {
			gs = Hit(gs)
		} else {
			gs = Stand(gs)
		}
	}
	gs = EndHand(gs)

}

type Move func(*Game)

func MoveHit(gs *Game) {
	return gs
}

func MoveStand(gs *Game) {
	return gs
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}
