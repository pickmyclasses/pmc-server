package logic

import (
	"errors"
	classDao "pmc_server/dao/postgres/class"
	courseDao "pmc_server/dao/postgres/course"
	dao "pmc_server/dao/postgres/schedule"

	"pmc_server/model"
	"pmc_server/model/dto"
)

func CreateSchedule(param model.PostScheduleParams) error {
	exist, err := dao.CheckIfUserExist(param.UserID)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("user does not exist")
	}

	exist, err = dao.CheckIfClassExist(param.ClassID)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("class does not exist")
	}

	exist, err = dao.CheckIfScheduleExist(param.ClassID, param.UserID, param.SemesterID)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("schedule already exist")
	}

	err = dao.CreateSchedule(param.ClassID, param.UserID, param.SemesterID)
	if err != nil {
		return err
	}
	return nil
}

func GetSchedule(param model.GetScheduleParams) (*dto.Schedule, error) {
	exist, err := dao.CheckIfUserExist(param.UserID)
	if err != nil {
		return nil, err
	}

	if !exist {
		return nil, errors.New("user does not exist")
	}

	scheduleList, err := dao.GetScheduleByUserID(param.UserID)

	if err != nil {
		return nil, err
	}

	scheduleRes := &dto.Schedule{
		ScheduledClassList: make([]dto.ClassInfo, 0),
	}

	for _, schedule := range scheduleList {
		class, err := classDao.GetClassInfoByID(int(schedule.ClassID))
		if err != nil {
			return nil, err
		}
		course, err := courseDao.GetCourseByID(int(class.CourseID))
		if err != nil {
			return nil, err
		}
		scheduleClassInfo := &dto.ClassInfo{
			ClassData:  *class,
			CourseData: *course,
		}
		scheduleRes.ScheduledClassList = append(scheduleRes.ScheduledClassList, *scheduleClassInfo)
	}

	return scheduleRes, nil
}

func DeleteSchedule(param model.DeleteScheduleParams) error {
	err := dao.DeleteUserSchedule(param.UserID, param.SemesterID, param.ClassID)
	if err != nil {
		return err
	}
	return nil
}
