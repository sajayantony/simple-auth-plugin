package plugin

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Sirupsen/logrus"
	dockerConfigTypes "github.com/docker/docker/api/types"
	"github.com/docker/go-plugins-helpers/authorization"
)

//Plugin state for simple auth
type Plugin struct {
}

//CreatePlugin initializes the plugin
func CreatePlugin() (*Plugin, error) {
	return &Plugin{}, nil
}

//AuthZReq implements authorization interface
func (p *Plugin) AuthZReq(req authorization.Request) authorization.Response {
	logrus.Warning("Receiving request...")
	logRequest(&req)
	return authorization.Response{Allow: true}
}

//AuthZRes implements authorization interface
// ref: https://docs.docker.com/engine/extend/plugins_authorization/#request-authorization
func (p *Plugin) AuthZRes(req authorization.Request) authorization.Response {
	return authorization.Response{Allow: true}
}

func logRequest(req *authorization.Request) {
	//DEMO: The plugin does not get any payload of the auth request.
	logrus.Infof("REQUEST:%s:%s:%s",
		req.RequestMethod,
		req.RequestURI,
		req.UserAuthNMethod)

	keys := make([]string, len(req.RequestHeaders))
	for k := range req.RequestHeaders {
		keys = append(keys, k+":"+req.RequestHeaders[k])
	}
	logrus.Infof("HEADERS: %s", strings.Join(keys, ","))

	if val, ok := req.RequestHeaders["X-Registry-Auth"]; ok {
		logrus.Info("AuthHeader" + val)
	} else {
		logrus.Info("NO_AUTH_HEADER")
	}

	// https://docs.docker.com/engine/api/v1.37/#tag/System
	if strings.HasSuffix(req.RequestURI, "/auth") {
		var a dockerConfigTypes.AuthConfig
		json.Unmarshal(req.RequestBody, &a)
		//logrus.Infof("%s", authReq.Username)
		var auth = fmt.Sprintf("%s:%s:%s:%s:%s:",
			a.Username,
			a.Password,
			a.RegistryToken,
			a.IdentityToken,
			a.ServerAddress)
		logrus.Info("Authentication API invoked.....")
		logrus.Infof(auth)
	}
}
