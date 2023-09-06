package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/repository"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	dailyTicketReport, err := repository.LoadTicketsReport("./tickets.csv")

	if err != nil {
		panic("Error: falla al cargar el archivo csv")
	}

	totalTickets, _ := dailyTicketReport.GetTotalTickets("Brazil")
	fmt.Println("Total tickets to Brazil:", totalTickets)

	periodCount, _ := dailyTicketReport.GetCountByPeriod(tickets.PeriodMorning)
	fmt.Println("Period tickets count:", periodCount)

	averageTicketsToDest, _ := dailyTicketReport.AverageDestination("Brazil", len(dailyTicketReport.Tickets))
	fmt.Println("Average tickets to Brazil:", averageTicketsToDest)
}
