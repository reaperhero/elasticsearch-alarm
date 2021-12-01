package model

import "github.com/jinzhu/gorm"

type AlarmConfig struct {
	gorm.Model
	EsIndex       string   `gorm:"es_index" json:"es_index"`
	MsgType       string   `gorm:"msg_type" json:"msg_type"`
	MsgDefine     string   `gorm:"msg_type" json:"msg_define"`
	CheckInterval int      `gorm:"check_interval" json:"check_interval"`
	IsRunning     bool     `gorm:"is_running" json:"is_running"`
	MailUser      string   `gorm:"mail_user"  json:"mail_user"`
	DingUrl       string   `gorm:"ding_url" json:"ding_url"`
	DingMobiles   []string `gorm:"ding_mobiles" json:"ding_mobiles"`
}


type AlarmConfigInstance struct {
	gorm.Model
	ConfigID   int `gorm:"config_id" json:"config_id"`
	InstanceId int `gorm:"instance_id" json:"instance_id"`
}

type AlarmInstance struct {
	gorm.Model
	EsName string `gorm:"es_name" json:"es_name"`
	EsUrl  string `gorm:"es_url" json:"es_url"`
	EsUser string `gorm:"es_user" json:"es_user"`
	EsPass string `gorm:"es_pass" json:"es_pass"`
}
