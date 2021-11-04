package list

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestList_AddNodeHead(t *testing.T) {
	Convey("add node to head", t, func() {
		l := List{}
		Convey("add one node to head", func() {
			l.AddNodeHead(1)
			So(l.GetLen(), ShouldEqual, 1)
			So(l.head.value, ShouldEqual, 1)
			So(l.tail.value, ShouldEqual, 1)
		})
		Convey("add tow node to head", func() {
			l.AddNodeHead(1)
			l.AddNodeHead(2)
			So(l.GetLen(), ShouldEqual, 2)
			So(l.head.value, ShouldEqual, 2)
			So(l.tail.value, ShouldEqual, 1)
		})
	})
}

func TestList_AddNodeTail(t *testing.T) {
	Convey("add node to tail", t, func() {
		l := List{}
		Convey("add one node to tail", func() {
			l.AddNodeTail(1)
			So(l.GetLen(), ShouldEqual, 1)
			So(l.head.value, ShouldEqual, 1)
			So(l.tail.value, ShouldEqual, 1)
		})
		Convey("add tow node to tail", func() {
			l.AddNodeTail(1)
			l.AddNodeTail(2)
			So(l.GetLen(), ShouldEqual, 2)
			So(l.head.value, ShouldEqual, 1)
			So(l.tail.value, ShouldEqual, 2)
		})
	})

}

func TestList_DelNode(t *testing.T) {
	Convey("del node", t, func() {
		l := List{}
		l.AddNodeTail(1)
		Convey("del nil", func() {
			l.DelNode(nil)
			So(l.GetLen(), ShouldEqual, 1)
		})
		Convey("del one", func() {
			node := l.GetNode(0)
			l.DelNode(node)
			So(l.GetLen(), ShouldEqual, 0)
		})
	})

}

func TestList_GetIdx(t *testing.T) {
	Convey("get val by index", t, func() {
		l := List{}
		l.AddNodeTail(1)
		l.AddNodeTail(2)
		l.AddNodeTail(3)

		Convey("get val index 0", func() {
			So(l.GetIdx(0), ShouldEqual, 1)
		})
		Convey("get val index 1", func() {
			So(l.GetIdx(1), ShouldEqual, 2)
		})
		Convey("get val index 2", func() {
			So(l.GetIdx(2), ShouldEqual, 3)
		})
	})
}

func TestList_GetLen(t *testing.T) {

	list := List{}
	Convey("get list node size", t, func() {
		Convey("list is empty", func() {
			So(list.GetLen(), ShouldEqual, 0)
		})
		Convey("list is not empty", func() {
			list.AddNodeHead(1)
			list.AddNodeHead(2)
			list.AddNodeHead(3)
			So(list.GetLen(), ShouldEqual, 3)
		})
	})
}

func TestList_GetNode(t *testing.T) {
	Convey("get node", t, func() {
		list := List{}
		list.AddNodeTail(1)
		list.AddNodeTail(2)
		list.AddNodeTail(3)
		Convey("get node by index 1", func() {
			So(list.GetNode(0).GetValue(), ShouldEqual, 1)
		})
		Convey("get node by index 2", func() {
			So(list.GetNode(1).GetValue(), ShouldEqual, 2)
		})
		Convey("get node by index 3", func() {
			So(list.GetNode(2).GetValue(), ShouldEqual, 3)
		})
	})
}
