package tests

import (
	"github.com/goravel/framework/testing"

	"cloud-compute/bootstrap"
)

func init() {
	bootstrap.Boot()
}

type TestCase struct {
	testing.TestCase
}
