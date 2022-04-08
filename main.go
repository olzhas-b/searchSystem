package main

import (
	"fmt"
	"github.com/olzhas-b/SearchSystem/pkg"
	"log"
)

func main() {
	defaultTrie := pkg.NewTrie()
	searchEngine := pkg.NewSearchSystem(defaultTrie)
	dictionary := []string{
		"Vk", "Vk search", "vk search", "what is that",
		"vk", "vk search", "vk what's up", "vk why is so hard",
		"vk emotional damage",
		"vk why", "vk why", "vk why",
	}

	// feed our searchEngine by dictionary
	for _, word := range dictionary {
		if err := searchEngine.InsertRequest(word); err != nil {
			log.Printf("got unexpected need to fix error: %v\n", err)
		}
	}

	// get top 5 suggestions by how many people searched it before
	suggestions, err := searchEngine.FindTopFiveSuggestion("vk")
	if err != nil {
		log.Fatalf("something wrong got error: %v", err)
	}

	// just show suggestions list on the cmd
	for i, suggestion := range suggestions {
		fmt.Printf("%d)%s\n", i+1, suggestion)
	}
}
