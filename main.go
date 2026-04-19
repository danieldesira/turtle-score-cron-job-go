package main

import (
	"fmt"

	"github.com/danieldesira/turtle-score-cron-job-go/lib"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	fmt.Println("Starting up scores cheat detection...")

	redisClient, err := lib.ConnectRedis()
	rulesheet := lib.LoadRulesheet()
	fmt.Println("Loaded rulesheet", rulesheet)
	finalLevel := lib.GetFinalLevel(rulesheet)
	fmt.Println("Final level is", finalLevel)

	for entry := lib.GetNextScoreEntry(redisClient); entry != ""; entry = lib.GetNextScoreEntry(redisClient) {
		score, err := lib.ParseRawScore(entry)
		if err != nil {
			fmt.Println("Failed to parse raw score from entry:", entry, "Error:", err)
			continue
		}
		fmt.Println("Processing new score:", score)
		processedScore := lib.ProcessScore(score, rulesheet)
		fmt.Println("Processed score:", processedScore)
	}

}
