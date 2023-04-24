package main

import (
	"fmt"
	"time"
)

type User interface {
	Player | Refree
}
type Player struct {
	Name string
	Play func()
}

type Refree struct {
	Name string
}

func Swim() {
	fmt.Println("Start Play swim!!")
	time.Sleep(1 * time.Second)
	fmt.Println("Finished Play swim!!")
}

func Dance() {
	fmt.Println("Start Play dance!!")
	time.Sleep(1 * time.Second)
	fmt.Println("Finished Play dance!!")
}

func Game() {
	fmt.Println("Start Play game!!")
	time.Sleep(1 * time.Second)
	fmt.Println("Finished Play game!!")
}

func main() {

	/* make channel*/
	chPlayer := make(chan Player)
	chRefree := make(chan Refree)

	/* receive player from channel */
	go PlayChannels(chPlayer, chRefree)

	/* make player */
	player1 := Player{"player1", Swim}
	player2 := Player{"player2", Dance}
	player3 := Player{"player3", Game}

	/* send player to channel */
	chPlayer <- player1
	chPlayer <- player2
	chPlayer <- player3

	/* make refree */
	refree1 := Refree{"refree1"}
	refree2 := Refree{"refree2"}

	chRefree <- refree1
	chRefree <- refree2
}

/* Channel에 들어온 데이터 처리 함수 */
/* Goroutine 으로 비동기 처리 */
/* Case 순서대로 FIFO (first In first Out 처리) */
func PlayChannels(chPlayer chan Player, chRefree chan Refree) {
	for {
		select {
		/* chPlayer에 들어온 데이터 처리 */
		case player := <-chPlayer:
			player.Play()

			/* chRefree에 들어온 데이터 처리 */
		case refree := <-chRefree:
			fmt.Println(refree.Name + " is watching!!")

		}
	}
}
