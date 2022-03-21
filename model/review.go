package model

type Review struct {
	Base
	Rating                  float32 `gorm:"column:rating" json:"rating"`
	Anonymous               bool    `gorm:"column:anonymous" json:"anonymous"`
	Recommended             bool    `gorm:"column:recommended" json:"recommended"`
	Pros                    string  `gorm:"column:pros" json:"pros"`
	Cons                    string  `gorm:"column:cons" json:"cons"`
	Comment                 string  `gorm:"column:comment" json:"comment"`
	CourseID                int64   `gorm:"column:course_id" json:"courseID"`
	UserID                  int64   `gorm:"column:user_id" json:"userID"`
	LikeCount               int32   `gorm:"column:like_count" json:"likeCount"`
	DislikeCount            int32   `gorm:"column:dislike_count" json:"dislikeCount"`
	HourSpentBeyondCredit  bool    `gorm:"hour_spent_beyond_credit" json:"hourSpentBeyondCredit"`
	GradeReceived           string  `gorm:"grade_received" json:"gradeReceived"`
	ExamHeavy               bool    `gorm:"exam_heavy" json:"examHeavy"`
	HomeworkHeavy           bool    `gorm:"homework_heavy" json:"homeworkHeavy"`
	ExtraCreditOffered      bool    `gorm:"extra_credit_offered" json:"ExtraCreditOffered"`
	HourSpentLessThanCredit bool    `gorm:"hour_spent_less_than_credit" json:"hourSpentLessThanCredit"`
}
