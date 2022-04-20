package api

import (
	"sec_kill/dto"
	sec_kills "sec_kill/service/seckill"

	"github.com/gin-gonic/gin"
)

func SetKill(c *gin.Context) {
	req := &dto.Seckill{}
	if err := c.Bind(req); err != nil {
		c.String(400, "error:"+err.Error())
		return
	}
	srv := sec_kills.NewSecKillService()
	err := srv.SecKill(req)
	if err != nil {
		c.String(400, "error:"+err.Error())
		return
	}
	c.JSON(200, gin.H{
		"msg": "success",
	})
}
