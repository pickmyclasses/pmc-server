package logic

import (
	classDao "pmc_server/dao/postgres/class"
	courseDao "pmc_server/dao/postgres/course"
	reviewDao "pmc_server/dao/postgres/review"
	dao "pmc_server/dao/postgres/schedule"
	tagDao "pmc_server/dao/postgres/tag"
	"pmc_server/model"
	"pmc_server/model/dto"
	"pmc_server/shared"
)

func CreateSchedule(param model.PostScheduleParams) error {
	exist, err := dao.CheckIfUserExist(param.UserID)
	if err != nil {
		return err
	}
	if !exist {
		return shared.ContentNotFoundErr{}
	}

	exist, err = dao.CheckIfClassExist(param.ClassID)
	if err != nil {
		return err
	}
	if !exist {
		return shared.ContentNotFoundErr{}
	}

	exist, err = dao.CheckIfScheduleExist(param.ClassID, param.UserID, param.SemesterID)
	if err != nil {
		return err
	}
	if exist {
		return shared.ResourceConflictErr{}
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
		return nil, shared.ContentNotFoundErr{}
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
		tagList, err := tagDao.GetTagListByCourseID(class.CourseID)
		if err != nil {
			return nil, err
		}

		rating, err := reviewDao.GetCourseOverallRating(class.CourseID)
		if err != nil {
			return nil, err
		}

		scheduleClassInfo := &dto.ClassInfo{
			ClassData:  *class,
			CourseInfo: dto.CourseInfo{
				OverallRating: rating.OverAllRating,
				CourseData: *course,
				CourseTags: tagList,
			},
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
