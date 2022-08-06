package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	StartTime string `toml:"start-time"`
	EndTime string `toml:"end-time"`
	Format string `toml:"format"`
	//Interval time.Duration `toml:"internal"`
	MetricBatchSize int `toml:"metric-batch-size"`
	InfluxDB InfluxDB `toml:"influxdb"`
	Streams []Stream `toml:"stream"`
}

type InfluxDB struct {
	Urls []string `toml:"urls"`
	Database string `toml:"database"`
	RetentionPolicy string `toml:"retention-policy"`
	Timeout int `toml:"timeout"`
}

type Stream struct {
	//Interval time.Duration `toml:"internal"`
	Measurement string `toml:"measurement"`
	Tag []Tag `toml:"tag"`
	Field []Field `field:"field"`
}

type Tag struct {
	Key string
	Value []string
	ValuesTag int
	
}

type Field struct {
	Key string
	Value []interface{}
	ValuesTag int
	ValueType string
}



func main() {
	var config Config
	if _, err := toml.DecodeFile("etc/sample.conf", &config); err != nil {

	}
	fmt.Printf("startTime: %v\n", config.StartTime)
}