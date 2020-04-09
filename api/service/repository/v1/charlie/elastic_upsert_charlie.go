package charlie

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/opentracing/opentracing-go"

	"github.com/Bhinneka/alpha/api/lib/tracer"
	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
)

var mutexElastic = &sync.Mutex{}

func (impl elastic) UpsertCharlie(ctx context.Context, db *elasticsearch.Client, param domainCharlie.Domain) {
	mutexElastic.Lock()

	var (
		retry         int
		operationName string = "Repository_Elastic_UpsertCharlie"
	)

	span := opentracing.StartSpan(operationName)
	defer span.Finish()

	byteParam, _ := param.MarshalElasticValue()
	docID := strconv.FormatUint(param.CharlieID, 10)
	doc := `{
		"doc" : ` + string(byteParam) + `,
		"doc_as_upsert" : true
	}`

	span.SetTag("doc.id", docID)
	span.SetTag("doc.body", doc)

	for {
		retry++
		span.LogKV("retry.count", fmt.Sprintf("%v", retry))
		res, err := db.Update(
			domainCharlie.TableName,
			docID,
			strings.NewReader(doc),
		)
		if err == nil {
			log.Println(fmt.Sprintf("success upsert elastic charlie : %s", res))
			break
		}

		tracer.SetErrorOpentracing(span, "upsert_elastic", err.Error())
		log.Println(fmt.Sprintf("error upsert elastic charlie : %s, retrying in %v seconds", err.Error(), retry))
		time.Sleep(time.Duration(retry) * time.Second)

	}
	mutexElastic.Unlock()
}
