package bruteforcer

import (
	"fmt"
)

type Stats struct {
	From      Movers
	To        Movers
	Time      int
	moveCount int
	MovesLog  []movesLog
}

func (stats *Stats) String() string {

	fmt.Println(stats.MovesLog)
	time := fmt.Sprintf("time: {%v}", stats.Time)
	from := fmt.Sprintf("from: {%v}", stats.From)
	to := fmt.Sprintf("to: {%v}", stats.To)
	log := fmt.Sprintf("log: {%v}", stats.MovesLog)
	return fmt.Sprintf("{%v, %v, %v, %v}", time, from, to, log)
}

func (stats *Stats) move(movers Movers, isForward bool) {
	if isForward {
		stats.moveForward(movers)
	} else {
		stats.moveBackward(movers)
	}
	stats.Time += movers.getTime()
	stats.moveCount += 1

	stats.MovesLog = append(stats.MovesLog, movesLog{isForward, movers})
}

func (stats *Stats) moveForward(movers Movers) {

	for _, mover := range movers {
		stats.From = stats.From.remove(mover)
		stats.To = stats.To.add(mover)
	}
}

func (stats *Stats) moveBackward(movers Movers) {
	for _, mover := range movers {
		stats.To = stats.To.remove(mover)
		stats.From = stats.From.add(mover)
	}
}

func (stats *Stats) isNeedNewMove() bool {
	if stats.moveCount >= MaxMoveCount || stats.Time >= MaxTime {
		return false
	}
	return true
}

func (stats *Stats) isFinished() bool {
	if len(stats.To) == 4 && len(stats.From) == 0 {
		return true
	}

	return false
}
