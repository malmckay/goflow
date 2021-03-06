package utils_test

import (
	"testing"

	"github.com/nyaruka/goflow/utils"

	"github.com/stretchr/testify/assert"
)

func TestJSONMarshaling(t *testing.T) {
	j, err := utils.JSONMarshal(nil)
	assert.NoError(t, err)
	assert.Equal(t, []byte(`null`), j)

	j, err = utils.JSONMarshal("Rwanda > Kigali")
	assert.NoError(t, err)
	assert.Equal(t, []byte(`"Rwanda > Kigali"`), j)

	j, err = utils.JSONMarshal(map[string]string{"foo": "bar"})
	assert.NoError(t, err)
	assert.Equal(t, []byte(`{"foo":"bar"}`), j)

	j, err = utils.JSONMarshalPretty(map[string]string{"foo": "bar"})
	assert.NoError(t, err)
	assert.Equal(t, []byte("{\n    \"foo\": \"bar\"\n}"), j)
}
