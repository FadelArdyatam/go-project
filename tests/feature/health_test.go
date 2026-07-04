package feature

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"cloud-compute/tests"
)

type HealthTestSuite struct {
	suite.Suite
	tests.TestCase
}

func TestHealthTestSuite(t *testing.T) {
	suite.Run(t, new(HealthTestSuite))
}

func (s *HealthTestSuite) TestHealth() {
	response, err := s.Http(s.T()).Get("/health")

	s.Require().NoError(err)
	response.AssertOk().AssertExactJson(map[string]any{
		"status": "healthy",
	})
}
