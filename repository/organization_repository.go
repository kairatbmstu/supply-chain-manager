package repository

import (
	"database/sql"
	"example.com/m/v2/database"
	"example.com/m/v2/domain"
	"github.com/huandu/go-sqlbuilder"
	"log"
)

var OrganizationRepositoryInstance OrganizationRepository

type OrganizationQueryCriteria struct {
	BIN  string
	GLN  string
	IBAN string
}

type OrganizationRepository struct {
}

func (OrganizationRepository) FindByCriteria(criteria OrganizationQueryCriteria) (*domain.Organization, error) {
	organization := new(domain.Organization)
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()

	sb.Select("id", "external_id", "bin", "name_ru",
		"name_kz", "name_en", "is_resident", "is_blocked", "corp_phone",
		"contact_fio",
		"website", "gln",
		"is_stm", "is_nds",
		"iban", "bic", "kbe_code", "address", "form_of_law_id")
	sb.From("app_user")

	if criteria.BIN != "" {
		sb.Where(sb.Equal("bin", criteria.BIN))
	} else if criteria.GLN != "" {
		sb.Where(sb.Equal("gln", criteria.GLN))
	} else if criteria.IBAN != "" {
		sb.Where(sb.Equal("iban", criteria.IBAN))
	}

	query, args := sb.Build()
	rows, err := database.DB.Query(query, args...)

	if err != nil {
		log.Println("error : ", err)
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}
	err = rows.Scan(&organization.ID, &organization.ExternalId, &organization.Bin,
		&organization.NameRu, &organization.NameKz,
		&organization.NameEn, &organization.IsResident,
		&organization.IsResident, &organization.Blocked,
		&organization.CorpPhone,
		&organization.ContactFio, &organization.Website,
		&organization.Gln,
		&organization.IsStm,
		&organization.IsNds,
		&organization.Iban,
		&organization.KBE,
		&organization.Address,
		&organization.FormOfLaw.ID,
	)
	if err != nil {
		log.Println("err : ", err)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return organization, nil
}

func (OrganizationRepository) Create(organization *domain.Organization) (*domain.Organization, error) {
	return nil, nil
}

func (OrganizationRepository) Update(organization *domain.Organization) (*domain.Organization, error) {
	return nil, nil
}

func (OrganizationRepository) Delete(organization *domain.Organization) error {
	return nil
}

func (c OrganizationRepository) nextID() (int, error) {
	var id int
	query := "select nextval('s_organization')"
	rows, err := database.DB.Query(query, nil)

	if err != nil {
		log.Println("error  : ", err)
		return 0, err
	}
	if !rows.Next() {
		return 0, nil
	}

	err = rows.Scan(&id)
	if err != nil {
		log.Println("error  : ", err)
		return 0, err
	}

	return id, nil
}
