package set

import (
	"testing"

	"github.com/smarty/assertions/should"
)

func Test(t *testing.T) {
	s := New(1, 2, 3, 3)
	should.So(t, s, should.Equal, New(1, 2, 3))
	s.Add(4)
	should.So(t, s, should.Equal, New(1, 2, 3, 4))
	should.So(t, s.Contains(4), should.BeTrue)
	should.So(t, s.Contains(5), should.BeFalse)
}
