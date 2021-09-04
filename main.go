package main

// import "fmt"

func main() {
	//variable declaration
	//[type 1]
	// var card string = "Ace of Spade"

	//[type 2]
	// card1 := "Ace"

	// fmt.Println(card)
	// fmt.Println("\n", card1)

	//[]string => [deck] files in same package can freely call another file's function
	// cards := newDeck()

	// //deal(d deck, handSize int) ==> so d=cards , handSize=5
	// //will return 2 value assigning to hand, remainingCard
	// hand, remainingCard := deal(cards, 5)

	// // cards.Print()
	// hand.Print()
	// remainingCard.Print()

	// //to see the coverted value
	// fmt.Println(cards.toString())

	// cards.saveToFile("Deck_cards")

	//now any var of type [deck] gets access to print() ==> so here [deck] assigned to [cards]
	// cards.Print()

	// for i, eachCard := range cards {
	// 	fmt.Println(i, eachCard)
	// }

	//calling file from harddrive
	cards := newDeckFromFile("Deck_cards")
	// cards.Print()
	cards.shuffle()
	cards.Print()
}
