package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
)

type Trial struct {
	db     *sql.DB
	config *Config
}

func trials(db *sql.DB, config *Config) *Trial {
	return &Trial{
		db:     db,
		config: config,
	}
}

func (t *Trial) rooms() {
	for i := 0; i < t.config.Trials.Rooms; i++ {
		_, err := t.db.Exec("INSERT INTO juror_mod.courtroom (id, loc_code, room_number, description) VALUES ($1, $2, $3, $4)",
			i+1, t.config.LocCode[0], i+1, "Room "+fmt.Sprint(i+1))

		if err != nil {
			log.Errorf("Error inserting courtroom: %s", err.Error())
			os.Exit(1)
		}
	}
}

func (t *Trial) judges() {
	for i := 0; i < t.config.Trials.Judges; i++ {
		_, err := t.db.Exec("INSERT INTO juror_mod.judge (id, owner, code, description, is_active) VALUES ($1, $2, $3, $4, $5)",
			i+1, t.config.LocCode[0], i+1, "Judge "+fmt.Sprint(i+1), true)

		if err != nil {
			log.Errorf("Error inserting judge: %s", err.Error())
			os.Exit(1)
		}
	}
}

func (t *Trial) createTrials() {
	for i := 0; i < t.config.Trials.Total; i++ {
		trialNumber := fmt.Sprintf("T%d", i+1)

		_, err := t.db.Exec("INSERT INTO juror_mod.trial (trial_number, loc_code, description, courtroom, judge, trial_type, trial_start_date, juror_requested) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
			trialNumber, t.config.LocCode[0], "et el", 1, 1, "CRI", "2021-01-01", 10)

		log.Infof("Creating trial %s", trialNumber)

		if err != nil {
			log.Errorf("Error inserting trial: %s", err.Error())
			os.Exit(1)
		}
	}
}
