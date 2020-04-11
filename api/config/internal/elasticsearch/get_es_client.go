package elasticsearch

import (
	"log"
	"os"
	"strconv"

	"github.com/olivere/elastic/v7"
)

// GetESClient : get elastic search client
func GetESClient() *elastic.Client {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(
			os.Getenv("ELASTIC_HOST_1"),
		),
		// elastic.SetBasicAuth(
		// 	os.Getenv("ELASTIC_USERNAME"),
		// 	os.Getenv("ELASTIC_PASSWORD"),
		// ),
	)
	if err != nil {
		panic(err)
	}

	_, err = strconv.ParseBool(os.Getenv("SERVER_LOG_MODE"))
	if err == nil {
		elastic.SetTraceLog(log.New(os.Stdout, "", log.LstdFlags))
	}

	return client
}
