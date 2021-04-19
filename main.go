package main

import (
	"errors"
	"fmt"
)

var countDay, calculatedUnits  int = 0, 24 * 60
var units                      string = "минут"

func main() {
	calculate, err := calculateUnits(countDay, calculatedUnits)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%d день, содержит %d %s", countDay, calculate, units)
}

func calculateUnits(day int, unit int) (int, error) {
	if day <= 0 {
		return 0, errors.New("количество дней не может быть нулевым или отрицательным")
	}
	return day * unit, nil
}


