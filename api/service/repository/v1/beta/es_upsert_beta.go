package beta

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
	domainBeta "github.com/Bhinneka/alpha/api/service/domain/v1/beta"
)

func (impl es) UpsertBeta(ctx context.Context, client *elastic.Client, param domainBeta.Domain) error {
	var (
		operationName string = "Repository_ES_UpsertBeta"
		doc           interface{}
	)

	span, _ := opentracing.StartSpanFromContext(ctx, operationName)
	defer span.Finish()

	docID := strconv.FormatUint(param.BetaID, 10)
	docByte, _ := param.MarshalJSON()
	json.Unmarshal(docByte, &doc)
	span.SetTag("doc.id", docID)
	span.SetTag("doc.body", string(docByte))

	body, err := client.Update().
		Index(domainBeta.TableName).
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

	log.Println(fmt.Sprintf("success upsert elastic beta : %s", body.Result))
	return nil

}
