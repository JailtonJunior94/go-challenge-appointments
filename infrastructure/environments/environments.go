package environments

import (
	"log"

	"github.com/spf13/viper"
)

var (
	Environment                = ""
	Port                       = 0
	SqlConnectionString        = ""
	ServiceBusConnectionString = ""
	QueueName                  = ""
	ExamsBaseUrl               = ""
)

func NewConfig() {
	var err error

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	Environment = viper.GetString("environment")
	Port = viper.GetInt("api.port")
	SqlConnectionString = viper.GetString("api.connectionString")
	ExamsBaseUrl = viper.GetString("api.examsBaseUrl")
	ServiceBusConnectionString = viper.GetString("notifier.serviceBus.connectionString")
	QueueName = viper.GetString("notifier.serviceBus.queueName")
}
