package smoke_test

import (
	"flag"
	"testing"

	. "github.com/onsi/gomega"
)

var (
	Builder string
)

func init() {
	flag.StringVar(&Builder, "name", "", "")
}

func TestSmoke(t *testing.T) {
	Expect := NewWithT(t).Expect

	flag.Parse()

	Expect(Builder).NotTo(Equal(""))

}
