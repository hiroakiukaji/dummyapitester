package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestApiTest(t *testing.T) {
	Convey("Provide integer", t, func() {
		Convey("It should return json", func() {
			So(apitest(10), ShouldEqual, "{\"root\":10,\"square\":100}")
		})
	})
}
