package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIt(t *testing.T) {
	Convey("It returns the number of Towers", t, func() {
		houses := []int{1, 2, 3, 4, 5}
		towerRange := 1
		required := planTowers(houses, towerRange)
		So(required, ShouldEqual, 2)
	})

	Convey("It returns the number of Towers", t, func() {
		houses := []int{1, 3, 5, 7, 9}
		towerRange := 1
		required := planTowers(houses, towerRange)
		So(required, ShouldEqual, 5)
	})

	Convey("It returns the number of Towers", t, func() {
		houses := []int{1, 2, 3, 7, 8, 10}
		towerRange := 1
		required := planTowers(houses, towerRange)
		So(required, ShouldEqual, 3)
	})

	Convey("It returns the number of Towers", t, func() {
		houses := []int{1, 2, 3, 5, 9}
		towerRange := 2
		required := planTowers(houses, towerRange)
		So(required, ShouldEqual, 2)
	})
}
