package list

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNode_GetNext(t *testing.T) {
	Convey("get next", t, func() {
		Convey("dont has next", func() {
			node := NewNode(1)
			So(node.GetNext(), ShouldEqual, nil)
		})

		Convey("get next has value", func() {
			node := NewNode(1)
			node.next = NewNode(2)
			So(node.GetNext(), ShouldEqual, node.next)
		})
	})
}

func TestNode_GetValue(t *testing.T) {

	Convey("get value", t, func() {
		Convey("dont has value", func() {
			node := NewNode(1)
			So(node.GetValue(), ShouldEqual, 1)
		})

		Convey("value is nil", func() {
			node := NewNode(nil)
			So(node.GetValue(), ShouldEqual, nil)
		})
	})
}
