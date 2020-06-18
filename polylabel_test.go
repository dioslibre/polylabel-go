package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	t.Errorf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}

func loadData(filename string) (polygon [][][]float64) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		panic("failed to open json file")
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &polygon)
	if err != nil {
		panic("failed to parse json file")
	}

	return polygon
}

func TestPolylabelWater1(t *testing.T) {
	polygon := loadData("test_data/water1.json")
	var x, y float64

	x, y = polylabel(polygon, 1.0)
	AssertEqual(t, x, 3865.85009765625)
	AssertEqual(t, y, 2124.87841796875)

	x, y = polylabel(polygon, 50.0)
	AssertEqual(t, x, 3854.296875)
	AssertEqual(t, y, 2123.828125)
}

func TestPolylabelWater2(t *testing.T) {
	polygon := loadData("test_data/water2.json")

	x, y := polylabel(polygon, 1.0)
	AssertEqual(t, x, 3263.5)
	AssertEqual(t, y, 3263.5)
}

func TestDegeneratePolygons(t *testing.T) {
	var x, y float64

	polygon := [][][]float64{{{0, 0}, {1, 0}, {2, 0}, {0, 0}}}
	x, y = polylabel(polygon, 1.0)
	AssertEqual(t, x, 0.0)
	AssertEqual(t, y, 0.0)

	polygon = [][][]float64{{{0, 0}, {1, 0}, {1, 1}, {1, 0}, {0, 0}}}
	x, y = polylabel(polygon, 1.0)
	AssertEqual(t, x, 0.0)
	AssertEqual(t, y, 0.0)
}
