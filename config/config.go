package config

var (
	envConfig *EnvConfig
)

type EnvConfig struct {
	Endpoint string
	NfsServerAddress string
	NodeId string
	LogFile string
	LogFileMaxSize int
}

func InitConfig(endpoint, nfsServer, nodeId, logFile string, logFileMaxSize int) {
	envConfig = &EnvConfig{
		Endpoint: endpoint,
		NfsServerAddress: nfsServer,
		NodeId: nodeId,
		LogFile: logFile,
		LogFileMaxSize: logFileMaxSize,
	}
}

func GetConfig() *EnvConfig {
	return envConfig
}
