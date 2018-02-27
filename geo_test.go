package georadis

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// TODO: mock redis

func TestGeo(t *testing.T) {
	Convey("test geo basic functions", t, func() {
		pool := NewPool()
		gc := NewGeo(pool)

		k := "yoyo"
		data := []*MetaData{
			NewMetaData("a1", 23.1, 100.7),
			NewMetaData("a2", 23.9, 100.8),
		}
		err := gc.Set(k, data)
		So(err, ShouldBeNil)

		actualData, err := gc.Get(k, []string{"a1", "a2", "b1"})
		So(err, ShouldBeNil)
		So(len(actualData), ShouldEqual, 3)
		So(actualData[2], ShouldBeNil)

		loc := Coordinate{
			Lat: 23.09,
			Lon: 100.69,
		}
		results, err := gc.Neighbors(k, loc, 10, KM)
		So(err, ShouldBeNil)
		So(results, ShouldNotBeNil)
		So(len(results), ShouldEqual, 1)

		results, err = gc.Neighbors(k, loc, 100, KM)
		So(err, ShouldBeNil)
		So(results, ShouldNotBeNil)
		So(len(results), ShouldEqual, 2)

		results, err = gc.Neighbors(k, loc, 10, KM, WithHash, WithDist, WithCoord)
		So(err, ShouldBeNil)
		So(results, ShouldNotBeNil)
		So(len(results), ShouldEqual, 1)
		So(results[0].Name, ShouldEqual, "a1")
		So(results[0].Coord.Lat, ShouldAlmostEqual, 23.1, .001)
		So(results[0].Coord.Lon, ShouldAlmostEqual, 100.7, .001)

	})
}