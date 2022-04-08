package pkg

type ISearchSystemService interface {
	InsertRequest(word string) (err error)
	FindTopFiveSuggestion(request string) (result []string, err error)
}
