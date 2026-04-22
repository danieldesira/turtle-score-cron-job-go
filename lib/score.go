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
	Outcome      string
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
		Outcome:      deduceOutcome(rawScore, GetFinalLevel(rulesheet)),
	}
	if validateInteractions(rawScore.Level, interactions, rulesheet) {
		processedScore.TotalScore = calculateTotalScore(processedScore, rulesheet)
		return processedScore
	} else {
		return nil
	}
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

func deduceOutcome(score *RawScore, finalLevel int) string {
	if score.Level > finalLevel {
		return "WIN"
	} else {
		return "LOSS"
	}
}

func validateInteractions(scoreLevel int, interactions map[string]int, rulesheet *Rulesheet) bool {
	for interaction, count := range interactions {
		if rulesheet.InteractionRewards[interaction] > 0 {
			maxAllowedOcurrances := 0
			for level := 1; level <= scoreLevel; level++ {
				maxAllowedOcurrances += rulesheet.LevelMaxInteractions[strconv.Itoa(level)][interaction]
			}
			if count > maxAllowedOcurrances {
				return false
			}
		}
	}
	return true
}

func calculateTotalScore(score *ProcessedScore, rulesheet *Rulesheet) int {
	return getDurationRewardIfApplicable(score, rulesheet) + getLevelPassRewards(score, rulesheet) + getInteractionRewards(score, rulesheet) + getResetsRewards(score, rulesheet)
}

func getDurationRewardIfApplicable(score *ProcessedScore, rulesheet *Rulesheet) int {
	if score.Duration <= rulesheet.DurationRewards.DurationLimit && score.Outcome == "WIN" {
		return rulesheet.DurationRewards.Reward
	} else {
		return 0
	}
}

func getLevelPassRewards(score *ProcessedScore, rulesheet *Rulesheet) int {
	points := 0
	for level := 1; level < score.Level; level++ {
		points += rulesheet.LevelRewards[strconv.Itoa(level)]
	}
	return points
}

func getInteractionRewards(score *ProcessedScore, rulesheet *Rulesheet) int {
	points := 0
	for key, count := range score.Interactions {
		points += rulesheet.InteractionRewards[key] * count
	}
	return points
}

func getResetsRewards(score *ProcessedScore, rulesheet *Rulesheet) int {
	points := 0
	if score.ResetsUsed == 0 {
		points += rulesheet.Resets.RewardForPerfect
	}
	points += score.ResetsUsed * rulesheet.Resets.RewardPerRemaining
	return points
}
