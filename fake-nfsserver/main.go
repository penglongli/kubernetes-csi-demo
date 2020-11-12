package main

import (
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/google/uuid"
	"github.com/orcaman/concurrent-map"
)

var (
	NfsMap = cmap.New()

	// NFS volume base dir.
	NfsExport = "/data/nfs_export"
)

func main() {
	r := gin.Default()

	r.POST("/api/nfs", CreateNfsVolume)
	r.DELETE("/api/nfs/:id", DeleteNfsVolume)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}

type CreateRequest struct {
	// NFS volume region, fake.
	Region string `json:"region"`
	// NFS volume zone, fake.
	Zone string `json:"zone"`
	// NFS volume request size, uint GB.
	RequestSize int `json:"requestSize"`

	// CreateVolume idempotent key. Default namespace-pvc
	IdempotentKey string `json:"idempotentKey"`

	// NFS volume display name, fake.
	DisplayName string `json:"displayName"`
	// NFS volume remark, fake.
	Remark string `json:"remark"`
}

func CreateNfsVolume(ctx *gin.Context) {
	requestId := uuid.New().String()

	request := new(CreateRequest)
	if err := ctx.BindJSON(request); err != nil {
		glog.Errorf("[CreateVolume] binding json failed, err: %s", err.Error())
		Failed(ctx, err.Error())
		return
	}

	glog.Infof("[CreateVolume] received request, req: %#v", request)
	if _, ok := NfsMap.Get(request.IdempotentKey); !ok {
		NfsMap.Set(request.IdempotentKey, requestId)

		cmd := exec.Command("mkdir", "-p", NfsExport+"/"+requestId)
		err := cmd.Run()
		if err != nil {
			glog.Errorf("[CreateVolume] create nfs volume failed, err: %s", err.Error())
			Failed(ctx, err.Error())
			return
		}
		glog.Infof("[CreateVolume] create nfs volume success, name: %s, remark: %s, volumeId: %s",
			request.DisplayName, request.Remark, requestId)
	} else {
		glog.Warningf("[CreateVolume] nfs volume is created. Don't need create again.")
		Success(ctx, requestId)
		return
	}

	Success(ctx, requestId)
}

func DeleteNfsVolume(ctx *gin.Context) {
	volumeId := ctx.Param("id")
	glog.Infof("[DeleteVolume] received request, volumeId: %s")

	var key string
	for k, v := range NfsMap.Items() {
		if v.(string) == volumeId {
			key = k
		}
	}
	if key == "" {
		glog.Warningf("[DeleteNfsVolume] cannot find volumeId.")
	} else {
		glog.Infof("[DeleteNfsVolume] found key: %s", key)
	}

	cmd := exec.Command("rm", "-rf", NfsExport+"/"+volumeId)
	err := cmd.Run()
	if err != nil {
		glog.Errorf("[DeleteNfsVolume] delete nfs volume failed, err: %s", err.Error())
		Failed(ctx, err.Error())
		return
	} else {
		glog.Infof("[DeleteNfsVolume] delete nfs success.")
		NfsMap.Remove(key)
		Success(ctx, volumeId)
	}
}

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, map[string]interface{}{
		"code":    0,
		"message": "",
		"data":    data,
	})
}

func Failed(ctx *gin.Context, message string) {
	ctx.JSON(200, map[string]interface{}{
		"code":    1,
		"message": message,
		"data":    nil,
	})
}
