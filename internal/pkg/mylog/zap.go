package mylog

import (
	"bytes"
	"fmt"
	"go.uber.org/zap/zapio"
	"log"
	"mini_game_balance/configs"
	"mini_game_balance/internal/pkg/utils"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/buffer"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type customZap struct {
	c *configs.Zap
}

func Init() {
	NewZap()

	// 设置log 输出到zap
	c := &configs.ServerConfig.Zap
	// 设置日志输出
	zapWriter := &zapio.Writer{
		Log:   zap.L(),
		Level: c.TransportLevel(),
	}
	log.SetOutput(zapWriter)
	log.SetFlags(log.Llongfile | log.LstdFlags)

}
func NewZap() (logger *zap.Logger) {
	c := &configs.ServerConfig.Zap
	if ok, _ := utils.PathExists(c.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", c.Director)
		_ = os.Mkdir(c.Director, os.ModePerm)
	}
	var customZap = &customZap{c}
	cores := customZap.getZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if c.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	zap.ReplaceGlobals(logger)

	return logger
}

// getEncoder 获取 zapcore.Encoder
func (z *customZap) getEncoder() zapcore.Encoder {
	if z.c.Format == "json" {
		return zapcore.NewJSONEncoder(z.getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(z.getEncoderConfig())
}

// getEncoderConfig 获取zapcore.EncoderConfig
func (z *customZap) getEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  z.c.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    z.c.ZapEncodeLevel(),
		EncodeTime:     z.customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   CallerEncoder,
	}
}

// getEncoderCore 获取Encoder的 zapcore.Core
func (z *customZap) getEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	writer, err := z.getWriteSyncer(l.String()) // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return nil
	}

	return zapcore.NewCore(&EscapeSeqJSONEncoder{z.getEncoder()}, writer, level)
}

// customTimeEncoder 自定义日志输出时间格式
func (z *customZap) customTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(t.Format(z.c.Prefix + "2006/01/02 - 15:04:05.000"))
}

// getZapCores 根据配置文件的Level获取 []zapcore.Core
func (z *customZap) getZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := z.c.TransportLevel(); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, z.getEncoderCore(level, z.getLevelPriority(level)))
	}
	return cores
}
func (z *customZap) getWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	fileWriter, err := rotatelogs.New(
		path.Join(z.c.Director, "%Y-%m-%d", level+".log"),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(z.c.MaxAge)*24*time.Hour), // 日志留存时间
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if z.c.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}

// GetLevelPriority 根据 zapcore.Level 获取 zap.LevelEnablerFunc
func (z *customZap) getLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool { // 日志级别
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool { // 警告级别
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool { // 错误级别
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool { // dpanic级别
			return level == zap.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool { // panic级别
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool { // 终止级别
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	}
}

// FuncName 返回调用本函数的函数名称
// pc runtime.Caller 返回的第一个值
func getFuncName(pc uintptr) string {
	funcName := runtime.FuncForPC(pc).Name()
	sFuncName := strings.Split(funcName, ".")
	return sFuncName[len(sFuncName)-1]
}

// CallerEncoder serializes a caller in package/file:funcname:line format
func CallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	shortCaller := caller.TrimmedPath()
	shortCallerSplited := strings.Split(shortCaller, ":")
	funcName := getFuncName(caller.PC)
	result := shortCallerSplited[0] + ":" + funcName + ":" + shortCallerSplited[1]
	enc.AppendString(result)
}

type EscapeSeqJSONEncoder struct {
	zapcore.Encoder
}

// EncodeEntry 将方法zap.error中的errorVerbose的堆栈换行符修改
func (enc *EscapeSeqJSONEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	b, err := enc.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, err
	}
	newBuffer := buffer.NewPool().Get()

	b1 := bytes.Replace(b.Bytes(), []byte("\\n"), []byte("\n"), -1)
	b2 := bytes.Replace(b1, []byte("\\t"), []byte("\t"), -1)
	_, err = newBuffer.Write(b2)
	if err != nil {
		return nil, err
	}
	return newBuffer, nil
}
