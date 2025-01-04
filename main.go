package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	// InfluxDB连接配置
	token := "4t1azdOBhIcpGaOJ_5tK82luGKkxOuIRAMngMcCT1LLVZ_L7dAJK7DlJ896F0z-mTHM6OwCkB8jW2hQI3EF-ug=="
	url := "http://localhost:8086"
	org := "your-org"
	bucket := "your-bucket"

	// 创建客户端
	client := influxdb2.NewClient(url, token)
	defer client.Close()

	// 获取写入客户端
	writeAPI := client.WriteAPIBlocking(org, bucket)

	// 创建数据点
	p := influxdb2.NewPoint(
		"system_metrics",
		map[string]string{
			"host": "host1",
		},
		map[string]interface{}{
			"cpu_usage": 80.0,
			"memory":    70.5,
		},
		time.Now(),
	)

	// 写入数据
	if err := writeAPI.WritePoint(context.Background(), p); err != nil {
		log.Fatal(err)
	}

	// 查询数据
	queryAPI := client.QueryAPI(org)
	query := `from(bucket:"your-bucket")
		|> range(start: -1h)
		|> filter(fn: (r) => r["_measurement"] == "system_metrics")`

	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	// 处理查询结果
	for result.Next() {
		fmt.Printf("值: %v, 时间: %v\n", result.Record().Value(), result.Record().Time())
	}
}
