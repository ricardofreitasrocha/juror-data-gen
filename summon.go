package main

import (
	"database/sql"
	"encoding/json"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/charmbracelet/log"
)

var poolsToSummon []string
var waitForResponses sync.WaitGroup

type Summon struct {
	db     *sql.DB
	config *Config
}

type SummonPayload struct {
	PoolNumber       string   `json:"poolNumber"`
	StartDate        string   `json:"startDate"`
	AttendTime       string   `json:"attendTime"`
	JurorsRequested  int      `json:"noRequested"`
	BureauDeferrals  int      `json:"bureauDeferrals"`
	JurorsRequired   int      `json:"numberRequired"`
	CitizensToSummon int      `json:"citizensToSummon"`
	CatchmentArea    string   `json:"catchmentArea"`
	Postcodes        []string `json:"postcodes"`
}

func summon(db *sql.DB, config *Config) *Summon {
	return &Summon{
		db:     db,
		config: config,
	}
}

func (s *Summon) summon() {

	attendanceDate := time.Now().AddDate(0, 0, s.config.DaysToAdd)

	summonPayload := &SummonPayload{
		BureauDeferrals:  0,
		AttendTime:       attendanceDate.Format("2006-01-02 15:04"),
		StartDate:        attendanceDate.Format("2006-01-02"),
		JurorsRequested:  s.config.VotersPerPool,
		JurorsRequired:   s.config.VotersPerPool,
		CitizensToSummon: s.config.VotersPerPool,
		CatchmentArea:    s.config.LocCode[0],
		Postcodes:        []string{"CH1"},
	}

	for _, pool := range poolsToSummon {
		summonPayload.PoolNumber = pool

		log.Infof("Summoning pool: %s", pool)

		payload, _ := json.Marshal(summonPayload)
		_, err := request("POST", summonVotersUrl.String(), payload, true)
		if err != nil {
			log.Errorf("Errored out on summoning pool: %s", err.Error())
			os.Exit(1)
		}

		log.Debugf("Summoned pool: %s", pool)

		if s.config.AddResponses {
			waitForResponses.Add(1)
			go s.addResponses(pool)
		}
	}

	if s.config.AddResponses {
		waitForResponses.Wait()
	}
}

