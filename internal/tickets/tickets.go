package tickets

import (
	"strconv"
	"strings"
)

const (
	PeriodDawn      = "0-6"
	PeriodMorning   = "7-12"
	PeriodAfternoon = "13-19"
	PeriodNight     = "20-24"
)

type Ticket struct {
	Id          int
	Name        string
	Email       string
	DestCountry string
	Time        string
	Price       int
}

type TicketReport struct {
	Tickets []Ticket
}

func (r *TicketReport) GetTotalTickets(destination string) (int, error) {
	var count int
	for _, ticket := range r.Tickets {
		if ticket.DestCountry == destination {
			count++
		}
	}

	return count, nil
}

func (r *TicketReport) GetCountByPeriod(period string) (int, error) {
	var count int
	minHour, maxHour := getPeriodHourRange(period)

	for _, ticket := range r.Tickets {
		if isTicketInHourRange(ticket, minHour, maxHour) {
			count++
		}
	}
	return count, nil
}

func (r *TicketReport) AverageDestination(destination string, total int) (float64, error) {
	count, err := r.GetTotalTickets(destination)
	return float64(count) / float64(total), err
}

func isTicketInHourRange(ticket Ticket, minHour, maxHour int) bool {
	ticketHour, _ := strconv.Atoi(strings.Split(ticket.Time, ":")[0])
	if ticketHour >= minHour && ticketHour <= maxHour {
		return true
	}
	return false
}

func getPeriodHourRange(period string) (minHour int, maxHour int) {
	hourRange := strings.Split(period, "-")
	minHour, _ = strconv.Atoi(hourRange[0])
	maxHour, _ = strconv.Atoi(hourRange[1])
	return
}
