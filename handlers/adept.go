package handlers

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ConvertAdept()  {
	// parse file line by line and build object model
	file, err := os.Open("../data/adept.dat")
	if err != nil {
		log.Print(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var currentObj = make(map[string]interface{})
	var currentLvl = 0
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "!") { // comment
			continue
		}
		fields := strings.Split(text, "|")

	}
	if err := scanner.Err(); err != nil {
		log.Print(err)
		return
	}
}

