package app

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"go.uber.org/zap/zapcore"

	"github.com/penglongli/go-utils/zaplog"

	"github.com/penglongli/kubernetes-csi-demo/config"
	"github.com/penglongli/kubernetes-csi-demo/driver/nfs"
	"github.com/penglongli/kubernetes-csi-demo/version"
)

var (
	endpoint       string
	nfsServerAddr  string
	nodeID         string
	logFile        string
	logFileMaxSize int
)

func NewCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "csi-storage-nfs",
		Run: func(cmd *cobra.Command, args []string) {
			if nfsServerAddr == "" {
				log.Fatalf("--nfs-server-address must provide.")
				return
			}

			// init config
			config.InitConfig(endpoint, nfsServerAddr, nodeID, logFile, logFileMaxSize)

			// init log
			zlog.InitLog(logFile, zapcore.InfoLevel)

			// run server
			nfs.NewDriver().Run()
		},
	}
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version info.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(version.String())
		},
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().StringVarP(&endpoint, "endpoint", "", "unix://plugin/csi.sock", "CSI endpoint")
	rootCmd.PersistentFlags().StringVarP(&nfsServerAddr, "nfs-server-address", "", "", "nfs server address. Which used to create nfs volume.")
	rootCmd.PersistentFlags().StringVarP(&nodeID, "node-id", "", "", "node ID")
	rootCmd.PersistentFlags().StringVarP(&logFile, "log-file", "", "/var/log/csi-storage-nfs.log", "log file")
	rootCmd.PersistentFlags().IntVarP(&logFileMaxSize, "log-file-max-size", "", 1000, "log size")
	return rootCmd
}
