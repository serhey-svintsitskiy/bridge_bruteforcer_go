package bruteforcer

import "fmt"

type movesLog struct {
	isForward bool
	movers    Movers
}

func (moveLog movesLog) String() string {
	direction := "<-"
	if moveLog.isForward {
		direction = "->"
	}
	var movers []string
	for _, mover := range moveLog.movers {
		movers = append(movers, fmt.Sprintf("%v", mover.Time))
	}

	return fmt.Sprintf("[%s, %v]", direction, movers)
}
