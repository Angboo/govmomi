package simulator_test


import (
"crypto/tls"
	"github.com/vmware/govmomi/simulator"
	"log"
"net/url"

)

type CloseFunc func()

func StartTestVc(tlsConfig *tls.Config, multiDc ...bool) (url *url.URL, closer func()) {
	model := simulator.VPX()
	if len(multiDc) > 0 && multiDc[0] {
		model.Datacenter = 2
		model.Datastore = 1
		model.Cluster = 1
		model.Host = 0
	}
	err := model.Create()
	if err != nil {
		log.Fatal(err)
	}

	return serve(tlsConfig, model)
}

func StartTestVcWithDir(tlsConfig *tls.Config, dir string) (url *url.URL, closer func()) {
	model := simulator.VPX()

	err := model.Load(dir)
	if err != nil {
		log.Fatal(err)
	}

	return serve(tlsConfig, model)
}

func serve(tlsConfig *tls.Config, model *simulator.Model) (url *url.URL, closer func()) {
	model.Service.TLS = tlsConfig
	s := model.Service.NewServer()

	return s.URL, func() {
		s.CloseClientConnections()
		s.Close()
		model.Remove()
	}
}
