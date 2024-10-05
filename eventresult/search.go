package eventresult

import "github.com/ysmilda/speedhive-go"

type searchService struct {
	c *speedhive.Client
}

type SearchFreesearchOptions struct {
	Sport         *speedhive.SportValue         `url:"sport,omitempty"`
	SportCategory *speedhive.SportCategoryValue `url:"sportCategory,omitempty"`
	Country       *speedhive.CountryValue       `url:"country,omitempty"`
	Filter        *SearchFilterValue            `url:"filter,omitempty"`
	SearchMode    *SearchModeValue              `url:"searchMode,omitempty"`
	SortByDate    *bool                         `url:"sortByDate,omitempty"`
	Type          *SearchTypeValue              `url:"type,omitempty"`
	FuzzySearch   *bool                         `url:"fuzzySearch,omitempty"`
	Count         *int                          `url:"count,omitempty"`
	Offset        *int                          `url:"offset,omitempty"`
}

// Freesearch searches for events based on a free search query.
// The search can be filtered by passing in the SearchFreesearchOptions.
func (s searchService) Freesearch(query string, opt *SearchFreesearchOptions) ([]Search, error) {
	u := "/api/v0.2.3/search/freesearch"
	res, err := speedhive.Get[SearchResults](s.c, u, opt)
	if err != nil {
		return nil, err
	}
	return res.Results, nil
}
