package miniredis

import (
	"testing"
	"time"

	"github.com/s-barr-fetch/miniredis/proto"
)

// Test OBJECT IDLETIME.
func TestObjectIdletime(t *testing.T) {
	s, err := Run()
	ok(t, err)
	defer s.Close()
	c, err := proto.Dial(s.Addr())
	ok(t, err)
	defer c.Close()

	{
		start := time.Now()
		s.SetTime(start)

		mustOK(t, c,
			"SET", "foo", "bar",
		)

		mustDo(t, c,
			"OBJECT", "IDLETIME", "foo",
			proto.Int(0),
		)

		s.SetTime(start.Add(time.Minute))
		mustDo(t, c,
			"OBJECT", "IDLETIME", "foo",
			proto.Int(60),
		)

		s.Get("foo")
		mustDo(t, c,
			"object", "idletime", "foo",
			proto.Int(0),
		)

		s.Del("foo")
		mustDo(t, c,
			"OBJECT", "IDLETIME", "foo",
			proto.Nil,
		)
	}
}
