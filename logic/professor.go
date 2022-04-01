package logic

import (
	classDao "pmc_server/dao/postgres/class"
	dao "pmc_server/dao/postgres/professor"
	"pmc_server/model/dto"
)

func GetProfessorList() ([]dto.Professor, error) {
	professorList, err := dao.GetProfessorList()
	if err != nil {
		return nil, err
	}
	professorDtoList := make([]dto.Professor, 0)
	for _, professor := range professorList {
		professorDto := dto.Professor{
			ProfessorName: professor.Name,
		}
		professorDtoList = append(professorDtoList, professorDto)
	}
	return professorDtoList, nil
}

func GetProfessorListByCourseID(courseID int64) ([]dto.Professor, error) {
	classList, err := classDao.GetClassByCourseID(courseID)
	if err != nil {
		return nil, err
	}

	professorList := make([]dto.Professor, 0)
	for _, class := range *classList {
		if class.Instructors != "" {
			professor := dto.Professor{ProfessorName: class.Instructors}
			professorList = append(professorList, professor)
		}
	}

	return professorList, nil
}
