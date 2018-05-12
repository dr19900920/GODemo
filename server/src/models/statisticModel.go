package models

type Statistic struct {
	SpecialCount   int `json:"specialCount"`   //特殊工单
	NoTrackCount   int `json:"noTrackCount"`   //待处理工单
	NeedTrackCount int `json:"needTrackCount"` //待跟踪工单
	ReleaseCount   int `json:"releaseCount"`   //即将到期工单数量
	PublicCount    int `json:"publicCount"`    //待认领学员
	DynamicCount   int `json:"dynamicCount"`   //七日动态学员数量
	FollowCount    int `json:"followCount"`    //关注学员
}

const (
	WORK_ORDER_COUNT_SQL     = "SELECT count(DISTINCT work_order.id) FROM work_order"
	WORK_ORDER_LINK_USER_SQL = " INNER JOIN user ON user.id = work_order.user_id"

	SPECIAL_COUNT_SQL = " AND activate_status = 1 AND assign_status = 0 AND deal_status = 0 AND sale_filter_id > 0"
	NOTRACK_COUNT_SQL = " AND activate_status = 1 AND assign_status = 0 AND dispose_status = 0 AND deal_status = 0"
)

func GetStatisticCount() *Statistic {

	var special int
	row := DataBase.QueryRow(WORK_ORDER_COUNT_SQL+" where user_id = ?"+SPECIAL_COUNT_SQL, 28405)
	row.Scan(&special)

	var noTrack int
	workOrderFilter := " AND" + buildWorkOrderFilter()
	row = DataBase.QueryRow(WORK_ORDER_COUNT_SQL + WORK_ORDER_LINK_USER_SQL + " WHERE 1" + workOrderFilter + NOTRACK_COUNT_SQL)
	row.Scan(&noTrack)

	return &Statistic{
		special,
		noTrack,
		0,
		0,
		0,
		0,
		0,
	}
}

//工单权限
func buildWorkOrderFilter() string {
	return ""
}
