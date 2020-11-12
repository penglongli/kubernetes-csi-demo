package main

import "github.com/penglongli/kubernetes-csi-demo/cmd/nfs/app"

func main() {
	if err := app.NewCommand().Execute(); err != nil {
		panic(err)
	}
}
