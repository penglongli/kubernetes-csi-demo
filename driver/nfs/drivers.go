package nfs

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	csicommon "github.com/kubernetes-csi/drivers/pkg/csi-common"

	"github.com/penglongli/kubernetes-csi-demo/config"
)

const (
	DriverName = "com.storage.csi.nfs"
	DriverVersion = "1.0.0"
)

type driver struct {
	csiDriver *csicommon.CSIDriver
}

func NewDriver() *driver {
	csiDriver := csicommon.NewCSIDriver(DriverName, DriverVersion, config.GetConfig().NodeId)
	csiDriver.AddVolumeCapabilityAccessModes([]csi.VolumeCapability_AccessMode_Mode{
		csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER,
	})

	// CFS plugin does not support ControllerServiceCapability now.
	// If support is added, it should set to appropriate
	// ControllerServiceCapability RPC types.
	csiDriver.AddControllerServiceCapabilities([]csi.ControllerServiceCapability_RPC_Type{
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
	})

	return &driver{
		csiDriver: csiDriver,
	}
}

func (driver *driver) Run() {
	s := csicommon.NewNonBlockingGRPCServer()

	envConfig := config.GetConfig()
	s.Start(
		envConfig.Endpoint,
		csicommon.NewDefaultIdentityServer(driver.csiDriver),
		&controllerServer{
			csicommon.NewDefaultControllerServer(driver.csiDriver),
		},
		&nodeServer{
			csicommon.NewDefaultNodeServer(driver.csiDriver),
		})
	s.Wait()
}
