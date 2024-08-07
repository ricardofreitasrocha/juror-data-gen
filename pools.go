package main

import (
	"database/sql"
	"encoding/json"
	"math/rand"
	"time"

	"github.com/charmbracelet/log"
)

type Pools struct {
	db     *sql.DB
	config *Config
}

type Pool struct {
	PoolNumber      string `json:"poolNumber"`
	CourtCode       string `json:"courtCode"`
	AttendanceDate  string `json:"attendanceDate"`
	NumberRequested int    `json:"numberRequested"`
	PoolType        string `json:"poolType"`
	AttendanceTime  string `json:"attendanceTime"`
	DeferralsUsed   int    `json:"deferralsUsed"`
	CourtOnly       bool   `json:"courtOnly"`
}

func pools(db *sql.DB, config *Config) *Pools {
	return &Pools{
		db:     db,
		config: config,
	}
}

func (p *Pools) request() {
	pool := &Pool{}

	for _, locCode := range p.config.LocCode {
		for i := 0; i < p.config.Pools; i++ {
			poolNumber, _ := generatePoolNumber(locCode)

			attendanceDate := time.Now().AddDate(0, 0, p.config.DaysToAdd)

			pool.PoolNumber = poolNumber
			pool.CourtCode = locCode
			pool.AttendanceDate = attendanceDate.Format("2006-01-02")
			pool.AttendanceTime = attendanceDate.Format("15:04")
			pool.NumberRequested = p.config.VotersPerPool
			pool.PoolType = poolType()
			pool.DeferralsUsed = 0
			pool.CourtOnly = false

			log.Infof("Requesting pool: %s - Pool type: %s", pool.PoolNumber, pool.PoolType)

			payload, _ := json.Marshal(pool)

			if _, err := request("POST", requestPoolUrl.String(), payload, true); err != nil {
				log.Error(err)
				panic("Something went wrong")
			}

			if p.config.Summon {
				poolsToSummon[locCode] = append(poolsToSummon[locCode], poolNumber)
			}
		}
	}
}

func poolType() string {
	random := rand.Intn(3)

	switch random {
	case 0:
		return "CRO"
	case 1:
		return "CIV"
	case 2:
		return "HGH"
	}

	return ""
}
