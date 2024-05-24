package dao

import (
	"os"
	"testing"

	"github.com/filiquid/listener/config"
)

var testdao *Dao

func TestMain(m *testing.M) {
	var err error
	testdao, err = NewDao(&config.DbConfig{
		User:     "tlistener",
		Password: "123456",
		Url:      "127.0.0.1:3306",
		DbName:   "filevents",
	})
	if err != nil {
		panic(err)
	}
	testdao.Migrate()
	os.Exit(m.Run())
}

// go test -v -count=1 github.com/filiquid/listener/dao -run TestHeight
func TestHeight(t *testing.T) {
	var heightList = []uint64{1, 1, 2, 2, 3}
	for _, src := range heightList {
		if err := testdao.UpdateLastHeight(src); err != nil {
			t.Fatal(err)
		}
		dest, err := testdao.GetLastHeight()
		if err != nil {
			t.Fatal(err)
		}
		if src != dest {
			t.Fatalf("expect %d, got %d", src, dest)
		}
	}
}
