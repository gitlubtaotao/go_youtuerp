package models



import (
        "time"
)



type BargainSurchargeFees struct {

     Id    int64 `xorm:"pk autoincr BIGINT(20)"`
     Name    string `xorm:"comment('费用简称') VARCHAR(255)"`
     NameCn    string `xorm:"comment('费用英文名') VARCHAR(255)"`
     NameEn    string `xorm:"comment('费用中文名') VARCHAR(255)"`
     UnitPrice    string `xorm:"default 0.0000 DECIMAL(15,4)"`
     Account    int `xorm:"default 0 comment('数量') INT(11)"`
     FinanceCurrency    string `xorm:"default 'CNY' comment('币种') VARCHAR(64)"`
     Remarks    string `xorm:"VARCHAR(255)"`
     Parcel    string `xorm:"default 0.0000 DECIMAL(15,4)"`
     SupplyId    int `xorm:"default 0 comment('供应商') index INT(11)"`
     CreatedAt    time.Time `xorm:"not null DATETIME"`
     UpdatedAt    time.Time `xorm:"not null DATETIME"`
     SourceType    string `xorm:"index(index_bargain_surcharge_fees_on_source_type_and_source_id) VARCHAR(255)"`
     SourceId    int64 `xorm:"index(index_bargain_surcharge_fees_on_source_type_and_source_id) BIGINT(20)"`
     BargainMainId    int64 `xorm:"index BIGINT(20)"`
     20fr    string `xorm:"DECIMAL(14,4)"`
     20gp    string `xorm:"DECIMAL(14,4)"`
     20ot    string `xorm:"DECIMAL(14,4)"`
     20rf    string `xorm:"DECIMAL(14,4)"`
     40gp    string `xorm:"DECIMAL(14,4)"`
     40fr    string `xorm:"DECIMAL(14,4)"`
     40ot    string `xorm:"DECIMAL(14,4)"`
     40hq    string `xorm:"DECIMAL(14,4)"`
     40rf    string `xorm:"DECIMAL(14,4)"`
     45gp    string `xorm:"DECIMAL(14,4)"`
     45hq    string `xorm:"DECIMAL(14,4)"`
     20rh    string `xorm:"DECIMAL(14,4)"`
     20hq    string `xorm:"DECIMAL(14,4)"`
     20ht    string `xorm:"DECIMAL(14,4)"`
     40ht    string `xorm:"DECIMAL(14,4)"`
     40rh    string `xorm:"DECIMAL(14,4)"`
     DeletedAt    time.Time `xorm:"default '1969-12-31 16:00:00' DATETIME"`

}

