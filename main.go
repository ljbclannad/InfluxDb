package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"

	"influxdb-demo/config"
)

func main() {
	cfg := config.NewConfig() // 创建配置实例
	// InfluxDB连接配置
	token := cfg.InfluxDB.Token

	url := cfg.InfluxDB.URL
	client := influxdb2.NewClient(url, token)

	defer client.Close()
	org := cfg.InfluxDB.Org
	bucket := cfg.InfluxDB.Bucket
	writeAPI := client.WriteAPIBlocking(org, bucket)
	for value := 0; value < 5; value++ {
		tags := map[string]string{
			"tagname1": "tagvalue1",
		}
		fields := map[string]interface{}{
			"field1": value,
		}
		point := write.NewPoint("measurement1", tags, fields, time.Now())
		time.Sleep(1 * time.Second) // separate points by 1 second

		if err := writeAPI.WritePoint(context.Background(), point); err != nil {
			log.Fatal(err)
		}
	}

	// 查询数据
	queryAPI := client.QueryAPI(org)
	query := `from(bucket:"ljb")
		|> range(start: -1h)
		|> filter(fn: (r) => r["_measurement"] == "measurement1")`

	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	// 处理查询结果
	for result.Next() {
		fmt.Printf("值: %v, 时间: %v\n", result.Record().Value(), result.Record().Time())
	}
}
