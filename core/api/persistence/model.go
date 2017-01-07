package persistence

import "time"

//系统中全部实体类的基础实体类
type PersistentEntity struct {
	Id                       uint64    `json:"id"`          // id
	Version                  uint64   `json:"creationDate"` // 版本号
	CreationDate, UpdateDate time.Time                      // 创建日期 更新时间
	CreateBy, UpdateBy       string                         // 创建者  更新者
	CreateById, UpdateById   uint64                         // 创建者Id 更新者Id
}

type JsonData struct {
	TotalCount uint64
	Data interface{}
}