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
	"strconv"
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

	id, err := dao.CheckIfScheduleExist(param.ClassID, param.UserID, param.SemesterID)
	if err != nil {
		return err
	}
	// upsert the schedule
	if exist {
		err = dao.UpdateScheduleByID(id, param.ClassID, param.SemesterID)
		if err != nil {
			return err
		}
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
		CustomEvents: make([]dto.CustomEvent, 0),
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

		maxCredit, err := strconv.ParseFloat(course.MaxCredit, 32)
		if err != nil {
			maxCredit = 0.0
		}
		minCredit, err := strconv.ParseFloat(course.MinCredit, 32)
		if err != nil {
			minCredit = 0.0
		}

		classList, err := classDao.GetClassByCourseID(class.CourseID)
		if err != nil {
			return nil, err
		}

		scheduleClassInfo := &dto.ClassInfo{
			ClassData: *class,
			CourseInfo: dto.Course{
				CourseID:           course.ID,
				IsHonor:            course.IsHonor,
				FixedCredit:        course.FixedCredit,
				DesignationCatalog: course.DesignationCatalog,
				Description:        course.Description,
				Prerequisites:      course.Prerequisites,
				Title:              course.Title,
				CatalogCourseName:  course.CatalogCourseName,
				Component:          course.Component,
				MaxCredit:          maxCredit,
				MinCredit:          minCredit,
				Classes:            *classList,
				OverallRating:      rating.OverAllRating,
				Tags:               tagList,
			},
		}
		scheduleRes.ScheduledClassList = append(scheduleRes.ScheduledClassList, *scheduleClassInfo)
	}

	customEventList, err := dao.GetCustomEventByUserID(param.UserID)
	for _, event := range customEventList {
		customEvent := &dto.CustomEvent{
			ID: int32(event.ID),
			Title:       event.Title,
			Description: event.Description,
			Color:       event.Color,
			Days:        event.Days,
			StartTime:   event.StartTime,
			EndTime:     event.EndTime,
		}
		scheduleRes.CustomEvents = append(scheduleRes.CustomEvents, *customEvent)
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
