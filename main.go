package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/AlessandroPomponio/go-gibberish/consts"
	"github.com/AlessandroPomponio/go-gibberish/gibberish"
	"github.com/AlessandroPomponio/go-gibberish/persistence"
	"github.com/AlessandroPomponio/go-gibberish/training"
)

var (
	performTraining bool
)

func main() {

	flag.BoolVar(&performTraining, "train", false, "train")
	flag.Parse()

	if performTraining {
		training.TrainModel(consts.AcceptedCharacters, "big.txt", "good.txt", "bad.txt", "knowledge.json")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	data := persistence.LoadKnowledgeBase("knowledge.json")
	for {

		fmt.Print("Insert something to check: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		isGibberish := gibberish.IsGibberish(input, data)
		fmt.Println(fmt.Sprintf("Input: %s: is gibberish? %v\n", input, isGibberish))

	}

}