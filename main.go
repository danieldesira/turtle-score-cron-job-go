package main

import (
	"context"
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

	for newScore := redisClient.RPop(context.Background(), "scoreQueue"); newScore.Val() != ""; newScore = redisClient.RPop(context.Background(), "scoreQueue") {
		fmt.Println("Processing new score:", newScore)
	}

}
