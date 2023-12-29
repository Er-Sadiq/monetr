package recurring

import (
	"testing"

	"github.com/monetr/monetr/server/models"
	"github.com/stretchr/testify/assert"
)

func TestCleanNameRegex(t *testing.T) {
	t.Run("long no merchant", func(t *testing.T) {
		txn := models.Transaction{
			OriginalName:         "WHEN I WORK INC:1233303024 57:COURANT,ELLIOT; 798080132284EPJ. Merchant name: WHEN I WORK INC",
			OriginalMerchantName: "",
		}

		lower, _ := CleanNameRegex(&txn)
		assert.EqualValues(t, []string{
			"when", "work", "inc", "courant", "elliot", "merchant", "name", "when", "work", "inc",
		}, lower, "should match the cleaned string")
	})

	t.Run("github", func(t *testing.T) {
		txn := models.Transaction{
			OriginalName:         "GITHUB. Merchant name: GITHUB",
			OriginalMerchantName: "GitHub",
		}

		lower, _ := CleanNameRegex(&txn)
		assert.EqualValues(t, []string{
			"github", "merchant", "name", "github", "github",
		}, lower, "should match the cleaned string")
	})
}

func BenchmarkCleanNameRegex(b *testing.B) {
	b.StopTimer()
	txn := models.Transaction{
		OriginalName:         "WHEN I WORK INC:1233303024 57:COURANT,ELLIOT; 798080132284EPJ. Merchant name: WHEN I WORK INC",
		OriginalMerchantName: "",
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CleanNameRegex(&txn)
	}
}
