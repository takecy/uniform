package aws

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestListAddresses(t *testing.T) {
	t.SkipNow()

	Convey("Given aws environment", t, func() {
		key := "your key"
		secret := "your secret"
		region := "ap-northeast-1"
		tags := map[string]string{
			"Environment": "Development",
			"Type":        "API",
		}
		ai := NewAWS(key, secret, region, tags)

		Convey("When call", func() {
			ads, err := ai.ListAddresses()

			Convey("Then return ads", func() {
				So(err, ShouldBeNil)
				So(len(ads), ShouldBeGreaterThan, 0)
			})
		})
	})
}
