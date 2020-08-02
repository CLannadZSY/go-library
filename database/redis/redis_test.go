package redis

import "testing"

func TestConnectRedis(t *testing.T) {
	pool := ConnectRedis("localhost", "", "")
	t.Logf("%v", pool)
	r1 := pool.Get()
	defer r1.Close()
	res, err := r1.Do("GET", "test")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s", res)
}
