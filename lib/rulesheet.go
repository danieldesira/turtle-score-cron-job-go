package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
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
