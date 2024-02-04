package lexer

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestLexer_CreateList(t *testing.T) {
	convey.Convey("TestParse", t, func() {
		l, err := Parse("1 + 2 - 3 + (4 * 5 )")
		convey.So(err, convey.ShouldBeNil)
		err = l.CreateList()
		convey.So(err, convey.ShouldBeNil)
		for e := l.List.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value)
		}
	})
}

func TestLexer_ConstructAst(t *testing.T) {
	convey.Convey("TestParse", t, func() {
		l, err := Parse("1 + 2 - 3 + (4 * 5 )")
		convey.So(err, convey.ShouldBeNil)
		err = l.CreateList()
		convey.So(err, convey.ShouldBeNil)
		err = l.ConstructAst()
		convey.So(err, convey.ShouldBeNil)
		data, _ := json.Marshal(l.Node)
		fmt.Println(string(data))
	})
}

func TestLexer_Interpreter(t *testing.T) {
	convey.Convey("TestParse", t, func() {
		l, err := Parse("1 + 2 - 3 + (4 * (5 -2) ) - 1")
		convey.So(err, convey.ShouldBeNil)
		err = l.CreateList()
		convey.So(err, convey.ShouldBeNil)
		err = l.ConstructAst()
		convey.So(err, convey.ShouldBeNil)
		data, _ := json.Marshal(l.Node)
		fmt.Println(string(data))
		result, err := l.Interpreter()
		convey.So(err, convey.ShouldBeNil)
		fmt.Println(result)
	})
}
