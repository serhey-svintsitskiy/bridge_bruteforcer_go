package bruteforcer

type Mover struct {
	Time int
}

type Movers []Mover

func (movers Movers) getTime() int {
	maxTime := 0
	for _, mover := range movers {
		if mover.Time > maxTime {
			maxTime = mover.Time
		}
	}

	return maxTime
}

func (movers Movers) remove(mover Mover) Movers {
	var newMovers Movers
	for _, existingMover := range movers {
		if existingMover.Time == mover.Time {
			continue
		}
		newMovers = append(newMovers, existingMover)
	}

	return newMovers
}

func (movers Movers) add(mover Mover) Movers {
	return append(movers, mover)
}
