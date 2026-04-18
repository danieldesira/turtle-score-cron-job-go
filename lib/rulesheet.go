package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"maps"
	"os"
	"strconv"
)

type Rulesheet struct {
	LevelRewards         map[string]int            `json:"levelRewards"`
	InteractionRewards   map[string]int            `json:"interactionRewards"`
	LevelMaxInteractions map[string]map[string]int `json:"levelMaxInteractions"`
	DurationRewards      struct {
		DurationLimit int `json:"durationLimit"`
		Reward        int `json:"reward"`
	} `json:"durationRewards"`
	Resets struct {
		Max                int `json:"max"`
		RewardPerRemaining int `json:"rewardPerRemaining"`
		RewardForPerfect   int `json:"rewardForPerfect"`
	} `json:"resets"`
}

func LoadRulesheet() *Rulesheet {
	file, err := os.Open("rulesheets/turtle-score-sheet.json")
	if err != nil {
		fmt.Println("Error opening rulesheet...")
		return nil
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading rulesheet...")
		return nil
	}

	var rulesheet Rulesheet
	err = json.Unmarshal(content, &rulesheet)
	if err != nil {
		fmt.Println("Error parsing rulesheet...")
		return nil
	}

	return &rulesheet
}

func GetFinalLevel(rulesheet *Rulesheet) int {
	levelKeys := maps.Keys(rulesheet.LevelRewards)
	levels := []int{}
	for key := range levelKeys {
		val, err := strconv.Atoi(key)
		if err != nil {
			fmt.Printf("Error converting level key %s to int: %v\n", key, err)
			continue
		}
		levels = append(levels, val)
	}
	maxLevel := 0
	for index := range levels {
		if levels[index] > maxLevel {
			maxLevel = levels[index]
		}
	}
	return maxLevel
}