type JurorPool struct {
	JurorNumber  string `json:"juror_number"`
	Title        string `json:"title"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine4 string `json:"address_line_4"`
	Postcode     string `json:"postcode"`
}

type Eligibility struct {
	LivedConsecutive     bool `json:"livedConsecutive"`
	MentalHealthAct      bool `json:"mentalHealthAct"`
	MentalHealthCapacity bool `json:"mentalHealthCapacity"`
	OnBail               bool `json:"onBail"`
	Convicted            bool `json:"convicted"`
}

type SummonResponse struct {
	JurorNumber           string      `json:"jurorNumber"`
	Title                 string      `json:"title"`
	FirstName             string      `json:"firstName"`
	LastName              string      `json:"lastName"`
	AddressLineOne        string      `json:"addressLineOne"`
	AdressTown            string      `json:"addressTown"`
	AddressPostcode       string      `json:"addressPostcode"`
	CjsEmployment         []string    `json:"cjsEmployment"`
	DateOfBirth           string      `json:"dateOfBirth"`
	Deferral              bool        `json:"deferral"`
	Eligibility           Eligibility `json:"eligibility"`
	EmailAddress          string      `json:"emailAddress"`
	Excusal               bool        `json:"excusal"`
	Signed                bool        `json:"signed"`
	SpecialNeeds          []string    `json:"specialNeeds"`
	Relationship          string      `json:"relationship"`
	Welsh                 bool        `json:"welsh"`
	CanServeOnSummonsDate bool        `json:"canServeOnSummonsDate"`
}

func (s *Summon) addResponses(pool string) {
	defer waitForResponses.Done()

	log.Infof("Adding responses to pool %s", pool)

	rows, err := s.db.Query(`
		SELECT juror_pool.juror_number, juror.title, juror.first_name, juror.last_name, juror.address_line_1, juror.address_line_4, juror.postcode
		FROM juror_mod.juror_pool
		JOIN juror_mod.juror ON juror_pool.juror_number = juror.juror_number
		WHERE pool_number = $1
	`, pool)
	if err != nil {
		log.Errorf("Errored out on getting pool: %s", err.Error())
		os.Exit(1)
	}

	for rows.Next() {
		jurorPool := &JurorPool{}

		err = rows.Scan(&jurorPool.JurorNumber, &jurorPool.Title, &jurorPool.FirstName, &jurorPool.LastName, &jurorPool.AddressLine1, &jurorPool.AddressLine4, &jurorPool.Postcode)
		if err != nil {
			log.Errorf("Errored out on scanning row: %s", err.Error())
			os.Exit(1)
		}

		summonResponse := &SummonResponse{
			JurorNumber:     jurorPool.JurorNumber,
			Title:           jurorPool.Title,
			FirstName:       jurorPool.FirstName,
			LastName:        jurorPool.LastName,
			AddressLineOne:  jurorPool.AddressLine1,
			AdressTown:      jurorPool.AddressLine4,
			AddressPostcode: jurorPool.Postcode,
			CjsEmployment:   []string{},
			DateOfBirth:     "1980-01-01",
			Deferral:        false,
			Eligibility: Eligibility{
				LivedConsecutive:     true,
				MentalHealthAct:      false,
				MentalHealthCapacity: false,
				OnBail:               false,
				Convicted:            false,
			},
			EmailAddress:          "",
			Excusal:               false,
			Signed:                true,
			SpecialNeeds:          []string{},
			Relationship:          "",
			Welsh:                 false,
			CanServeOnSummonsDate: true,
		}

		log.Debugf("Posting response for juror %s", jurorPool.JurorNumber)

		payload, _ := json.Marshal(summonResponse)
		_, err := request("POST", addResponseUrl.String(), payload, true)

		if err != nil {
			log.Errorf("Errored out on adding response: %s", err.Error())
			os.Exit(1)
		}

		processResponse(&s.config.Ranges, jurorPool.JurorNumber)

		log.Debugf("Posted response for juror %s", jurorPool.JurorNumber)
	}
}

func processResponse(r *Ranges, jurorNumber string) {
	random := rand.Float32()

	ranges := map[string]float32{
		"responded":     r.Responded,
		"excused":       r.Excused,
		"deferred":      r.Deferred,
		"disqualified":  r.Disqualified,
		"undeliverable": r.Undeliverable,
	}

	var process string
	cumulative := float32(0)
	for key, value := range ranges {
		cumulative += value
		if random <= cumulative {
			process = key
			break
		}
	}

	// TODO: add fails into an list?

	switch process {
	case "responded":
		_, _ = request("PUT", markAsRespondedUrl.String(jurorNumber), nil, true)
		log.Debugf("Marked %s as responded", jurorNumber)
	case "excused":
		payload, _ := json.Marshal(map[string]string{"excusalReasonCode": "PE", "replyMethod": "PAPER", "excusalDecision": "GRANT"})
		_, err := request("PUT", excusalUrl.String(jurorNumber), payload, true)

		if err != nil {
			log.Errorf("Errored out on excusing %s: %s", jurorNumber, err.Error())
			return
		}

		log.Debugf("Marked %s as excused", jurorNumber)
	case "deferred":
		in6Months := time.Now().AddDate(0, 6, 0).Format("02/01/2006")

		payload, _ := json.Marshal(map[string]string{"deferralReason": "PE", "deferralDecision": "GRANT", "jurorNumber": jurorNumber, "deferralDate": in6Months, "allow_multiple_deferral": "true"})
		_, err := request("PUT", deferralUrl.String(jurorNumber), payload, true)

		if err != nil {
			log.Errorf("Errored out on deferring %s: %s", jurorNumber, err.Error())
			os.Exit(1)
		}

		log.Debugf("Marked %s as deferred", jurorNumber)
	case "disqualified":
		payload, _ := json.Marshal(map[string]string{"code": "C", "replyMethod": "PAPER"})
		_, err := request("PATCH", disqualifyUrl.String(jurorNumber), payload, true)

		if err != nil {
			log.Errorf("Errored out on disqualifying %s: %s", jurorNumber, err.Error())
			return
		}

		log.Debugf("Marked %s as disqualified", jurorNumber)
	case "undeliverable":
		_, _ = request("PUT", markUndeliverable.String(jurorNumber), nil, true)
		log.Debugf("Marked %s as undeliverable", jurorNumber)
	}

}
