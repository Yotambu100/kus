package cobertura

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseCoberturaXML(t *testing.T) {
	c, err := ParseCoberturaXML("./base_test.xml")
	assert.NoError(t, err)
	assert.NotNil(t, c)
}
