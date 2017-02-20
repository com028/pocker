// gotst
package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

const (
	CARD_NUM   = 7
	TARGET_NUM = 5
)

func Combine(cards []int, index int) {
	rst := ""
	for i := 0; i < CARD_NUM; i++ {
		if (index>>uint32(i))&1 == 1 {
			rst += strconv.Itoa(cards[i])
		}
	}
	if len(rst) == TARGET_NUM {
		fmt.Println(rst)
	}
}

func CalcCombine(cards []int) {
	for i := 31; i < 1<<CARD_NUM; i++ {
		Combine(cards, i)
	}
}

func CalcSecond(cards []int) {
	four, three, two, tmp := 0, 0, 0, 1
	for i := 0; i < len(cards)-1; i++ {
		if cards[i] == cards[i+1] {
			tmp += 1
			if i+2 != len(cards) {
				continue
			}
		}
		fmt.Println("tmp:", tmp)
		switch tmp {
		case 2:
			two += 1
			break
		case 3:
			three = 1
			break
		case 4:
			four = 1
			break
		}
		tmp = 1
	}
	fmt.Println("four:, three :, two:", four, three, two)
}

func Shuffle(cards []int) {
	end := len(cards) - 2
	for i := end; i > 0; i-- {
		index := rand.Intn(i)
		fmt.Println("index:", index)
		tmp := cards[index]
		cards[index] = cards[i]
		cards[i] = tmp
	}
	fmt.Println("%v", cards)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	cards := []int{2, 2, 0, 1, 1}
	//	CalcCombine(cards)
	CalcSecond(cards)
	//	a := 10
	b := int64(10)
	fmt.Println(reflect.TypeOf(b))

	go Shuffle(cards)
	time.Sleep(1000 * time.Millisecond)

	fmt.Println("rand:", rand.Intn(10))
}
