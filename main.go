package main

import (
	"os/user"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/docker/go-plugins-helpers/authorization"
	"github.com/sajayantony/simple-auth-plugin/plugin"
)

const (
	pluginSocket = "/run/docker/plugins/simple-auth-plugin.sock"
)

func main() {
	logrus.Info("Starting Simple Auth Plugin")

	//logrus.SetLevel(logrus.DebugLevel)

	p, err := plugin.CreatePlugin()
	if err != nil {
		logrus.Fatal(err)
	}

	// Initialize the handler using go-plugin-helpers
	// You could also just implement the entire socket listener
	h := authorization.NewHandler(p)
	u, _ := user.Lookup("root")
	gid, _ := strconv.Atoi(u.Gid)

	logrus.Infof("Starting socket %s", pluginSocket)

	if err := h.ServeUnix(pluginSocket, gid); err != nil {
		logrus.Fatal(err)
	}
}
