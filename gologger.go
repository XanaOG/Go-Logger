package gologger

import (
	"fmt"
)

var LevelMap = map[string]string{
	"SUCCESS": "\033[1;92m",
	"ERROR":   "\033[1;91m",
	"WARNING": "\033[1;93m",
	"INFO":    "\033[1;94m",
}

func AddColor(Name, ColorCode string) {
	LevelMap[Name] = ColorCode
}

func Display(Tag, Content string) {
	ColorCode, exists := LevelMap[Tag]
	if !exists {
		fmt.Println("Color code does not exist - " + Content)
		return
	}
	fmt.Printf("[%s%s\033[0m] %s\n", ColorCode, Tag, Content)
}
