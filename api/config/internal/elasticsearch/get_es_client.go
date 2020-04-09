package elasticsearch

import (
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
)

// GetESClient : get elastic search client
func GetESClient() *elasticsearch.Client {
	hosts := strings.Split(os.Getenv("ELASTIC_7_HOST"), ",")
	cfg := elasticsearch.Config{
		Addresses: hosts,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	_, err = es.Info()
	if err != nil {
		panic(err)
	}
	return es
}
