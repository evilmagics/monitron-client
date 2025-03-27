package tests

import (
	"testing"

	"github.com/shirou/gopsutil/v4/disk"
)

func TestStatCollection(t *testing.T) {
	c, err := disk.IOCounters()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(c)
	t.Fail()
}
