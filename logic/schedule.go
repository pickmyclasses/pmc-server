package logic

import (
	dao "pmc_server/dao/schedule"
	"pmc_server/model"
)

func CreateSchedule(param model.PostScheduleParams) error {
	err := dao.PostUserSchedule(param.UserID, param.ClassID, param.SemesterID)
	if err != nil {
		return err
	}
	return nil
}

func GetSchedule(param model.GetScheduleParams) (*model.Schedule, error) {
	schedule, err := dao.GetUserSchedule(param.UserID)
	if err != nil {
		return nil, err
	}
	return schedule, nil
}

func DeleteSchedule(param model.DeleteScheduleParams) error {
	err := dao.DeleteUserSchedule(param.UserID, param.SemesterID, param.ClassID)
	if err != nil {
		return err
	}
	return nil
}
