package simulator_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"log"
	"testing"
	"time"
)

func TestHostDateTimeSystem_QueryDateTime(t *testing.T) {
	ctx := context.Background()
	vcURL, closeServer := StartTestVc(nil, false)
	defer closeServer()

	c, err := govmomi.NewClient(ctx, vcURL, true)
	if err != nil {
		log.Println(err)
	}
	system := &mo.HostDateTimeSystem{Self: types.ManagedObjectReference{
		Type:  "HostDateTimeSystem",
		Value: "dateTimeSystem"}}

	system1 := object.NewHostDateTimeSystem(c.Client, system.Reference())
	_, _ = system1.Query(ctx)

	assert.NoError(t, err)
}

func TestHostDateTimeSystem_UpdateDateTime(t *testing.T) {
	ctx := context.Background()
	vcURL, closeServer := StartTestVc(nil, false)
	defer closeServer()
	c, err := govmomi.NewClient(ctx, vcURL, true)
	if err != nil {
		log.Println(err)
	}
	system := &mo.HostDateTimeSystem{Self: types.ManagedObjectReference{
		Type:  "HostDateTimeSystem",
		Value: "dateTimeSystem"}}

	system1 := object.NewHostDateTimeSystem(c.Client, system.Reference())
	now := time.Now()
	_ = system1.Update(ctx, now)
	query, _ := system1.Query(ctx)

	assert.NoError(t, err)
	assert.True(t, query.Equal(now))
}

