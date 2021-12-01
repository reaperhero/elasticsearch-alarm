package model

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"time"
)

type HistoryAccessPoint struct {
	Timestamp time.Time `json:"@timestamp"`
	Message   string    `json:"message,omitempty"`
	Amessage  AccessMeaage
}

func (a *HistoryAccessPoint) CreateAmessage() error {
	if err := json.Unmarshal([]byte(a.Message), &a.Amessage); err != nil {
		if !strings.Contains(err.Error(), "invalid character") {
			logrus.WithField("error", err).Infof("access get [%s]", a.Message)
		}
		return err
	}
	return nil
}

type HistoryGrpcPoint struct {
	Timestamp time.Time `json:"@timestamp"`
	Message   string    `json:"message,omitempty"`
	Gmessage  GrpcMessage
}

func (g *HistoryGrpcPoint) CreateGmessage() error {
	if err := json.Unmarshal([]byte(g.Message), &g.Gmessage); err != nil {
		logrus.WithField("error", err).Infof("grpc get [%s]", g.Message)
		return err
	}
	return nil
}

type AccessMeaage struct {
	Time         time.Time `json:"time"`
	ID           string    `json:"id"`
	RemoteIP     string    `json:"remote_ip"`
	Host         string    `json:"host"`
	Method       string    `json:"method"`
	URI          string    `json:"uri"`
	UserAgent    string    `json:"user_agent"`
	Status       int       `json:"status"`
	Error        string    `json:"error"`
	Latency      int64     `json:"latency"`
	LatencyHuman string    `json:"latency_human"`
	BytesIn      int       `json:"bytes_in"`
	BytesOut     int       `json:"bytes_out"`
}

type GrpcMessage struct {
	Level string `json:"level"`
	Error string `json:"error"`
	Msg   string `json:"msg"`
}