package mongodb

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ = func() (_ struct{}) {
	re := regexp.MustCompile(`^(.*citizenship-api)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		fmt.Println(".env is not loaded properly")

		os.Exit(2)
	}
	return
}()

func TestGetDB(t *testing.T) {
	t.Run("POSITIVE_GET_DB", func(t *testing.T) {
		ctx := context.Background()
		db := GetDB(ctx)
		defer db.Client().Disconnect(ctx)
		assert.NotNil(t, db)
	})
}

func Test_getClient(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want *mongo.Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getClient(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
