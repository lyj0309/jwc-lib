package dbLib

import (
	"time"
)

//Classroom 表结构体
type Classroom struct {
	Week    int
	Day     int
	Id      int `gorm:"primaryKey"`
	Empty   int
	Room    string
	Time    int
	House   string
	Course  string
	Class   string
	Teacher string
	BuildId int //某个学院楼教室ID,遍历要
}

// CultivateScheme 培养方案表， Department 学院 Major 专业 Grade 年级
type CultivateScheme struct {
	Department string `gorm:"primaryKey"`
	Major      string `gorm:"primaryKey"`
	Grade      string `gorm:"primaryKey"`
	Data       string
}

// ExamAlarm 考试提醒
type ExamAlarm struct {
	Subject   string     `gorm:"primaryKey"`
	StartTime *time.Time //开始时间
	Time      string     `gorm:"primaryKey"` //考试时间
	Class     string     `gorm:"primaryKey"`
	Users     string
}

type SessionType struct {
	Session string
	ID      string
	IP      string
	Time    int
}

type CsList map[string]map[string][]string

// Cs 培养方案 Majors 专业 Grade 年级
type Cs struct {
	Majors map[string][]string
}

//*
// 信息学院 map[ 自动化map[19，20] ]
//
//*/

// 课程提醒

type ClassRemind struct {
	Time    time.Time `gorm:"primaryKey"`
	Class   string    `gorm:"primaryKey"`
	When    string
	Teacher string `gorm:"primaryKey"`
	Name    string
	Users   string
}

//TempUser 临时用户，关注了公众号但是没登录小程序的
type TempUser struct {
	UnionId string `gorm:"primaryKey"`
	OpenId  string
}

type CourseTimetableRemindForm struct {
	Week    string `json:"week"`
	Day     string `json:"day"`
	Start   string `json:"start"`
	End     string `json:"end"`
	Name    string `json:"name"`
	Teacher string `json:"teacher"`
	Class   string `json:"class"`
}
