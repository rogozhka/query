package query

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {

	rows, err := Fetch("SELECT * FROM `_x_query_test_table1`", db)

	assert.Nil(t, err, "SELECT result")
	assert.True(t, len(rows) == 50, "Rows number")

	first := rows[0]

	reqColumns := map[string]bool{
		"test_uid":       true,
		"test_datetime":  true,
		"test_expired":   true,
		"test_is_active": true,
		"test_number":    true,
	}

	checkKeysMatch(t, &reqColumns, &first)

	assert.Equal(t, "cron", rows[0]["test_uid"], "0:test_uid")
	assert.Equal(t, "memset", rows[1]["test_uid"], "1:test_uid")
	assert.Equal(t, "update8k", rows[2]["test_uid"], "2:test_uid")

	assert.Equal(t, "cron:text:", rows[0]["test_text"], "0:test_text")
	assert.Equal(t, "memset:text:", rows[1]["test_text"], "1:test_text")
	assert.Equal(t, "update8k:text:", rows[2]["test_text"], "2:test_text")

	assert.Equal(t, "1", rows[0]["test_number"], "0:test_number")
	assert.Equal(t, "2", rows[1]["test_number"], "1:test_number")
	assert.Equal(t, "3", rows[2]["test_number"], "2:test_number")

}

func TestFetchLimited(t *testing.T) {

	rows, err := FetchLimited("SELECT * FROM `_x_query_test_table1`", 2, db)

	assert.Nil(t, err, "SELECT result")
	assert.True(t, len(rows) == 2, "Rows number")

	first := rows[0]

	reqColumns := map[string]bool{
		"test_uid":       true,
		"test_datetime":  true,
		"test_expired":   true,
		"test_is_active": true,
		"test_number":    true,
	}

	checkKeysMatch(t, &reqColumns, &first)

	assert.Equal(t, "cron", rows[0]["test_uid"], "0:test_uid")
	assert.Equal(t, "memset", rows[1]["test_uid"], "1:test_uid")

	assert.Equal(t, "cron:text:", rows[0]["test_text"], "0:test_text")
	assert.Equal(t, "memset:text:", rows[1]["test_text"], "1:test_text")

	assert.Equal(t, "1", rows[0]["test_number"], "0:test_number")
	assert.Equal(t, "2", rows[1]["test_number"], "1:test_number")

}

func checkKeysMatch(t *testing.T, reqColumns *map[string]bool, actual *Row) {

	for key, _ := range *reqColumns {
		_, is := (*actual)[key]

		assert.True(t, is, fmt.Sprintf("Key present:%s", key))
	}
}
