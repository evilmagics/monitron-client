package routes

import (
	"monitron-client/internal/stats"
	"monitron-client/utils"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type WSTask[T any] func() (T, error)

func HandleCPUWS(interval ...time.Duration) fiber.Handler {
	return HandleStatWS(stats.StatCPU, interval...)
}
func HandleMemoryWS(interval ...time.Duration) fiber.Handler {
	return HandleStatWS(stats.StatMemory, interval...)
}
func HandleDiskWS(interval ...time.Duration) fiber.Handler {
	return HandleStatWS(stats.StatDisk, interval...)
}
func HandleNetworkWS(interval ...time.Duration) fiber.Handler {
	return HandleStatWS(stats.StatNetwork, interval...)
}

func HandleStatWS[T any](job WSTask[T], interval ...time.Duration) fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		if len(interval) == 0 {
			interval = []time.Duration{1 * time.Second}
		}
		for {
			stat, _ := job()
			data, _ := json.Marshal(utils.SuccessResponse(stat))

			err := c.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				break
			}
			time.Sleep(interval[0])
		}
	})
}
