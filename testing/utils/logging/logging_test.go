package logging

import (
	"hanson/utils/logging"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInitLog(t *testing.T) {
	Convey("Initialize Log", t, func() {
		So(logging.InitLog(), ShouldBeTrue)
	})
}
