package awswrapper

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
)

// AWSLogs contains the CloudWatch Logs message
// event data that is a Base64-encoded.gzip file
// archive.
type AWSLogs struct {
	Data string `json:"data"`
}

// CloudWatchEvent is the CloudWatch Logs message
// event data.
type CloudWatchEvent struct {
	AWSLogs AWSLogs `json:"awslogs"`
}

// LogEvents contains the log event information.
type LogEvents struct {
	ID        string `json:"id"`        // A unique identifier for every log event.
	Timestamp int64  `json:"timestamp"` // A timestamp when the log is created in unix epoch format.
	Message   string `json:"message"`   // The log event message.
}

// CloudWatchData is the CloudWatch Logs message
// data (decoded).
type CloudWatchData struct {
	MessageType         string      `json:"messageType"`         // Data message will use the "DATA_MESSAGE" type.
	Owner               string      `json:"owner"`               // The AWS Account ID of the originating log data.
	LogGroup            string      `json:"logGroup"`            // The log group name of the originating log data.
	LogStream           string      `json:"logStream"`           // The log stream name of the originating log data.
	SubscriptionFilters []string    `json:"subscriptionFilters"` // The list of subscription filter names that matched the originating log data.
	LogEvents           []LogEvents `json:"logEvents"`           // The actual log data, represented as an array of log records.
}

// DecodeData returns the decoded CloudWatch Logs Data.
func (cw *CloudWatchEvent) DecodeData() (*CloudWatchData, error) {
	var data = new(CloudWatchData)

	// Decode the Base64-encoded data
	decoded, err := base64.StdEncoding.DecodeString(cw.AWSLogs.Data)
	if err != nil {
		return nil, err
	}

	// Decompress the data by using the decoded datra
	buffer := bytes.NewReader(decoded)
	reader, err := gzip.NewReader(buffer)
	if err != nil {
		return nil, err
	}

	// Decode the decompressed data into the CloudWatch struct
	err = json.NewDecoder(reader).Decode(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
