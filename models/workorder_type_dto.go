package models

// WorkOrderTypeDto model
type WorkOrderTypeDto struct {
	ID       int64  `orm:"auto;pk;type(bigint);column(id)" json:"id"`            // ID
	Title    string `orm:"unique;type(varchar);null;column(title)" json:"title"` // 类型名称
	Count    int64  `orm:"type(bigint);column(count)" json:"count"`              // 类型名称
	CreateAt int64  `orm:"type(bigint);column(create_at)" json:"create_at"`      // 提交时间
}
