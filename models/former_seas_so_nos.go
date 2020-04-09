package models

import (
	"time"
)

type FormerSeasSoNos struct {
	Id                       int64     `xorm:"pk autoincr BIGINT(20)"`
	CreatedAt                time.Time `xorm:"not null DATETIME"`
	UpdatedAt                time.Time `xorm:"not null DATETIME"`
	Vessel                   string    `xorm:"VARCHAR(255)"`
	Voyage                   string    `xorm:"comment('航次') VARCHAR(255)"`
	SeaPortPodId             int       `xorm:"index INT(11)"`
	SeaPortPolId             int       `xorm:"index INT(11)"`
	SeaPortViaId             int       `xorm:"index INT(11)"`
	SoNo                     string    `xorm:"TEXT"`
	PodAgentType             string    `xorm:"index(index_former_seas_so_nos_on_pod_agent_type_and_pod_agent_id) VARCHAR(255)"`
	PodAgentId               int64     `xorm:"index(index_former_seas_so_nos_on_pod_agent_type_and_pod_agent_id) BIGINT(20)"`
	PodAgentContent          string    `xorm:"TEXT"`
	ShippedOnBoardDate       time.Time `xorm:"comment('装船日期') DATETIME"`
	OceanChangesPaytypeId    int       `xorm:"comment('运费付款方式') index INT(11)"`
	OtherChangesPaytypeId    int       `xorm:"comment('其他运费付款方式') index INT(11)"`
	CyOpenDate               time.Time `xorm:"comment('开舱时间') DATETIME"`
	VoucherCutOff            time.Time `xorm:"comment('截放行条时间') DATETIME"`
	VgmSubmissionDeadline    time.Time `xorm:"comment('VGM截止时间') DATETIME"`
	SiCutOff                 time.Time `xorm:"comment('截补料时间') DATETIME"`
	OrderMasterId            int       `xorm:"index INT(11)"`
	SerialNo                 string    `xorm:"VARCHAR(255)"`
	Status                   string    `xorm:"VARCHAR(255)"`
	BoatCompanyId            int       `xorm:"INT(11)"`
	CutOffDate               time.Time `xorm:"DATETIME"`
	EstimatedTimeOfDeparture time.Time `xorm:"DATETIME"`
	EstimatedTimeOfArrival   time.Time `xorm:"DATETIME"`
	DeletedAt                time.Time `xorm:"default '1969-12-31 16:00:00' index DATETIME"`
}
