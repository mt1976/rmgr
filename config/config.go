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
	MQContentType       string
	MQExchange          string
	DateTimeFormat      string
	MQAddressFormat     string
	DefaultStatus       string
	StaleAfterMS        int
	InterestType        string
	RateText            string
	FXText              string
	RoutingKeyFormat    string
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
	Configuration.MQContentType = "application/json"
	Configuration.DateTimeFormat = "2006-01-02T15:04:05.9999999999Z"
	Configuration.MQAddressFormat = "amqp://%v:%v@%v:%v/"
	Configuration.DefaultStatus = "OK"
	Configuration.MQExchange = "TestExchange"
	Configuration.StaleAfterMS = 100 // 100ms
	Configuration.InterestType = "INTEREST"
	Configuration.RateText = "RATE"
	Configuration.FXText = "FX"
	Configuration.RoutingKeyFormat = "%v.%v.%v.%v"
}
