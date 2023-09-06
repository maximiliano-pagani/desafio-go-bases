package tickets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalTickets(t *testing.T) {
	t.Run("Get total tickets to Brazil", func(t *testing.T) {
		// Arrange
		ticketsReport := TicketReport{
			Tickets: []Ticket{
				{DestCountry: "Finland", Time: "16:11"},
				{DestCountry: "China", Time: "20:19"},
				{DestCountry: "China", Time: "19:11"},
				{DestCountry: "Mongolia", Time: "3:16"},
			},
		}
		expectedResult := 0

		// Act
		result, _ := ticketsReport.GetTotalTickets("Brazil")

		// Assert
		assert.Equal(t, expectedResult, result, fmt.Sprintf("Total tickets is not %d", expectedResult))
	})

	t.Run("Get total tickets to China", func(t *testing.T) {
		// Arrange
		ticketsReport := TicketReport{
			Tickets: []Ticket{
				{DestCountry: "Finland", Time: "16:11"},
				{DestCountry: "China", Time: "20:19"},
				{DestCountry: "China", Time: "19:11"},
				{DestCountry: "Mongolia", Time: "3:16"},
			},
		}
		expectedResult := 2

		// Act
		result, _ := ticketsReport.GetTotalTickets("China")

		// Assert
		assert.Equal(t, expectedResult, result, fmt.Sprintf("Total tickets is not %d", expectedResult))
	})
}

func TestGetCountByPeriod(t *testing.T) {
	t.Run("Get count by period afternoon", func(t *testing.T) {
		// Arrange
		ticketsReport := TicketReport{
			Tickets: []Ticket{
				{DestCountry: "Finland", Time: "16:11"},
				{DestCountry: "China", Time: "20:19"},
				{DestCountry: "China", Time: "19:11"},
				{DestCountry: "Mongolia", Time: "3:16"},
			},
		}
		expectedResult := 2

		// Act
		result, _ := ticketsReport.GetCountByPeriod(PeriodAfternoon)

		// Assert
		assert.Equal(t, expectedResult, result, fmt.Sprintf("Period count is not %d", expectedResult))
	})

	t.Run("Get count by period morning", func(t *testing.T) {
		// Arrange
		ticketsReport := TicketReport{
			Tickets: []Ticket{
				{DestCountry: "Finland", Time: "16:11"},
				{DestCountry: "China", Time: "20:19"},
				{DestCountry: "China", Time: "19:11"},
				{DestCountry: "Mongolia", Time: "3:16"},
			},
		}
		expectedResult := 0

		// Act
		result, _ := ticketsReport.GetCountByPeriod(PeriodMorning)

		// Assert
		assert.Equal(t, expectedResult, result, fmt.Sprintf("Period count is not %d", expectedResult))
	})
}

func TestAverageDestination(t *testing.T) {
	t.Run("Average destination to Brazil", func(t *testing.T) {
		// Arrange
		ticketsReport := TicketReport{
			Tickets: []Ticket{
				{DestCountry: "Finland", Time: "16:11"},
				{DestCountry: "China", Time: "20:19"},
				{DestCountry: "China", Time: "19:11"},
				{DestCountry: "Mongolia", Time: "3:16"},
			},
		}
		expectedResult := 0.0

		// Act
		result, _ := ticketsReport.AverageDestination("Brazil", len(ticketsReport.Tickets))

		// Assert
		assert.Equal(t, expectedResult, result, fmt.Sprintf("Average destination is not %v", expectedResult))
	})
}

func TestIsTicketInHourRange(t *testing.T) {
	t.Run("Is ticket in hour range", func(t *testing.T) {
		// Arrange
		ticket := Ticket{Time: "15:09"}
		expectedResult := false
		minHourRange := 5
		maxHourRange := 10

		// Act
		result := isTicketInHourRange(ticket, minHourRange, maxHourRange)

		// Assert
		assert.Equal(t, expectedResult, result, "Ticket is not in hour range but got a false positive")
	})

	t.Run("Is ticket in hour range", func(t *testing.T) {
		// Arrange
		ticket := Ticket{Time: "5:45"}
		expectedResult := true
		minHourRange := 5
		maxHourRange := 10

		// Act
		result := isTicketInHourRange(ticket, minHourRange, maxHourRange)

		// Assert
		assert.Equal(t, expectedResult, result, "Ticket is in hour range but failed to detect so")
	})
}

func TestGetPeriodHourRange(t *testing.T) {
	t.Run("Get period hour range", func(t *testing.T) {
		// Arrange
		periodTest := "9-11"
		minExpected := 9
		maxExpected := 11

		// Act
		minResult, maxResult := getPeriodHourRange(periodTest)

		// Assert
		assert.Equal(t, minExpected, minResult, fmt.Sprintf("Min hour range is not %d", minExpected))
		assert.Equal(t, maxExpected, maxResult, fmt.Sprintf("Max hour range is not %d", maxExpected))
	})
}
