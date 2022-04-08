package pkg

import (
	"github.com/olzhas-b/SearchSystem/tools"
	"sort"
)

type SearchSystemService struct {
	trie *Trie
}

func NewSearchSystem(trie *Trie) *SearchSystemService {
	return &SearchSystemService{
		trie: trie,
	}
}

func (srv *SearchSystemService) InsertRequest(request string) (err error) {
	request = tools.ConvertToLowercase(request)
	return srv.trie.insert(request)
}

func (srv *SearchSystemService) FindTopFiveSuggestion(request string) (result []string, err error) {
	allAcceptableSuggestions, err := srv.trie.search(request)
	sort.Slice(allAcceptableSuggestions, func(i, j int) bool {
		return allAcceptableSuggestions[i].NumberOfSearched > allAcceptableSuggestions[j].NumberOfSearched
	})
	for i := 0; i < tools.MinInt64(5, len(allAcceptableSuggestions)); i++ {
		result = append(result, allAcceptableSuggestions[i].Word)
	}
	return
}
