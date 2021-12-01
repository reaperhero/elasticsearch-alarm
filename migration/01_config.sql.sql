CREATE TABLE `alarm_config`
(
    id             INT(11)      NOT NULL AUTO_INCREMENT,
    es_index       VARCHAR(256) NOT NULL DEFAULT '*' COMMENT '搜索索引',
    msg_type       varchar(20)           DEFAULT 'json' COMMENT '监控数据类型json，text',
    msg_define     varchar(20)           DEFAULT 'error' COMMENT '告警标识 json:error text:time out',
    check_interval INT(11)      NOT NULL DEFAULT 60 COMMENT '轮询间隔时间',
    is_running     TINYINT(2)   NOT NULL DEFAULT 0 COMMENT '发送短信和dingding机器人是否关闭',
    mail_user      VARCHAR(2048)         DEFAULT '' COMMENT '接受邮件的用户，表示多个可以用分号隔开',
    ding_token       VARCHAR(8192)         DEFAULT '' COMMENT 'dingding群机器人地址',
    ding_mobiles   VARCHAR(256)          DEFAULT '' COMMENT 'dingding机器人@的用户手机号',
    create_at      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;


CREATE TABLE `alarm_instance`
(
    id      INT(11)      NOT NULL AUTO_INCREMENT,
    es_name VARCHAR(128) NOT NULL COMMENT '名称',
    es_user VARCHAR(128) default '' COMMENT 'es账号名',
    es_pass VARCHAR(128) default '' COMMENT 'es密码',
    es_url  VARCHAR(128) COMMENT '请求的es url地址, eg: http://xxxx:9200',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;



CREATE TABLE `alarm_config_instance`
(
    id          INT(11) NOT NULL AUTO_INCREMENT,
    config_id   INT(11) NOT NULL COMMENT '规则id',
    instance_id INT(11) NOT NULL COMMENT '管理实例',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8