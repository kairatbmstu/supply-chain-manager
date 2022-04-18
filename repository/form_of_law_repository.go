package repository

import (
	sql "database/sql"
	"log"

	"example.com/m/v2/database"
	"example.com/m/v2/domain"
	sqlbuilder "github.com/huandu/go-sqlbuilder"
)

var FormOfLawRepositoryInstance FormOfLawRepository

type FormOfLawRepository struct {
}

func (c FormOfLawRepository) FindAll() ([]domain.FormOfLaw, error) {

	sb := sqlbuilder.NewSelectBuilder()

	sb.Select("id", "code", "is_resident", "name_ru", "name_kz", "name_en")
	sb.From("form_of_law")
	query, args := sb.Build()
	rows, err := database.DB.Query(query, args...)

	if err != nil {
		log.Println("error : ", err)
		return nil, err
	}
	defer rows.Close()

	result := make([]domain.FormOfLaw, 0)
	for rows.Next() {
		formoflaw := new(domain.FormOfLaw)
		err = rows.Scan(&formoflaw.ID, &formoflaw.Code, &formoflaw.IsResident,
			&formoflaw.NameRu, &formoflaw.NameKz, &formoflaw.NameEn)
		if err != nil {
			log.Println("err : ", err)
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}

		result = append(result, *formoflaw)
	}

	return result, nil
}

func (c FormOfLawRepository) FindResident() ([]domain.FormOfLaw, error) {

	sb := sqlbuilder.NewSelectBuilder()

	sb.Select("id", "code", "is_resident", "name_ru", "name_kz", "name_en")
	sb.From("form_of_law")
	sb.Where("is_resident = 'true'")
	query, args := sb.Build()
	rows, err := database.DB.Query(query, args...)

	if err != nil {
		log.Println("error : ", err)
		return nil, err
	}
	defer rows.Close()

	result := make([]domain.FormOfLaw, 0)
	for rows.Next() {
		formoflaw := new(domain.FormOfLaw)
		err = rows.Scan(&formoflaw.ID, &formoflaw.Code, &formoflaw.IsResident,
			&formoflaw.NameRu, &formoflaw.NameKz, &formoflaw.NameEn)
		if err != nil {
			log.Println("err : ", err)
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}

		result = append(result, *formoflaw)
	}

	return result, nil
}

func (c FormOfLawRepository) FindNonResident() ([]domain.FormOfLaw, error) {

	sb := sqlbuilder.NewSelectBuilder()

	sb.Select("id", "code", "is_resident", "name_ru", "name_kz", "name_en")
	sb.From("form_of_law")
	sb.Where("is_resident = 'false'")
	query, args := sb.Build()
	rows, err := database.DB.Query(query, args...)

	if err != nil {
		log.Println("error : ", err)
		return nil, err
	}
	defer rows.Close()

	result := make([]domain.FormOfLaw, 0)
	for rows.Next() {
		formoflaw := new(domain.FormOfLaw)
		err = rows.Scan(&formoflaw.ID, &formoflaw.Code, &formoflaw.IsResident,
			&formoflaw.NameRu, &formoflaw.NameKz, &formoflaw.NameEn)
		if err != nil {
			log.Println("err : ", err)
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}

		result = append(result, *formoflaw)
	}

	return result, nil
}
