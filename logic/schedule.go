package logic

import (
	dao "pmc_server/dao/schedule"
	"pmc_server/model"
)

func CreateSchedule(param model.ScheduleParams) error {
	err := dao.PostUserSchedule(param.UserID, param.ClassID, param.SemesterID)
	if err != nil {
		return err
	}
	return nil
}
