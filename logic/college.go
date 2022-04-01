package logic

import (
	dao "pmc_server/dao/postgres/college"
	"pmc_server/model/dto"
)

func GetCollegeSemesterList(collegeID int32) ([]dto.Semester, error) {
	semesterList, err := dao.GetCollegeSemesterList(collegeID)
	if err != nil {
		return nil, err
	}

	semesterDtoList := make([]dto.Semester, 0)
	for _, semester := range semesterList {
		college, err := dao.GetCollegeByID(collegeID)
		if err != nil {
			return nil, err
		}
		semesterDto := dto.Semester{
			CollegeName: college.Name,
			Year:        semester.Year,
			Season:      semester.Season,
		}
		semesterDtoList = append(semesterDtoList, semesterDto)
	}
	return semesterDtoList, nil
}
