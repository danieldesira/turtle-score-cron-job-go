package lib

import (
	"encoding/json"
	"strconv"
	"strings"
)

type RawScore struct {
	Interactions    string `json:"interactions"`
	Duration        int    `json:"duration"`
	Level           int    `json:"level"`
	PlayerID        int    `json:"playerId"`
	Timestamp       string `json:"timestamp"`
	RemainingResets int    `json:"remainingResets"`
}

func ParseRawScore(entry string) (*RawScore, error) {
	var score RawScore
	err := json.Unmarshal([]byte(entry), &score)
	if err != nil {
		return nil, err
	}
	return &score, nil
}

type ProcessedScore struct {
	Interactions map[string]int
	Duration     int
	Level        int
	PlayerID     int
	Timestamp    string
	ResetsUsed   int
	TotalScore   int
}

func ProcessScore(rawScore *RawScore, rulesheet *Rulesheet) *ProcessedScore {
	interactions := parseInteractions(rawScore.Interactions)
	resetsUsed := rulesheet.Resets.Max - rawScore.RemainingResets
	processedScore := &ProcessedScore{
		Interactions: interactions,
		Duration:     rawScore.Duration,
		Level:        rawScore.Level,
		PlayerID:     rawScore.PlayerID,
		Timestamp:    rawScore.Timestamp,
		ResetsUsed:   resetsUsed,
	}
	//processedScore.TotalScore = calculateTotalScore(processedScore, rulesheet)
	return processedScore
}

func parseInteractions(interactionsStr string) map[string]int {
	interactions := make(map[string]int)
	entries := strings.SplitSeq(interactionsStr, "|")
	for entry := range entries {
		parts := strings.Split(entry, ",")
		if len(parts) == 2 {
			val, err := strconv.Atoi(parts[1])
			if err == nil {
				interactions[parts[0]] = val
			}
		}
	}
	return interactions
}

// func validateInteractions(interactions map[string]int, rulesheet *Rulesheet) bool {
// 	for interaction, count := range interactions {
// 		if rulesheet.InteractionRewards[interaction] > 0 {
// 			maxAllowedOcurrances := 0
// 			for level := 1; level <= GetFinalLevel(rulesheet); level++ {

// 			}
// 		}
// 	}
// }
