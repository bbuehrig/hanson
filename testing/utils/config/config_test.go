package config

import (
	"hanson/utils/config"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReadConfig(t *testing.T) {
	Convey("Read Config", t, func() {
		So(config.ReadConfig(), ShouldBeTrue)
	})
}

func TestGetStringValue(t *testing.T) {
	Convey("Get an existing String-Value from Config", t, func() {
		So(config.GetStringValue("TestingString"), ShouldNotBeBlank)
	})

	Convey("Config-Value is equal to 'yes'", t, func() {
		So(config.GetStringValue("TestingString"), ShouldEqual, "yes")
	})

	Convey("Get a not existing Config-Value", t, func() {
		So(config.GetStringValue("NotExisting"), ShouldBeBlank)
	})
}

func TestGetBoolValue(t *testing.T) {
	Convey("Get an existing Bool-Value from Config", t, func() {
		So(config.GetBoolValue("TestingBool"), ShouldBeTrue)
	})

	Convey("Get a not existing Bool-Value from Config", t, func() {
		So(config.GetBoolValue("NotExisting"), ShouldBeFalse)
	})
}
