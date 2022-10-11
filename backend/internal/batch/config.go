package batch

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/klaytn/klaytn/common"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ConfigConnection ...
func (c *Client) ConfigConnection() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	KasAccessKey := os.Getenv("KAS_ACCESS_KEY")
	KasSecretKey := os.Getenv("KAS_SECRET_ACCESS_KEY")
	ContractAddr := os.Getenv("CONTRACT_ADDR")
	LivepeerAPIKey := os.Getenv("LIVEPEER_API_KEY")

	c.Environ.LivepeerAPIKey = LivepeerAPIKey
	c.Environ.KasAccessKey = KasAccessKey
	c.Environ.KasSecretKey = KasSecretKey
	c.ContractAddr = common.HexToAddress(ContractAddr)
	c.httpClient = http.DefaultClient

	c.Channel.MetadataChan = make(chan MappedData)
	c.Channel.TokenChan = make(chan *TokenInfo)

	c.Database.ConnectDatabase()
}

func (c *Client) SetBatchAccessLog() {
	logFile := "log/batch/access/batch-%Y-%m-%d-%H:%M.log"
	rotator, err := rotatelogs.New(
		logFile,
		rotatelogs.WithMaxAge(time.Minute*5),
		rotatelogs.WithRotationTime(time.Minute))
	FatalWithError(err)
	encoderConfig := map[string]string{
		"levelEncoder": "capital",
		"timeKey":      "date",
		"timeEncoder":  "iso8601",
	}
	data, _ := json.Marshal(encoderConfig)
	var encCfg zapcore.EncoderConfig
	err = json.Unmarshal(data, &encCfg)
	FatalWithError(err)
	w := zapcore.AddSync(rotator)
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encCfg), w, zap.InfoLevel)
	accessLoger := zap.New(core)
	accessLoger.Info("Now logging in a rotated file")
	c.AccessLogger = accessLoger
}

func (c *Client) SetBatchErrorLog() {
	logFile := "log/batch/error/batch-%Y-%m-%d-%H:%M.log"
	rotator, err := rotatelogs.New(
		logFile,
		rotatelogs.WithMaxAge(time.Minute*5),
		rotatelogs.WithRotationTime(time.Minute))
	FatalWithError(err)
	encoderConfig := map[string]string{
		"levelEncoder": "capital",
		"timeKey":      "date",
		"timeEncoder":  "iso8601",
	}
	data, _ := json.Marshal(encoderConfig)
	var encCfg zapcore.EncoderConfig
	err = json.Unmarshal(data, &encCfg)
	FatalWithError(err)
	w := zapcore.AddSync(rotator)
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encCfg), w, zap.ErrorLevel)
	errorLoger := zap.New(core)
	c.ErrorLogger = errorLoger
}
