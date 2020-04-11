package beta

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/opentracing/opentracing-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Bhinneka/alpha/api/lib/tracer"
	domainBeta "github.com/Bhinneka/alpha/api/service/domain/v1/beta"
)

func (impl mongoDB) InsertBetaHistory(db *mongo.Database, param domainBeta.Domain) {
	var (
		retry         int
		opertaionName string = "Repository_Mongodb_InsertBetaHistory"
		doc           interface{}
	)

	ctx := context.Background()

	span := opentracing.StartSpan(opertaionName)
	defer span.Finish()

	docByte, _ := param.MarshalBSONValue()
	bson.Unmarshal(docByte, &doc)

	for {
		retry++
		collection := db.Collection(domainBeta.TableName)
		res, err := collection.InsertOne(ctx, &doc)
		if err == nil {
			log.Println(fmt.Sprintf("success mongodb insert beta : %s", res))
			break
		}

		tracer.SetErrorOpentracing(span, "insert_mongodb", err.Error())
		log.Println(fmt.Sprintf("error mongodb insert beta : %s, retrying in %v seconds", err.Error(), retry))
		time.Sleep(time.Duration(retry) * time.Second)
	}
}
