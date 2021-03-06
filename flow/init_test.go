package flow

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/vito/spiff/yaml"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Flowing")
}

func parseYAML(source string) yaml.Node {
	parsed, err := yaml.Parse([]byte(source))
	if err != nil {
		panic(err)
	}

	return parsed
}
