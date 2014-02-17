package learning_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLearning(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Learning Suite")
}
