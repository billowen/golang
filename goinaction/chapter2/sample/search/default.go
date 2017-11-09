package search

type defaultMatcher struct{}

func init() {
	var matcher defaultMatcher
	Reister("default", matcher)
}

func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
