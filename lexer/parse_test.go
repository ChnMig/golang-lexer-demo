package lexer

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestParse(t *testing.T) {
	convey.Convey("TestParse", t, func() {
		l, err := Parse("1 + 2 - 3 ++ (4 * 5 )")
		convey.So(err, convey.ShouldBeNil)
		fmt.Println(l.Tokens)
	})
}
