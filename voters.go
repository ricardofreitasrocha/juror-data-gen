package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/charmbracelet/log"
)

type Voters struct {
	db       *sql.DB
	locCode  string
	count    int
	postcode []string
}

type Person struct {
	Gender string `json:"gender"`
	Name   struct {
		Title string `json:"title"`
		First string `json:"first"`
		Last  string `json:"last"`
	}
	Location struct {
		Street struct {
			Name string `json:"name"`
		}
		City     string `json:"city"`
		Postcode string `json:"postcode"`
	}
}

type Results struct {
	Results []Person `json:"results"`
}

func voters(db *sql.DB, locCode string, count int, postcode []string) *Voters {
	return &Voters{
		db:       db,
		locCode:  locCode,
		count:    count,
		postcode: postcode,
	}
}

func (v *Voters) insert() {
	people, _ := http.Get(fmt.Sprintf("https://randomuser.me/api/?nat=gb&results=%d", v.count))

	results := &Results{}
	json.NewDecoder(people.Body).Decode(results)

	log.Infof("Will insert %d voters", v.count)

	for i, person := range results.Results {
		title := person.Name.Title
		first := person.Name.First
		last := person.Name.Last
		street := person.Location.Street.Name
		city := person.Location.City
		postcode := v.postcode[rand.Intn(len(v.postcode))]

		jn := jurorNumber("0"+v.locCode, strconv.Itoa(i))

		query := fmt.Sprintf("INSERT INTO juror_mod.VOTERS (LOC_CODE,PART_NO,REGISTER_LETT,POLL_NUMBER,NEW_MARKER,TITLE,FNAME,LNAME,ADDRESS,ADDRESS4,ZIP,REC_NUM) VALUES ('%s','%s','%d','%d',NULL,'%s','%s','%s','%s','%s','%s',%d);", v.locCode, jn, i, i, title, first, last, street, city, postcode, i)

		log.Infof("Inserting voter %s", jn)

		v.db.Exec(query)
	}

	waitForVoters.Done()
}

func jurorNumber(locCode, i string) string {
	pad := 9 - len(locCode) - len(i)

	for i := 0; i < pad; i++ {
		locCode = locCode + "0"
	}

	return locCode + i
}
