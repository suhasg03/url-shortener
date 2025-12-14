package urlshortener

import (
	"math/rand"
)

type TicketServer struct {
	ID      int64
	Start   int64
	End     int64
	Current int64
}

var ticketServer []TicketServer = []TicketServer{
	{
		ID:      1,
		Start:   1000000000,
		End:     3000000000,
		Current: 1000000000,
	},
	{
		ID:      2,
		Start:   3000000000,
		End:     6000000000,
		Current: 3000000000,
	},
	{
		ID:      3,
		Start:   6000000000,
		End:     9000000000,
		Current: 6000000000,
	},
	{
		ID:      4,
		Start:   9000000000,
		End:     12000000000,
		Current: 9000000000,
	},
}

func GetPsuedoRandomIntegerId() int64 {
	index := generateRandomNumber(len(ticketServer))
	current := ticketServer[index].Current
	ticketServer[index].Current += 1
	return current
}

func generateRandomNumber(count int) int {
	randomNum := rand.Intn(count)
	return randomNum
}
