package simulator

import "github.com/vmware/govmomi/vim25/mo"

type HostServiceSystem struct {
	mo.HostServiceSystem
}

func NewHostServiceSystem(hs *mo.HostSystem) *HostServiceSystem {
	return &HostServiceSystem{
		HostServiceSystem: mo.HostServiceSystem{
			ExtensibleManagedObject: mo.ExtensibleManagedObject{
				Self: hs.ConfigManager.ServiceSystem.Reference(),
			},
		},
	}
}
func (s HostServiceSystem) name() {
}
