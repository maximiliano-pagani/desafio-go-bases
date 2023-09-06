package repository

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func LoadTicketsReport(filepath string) (*tickets.TicketReport, error) {
	rawData, err := loadCsvData(filepath)

	if err != nil {
		panic(err)
	}

	ticketsData, err := parseTicketsData(rawData)

	if err != nil {
		panic(err)
	}

	ticketReport := tickets.TicketReport{Tickets: ticketsData}

	return &ticketReport, nil
}

func loadCsvData(filepath string) ([][]string, error) {
	var rawData [][]string

	file, err := os.Open(filepath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	rawData, err = csv.NewReader(file).ReadAll()

	if err != nil {
		panic(err)
	}

	return rawData, err
}

func parseTicketsData(rawData [][]string) ([]tickets.Ticket, error) {
	var (
		ticketsData []tickets.Ticket
		err         error
	)

	for _, ticketLine := range rawData {
		ticket := tickets.Ticket{}

		for attrIndex, attr := range ticketLine {
			switch attrIndex {
			case 0:
				ticket.Id, err = strconv.Atoi(attr)
			case 1:
				ticket.Name = attr
			case 2:
				ticket.Email = attr
			case 3:
				ticket.DestCountry = attr
			case 4:
				ticket.Time = attr
			case 5:
				ticket.Price, err = strconv.Atoi(attr)
			}

			if err != nil {
				return nil, err
			}
		}

		ticketsData = append(ticketsData, ticket)
	}

	return ticketsData, nil
}
