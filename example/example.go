package main

import "github.com/tobiashort/clog-go"

func main() {
	clog.Level = clog.LevelDebug
	clog.Debug("Hey")
	clog.Debugf("Hey %s", "you")
	clog.Debugs("What", "foo", true, "bar", true)
	clog.Info("Hey")
	clog.Infof("Hey %s", "you")
	clog.Infos("What", "foo", true, "bar", true)
	clog.Warn("Hey")
	clog.Warnf("Hey %s", "you")
	clog.Warns("What", "foo", true, "bar", true)
	clog.Error("Hey")
	clog.Errorf("Hey %s", "you")
	clog.Errors("What", "foo", true, "bar", true)
}
