package stats

import (
	"monitron-client/utils"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type WebSocketParams struct {
	Interval *time.Duration `params:"interval"`
}

func HandleCPUUsageWS(interval ...time.Duration) fiber.Handler {
	return HandleStatWS(GetCachedCPUUsage, interval...)
}
func HandleMemoryUsageWS(interval ...time.Duration) fiber.Handler {
	return HandleStatWS(GetCachedMemoryUsage, interval...)
}
func HandleDiskUsageWS(interval ...time.Duration) fiber.Handler {
	return HandleStatWS(GetCachedDiskUsage, interval...)
}
func HandleNetworkUsageWS(interval ...time.Duration) fiber.Handler {
	return HandleStatWS(GetCachedNetworkUsage, interval...)
}

func HandleStatWS[T any](job GetCacheFunc[T], interval ...time.Duration) fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		if val, err := time.ParseDuration(c.Query("interval")); err == nil && len(interval) == 0 {
			interval = []time.Duration{val}
		} else if err != nil || len(interval) == 0 {
			interval = []time.Duration{1 * time.Second}
		}

		ticker := time.NewTicker(interval[0])
		defer ticker.Stop()
		defer c.Close()

		for range ticker.C {
			stat := job()
			if stat == nil {
				stat = new(T)
			}
			data, _ := json.Marshal(utils.SuccessResponse(stat))

			err := c.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				ticker.Stop()
				c.Close()
				break
			}
		}
	})
}
