package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
)

type AlarmConfig struct {
	gorm.Model
	EsIndiceName  string `gorm:"es_indice_name" json:"es_indice_name"`
	MsgType       string `gorm:"msg_type" json:"msg_type"`
	MsgDefine     string `gorm:"msg_type" json:"msg_define"`
	CheckInterval int    `gorm:"check_interval" json:"check_interval"`
	IsRunning     bool   `gorm:"is_running" json:"is_running"`
	MailUser      string `gorm:"mail_user"  json:"mail_user"`
	DingToken     string `gorm:"ding_token" json:"ding_token"`
	DingMobiles   string `gorm:"ding_mobiles" json:"ding_mobiles"`
}

type AlarmConfigInstance struct {
	gorm.Model
	InstanceId uint   `gorm:"instance_id" json:"instance_id"`
	ConfigIds  string `gorm:"config_ids" binding:"dive" json:"config_ids"`
	Ids        []uint `gorm:"-" json:"ids"`
}

func (a *AlarmConfigInstance) BeforeCreate(tx *gorm.DB) (err error) {
	a.ConfigIds = strings.Trim(strings.Join(strings.Split(fmt.Sprint(a.Ids), " "), ","), "[]")
	return
}
func (a *AlarmConfigInstance) AfterFind(tx *gorm.DB) (err error) {
	ls := strings.Split(a.ConfigIds, ",")
	for _, l := range ls {
		if intl, err := strconv.ParseUint(l, 10, 64); err != nil {
			a.Ids = append(a.Ids, uint(intl))
		}
	}
	return
}

type AlarmInstance struct {
	gorm.Model
	EsName string `gorm:"es_name" json:"es_name"`
	EsUrl  string `gorm:"es_url" json:"es_url"`
	EsUser string `gorm:"es_user" json:"es_user"`
	EsPass string `gorm:"es_pass" json:"es_pass"`
}

type AlarmMsg struct {
	Type string
	Msg  string
}
