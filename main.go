package main

import (
	"github.com/cosmos/relayer/v2/cmd"
	specyconfig "github.com/cosmos/relayer/v2/specy/config"
)

func init() {
	specyconfig.ReadSpecyConfig()
}

func main() {
	cmd.Execute()
}
