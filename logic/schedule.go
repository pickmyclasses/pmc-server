package logic

import (
	majorDao "pmc_server/dao/postgres/major"
	userDao "pmc_server/dao/postgres/user"
	"pmc_server/init/postgres"
	"strconv"

	classDao "pmc_server/dao/postgres/class"
	courseDao "pmc_server/dao/postgres/course"
	historyDao "pmc_server/dao/postgres/history"
	reviewDao "pmc_server/dao/postgres/review"
	dao "pmc_server/dao/postgres/schedule"
	tagDao "pmc_server/dao/postgres/tag"
	"pmc_server/model"
	"pmc_server/model/dto"
	"pmc_server/shared"
)

func CreateSchedule(param model.PostEventParam) error {
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

	id, err := dao.CheckIfScheduleExist(param.ClassID, param.UserID, 2)
	if err != nil {
		return err
	}
	// upsert the schedule
	if exist {
		err = dao.UpdateScheduleByID(id, param.ClassID, 2)
		if err != nil {
			return err
		}
	}

	err = dao.CreateSchedule(param.ClassID, param.UserID, 2)
	if err != nil {
		return err
	}

	// update user course history
	class, err := classDao.GetClassByID(int(param.ClassID))
	if err != nil {
		return err
	}

	// TODO: fix this semesterID
	err = historyDao.CreateSingleUserCourseHistory(param.UserID, class.CourseID, 2, class.Instructors)
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
		CustomEvents:       make([]dto.CustomEvent, 0),
	}

	for _, schedule := range scheduleList {
		class, err := classDao.GetClassByID(int(schedule.ClassID))
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

		courseDto := dto.Course{
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
		}
		// check if the course is in user's major, if yes, add an extra attachment to it
		if param.UserID != 0 {
			user, err := userDao.GetUserByID(param.UserID)
			if err != nil {
				return nil, err
			}

			majorQuery := majorDao.Major{
				CollegeID: int32(user.CollegeID),
				Querier:   postgres.DB,
			}

			major, err := majorQuery.QueryMajorByName(user.Major)
			if err != nil {
				return nil, err
			}

			if major.Name == "" {
				courseDto.DegreeCatalogs = make([][]string, 0)
			}

			courseSetQuery := courseDao.CourseSet{
				MajorID: int32(major.ID),
				Querier: postgres.DB,
			}

			majorSetList, err := courseSetQuery.QueryMajorCourseSets()
			if err != nil {
				return nil, err
			}

			degreeCatalogList := make([][]string, 0)
			for _, set := range majorSetList {
				for _, cid := range set.CourseIDList {
					catalogTuple := make([]string, 0, 2)
					if cid == course.ID {
						if set.ParentSetID != -1 {
							parentSet, err := courseSetQuery.QueryCourseSetByID(set.ParentSetID)
							if err == nil && parentSet.Name != "" {
								catalogTuple = append(catalogTuple, parentSet.Name)
							}
						}
						catalogTuple = append(catalogTuple, set.Name)
						degreeCatalogList = append(degreeCatalogList, catalogTuple)
					}
				}
			}
		}

		scheduleClassInfo := &dto.ClassInfo{
			ClassData:  *class,
			CourseInfo: courseDto,
		}
		scheduleRes.ScheduledClassList = append(scheduleRes.ScheduledClassList, *scheduleClassInfo)
	}

	customEventList, err := dao.GetCustomEventByUserID(param.UserID)
	for _, event := range customEventList {
		customEvent := &dto.CustomEvent{
			ID:          int32(event.ID),
			Title:       event.Title,
			Description: event.Description,
			Color:       event.Color,
			Days:        event.Days,
			StartTime:   event.StartTime,
			EndTime:     event.EndTime,
			Kind:        event.Kind,
		}
		scheduleRes.CustomEvents = append(scheduleRes.CustomEvents, *customEvent)
	}

	return scheduleRes, nil
}

func DeleteSchedule(userID, classID int64) error {
	err := dao.DeleteUserSchedule(userID, classID)
	if err != nil {
		return err
	}

	class, err := classDao.GetClassByID(int(classID))
	if err != nil {
		return err
	}
	// TODO: fix this semesterID
	err = historyDao.DeleteSingleUserCourseHistory(userID, class.CourseID, 2)
	if err != nil {
		return err
	}
	return nil
}

func CreateCustomEvent(param model.PostEventParam) error {
	exist, err := dao.CheckIfUserExist(param.UserID)
	if err != nil {
		return err
	}
	if !exist {
		return shared.ContentNotFoundErr{}
	}

	if !param.IsNew {
		if param.Event.EventID == 0 {
			return shared.ParamIncompatibleErr{}
		}
		exist, err := dao.CheckIfCustomEventExist(param.Event.EventID)
		if err != nil {
			return err
		}
		if !exist {
			return shared.ContentNotFoundErr{}
		}

		err = dao.UpdateCustomEventByID(param.Event.EventID, param.UserID, 2, param.Event.Title, param.Event.Description,
			param.Event.Color, param.Event.Days, param.Event.StartTime, param.Event.EndTime, param.Event.Kind)

		if err != nil {
			return err
		}
		return nil
	} else {
		err = dao.CreateCustomEventByUserID(param.UserID, 2, param.Event.Title, param.Event.Description,
			param.Event.Color, param.Event.Days, param.Event.StartTime, param.Event.EndTime, param.Event.Kind)
		if err != nil {
			return err
		}
		return nil
	}
}

func DeleteCustomEvent(id int64) error {
	err := dao.DeleteCustomEvent(id)
	if err != nil {
		return err
	}
	return nil
}
