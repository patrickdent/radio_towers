package main

import "errors"

func main() {
}

// assume that the houses are in order
func planTowers(houses []int, towerRange int) int {
	towers := []tower{}

	towers = append(towers, getNextTower(0, towerRange, houses))

	towers = planRemainingTowers(towers, houses, towerRange)

	return len(towers)
}

func planRemainingTowers(towers []tower, houses []int, towerRange int) []tower {
	currentTower := towers[len(towers)-1]
	nextIndex, err := findIndexOfNextHouseOutOfRange(currentTower.end, currentTower.index, houses)

	if err != nil {
		return towers
	}

	towers = append(towers, getNextTower(nextIndex, towerRange, houses))

	return planRemainingTowers(towers, houses, towerRange)
}

func getNextTower(startIndex, towerRange int, houses []int) tower {
	startLocation := houses[startIndex]
	indexOfHouseAtMaxRange := findIndexOfHouseAtMaxRange(startLocation, startIndex, towerRange, houses)

	return tower{
		start: startLocation,
		end:   houses[indexOfHouseAtMaxRange] + towerRange,
		index: indexOfHouseAtMaxRange,
	}
}

func findIndexOfHouseAtMaxRange(startLocation, previousIndex, towerRange int, houses []int) int {
	currentIndex := previousIndex + 1

	if currentIndex == len(houses) {
		return previousIndex
	}

	if houses[currentIndex]-startLocation <= towerRange {
		return findIndexOfHouseAtMaxRange(startLocation, currentIndex, towerRange, houses)
	}

	return previousIndex
}

func findIndexOfNextHouseOutOfRange(endOfRange, previousIndex int, houses []int) (int, error) {
	if previousIndex == len(houses)-1 {
		return 0, errors.New("you're covered")
	}

	currentIndex := previousIndex + 1
	if endOfRange < houses[currentIndex] {
		return currentIndex, nil
	}

	return findIndexOfNextHouseOutOfRange(endOfRange, currentIndex, houses)
}

// this describe a tower, with range reaching from start to end
type tower struct {
	start, end int
	index      int
}
