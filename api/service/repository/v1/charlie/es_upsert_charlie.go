package charlie

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/getsentry/sentry-go"

	"github.com/olivere/elastic/v7"
	"github.com/opentracing/opentracing-go"

	"github.com/Bhinneka/alpha/api/lib/tracer"
	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
)

func (impl es) UpsertCharlie(ctx context.Context, client *elastic.Client, param domainCharlie.Domain) error {

	var (
		operationName string = "Repository_ES_UpsertCharlie"
		doc           interface{}
	)

	span, _ := opentracing.StartSpanFromContext(ctx, operationName)
	defer span.Finish()

	docID := strconv.FormatUint(param.CharlieID, 10)
	docByte, _ := param.MarshalJSON()
	json.Unmarshal(docByte, &doc)
	span.SetTag("doc.id", docID)
	span.SetTag("doc.body", string(docByte))

	body, err := client.Update().
		Index(domainCharlie.TableName).
		Id(docID).
		Doc(doc).
		DocAsUpsert(true).
		Timeout("10s").
		Refresh("true").
		Do(ctx)

	if err != nil {
		tracer.SetErrorOpentracing(span, "upsert_elastic", err.Error())
		sentry.CaptureException(err)
		return err
	}

	log.Println(fmt.Sprintf("success upsert elastic charlie : %s", body.Result))
	return nil
}
