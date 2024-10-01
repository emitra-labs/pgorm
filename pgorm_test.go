package pgorm_test

import (
	"testing"

	"github.com/emitra-labs/pgorm"
)

func TestSimpleQuery(t *testing.T) {
	// Open database connection
	pgorm.Open()

	var result int64

	err := pgorm.DB.Raw("SELECT 1 + 1").Scan(&result).Error
	if err != nil {
		panic(err)
	}

	if result != 2 {
		t.Errorf("expected 2, got %d", result)
	}

	// Close the connection
	pgorm.Close()
}
