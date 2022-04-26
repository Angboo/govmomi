package simulator

import (
	"context"
	"crypto/tls"
	"github.com/stretchr/testify/assert"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/govc/host/esxcli"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
	"log"
	"net/url"
	"testing"
)

func Start(tlsConfig *tls.Config, multiDc ...bool) (url *url.URL, closer func()) {
	model := VPX()
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

func serve(tlsConfig *tls.Config, model *Model) (url *url.URL, closer func()) {
	model.Service.TLS = tlsConfig
	s := model.Service.NewServer()

	return s.URL, func() {
		s.CloseClientConnections()
		s.Close()
		model.Remove()
	}
}

func Test_EsxCliExecutorRun(t *testing.T) {
	ctx := context.Background()
	vcURL, closeServer := Start(nil, false)
	defer closeServer()
	c, err := govmomi.NewClient(ctx, vcURL, true)
	if err != nil {
		log.Println(err)
	}
	hs := object.NewHostSystem(c.Client, types.ManagedObjectReference{
		Type:  "HostSystem",
		Value: "host-50"})
	executor, _ := esxcli.NewExecutor(c.Client, hs)
	run, _ := executor.Run([]string{"vm", "appinfo", "get"})
	assert.NoError(t, err)
	assert.NotNil(t, run)
}
