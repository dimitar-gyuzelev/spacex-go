package v4

const PathHistory = "/history"

// HistoryAPI defines functions for working with the SpaceX History V4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/tree/master/docs/history/v4
type HistoryAPI interface {
	GetHistoryAll() ([]HistEvent, error)

	GetHistoryByID(id string) (HistEvent, error)

	PostHistoryQuery(query Query) (RespQuery[HistEvent], error)
}

// HistEvent represents the data returned by the Space X History V4 API.
// Docs: https://github.com/r-spacex/SpaceX-API/blob/master/docs/history/v4/schema.md#history-event-schema
type HistEvent struct {
	ID            string        `json:"id"`
	Title         string        `json:"title"`
	EventDateUTC  string        `json:"event_date_utc"`
	Details       string        `json:"details"`
	Links         HistEventLink `json:"links"`
	EventDateUnix int           `json:"event_date_unix"`
}

func (h HistEvent) String() string {
	return stringFn[HistEvent](h)
}

type HistEventLink struct {
	Article string `json:"article"`
}
