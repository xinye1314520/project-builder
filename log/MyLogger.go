package log

import (
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap"
)

var WebError *zap.Logger
var WebInfo *zap.Logger

//func UseInfo(msg string, fields ...zap.Field) {
//	WebInfo.Info(msg, fields)
//}

func init(){
	WebError = NewLogger("./logs/web_error.log", zapcore.InfoLevel, 1, 30, 7, true, "")
	WebInfo = NewLogger("./logs/web_info.log", zapcore.DebugLevel, 1, 30, 7, true, "")
}
