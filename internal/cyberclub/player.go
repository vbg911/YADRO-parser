package cyberclub

import (
	"fmt"
	"time"
)

type player struct {
	name string
}

func (p player) takeTableInClub(c *Cyberсlub, tableIndex int, eventTime time.Time) error {
	if tableIndex > c.TableAmount {
		return fmt.Errorf("TableIndexOutOfRange")
	}

	t := table{
		revenue:       c.tables[tableIndex].revenue,
		startRentTime: eventTime,
		rentDuration:  c.tables[tableIndex].rentDuration,
		isBusy:        true,
		usedBy:        p,
	}
	c.tables[tableIndex] = t
	return nil
}

func (p player) clearTableInClub(c *Cyberсlub, tableIndex int, eventTime time.Time) error {
	if tableIndex > c.TableAmount {
		return fmt.Errorf("TableIndexOutOfRange")
	}

	oldTable := c.tables[tableIndex]
	newRevenue := oldTable.calculateNewRevenue(eventTime, c.HourPrice)
	startTime := c.tables[tableIndex].startRentTime
	timeDif := eventTime.Sub(startTime)
	NewRentDuration := c.tables[tableIndex].rentDuration.Add(timeDif)
	t := table{
		revenue:       newRevenue,
		startRentTime: time.Time{},
		rentDuration:  NewRentDuration,
		isBusy:        false,
		usedBy:        player{},
	}
	c.tables[tableIndex] = t

	return nil
}

func (p player) takeQueueInClub(c *Cyberсlub) {
	c.queue = append(c.queue, p)
}
