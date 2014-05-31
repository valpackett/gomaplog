package gomaplog

import (
	"encoding/json"
)

type JSONFormatter struct {
}

func (formatter *JSONFormatter) Format(event LogEvent) ([]byte, error) {
	outMap := map[string]interface{}{}
	outMap["version"] = "1.1"
	outMap["host"] = event.Host
	outMap["short_message"] = event.Message
	outMap["full_message"] = event.LongMessage
	outMap["timestamp"] = event.Timestamp.Unix()
	outMap["level"] = event.Level
	for k, _ := range event.Extras {
		if k[0] == '_' {
			outMap[k] = event.Extras[k]
		} else {
			outMap["_"+k] = event.Extras[k]
		}
	}
	return json.Marshal(outMap)
}
