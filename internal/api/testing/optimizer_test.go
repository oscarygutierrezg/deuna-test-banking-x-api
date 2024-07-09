package testing

import (
	"banking-api/pkg/optimizer"
	"github.com/stretchr/testify/assert"
	"testing"
)

var optimizerEngine = optimizer.NewOptimizer("/home/abraham/falabella/python/altiro-optimizer/weboptimizer.py", "python3.8")

func TestInitOptimizerTest(t *testing.T) {
	assert := assert.New(t)

	assert.NotNil(optimizerEngine)
}
