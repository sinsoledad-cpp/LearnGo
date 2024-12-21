package main

import "fmt"

var (
	coins = 5000
	users = []string{
		"jfoiasjofi", "jsoijfoisaj", "jsoifjoisa", "jfaoisjfoi", "jfiosjfohjsinv",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() (left int) {
	for _, name := range users {
		for _, c := range name {
			switch c {
			case 'e', 'E':
				distribution[name]++
				coins--
			case 'i', 'I':
				distribution[name] += 2
				coins -= 2
			case 'o', 'O':
				distribution[name] += 3
				coins -= 3
			case 'u', 'U':
				distribution[name] += 4
				coins -= 4
			}
			
		}
	}
	left = coins
	return
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	for k, v := range distribution {
		fmt.Printf("%s:%d\n", k, v)
	}
}
