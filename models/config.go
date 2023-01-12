package models

type ConfigPostgres struct {
	PostgresDB        string `env:"POSTGRES_DB"`
	PostgresUser      string `env:"POSTGRES_USER"`
	PosttgresPassword string `env:"POSTGRES_PASSWORD"`
	HostDB            string `env:"HOST_DB"`
	PosrtDB           string `env:"PORT_DB"`
}

type ConfigKafka struct {
	KafkaAddr string `env:"KAFKA_ADDR"`
	Topic     string `env:"TOPIC"`
}

type ConfigGoods struct {
	HttpBind string `env:"HTTP_BIND"`
	*ConfigPostgres
	*ConfigKafka
}

type ConfigOrder struct {
	HttpBind string `env:"HTTP_BIND"`
	*ConfigPostgres
	*ConfigKafka
}

type ConfigOrderHistory struct {
	HttpBind string `env:"HTTP_BIND"`
	*ConfigPostgres
	*ConfigKafka
}
