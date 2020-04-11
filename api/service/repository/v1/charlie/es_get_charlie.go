package charlie

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/getsentry/sentry-go"
	"github.com/olivere/elastic/v7"
	"github.com/opentracing/opentracing-go"

	"github.com/Bhinneka/alpha/api/lib/constant"
	"github.com/Bhinneka/alpha/api/lib/tracer"
	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
)

func (impl es) GetCharlie(ctx context.Context, client *elastic.Client, param domainCharlie.ParamGet) (result []domainCharlie.Domain, totalData int64, err error) {
	const (
		operationName string = "Repository_ES_GetCharlie"
	)

	span, _ := opentracing.StartSpanFromContext(ctx, operationName)
	defer span.Finish()

	searchSource := getCharlieSearchSource(param)

	query, _ := searchSource.Source()
	byteQuery, _ := json.Marshal(query)
	span.SetTag("es.query", string(byteQuery))

	searchService := client.Search().
		Index(domainCharlie.TableName).
		Pretty(true).
		SearchSource(searchSource)

	res, err := searchService.Do(ctx)
	if err != nil {
		tracer.SetErrorOpentracing(span, constant.ErroTypeESQuery, err.Error())
		sentry.CaptureException(err)
		return nil, 0, err
	}

	if res.TotalHits() == 0 {
		return nil, 0, nil
	}

	for _, hit := range res.Hits.Hits {
		charlieDomain := domainCharlie.Domain{}
		if err := json.Unmarshal(hit.Source, &charlieDomain); err != nil {
			tracer.SetErrorOpentracing(span, constant.ErroTypeESQuery, err.Error())
			sentry.CaptureException(err)
			return nil, 0, err
		}
		result = append(result, charlieDomain)
	}

	return result, res.TotalHits(), err
}

func getCharlieSearchSource(param domainCharlie.ParamGet) *elastic.SearchSource {

	offset := (param.Page - 1) * param.Limit

	sourceQuery := elastic.NewSearchSource().
		From(offset).
		Size(param.Limit)

	orderBy := param.OrderBy
	if orderBy == "" {
		orderBy = "charlieID"
	}

	if param.Descending {
		sourceQuery = sourceQuery.SortBy(elastic.NewFieldSort(orderBy).Desc())
	} else {
		sourceQuery = sourceQuery.SortBy(elastic.NewFieldSort(orderBy).Asc())
	}

	query := elastic.NewBoolQuery()

	if param.CharlieID != 0 {
		query = query.Must(elastic.NewTermQuery("charlieID", param.CharlieID))
	}

	if strings.TrimSpace(param.CharlieName) != "" {
		query = query.Must(elastic.NewWildcardQuery("charlieName", fmt.Sprintf("*%s*", param.CharlieName)))
	}

	sourceQuery = sourceQuery.Query(query)

	return sourceQuery
}
