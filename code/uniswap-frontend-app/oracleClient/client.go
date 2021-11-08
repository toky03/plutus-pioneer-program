package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	krakenapi "github.com/beldur/kraken-go-api-client"
	"github.com/joho/godotenv"
)

const (
	restEndpoint = "/api/new/contract/instance"
)

type oracleAdapter struct {
	baseUrl string
}

func main() {

	godotenv.Load(".env")
	updateIntervallMinutes := os.Getenv("UPDATE_INTERVALL_MINUTES")
	if updateIntervallMinutes == "" {
		log.Printf("Environment Variable is %s /n", updateIntervallMinutes)
		log.Fatal("UPDATE_INTERVALL_MINUTES must be configured")
	}
	updateIntervallMinutesInt, err := strconv.Atoi(updateIntervallMinutes)
	if err != nil {
		log.Fatalf("could not convert %s into an Integer", updateIntervallMinutes)
	}

	oracleUuid := readFile("oracle.cid")

	adapter := createAdapter()
	for {
		lastPrice := readExchangeRate()
		adapter.sendUpdate(oracleUuid, lastPrice)
		time.Sleep(time.Duration(updateIntervallMinutesInt) * time.Second)

	}

}

func (a *oracleAdapter) sendUpdate(uuid, lastPrice string) {
	url := fmt.Sprintf("%s/%s/%s", a.baseUrl, uuid, "endpoint/update")
	payload := strings.NewReader(lastPrice)
	var response *http.Response
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("content-type", "application/json")

	response, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	} else if response.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(response.Body)
		err = errors.New("Status does not match expectation of 200 actual status is: " + response.Status + " content " + string(body))
		log.Println(err)
	} else {
		log.Printf("Updated oracle successfull to %s exchange", lastPrice)
	}
}

func readExchangeRate() string {
	apiKey := os.Getenv("KRAKEN_API_KEY")

	api := krakenapi.New(apiKey, "SECRET")

	// There are also some strongly typed methods available
	ticker, err := api.Ticker(krakenapi.ADAUSD)
	if err != nil {
		log.Fatal(err)
	}

	return ticker.ADAUSD.Close[0]

}

func createAdapter() *oracleAdapter {
	url := os.Getenv("ORACLE_URL")
	port := os.Getenv("ORACLE_PORT")
	if url == "" {
		url = "http://127.0.0.1"
	}
	if port == "" {
		port = "8080"
	}
	return &oracleAdapter{
		baseUrl: url + ":" + port + restEndpoint,
	}
}

func readFile(filename string) string {
	cidRootDirectory := os.Getenv("CID_ROOT")
	if cidRootDirectory == "" {
		cidRootDirectory = ".."
	}
	data, err := ioutil.ReadFile(cidRootDirectory + "/" + filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return ""
	}
	return string(data)
}
