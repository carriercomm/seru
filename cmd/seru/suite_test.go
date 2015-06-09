package main

import (
	"bytes"
	"github.com/megamsys/seru/cmd"
	"gopkg.in/check.v1"
	"os"
	"testing"
)

type S struct {
	recover []string
}

var _ = check.Suite(&S{})
var manager *cmd.Manager

func Test(t *testing.T) { check.TestingT(t) }

func (s *S) SetUpTest(c *check.C) {
	var stdout, stderr bytes.Buffer
	manager = cmd.NewManager("seru", version, header, &stdout, &stderr, os.Stdin)
}
