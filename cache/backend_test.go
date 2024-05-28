package cache

import (
	"os"
	"testing"

	"github.com/filiquid/listener/config"
	"github.com/filiquid/listener/dao"
)

var testcache *CacheData

func TestMain(m *testing.M) {
	db, err := dao.NewDao(&config.DbConfig{
		User:     "tliquid",
		Password: "123456",
		Url:      "127.0.0.1:3306",
		DbName:   "liquid",
	})
	if err != nil {
		panic(err)
	}
	testcache = NewCacheData(db, 10000)

	os.Exit(m.Run())
}

// go test -count=1 -v github.com/filiquid/listener/cache -run TestFetchAndSaveBasicSeniorData
func TestFetchAndSaveBasicSeniorData(t *testing.T) {
	if err := testcache.FetchAndSaveBasicSeniorData(); err != nil {
		t.Fatal(err)
	}
	// basic, senior, panel := testcache.GetBasicSeniorPanel()
	// t.Log(string(basic))
	// t.Log(string(senior))
	// t.Log(string(panel))
}
