package models

import (
	"strconv"
	"time"
)

type Problem struct {
	ProblemId    int32     `gorm:"problem_id"`
	Title        string    `gorm:"title"`
	Description  string    `gorm:"description"`
	Input        string    `gorm:"input"`
	Output       string    `gorm:"output"`
	SampleInput  string    `gorm:"sample_input"`
	SampleOutput string    `gorm:"sample_output"`
	Spj          string    `gorm:"spj"`
	Hint         string    `gorm:"hint"`
	Source       string    `gorm:"source"`
	InDate       time.Time `gorm:"in_date"`
	TimeLimit    int32     `gorm:"time_limit"`
	MemoryLimit  int32     `gorm:"memory_limit"`
	Defunct      string    `gorm:"defunct"`
	Accepted     int32     `gorm:"accepted"`
	Submit       int32     `gorm:"submit"`
	Solved       int32     `gorm:"solved"`
}

func (Problem) TableName() string {

	return "problem"
}

func stringToint32(str string) int32 {
	d, _ := strconv.Atoi(str)
	return int32(d)
}

//func (pr *Problem)QueryAllProblem() ([]*Problem, int64, error) {
//	var pro []*Problem
//	problem := new(Problem)
//	qs := dao.Orm.Model()
//	num, err := qs.OrderBy("-problem_id").All(&pro)
//	if err != nil {
//		return nil, num, err
//	}
//	return pro, num, nil
//}
//
//func QueryPageProblem(start, pageSize int) ([]*Problem, int64, int64, error) {
//	var pro []*Problem
//	problem := new(Problem)
//	qs := DB.QueryTable(problem)
//	totalNum, _ := qs.Count()
//	num, err := qs.OrderBy("-problem_id").Limit(pageSize, start).All(&pro)
//	if err != nil {
//		return nil, num, totalNum, err
//	}
//	return pro, num, totalNum, nil
//}
//
//func UpdateProStatus(pid int32) bool {
//	pro := Problem{ProblemId: pid}
//	if DB.Read(&pro) == nil {
//		if pro.Defunct == "Y" {
//			pro.Defunct = "N"
//		} else {
//			pro.Defunct = "Y"
//		}
//		if num, err := DB.Update(&pro); err == nil {
//			logs.Info(num)
//			return true
//		}
//	}
//	return false
//}
//
//func QueryUserProblem(uid int32) ([]*Problem, int64, error) {
//	var data []*Solution
//	Solutions := new(Solution)
//	qss := DB.QueryTable(Solutions)
//	_, err := qss.Filter("user_id", uid).Filter("result", 4).Limit(10).OrderBy("-solution_id").All(&data, "problem_id", "time", "in_date", "memory", "language")
//	if err != nil {
//		return nil, 0, err
//	}
//
//	var proIds []int32
//	for _, v := range data {
//		proIds = append(proIds, v.ProblemId)
//	}
//
//	var pro []*Problem
//	if proIds == nil {
//		return pro, 0, nil
//	}
//	problem := new(Problem)
//	qs := DB.QueryTable(problem)
//	num, err := qs.Filter("problem_id__in", proIds).OrderBy("-problem_id").All(&pro)
//	if err != nil {
//		return nil, num, err
//	}
//	return pro, num, nil
//}
//
//func QueryProblemById(id int32) (Problem, error) {
//	pro := Problem{ProblemId: id}
//	err := DB.Read(&pro, "ProblemId")
//	if err != nil {
//		return Problem{}, err
//	}
//	return pro, nil
//}
//
//func QueryProblemByTitle(title string) (Problem, error) {
//	pro := Problem{Title: title}
//	err := DB.Read(&pro, "title")
//	if err != nil {
//		return Problem{}, err
//	}
//	return pro, nil
//}
//
//func AddProblem(data []string, inDate time.Time) (int64, error) {
//	var pro Problem
//	pro.Title = data[0]
//	pro.TimeLimit = stringToint32(data[1])
//	pro.MemoryLimit = stringToint32(data[2])
//	pro.Description = data[3]
//	pro.Input = data[4]
//	pro.Output = data[5]
//	pro.SampleInput = data[6]
//	pro.SampleOutput = data[7]
//	pro.Hint = data[8]
//	pro.Spj = data[9]
//	pro.InDate = inDate
//	pro.Defunct = "N"
//
//	pid, err := DB.Insert(&pro)
//	if err != nil {
//		return pid, err
//	}
//	return pid, nil
//}
//
//func BatchAddProblem(problems []types.Problem) (map[int32]int64, error) {
//	//TOTO 加上事务
//	data := make(map[int32]int64)
//	for _, v := range problems {
//		var pro Problem
//		pro.Title = v.Title
//		pro.TimeLimit = v.TimeLimit
//		pro.MemoryLimit = v.MemoryLimit
//		pro.Description = v.Description
//		pro.Input = v.Input
//		pro.Output = v.Output
//		pro.SampleInput = v.SampleInput
//		pro.SampleOutput = v.SampleOutput
//		pro.Hint = v.Hint
//		pro.Spj = v.Spj
//		pro.InDate = time.Now()
//		pro.Defunct = pro.Defunct
//
//		pid, err := DB.Insert(&pro)
//		if err != nil {
//			continue
//		}
//
//		data[v.ProblemId] = pid
//	}
//
//	return data, nil
//}
//
//func UpdateProblemById(id int32, data []string, inDate time.Time) (bool, error) {
//	pro := Problem{ProblemId: id}
//	if DB.Read(&pro) == nil {
//		pro.Title = data[0]
//		pro.TimeLimit = stringToint32(data[1])
//		pro.MemoryLimit = stringToint32(data[2])
//		pro.Description = data[3]
//		pro.Input = data[4]
//		pro.Output = data[5]
//		pro.SampleInput = data[6]
//		pro.SampleOutput = data[7]
//		pro.Hint = data[8]
//		pro.Spj = data[9]
//		pro.InDate = inDate
//		if num, err := DB.Update(&pro); err == nil {
//			logs.Info(num)
//			return true, err
//		}
//	}
//	return false, syserror.UpdateProErr()
//}
//
//func DelProblemById(id int32) bool {
//	if num, err := DB.Delete(&Problem{ProblemId: id}); err == nil {
//		logs.Info(num)
//		return true
//	}
//	return false
//}