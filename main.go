package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/charmbracelet/log"
)

var (
	_config string
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
		LocCode       []string `json:"location_code"`
		Pools         int      `json:"pools"`
		Voters        int      `json:"voters"`
		Reset         bool     `json:"reset"`
		VotersPerPool int      `json:"voters_per_pool"`
		Summon        bool     `json:"summon"`
		AddResponses  bool     `json:"add_responses"`
		DaysToAdd     int      `json:"days_to_add"`
		Ranges        Ranges   `json:"ranges"`
		Trials        Trials   `json:"trials"`
	}
)

func main() {
	flag.StringVar(&_config, "c", "", "Use a config file")
	flag.Parse()

	_ = log.New(os.Stderr)
	log.SetLevel(log.DebugLevel)

	if _config == "" {
		log.Error("No config file provided")
		return
	}

	c, err := os.ReadFile(_config)
	if err != nil {
		log.Error("Could not read config file")
		return
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

	if len(config.LocCode) > 1 {
		log.Error("Only one location code is allowed for now")
		return
	}

	if len(config.LocCode[0]) != 3 {
		log.Error("Location code must be 3 characters")
		return
	}

	db := New(config.Reset)

	if config.Voters > 0 {
		v := voters(db.db, config.LocCode[0], config.Voters)
		v.insert()
	}

	if config.Pools > 0 {
		p := pools(db.db, &config)
		p.request()
	}

	if config.Summon {
		s := summon(db.db, &config)
		s.summon()
	}

	// always try to create rooms and judges
	t := trials(db.db, &config)
	t.rooms()
	t.judges()
	t.createTrials()
}
