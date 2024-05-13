package cyberclub

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	incomingClientEntered  = 1
	incomingClientTakeSeat = 2
	incomingClientWaiting  = 3
	incomingClientLeft     = 4
)

const (
	outcomingClientLeft     = 11
	outcomingClientTakeSeat = 12
)

const (
	clientAlreadyInClub   = "13 YouShallNotPass"
	clientComeAtWrongTime = "13 NotOpenYet"
	clientUnknown         = "13 clientUnknown"
	placeIsBusy           = "13 PlaceIsBusy"
	clubHaveFreeTable     = "13 IcanWaitNoLonger!"
)

const timeLayout = "15:04"

type Cyberсlub struct {
	TableAmount  int           // Количество столов в клубе
	WorkingStart time.Time     // Время открытия
	WorkingEnd   time.Time     // Время закрытия
	HourPrice    int           // Цена за 1 час
	tables       map[int]table // Все столы клуба
	queue        []player      // Очередь клиентов
	Scanner      *bufio.Scanner
}

func (c *Cyberсlub) createTables() {
	tables := make(map[int]table)
	for i := 1; i <= c.TableAmount; i++ {
		tables[i] = table{
			revenue:       0,
			startRentTime: time.Time{},
			rentDuration:  time.Time{},
			isBusy:        false,
			usedBy:        player{},
		}
	}
	c.tables = tables
}

func (c *Cyberсlub) createQueue() {
	var queue []player
	c.queue = queue
}

func (c *Cyberсlub) RunParsing() error {
	c.createTables()
	c.createQueue()

	fmt.Println(c.WorkingStart.Format(timeLayout))
	for c.Scanner.Scan() {
		eventText := c.Scanner.Text()
		fmt.Println(eventText)

		// Получаем время и id ивента
		eventSplit := strings.Split(eventText, " ")
		eventTime, err := time.Parse(timeLayout, eventSplit[0])
		if err != nil {
			return err
		}
		eventID, err := strconv.Atoi(eventSplit[1])
		user := player{name: eventSplit[2]}
		if err != nil {
			return err
		}

		switch eventID {
		case incomingClientEntered:
			if c.isPlayerInClub(user.name) {
				fmt.Println(eventSplit[0], clientAlreadyInClub)
				break
			}

			if !eventTime.After(c.WorkingStart) || !eventTime.Before(c.WorkingEnd) {
				fmt.Println(eventSplit[0], clientComeAtWrongTime)
				break
			}
		case incomingClientTakeSeat:
			tableIndex, err := strconv.Atoi(eventSplit[3])
			if err != nil {
				return err
			}

			free, err := c.isTableFree(tableIndex)
			if err != nil {
				return err
			}

			if !free {
				fmt.Println(eventSplit[0], placeIsBusy)
				break
			}

			pastTableIndex, err := c.getPlayerTable(user.name)
			if err != nil {
				// пользователь занимает стол первый раз, а не меняет его
				err := user.takeTableInClub(c, tableIndex, eventTime)
				if err != nil {
					return err
				}
			} else {
				// пользователь меняет стол
				err := user.clearTableInClub(c, pastTableIndex, eventTime)
				if err != nil {
					return err
				}

				err = user.takeTableInClub(c, tableIndex, eventTime)
				if err != nil {
					return err
				}
			}
		case incomingClientWaiting:
			haveFreeTable, err := c.haveFreeTable()
			if err != nil {
				return err
			}

			if haveFreeTable {
				fmt.Println(eventSplit[0], clubHaveFreeTable)
				break
			}

			if c.isQueueFull() {
				fmt.Println(eventSplit[0], outcomingClientLeft, user.name)
				break
			}
			user.takeQueueInClub(c)
		case incomingClientLeft:
			if !c.isPlayerInClub(user.name) {
				fmt.Println(eventSplit[0], clientUnknown)
				break
			}

			playerTable, err := c.getPlayerTable(user.name)
			if err != nil {
				return err
			}

			err = user.clearTableInClub(c, playerTable, eventTime)
			if err != nil {
				return err
			}

			if len(c.queue) > 0 {
				newUser := c.queue[0]
				err = newUser.takeTableInClub(c, playerTable, eventTime)
				if err != nil {
					return err
				}
				c.popQueue()
				fmt.Println(eventSplit[0], outcomingClientTakeSeat, newUser.name, playerTable)

			}
		}
	}

	err := c.close(c.WorkingEnd)
	if err != nil {
		return err
	}

	c.calculateAllRevenue()

	// Проверяем наличие ошибок при сканировании файла
	if err := c.Scanner.Err(); err != nil {
		return fmt.Errorf("ошибка при сканировании файла: %v", err)
	}
	return nil
}

func (c *Cyberсlub) isPlayerInClub(name string) bool {
	for _, t := range c.tables {
		if t.isBusy {
			if t.usedBy.name == name {
				return true
			}
		}
	}
	return false
}

func (c *Cyberсlub) getPlayerTable(name string) (int, error) {
	if !c.isPlayerInClub(name) {
		return 0, fmt.Errorf("UserNotInClub")
	}
	for i, t := range c.tables {
		if t.isBusy {
			if t.usedBy.name == name {
				return i, nil
			}
		}
	}
	return 0, fmt.Errorf("TableNotFound")
}

func (c *Cyberсlub) isTableFree(tableIndex int) (bool, error) {
	if tableIndex > c.TableAmount {
		return false, fmt.Errorf("TableIndexOutOfRange")
	}

	return !c.tables[tableIndex].isBusy, nil
}

func (c *Cyberсlub) haveFreeTable() (bool, error) {
	for i := 1; i < c.TableAmount; i++ {
		free, err := c.isTableFree(i)
		if err != nil {
			return false, err
		}
		if free {
			return true, nil
		}
	}
	return false, nil
}

func (c *Cyberсlub) isQueueFull() bool {
	return len(c.queue) == c.TableAmount
}

func (c *Cyberсlub) popQueue() {
	newQueue := c.queue[1:]
	c.queue = newQueue
}

func (c *Cyberсlub) close(eventTime time.Time) error {
	var names []string
	for i := 1; i <= c.TableAmount; i++ {
		free, err := c.isTableFree(i)
		if err != nil {
			return err
		}
		if !free {
			user := c.tables[i].usedBy
			err := user.clearTableInClub(c, i, eventTime)
			if err != nil {
				return err
			}
			names = append(names, user.name)
		}
	}
	sort.Strings(names)
	for _, v := range names {
		fmt.Println(eventTime.Format(timeLayout), outcomingClientLeft, v)
	}
	fmt.Println(eventTime.Format(timeLayout))
	return nil
}

func (c *Cyberсlub) calculateAllRevenue() {
	for i := 1; i <= c.TableAmount; i++ {
		fmt.Println(i, c.tables[i].revenue, c.tables[i].rentDuration.Format(timeLayout))
	}
}
