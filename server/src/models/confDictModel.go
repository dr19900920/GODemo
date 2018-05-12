package models

import "log"

type ConfDictModel struct {
	MobileProvince        []string        `json:"mobile_province"`         //号码归属地-省
	Grade                 []int           `json:"grade"`                   //学员年级
	TalkMethod            []string        `json:"talk_method"`             //沟通方式
	TalkStatus            []ConfBaseModel `json:"talk_status"`             //沟通状态/意向等级
	ProductType           []ConfBaseModel `json:"product_type"`            //产品类型/产品意向
	ScreeningStudentLevel []ConfBaseModel `json:"screening_student_level"` //初筛号源质量
	CrmRole               []ConfBaseModel `json:"crm_role"`                //学员角色
	StudentType           []ConfBaseModel `json:"student_type"`            //学员类型
	FullStudentType       []ConfBaseModel `json:"full_student_type"`       //已完成学员的学员类型
	DynamicType           []ConfBaseModel `json:"dynamic_type"`            //最新动态
	Sort                  ConfSortModel   `json:"sort"`                    //列表页面的筛选条件
}

type ConfSortModel struct {
	NeedTrack  [][]ConfBaseModel `json:"need_track"`
	NoTrack    [][]ConfBaseModel `json:"no_track"`
	Release    [][]ConfBaseModel `json:"release"`
	Dynamic    [][]ConfBaseModel `json:"dynamic"`
	Follow     [][]ConfBaseModel `json:"follow"`
	Fetch      [][]ConfBaseModel `json:"fetch"`
	Special    [][]ConfBaseModel `json:"special"`
	PublicList [][]ConfBaseModel `json:"public_list"`
}

type ConfBaseModel struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetConfDictFile() *ConfDictModel {

	var mobileProvincem []string
	rows, err := DataBase.Query("SELECT DISTINCT mobile_province name FROM student WHERE mobile_province != '' AND mobile_province IS NOT NULL AND mobile_province != '全国'")
	if err == nil {
		for rows.Next() {
			var name string
			rows.Scan(&name)
			mobileProvincem = append(mobileProvincem, name)
		}
	}

	grade := []int{1,2,3,4,5}

	talkMethod := []string{"电话", "微信", "QQ", "面谈", "其他"}
	talkStatus := []ConfBaseModel{
		ConfBaseModel{3, "A级"},
		ConfBaseModel{2, "B级"},
		ConfBaseModel{1, "C级"},
		ConfBaseModel{4, "无意向"},
	}

	var productType []ConfBaseModel
	rows, err = DataBase.Query("SELECT id, name FROM crm.product_type WHERE is_disabled = 0 ORDER BY id DESC")
	if err == nil {
		for rows.Next() {
			var product ConfBaseModel
			rows.Scan(&product.Id, &product.Name)
			productType = append(productType, product)
		}
	} else {
		log.Print(err)
	}

	studentLevel := []ConfBaseModel{
		ConfBaseModel{1, "A级"},
		ConfBaseModel{2, "B级"},
		ConfBaseModel{3, "C级"},
		ConfBaseModel{4, "无意向"},
		ConfBaseModel{5, "空错停"},
		ConfBaseModel{6, "分校号源下转"},
	}

	role := []ConfBaseModel{
		ConfBaseModel{0, "全部"},
		ConfBaseModel{1, "高校学员"},
		ConfBaseModel{2, "社会人员"},
	}

	studentType := []ConfBaseModel{
		ConfBaseModel{0, "全部"},
		ConfBaseModel{1, "非学员"},
		ConfBaseModel{2, "体验学员"},
	}

	fullstudentType := []ConfBaseModel{
		ConfBaseModel{0, "全部"},
		ConfBaseModel{1, "线下学员"},
		ConfBaseModel{2, "线上学员"},
	}

	dynamicType := []ConfBaseModel{
		ConfBaseModel{0, "全部"},
		ConfBaseModel{1, "有动态"},
		ConfBaseModel{2, "无动态"},
		ConfBaseModel{3, "完成注册"},
		ConfBaseModel{4, "最近登录"},
		ConfBaseModel{5, "学习新卡片"},
		ConfBaseModel{6, "刷题"},
	}

	return &ConfDictModel{
		mobileProvincem,
		grade,
		talkMethod,
		talkStatus,
		productType,
		studentLevel,
		role,
		studentType,
		fullstudentType,
		dynamicType,
		getSortConf(),
	}
}

