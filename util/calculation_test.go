package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDateToPanel(t *testing.T) {
	t.Run("Test DateToPanel same date", func(t *testing.T) {
		dateFormat := "2006-01-02"

		currDate, _ := time.Parse(dateFormat, "2019-11-18")
		inputDate, _ := time.Parse(dateFormat, "2019-11-18")

		expX, expY := 52, 1
		x, y := DateToPanel(currDate, inputDate)

		assert.Equal(t, expX, x)
		assert.Equal(t, expY, y)
	})

	t.Run("Test DateToPanel same week", func(t *testing.T) {
		dateFormat := "2006-01-02"

		currDate, _ := time.Parse(dateFormat, "2019-11-21")
		inputDate, _ := time.Parse(dateFormat, "2019-11-18")

		expX, expY := 52, 1
		x, y := DateToPanel(currDate, inputDate)

		assert.Equal(t, expX, x)
		assert.Equal(t, expY, y)
	})

	t.Run("Test DateToPanel different date, same day", func(t *testing.T) {
		dateFormat := "2006-01-02"

		currDate, _ := time.Parse(dateFormat, "2019-11-18")
		inputDate, _ := time.Parse(dateFormat, "2019-11-04")

		expX, expY := 50, 1
		x, y := DateToPanel(currDate, inputDate)

		assert.Equal(t, expX, x)
		assert.Equal(t, expY, y)
	})

	t.Run("Test DateToPanel different date", func(t *testing.T) {
		dateFormat := "2006-01-02"

		currDate, _ := time.Parse(dateFormat, "2019-11-18")
		inputDate, _ := time.Parse(dateFormat, "2019-01-02")

		expX, expY := 6, 3
		x, y := DateToPanel(currDate, inputDate)

		assert.Equal(t, expX, x)
		assert.Equal(t, expY, y)
	})
}

func TestPanelToDate(t *testing.T) {
	t.Run("Test PanelToDate same week", func(t *testing.T) {
		dateFormat := "2006-01-02"

		currDate, _ := time.Parse(dateFormat, "2019-11-21")
		expDate, _ := time.Parse(dateFormat, "2019-11-18")

		x, y := 52, 1
		actDate := PanelToDate(currDate, x, y)

		assert.Equal(t, expDate, actDate)
	})

	t.Run("Test PanelToDate different date, same day", func(t *testing.T) {
		dateFormat := "2006-01-02"

		currDate, _ := time.Parse(dateFormat, "2019-11-18")
		expDate, _ := time.Parse(dateFormat, "2019-11-04")

		x, y := 50, 1
		actDate := PanelToDate(currDate, x, y)

		assert.Equal(t, expDate, actDate)
	})

	t.Run("Test PanelToDate different date", func(t *testing.T) {
		dateFormat := "2006-01-02"

		currDate, _ := time.Parse(dateFormat, "2019-11-18")
		expDate, _ := time.Parse(dateFormat, "2019-01-02")

		x, y := 6, 3
		actDate := PanelToDate(currDate, x, y)

		assert.Equal(t, expDate, actDate)
	})
}
