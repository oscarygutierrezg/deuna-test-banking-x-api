package testing

import (
	"banking-api/pkg/mongodb"
	"github.com/stretchr/testify/assert"
	"testing"
)

var db mongodb.Database

const dbName = "test"

func TestInitMongoDBTest(t *testing.T) {
	assert := assert.New(t)

	errConn := db.ConnectStr("mongodb://127.0.0.1:27017", dbName)

	assert.Nil(errConn)

	_ = db.DeleteDB()
}
