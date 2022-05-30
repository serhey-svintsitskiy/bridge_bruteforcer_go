package bruteforcer

import (
	"fmt"
	"gonum.org/v1/gonum/stat/combin"
	"sync"
)

const MaxTime = 25
const MaxMoveCount = 5

type Bruteforcer struct {
	results []Stats
	c       chan Stats
	sync.Mutex
	jobs int
}

func (bruteforcer *Bruteforcer) addResult(stats Stats) {
	bruteforcer.results = append(bruteforcer.results, stats)
}

func (bruteforcer *Bruteforcer) getComb(movers Movers, quantity int) []Movers {

	cs := combin.Combinations(len(movers), quantity)
	var combinationList []Movers
	for _, c := range cs {
		combination := Movers{}
		for _, index := range c {
			combination = append(combination, movers[index])
		}
		combinationList = append(combinationList, combination)
	}

	return combinationList
}

func (bruteforcer *Bruteforcer) tryMove(stats Stats, isForward bool, movers Movers) {
	defer func() {
		bruteforcer.jobs--
	}()
	stats.move(movers, isForward)

	if stats.isFinished() {
		bruteforcer.c <- stats
		//bruteforcer.addResult(stats)
		return
	}

	if stats.isNeedNewMove() {
		bruteforcer.multiply(stats, !isForward)
	}
}

func (bruteforcer *Bruteforcer) multiply(stats Stats, isForward bool) {
	possibleMovers := stats.To
	moversCount := 1
	if isForward {
		possibleMovers = stats.From
		moversCount = 2
	}

	combList := bruteforcer.getComb(possibleMovers, moversCount)
	for _, movers := range combList {
		bruteforcer.jobs++
		go bruteforcer.tryMove(stats, isForward, movers)
	}
}

func (bruteforcer *Bruteforcer) Bruteforce() {

	mover1 := Mover{1}
	mover2 := Mover{2}
	mover3 := Mover{5}
	mover4 := Mover{10}

	bruteforcer.c = make(chan Stats)

	stats := Stats{
		Movers{mover1, mover2, mover3, mover4},
		Movers{},
		0,
		0,
		[]movesLog{},
	}

	bruteforcer.multiply(stats, true)

	for {
		fmt.Println(<-bruteforcer.c)
		if bruteforcer.jobs == 0 {
			break
		}
	}
}

func (bruteforcer *Bruteforcer) GetBestResult() Stats {
	if len(bruteforcer.results) == 0 {
		panic("results are empty")
	}
	minTimeResult := bruteforcer.results[0]
	for _, result := range bruteforcer.results {
		if result.Time < minTimeResult.Time {
			minTimeResult = result
		}
	}

	return minTimeResult
}
