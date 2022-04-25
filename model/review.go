package model

type Review struct {
	Base
	Rating             float32  `gorm:"column:rating" json:"rating"`
	Anonymous          bool     `gorm:"column:anonymous" json:"anonymous"`
	Recommended        bool     `gorm:"column:recommended" json:"recommended"`
	Pros               string   `gorm:"column:pros" json:"pros"`
	Cons               string   `gorm:"column:cons" json:"cons"`
	Comment            string   `gorm:"column:comment" json:"comment"`
	CourseID           int64    `gorm:"column:course_id" json:"courseID"`
	UserID             int64    `gorm:"column:user_id" json:"userID"`
	LikeCount          int32    `gorm:"column:like_count" json:"likeCount"`
	DislikeCount       int32    `gorm:"column:dislike_count" json:"dislikeCount"`
	HourSpent          int      `gorm:"column:hour_spent" json:"hourSpent"`
	GradeReceived      string   `gorm:"column:grade_received" json:"gradeReceived"`
	ExamHeavy          bool     `gorm:"column:exam_heavy" json:"examHeavy"`
	HomeworkHeavy      bool     `gorm:"column:homework_heavy" json:"homeworkHeavy"`
	ExtraCreditOffered bool     `gorm:"column:extra_credit_offered" json:"ExtraCreditOffered"`
	Tags               []string `gorm:"column:tags" json:"tags"`
}
