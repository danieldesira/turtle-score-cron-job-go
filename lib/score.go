package lib

type Score struct {
	Interactions    map[string]int
	Duration        int
	Level           int
	Outcome         string
	PlayerID        int
	Timestamp       int64
	RemainingResets int
}
