package main

import (
	"encoding/json"
	"fmt"
	_ "strconv"
	"time"

	rotatelogs "github.com/rawansuww/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {

	path := "./logs/log"
	rotator, err := rotatelogs.New(
		fmt.Sprintf("%s.%s", path, "%Y-%m-%d.%H:%M:%S.log"),
		rotatelogs.WithLinkName("./logs/log"),
		rotatelogs.WithMaxAge(time.Minute*2),
		rotatelogs.WithRotationTime(time.Second),
	)

	if err != nil {
		panic(err)
	}

	// initialize the JSON encoding config
	encoderConfig := map[string]string{
		"levelEncoder": "capital",
		"timeKey":      "date",
		"timeEncoder":  "iso8601",
	}
	data, _ := json.Marshal(encoderConfig)
	var encCfg zapcore.EncoderConfig
	if err := json.Unmarshal(data, &encCfg); err != nil {
		panic(err)
	}

	// add the encoder config and rotator to create a new zap logger
	w := zapcore.AddSync(rotator)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encCfg),
		w,
		zap.InfoLevel)
	logger := zap.New(core)

	for i := 0; i < 15; i++ {
		time.Sleep(time.Second * 30)
		logger.Info("Now logging in a rotated file")
	}

}
