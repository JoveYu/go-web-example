package middleware

import (
	"fmt"
	"time"
	"unicode"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type GORMLogger struct {
	Level zerolog.Level
}

func NewGORMLogger() *GORMLogger {
	return &GORMLogger{
		Level: zerolog.DebugLevel,
	}
}

func (l *GORMLogger) Print(v ...interface{}) {
	if len(v) < 2 {
		return
	}

	switch v[0] {
	case "sql":
		log.WithLevel(l.Level).
			Str("client", "gorm").
			// Str("line", v[1].(string)).
			Int64("duration", int64(v[2].(time.Duration)/time.Millisecond)).
			Str("sql", v[3].(string)).
			Strs("values", getFormattedValues(v[4].([]interface{}))).
			Int64("rows", v[5].(int64)).
			Msg("gorm")

	case "log":
		log.WithLevel(l.Level).
			Str("line", v[1].(string)).
			Interface("other", v[2:]).
			Msg("gorm")
	}
}

func isPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

func getFormattedValues(rawValues []interface{}) []string {
	formattedValues := make([]string, 0, len(rawValues))
	for _, value := range rawValues {
		switch v := value.(type) {
		case time.Time:
			formattedValues = append(formattedValues, fmt.Sprint(v))
		case []byte:
			if str := string(v); isPrintable(str) {
				formattedValues = append(formattedValues, fmt.Sprint(str))
			} else {
				formattedValues = append(formattedValues, "<binary>")
			}
		default:
			str := "NULL"
			if v != nil {
				str = fmt.Sprint(v)
			}
			formattedValues = append(formattedValues, str)
		}
	}
	return formattedValues
}
