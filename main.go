package main

import (
	"bufio"
	"log"
	"os"
)

func main1() {
	log.Printf("test.....")
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "." {
			break
		}
		//counts[input.Text()]++
		line := input.Text()
		log.Printf("counts[line] : %d", counts[line])
		counts[line] = counts[line] + 1
		log.Printf("Loop Map is %v ", counts)

	}
	log.Printf("Map is %v ", counts)
	log.Printf("Length is %d ", len(counts))
	for line, n := range counts {
		if n > 1 {
			log.Printf("%d\t%s\n", n, line)
		}
	}
}
