package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"syreclabs.com/go/faker"

	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"
)

type Logging struct {
	UUID   uuid.UUID `json:"uuid"`
	Status int64     `json:"status"`
	Body   struct {
		City       string `json:"city"`
		Department string `json:"department"`
		Company    string `json:"company"`
	} `json:"data"`
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/trigger", triggerHandler)

	http.ListenAndServe(":8080", mux)

}

func triggerHandler(w http.ResponseWriter, r *http.Request) {

	HTTPMethodList := []int64{404, 500, 200, 504, 502, 400}
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)

	loop := rand.Intn(10)
	for i := 1; i <= loop; i++ {

		var newLog Logging
		newLog.UUID = uuid.New()
		newLog.Status = HTTPMethodList[rand.Intn(cap(HTTPMethodList)-1)]
		newLog.Body.City = faker.Address().City()
		newLog.Body.Department = faker.Commerce().Department()
		newLog.Body.Company = faker.Company().Name()

		//log.WithField("data", newLog).Info()
		//fmt.Fprint(w, newLog)
		log.WithFields(log.Fields{
			"UUID":       uuid.New(),
			"Status":     HTTPMethodList[rand.Intn(cap(HTTPMethodList)-1)],
			"City":       faker.Address().City(),
			"Department": faker.Commerce().Department(),
			"Company":    faker.Company().Name(),
		}).Info("new log")
	}

	returnText := "Generate " + strconv.Itoa(loop) + " logs"
	fmt.Fprint(w, returnText)

}
