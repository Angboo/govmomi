package simulator

import (
	"fmt"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
	"time"
)

type HostDateTimeSystem struct {
	mo.HostDateTimeSystem
	current time.Time
}

func NewHostDateTimeSystem(hs *mo.HostSystem) *HostDateTimeSystem {
	return &HostDateTimeSystem{
		HostDateTimeSystem: mo.HostDateTimeSystem{
			Self: hs.ConfigManager.DateTimeSystem.Reference(),
		},
	}
}

func (s HostDateTimeSystem) init(r *Registry) {
}

func (s *HostDateTimeSystem) UpdateDateTime(req *types.UpdateDateTime) soap.HasFault {
	fmt.Println("start UpdateDateTime")
	s.current = req.DateTime
	return &methods.UpdateDateTimeBody{
		Res: &types.UpdateDateTimeResponse{},
	}
}

func (s *HostDateTimeSystem) QueryDateTime(req *types.QueryDateTime) soap.HasFault {
	fmt.Println("start QueryDateTime")
	return &methods.QueryDateTimeBody{
		Res: &types.QueryDateTimeResponse{
			Returnval: s.current,
		},
	}
}

func (s HostDateTimeSystem) UpdateDateTimeConfig(req *types.UpdateDateTimeConfig) soap.HasFault {

	return &methods.UpdateDateTimeConfigBody{}

}
