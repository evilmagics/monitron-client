package stats

import (
	"monitron-client/utils"
	"sync"

	"github.com/gofiber/fiber/v2"
)

var (
	HandleMemory  = handleCachedStatApi(GetCachedMemory)
	HandleCPU     = handleCachedStatApi(GetCachedCPU)
	HandleDisk    = handleCachedStatApi(GetCachedDisk)
	HandleNetwork = handleCachedStatApi(GetCachedNetwork)
)

func StatJob[T any](wg *sync.WaitGroup, job StatFunc[T], dst *T) {
	defer wg.Done()
	result, _ := job()
	*dst = result
}

func HandleAPI(c *fiber.Ctx) error {
	wg := new(sync.WaitGroup)
	wg.Add(5)

	stats := new(Stats)

	go StatJob(wg, StatHost, &stats.Host)
	go StatJob(wg, StatCPU, &stats.CPU)
	go StatJob(wg, StatMemory, &stats.Memory)
	go StatJob(wg, StatDisk, &stats.Disk)
	go StatJob(wg, StatNetwork, &stats.Network)
	wg.Wait()

	c.Status(fiber.StatusOK).JSON(utils.SuccessResponse(stats))
	return nil
}

func handleCachedStatApi[T any](fn GetCacheFunc[T]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		stat := fn()
		if fn == nil {
			stat = new(T)
		}
		c.Status(fiber.StatusOK).JSON(utils.SuccessResponse(stat))
		return nil
	}
}

func HandleHost(c *fiber.Ctx) error {
	stat, err := StatHost()
	if err != nil {
		return err
	}
	c.Status(fiber.StatusOK).JSON(utils.SuccessResponse(stat))
	return nil
}
