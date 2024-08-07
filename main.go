package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/charmbracelet/log"
)

var (
	_config        string
	waitForSummons sync.WaitGroup
	waitForVoters  sync.WaitGroup
)

type (
	Ranges struct {
		Responded     float32 `json:"responded"`
		Excused       float32 `json:"excused"`
		Deferred      float32 `json:"deferred"`
		Disqualified  float32 `json:"disqualified"`
		Undeliverable float32 `json:"undeliverable"`
	}

	Trials struct {
		Rooms  int `json:"rooms"`
		Judges int `json:"judges"`
		Total  int `json:"total"`
	}

	Config struct {
		LocCode       []string          `json:"location_code"`
		PostCodes     map[string]string `json:"postcodes"`
		Pools         int               `json:"pools"`
		Voters        int               `json:"voters"`
		Reset         bool              `json:"reset"`
		VotersPerPool int               `json:"voters_per_pool"`
		Summon        bool              `json:"summon"`
		AddResponses  bool              `json:"add_responses"`
		DaysToAdd     int               `json:"days_to_add"`
		Ranges        Ranges            `json:"ranges"`
		Trials        Trials            `json:"trials"`
	}
)

func main() {
	start := time.Now()

	flag.StringVar(&_config, "c", "", "Use a config file")
	flag.Parse()

	_ = log.New(os.Stderr)
	log.SetLevel(log.DebugLevel)

	if _config == "" {
		log.Error("No config file provided")
		os.Exit(1)
	}

	c, err := os.ReadFile(_config)
	if err != nil {
		log.Error("Could not read config file")
		os.Exit(1)
	}

	log.Info("Starting the data gen")

	// TODO: check if response ranges total is 100%

	// do health check
	// if health check fails, exit
	if _, err = request("GET", healthCheckUrl.String(), nil, false); err != nil {
		log.Errorf("Errored out on health check: %s", err.Error())
		os.Exit(1)
	}

	config := Config{}
	json.Unmarshal(c, &config)

	for _, locCode := range config.LocCode {
		if len(locCode) != 3 {
			log.Error("All location codes must be 3 characters")
			os.Exit(1)
		}
	}

	db := New(config.Reset)

	if config.Voters > 0 {
		for _, locCode := range config.LocCode {
			waitForVoters.Add(1)
			v := voters(db.db, locCode, config.Voters, config.PostCodes[locCode])
			go v.insert()
		}

		waitForVoters.Wait()
	}

	if config.Pools > 0 {
		p := pools(db.db, &config)
		p.request()
	}

	if config.Summon {
		for _, locCode := range config.LocCode {
			waitForSummons.Add(1)
			s := summon(db.db, &config)
			go s.summon(locCode)
		}

		waitForSummons.Wait()
	}

	// always try to create rooms and judges
	t := trials(db.db, &config)
	t.rooms()
	t.judges()
	t.createTrials()

	fmt.Printf("\nFinished in: %s\n", time.Since(start))
}
