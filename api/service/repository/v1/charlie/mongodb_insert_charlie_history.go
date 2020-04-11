package charlie

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/opentracing/opentracing-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Bhinneka/alpha/api/lib/tracer"
	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
)

func (impl mongoDB) InsertCharlieHistory(db *mongo.Database, param domainCharlie.Domain) {
	var (
		retry         int
		opertaionName string = "Repository_Mongodb_InsertCharlieHistory"
		doc           interface{}
	)

	ctx := context.Background()

	span := opentracing.StartSpan(opertaionName)
	defer span.Finish()

	docByte, _ := param.MarshalBSONValue()
	bson.Unmarshal(docByte, &doc)

	for {
		retry++
		collection := db.Collection(domainCharlie.TableName)
		res, err := collection.InsertOne(ctx, doc)
		if err == nil {
			log.Println(fmt.Sprintf("success mongodb insert charlie : %s", res))
			break
		}

		tracer.SetErrorOpentracing(span, "insert_mongodb", err.Error())
		log.Println(fmt.Sprintf("error mongodb insert charlie : %s, retrying in %v seconds", err.Error(), retry))
		time.Sleep(time.Duration(retry) * time.Second)
	}
}
