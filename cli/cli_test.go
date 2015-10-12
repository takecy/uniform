package cli

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAlive(t *testing.T) {
	Convey("Ginve ip addresses", t, func() {
		path := "/alive"
		prop := "version"
		port := 80
		cl := NewCli(path, prop, port)

		addresses := []string{
			"localhost",
		}

		Convey("When call", func() {
			i, err := cl.Alive(addresses)

			Convey("Then return ok", func() {
				So(err, ShouldBeNil)
				So(len(i), ShouldEqual, 1)
			})
		})
	})
}
