package gologger

import (
	"fmt"
)

var LevelMap = map[string]string{
	"SUCCESS": "\033[0;92m",
	"ERROR":   "\033[0;91m",
	"WARNING": "\033[0;93m",
	"INFO":    "\033[0;94m",
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