func getSortConf() ConfSortModel {
	return ConfSortModel{
		[][]ConfBaseModel{
			[]ConfBaseModel{
				ConfBaseModel{1, "按照最近沟通时间正序进行排列"},
				ConfBaseModel{2, "按照最近沟通时间倒序进行排列"},
			},
		},

		[][]ConfBaseModel{
			[]ConfBaseModel{
				ConfBaseModel{1, "按照录入时间倒序进行排列"},
				ConfBaseModel{2, "按照录入时间正序进行排列"},
			},
		},

		[][]ConfBaseModel {
			[]ConfBaseModel {
				ConfBaseModel {1, "按照到期时间正序进行排列"},
				ConfBaseModel {2, "按照到期时间倒序进行排列"},
				ConfBaseModel {3, "按照最近沟通时间正序进行排列"},
				ConfBaseModel {4, "按照最近沟通时间倒序进行排列"},
			},
		},
		[][]ConfBaseModel{
			[]ConfBaseModel {
				ConfBaseModel {1, "按照注册时间倒序进行排列"},
				ConfBaseModel {2, "按照注册时间正序进行排列"},
			},
			[]ConfBaseModel {
				ConfBaseModel {1, "按照登录时间倒序进行排列"},
				ConfBaseModel {2, "按照登录时间正序进行排列"},
			},
			[]ConfBaseModel {
				ConfBaseModel {1, "按照学习时间倒序进行排列"},
				ConfBaseModel {2, "按照学习时间正序进行排列"},
				ConfBaseModel {3, "按照学习卡片数倒序进行排列"},
				ConfBaseModel {4, "按照学习卡片数正序进行排列"},
			},
		},
		[][]ConfBaseModel{
			[]ConfBaseModel {
				ConfBaseModel {1, "按照最近沟通时间正序进行排列"},
				ConfBaseModel {2, "按照最近沟通时间倒序进行排列"},
			},
		},
		[][]ConfBaseModel{
			[]ConfBaseModel {
				ConfBaseModel {1, "按照录入时间倒序进行排列"},
				ConfBaseModel {2, "按照录入时间正序进行排列"},
				ConfBaseModel {3, "按照沟通时间倒序进行排列"},
				ConfBaseModel {4, "按照沟通时间正序进行排列"},
			},
			[]ConfBaseModel {
				ConfBaseModel {1, "按照录入时间倒序进行排列"},
				ConfBaseModel {2, "按照录入时间正序进行排列"},
				ConfBaseModel {3, "按照沟通时间倒序进行排列"},
				ConfBaseModel {4, "按照沟通时间正序进行排列"},
			},
		},
		[][]ConfBaseModel{
			[]ConfBaseModel {
				ConfBaseModel {1, "按照下发工单时间倒序进行排列"},
				ConfBaseModel {2, "按照下发工单时间正序进行排列"},
				ConfBaseModel {3, "按照沟通时间倒序进行排列"},
				ConfBaseModel {4, "按照沟通时间正序进行排列"},
			},
			[]ConfBaseModel {
				ConfBaseModel {1, "按照下发工单时间倒序进行排列"},
				ConfBaseModel {2, "按照下发工单时间正序进行排列"},
				ConfBaseModel {3, "按照沟通时间倒序进行排列"},
				ConfBaseModel {4, "按照沟通时间正序进行排列"},
			},
		},
		[][]ConfBaseModel{
			[]ConfBaseModel {
				ConfBaseModel {1, "按照沟通时间倒序进行排列"},
				ConfBaseModel {2, "按照沟通时间正序进行排列"},
			},
			[]ConfBaseModel {
				ConfBaseModel {1, "按照沟通时间倒序进行排列"},
				ConfBaseModel {2, "按照沟通时间正序进行排列"},
			},
		},
	}
}
