package store

import (
	"fisco-contract-call-service-provider/common/mysql"
	logging "fisco-contract-call-service-provider/common"
)

// init mysql
// created table if the table not exist

//tableName :
//	tb_irita_crosschain_tx
//	tb_irita_fabric_relayer

const (
	_TabName_cc_Tx   = "tb_irita_crosschain_tx"

	_Create_CrossChain_Tx_Sql = `CREATE TABLE tb_irita_crosschain_tx (
  funique_id bigint(20) NOT NULL AUTO_INCREMENT,
  request_id varchar(255) NOT NULL DEFAULT '' COMMENT '请求唯一id',
  from_chainid varchar(255) NOT NULL DEFAULT '' COMMENT '起始链ID',
  from_tx varchar(255) NOT NULL DEFAULT '' COMMENT '起始链交易ID',
  hub_req_tx varchar(255) NOT NULL DEFAULT '' COMMENT 'HUB请求交易ID',
  ic_request_id varchar(255) NOT NULL DEFAULT '' COMMENT 'HUB请求唯一id',
  to_chainid varchar(255) NOT NULL DEFAULT '' COMMENT '目标链ID',
  to_tx varchar(255) DEFAULT NULL COMMENT '目标链交易ID',
  hub_res_tx varchar(255) NOT NULL DEFAULT '' COMMENT 'HUB响应交易ID',
  from_res_tx varchar(255) DEFAULT NULL COMMENT '向起始链响应数据的交易ID',
  tx_status int(1) NOT NULL DEFAULT '0' COMMENT '交易状态 0：未知，1：成功 2：失败',
  tx_time datetime DEFAULT NULL COMMENT '交易完成时间',
  tx_createtime datetime NOT NULL DEFAULT '1999-01-01 00:00:00' COMMENT '交易创建时间',
  error text DEFAULT NULL COMMENT '异常',
  source_service int(1) NOT NULL DEFAULT '0' COMMENT '存储交易记录的来源服务0:表示relayer，1：表示provider',
  PRIMARY KEY (funique_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`
)

func InitMysql(conn string) {
	logging.Logger.Infof("初始化Mysql : %s", conn)
	mysql.Init(conn)
	checkTable(_Create_CrossChain_Tx_Sql, _TabName_cc_Tx)
}

func checkTable(sql, tabName string) {
	logging.Logger.Infof("检查数据库：%s", tabName)
	if mysql.TabIsExit(tabName) {
		logging.Logger.Infof("数据库已存在:%s", tabName)
	} else {
		logging.Logger.Info("数据库不存在:", tabName)
		mysql.CreateTable(sql, tabName)
	}
}
