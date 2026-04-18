# Game Scores Cheat Detection

Cron job in Go which consumes a Redis queue of game
scores and compares them against a rulesheet in a
file. Used for Mission Sea Turtle Nest but can be
adapted to support other games in the future through
simple architecture changes and adding JSON rulesheets
following the existing pattern.

## Starting up the program

In your terminal app, navigate to the project root and
enter: `go run main.go`.

You should see `Starting up scores cheat detection...`
and then you will see output whenever a score is submitted.

## Rule sheet structure

Please check out the following JSON document
as a reference:

```json
{
  "levelRewards": {
    "1": 10,
    "2": 10,
    "3": 50,
    "4": 75,
    "5": 100,
    "6": 100,
    "7": 100,
    "8": 100
  },
  "interactionRewards": {
    "Boat": -50,
    "GhostNest": -20,
    "JaggedPlastic": -3,
    "Nurdle": -1,
    "PlasticBag": -20,
    "Rope": -20,
    "Crab": 20,
    "NeptuneGrass": 1,
    "Sardine": 10,
    "Shrimp": 4,
    "MaleTurtle": 100
  },
  "levelMaxInteractions": {
    "1": {
      "Shrimp": 20,
      "Sardine": 10
    },
    "2": {
      "NeptuneGrass": 4,
      "Shrimp": 30,
      "Crab": 5
    },
    "3": {
      "Shrimp": 250,
      "Sardine": 100,
      "MaleTurtle": 1
    },
    "4": {
      "Sardine": 70,
      "Shrimp": 300
    },
    "5": {
      "Sardine": 30,
      "Shrimp": 300
    },
    "6": {
      "Crab": 1,
      "Shrimp": 100
    },
    "7": {
      "Shrimp": 120
    },
    "8": {
      "Shrimp": 200
    }
  },
  "durationReward": {
    "durationLimit": 300,
    "reward": 300
  },
  "resets": {
    "max": 3,
    "rewardPerRemaining": 50,
    "rewardForPerfect": 200
  }
}
```

`levelRewards`: Points awarding for passing
the respective level. E.g.: 75 points for
level 4.

`levelMaxInteractions`: Maximum number of
interactions for characters that translate to
positive scores. E.g.: Level 8 may not have more
than 200 shrimp interactions.

`durationReward`: Contains information
related to time-based reward systems.

`durationLimit`: Maximum time in seconds
for the reward to be granted.

`reward`: Points awarded following a win
in under the given `durationLimit`.

`resets`: Contains information related to
remaining resets and reset reward system.

`max`: Maximum number of possible resets.

`rewardPerRemaining`: Points rewarded for
each remaining reset.

`rewardForPerfect`: Points rewarded for not
using any resets.
