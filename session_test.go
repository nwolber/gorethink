package gorethink

import (
	"os"

	test "launchpad.net/gocheck"
)

func (s *RethinkSuite) TestSessionConnect(c *test.C) {
	session, err := Connect(ConnectOpts{
		Address:   url,
		AuthKey:   os.Getenv("RETHINKDB_AUTHKEY"),
		MaxIdle:   3,
		MaxActive: 3,
	})
	c.Assert(err, test.IsNil)

	row, err := Expr("Hello World").RunRow(session)
	c.Assert(err, test.IsNil)

	var response string
	err = row.Scan(&response)
	c.Assert(err, test.IsNil)
	c.Assert(response, test.Equals, "Hello World")
}

func (s *RethinkSuite) TestSessionConnectError(c *test.C) {
	var err error
	_, err = Connect(ConnectOpts{
		Address:   "nonexistanturl",
		MaxIdle:   3,
		MaxActive: 3,
	})
	c.Assert(err, test.NotNil)
}
