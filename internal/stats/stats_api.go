package stats

import (
	"monitron-client/utils"
	"sync"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type Stats struct {
	Host    HostStat    `json:"host"`
	CPU     CPUStat     `json:"cpu"`
	Memory  MemoryStat  `json:"memory"`
	Disk    DiskStat    `json:"disk"`
	Network NetworkStat `json:"network"`
}

func (s Stats) String() string {
	str, _ := json.Marshal(s)
	return string(str)
}

func StatJob[T any](wg *sync.WaitGroup, job func() (T, error), dst *T) {
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

func HandleMemory(c *fiber.Ctx) error {
	stat, err := StatMemory()
	if err != nil {
		return err
	}
	c.Status(fiber.StatusOK).JSON(utils.SuccessResponse(stat))
	return nil
}

func HandleCPU(c *fiber.Ctx) error {
	stat, err := StatCPU()
	if err != nil {
		return err
	}
	c.Status(fiber.StatusOK).JSON(utils.SuccessResponse(stat))
	return nil
}

func HandleHost(c *fiber.Ctx) error {
	stat, err := StatHost()
	if err != nil {
		return err
	}
	c.Status(fiber.StatusOK).JSON(utils.SuccessResponse(stat))
	return nil
}
func HandleDisk(c *fiber.Ctx) error {
	stat, err := StatDisk()
	if err != nil {
		return err
	}

	c.Status(fiber.StatusOK).JSON(utils.SuccessResponse(stat))

	return nil
}
func HandleNetwork(c *fiber.Ctx) error {
	stat, err := StatNetwork()
	if err != nil {
		return err
	}

	c.Status(fiber.StatusOK).JSON(utils.SuccessResponse(stat))

	return nil
}
