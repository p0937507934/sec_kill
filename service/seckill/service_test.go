package sec_kills

import (
	"sec_kill/driver"
	"sec_kill/dto"
	"sync"
	"testing"
)

func TestSecKill(t *testing.T) {
	srv := NewSecKillService()
	driver.StockServiceAddr = "localhost:8001"
	req := &dto.Seckill{Name: "iphone13"}
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		defer wg.Done()
		go srv.SecKill(req)
	}
	wg.Wait()
}
