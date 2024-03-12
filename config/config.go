package config

type Config struct {
	Port                int
	Target              string
	Protocol            string
	Role                int
	ApplicationName     string
	SimulatorSourceFile string
	SimulatorConfigFile string
	MQQueue             string
	MQUser              string
	MQPassword          string
	MQHost              string
	MQPort              int
}

const (
	SEND    int = iota
	RECEIVE int = iota
)

var Configuration = Config{}

func init() {
	Configuration.ApplicationName = "Rate Simulator"
	Configuration.Role = SEND
	Configuration.Port = 8080
	Configuration.Target = "localhost"
	Configuration.Protocol = "tcp"
	Configuration.SimulatorSourceFile = "data.csv"
	Configuration.SimulatorConfigFile = "config.csv"
	Configuration.MQQueue = "TestQueue"
	Configuration.MQUser = "guest"
	Configuration.MQPassword = "guest"
	Configuration.MQHost = "localhost"
	Configuration.MQPort = 5672

}
