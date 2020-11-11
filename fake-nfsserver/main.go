package main

import (
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/google/uuid"
)


var (
	NfsMap = make(map[string]string)

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
	Namespace string `json:"namespace"`
	PvcName   string `json:"pvcName"`
}

func CreateNfsVolume(ctx *gin.Context) {
	requestId := uuid.New().String()

	request := new(CreateRequest)
	if err := ctx.BindJSON(request); err != nil {
		glog.Errorf("[CreateVolume] binding json failed, err: %s", err.Error())
		ctx.JSON(200, map[string]string{
			"code": "1",
			"message": err.Error(),
		})
		return
	}

	glog.Infof("[CreateVolume] received request, namespace: %s, pvcName: %s", request.Namespace, request.PvcName)
	key := request.Namespace + "-" + request.PvcName
	if id, ok := NfsMap[key]; !ok {
		NfsMap[key] = requestId

		cmd := exec.Command("mkdir", "-p", NfsExport + "/" + requestId)
		err := cmd.Run()
		if err != nil {
			glog.Errorf("[CreateVolume] create nfs volume failed, err: %s", err.Error())
			ctx.JSON(200, map[string]string{
				"code": "1",
				"message": err.Error(),
			})
			return
		}
		glog.Infof("[CreateVolume] create nfs volume success, namespace: %s, pvcName: %s, volumeId: %s",
			request.Namespace, request.PvcName, requestId)
	} else {
		glog.Warningf("[CreateVolume] nfs volume is created. Don't need create again.")
		ctx.JSON(200, map[string]string{
			"code": "0",
			"message": "",
			"data": id,
		})
		return
	}

	ctx.JSON(200, map[string]string{
		"code": "0",
		"message": "",
		"data": requestId,
	})
}

func DeleteNfsVolume(ctx *gin.Context) {
	volumeId := ctx.Param("id")

	glog.Infof("[DeleteVolume] received request, volumeId: %s")
	var key string
	for k, v := range NfsMap {
		if v == volumeId {
			key = k
		}
	}
	if key == "" {
		glog.Warningf("[DeleteNfsVolume] cannot find volumeId.")
	} else {
		glog.Infof("[DeleteNfsVolume] found key: %s", key)
	}

	cmd := exec.Command("rm", "-rf", NfsExport + "/" + volumeId)
	err := cmd.Run()
	if err != nil {
		glog.Errorf("[DeleteNfsVolume] delete nfs volume failed, err: %s", err.Error())
		ctx.JSON(200, map[string]string{
			"code": "1",
			"message": err.Error(),
		})
		return
	} else {
		glog.Infof("[DeleteNfsVolume] delete nfs success.")
		ctx.JSON(200, map[string]string{
			"code": "0",
			"message": "",
			"data": volumeId,
		})
	}
}
