package list

import (
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func TestList_AddNodeHead(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		len  int
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *List
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				head: tt.fields.head,
				tail: tt.fields.tail,
				len:  tt.fields.len,
			}
			if got := l.AddNodeHead(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddNodeHead() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_AddNodeTail(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		len  int
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *List
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				head: tt.fields.head,
				tail: tt.fields.tail,
				len:  tt.fields.len,
			}
			if got := l.AddNodeTail(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddNodeTail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_DelNode(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		len  int
	}
	type args struct {
		node *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *List
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				head: tt.fields.head,
				tail: tt.fields.tail,
				len:  tt.fields.len,
			}
			if got := l.DelNode(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DelNode() = %v, want %v", got, tt.want)
			}
		})
	}
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
