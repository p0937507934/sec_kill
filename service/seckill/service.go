package sec_kills

import (
	"context"
	"sec_kill/driver"
	"sec_kill/dto"
	"sec_kill/proto/proto"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
)

type SecKillService struct {
	redis   *redis.Client
	luaHash string
}

var luaScripts = `
local value = redis.call("Get", KEYS[1])
print("当前值为 " .. value);
if( value - KEYS[2] >= 0 ) then
	local leftStock = redis.call("DecrBy" , KEYS[1],KEYS[2])
	print("剩余值为" .. leftStock );
	return leftStock
else
	print("数量不够，无法扣减");
	return value - KEYS[2]
end
return -1
	`

func NewSecKillService() *SecKillService {
	redis := driver.InitRedis()
	s, _ := redis.ScriptLoad(context.Background(), luaScripts).Result()
	return &SecKillService{
		redis:   redis,
		luaHash: s,
	}
}

func (s *SecKillService) SetActivity(req *dto.SetActivity) error {
	ctx := context.Background()
	result := s.redis.Set(ctx, req.Name, req.Stock, 0)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func (s *SecKillService) SecKill(req *dto.Seckill) error {
	ctx := context.Background()
	n, err := s.redis.EvalSha(ctx, s.luaHash, []string{req.Name, "1"}).Result()
	if err != nil {
		return err
	}
	if n.(int64) >= 0 {
		// grpc call stock service
		conn, err := grpc.Dial(driver.StockServiceAddr, grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer conn.Close()
		grpcClient := proto.NewStockServiceClient(conn)
		_, err = grpcClient.UpdateStock(context.Background(), &proto.UpdateStockRequest{
			ProductId: 1,
			Stock:     0,
			Lock:      1,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
