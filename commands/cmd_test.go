package commands

import (
	"testing"
	"testsort/05cache/simplekv/models"
)

func TestRun(t *testing.T) {
	models.InitCache()
	Run()
}
