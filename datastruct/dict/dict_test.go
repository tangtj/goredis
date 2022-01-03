package dict

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDict_Find(t *testing.T) {
	dict := NewDict(100)
	dict.Add("a", 1)

	val, _ := dict.Find("a")
	print(val)

}

func TestNewDict(t *testing.T) {
	convey.Convey("test create Dict", t, func() {
		convey.Convey("cap 1", func() {
			d := NewDict(1)
			convey.So(len(d.shards), convey.ShouldEqual, 128)
		})
		convey.Convey("cap 255", func() {
			d := NewDict(255)
			convey.So(len(d.shards), convey.ShouldEqual, 256)
		})
	})
}

func TestDict_Add(t *testing.T) {
	convey.Convey("test add item", t, func() {
		d := NewDict(255)
		convey.Convey("add one", func() {
			d.Add("a", 1)

			v, has := d.Find("a")
			convey.So(has, convey.ShouldBeTrue)
			convey.So(v, convey.ShouldEqual, 1)
		})
	})
}

func TestDict_Del(t *testing.T) {
	convey.Convey("test del item", t, func() {
		d := NewDict(100)

		convey.Convey("del no exits key", func() {
			convey.So(d.Del("a"), convey.ShouldBeFalse)
		})

		convey.Convey("del exits key", func() {
			d.Add("b", "1")
			convey.So(d.Del("b"), convey.ShouldBeTrue)
		})
	})
}
