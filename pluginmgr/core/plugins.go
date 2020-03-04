package core

import (
	"github.com/iotaledger/goshimmer/plugins/autopeering"
	"github.com/iotaledger/goshimmer/plugins/banner"
	"github.com/iotaledger/goshimmer/plugins/cli"
	"github.com/iotaledger/goshimmer/plugins/config"
	"github.com/iotaledger/goshimmer/plugins/gossip"
	"github.com/iotaledger/goshimmer/plugins/gracefulshutdown"
	"github.com/iotaledger/goshimmer/plugins/logger"
	"github.com/iotaledger/goshimmer/plugins/metrics"
	"github.com/iotaledger/goshimmer/plugins/portcheck"
	"github.com/iotaledger/goshimmer/plugins/tangle"
	"github.com/iotaledger/hive.go/node"
)

var PLUGINS = node.Plugins(
	banner.PLUGIN,
	config.PLUGIN,
	logger.PLUGIN,
	cli.PLUGIN,
	portcheck.PLUGIN,
	autopeering.PLUGIN,
	tangle.PLUGIN,
	gossip.PLUGIN,
	gracefulshutdown.PLUGIN,
	metrics.PLUGIN,
)