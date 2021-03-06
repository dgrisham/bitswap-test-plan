package main

import (
	test "github.com/ipfs/test-plans/bitswap-tuning/test"
	"github.com/testground/sdk-go/run"
	"github.com/testground/sdk-go/runtime"
)

func main() {
	run.Invoke(runf)
}

func runf(runenv *runtime.RunEnv) error {
	switch c := runenv.TestCase; c {
	case "transfer":
		return test.Transfer(runenv)
	case "fuzz":
		return test.Fuzz(runenv)
	default:
		panic("unrecognized test case")
	}
}
