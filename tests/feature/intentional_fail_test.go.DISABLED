package feature

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"cloud-compute/tests"
)

// ponytail: kept disabled on purpose — rename to .go (drop .DISABLED) for
// Tahap 9 to demonstrate a controlled CI failure, then delete it again after
// showing the failing pipeline log.
type IntentionalFailTestSuite struct {
	suite.Suite
	tests.TestCase
}

func TestIntentionalFailTestSuite(t *testing.T) {
	suite.Run(t, new(IntentionalFailTestSuite))
}

func (s *IntentionalFailTestSuite) TestBroken() {
	s.Equal(3, 1+1)
}
