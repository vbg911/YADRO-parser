package cyberclub

import (
	"math"
	"time"
)

type table struct {
	revenue       int       // Текущая выручка стола
	startRentTime time.Time // Время старта аренды
	rentDuration  time.Time // Промежуток времени когда стол занят
	isBusy        bool      // Занят ли стол игроком
	usedBy        player    // Игрок кем занят стол
}

func (t *table) calculateNewRevenue(endTime time.Time, rate int) int {
	startTime := t.startRentTime
	PastRev := t.revenue
	timeDif := endTime.Sub(startTime)
	currentRev := int(math.Ceil(timeDif.Hours())) * rate
	return PastRev + currentRev
}
