package repository

import (
	"fmt"
	"log"

	sql "database/sql"

	"example.com/m/v2/database"
	"example.com/m/v2/domain"
	sqlbuilder "github.com/huandu/go-sqlbuilder"
)

type UserRepository struct {
}

func (c UserRepository) GetOne(id int) (*domain.User, error) {
	user := new(domain.User)
	sb := sqlbuilder.NewSelectBuilder()

	sb.Select("id", "iin", "first_name", "middle_name",
		"last_name", "phone", "username", "email", "password", "is_active",
		"active_directory_link", "reg_datetime", "otdel_id", "organization_id",
		"allow_self_registration")
	sb.From("app_user")
	sb.Where(sb.Equal("id", id))
	query, args := sb.Build()
	rows, err := database.DB.Query(query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	err = rows.Scan(&user.ID, &user.IIN, &user.FirstName, &user.MiddleName, &user.Email,
		&user.Password, &user.IsActive, &user.ActiveDirectoryLink, &user.RegDatetime,
		&user.OtdelID, &user.OrganizationID, &user.AllowSelfRegistration)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	fmt.Println("user entity :", user)
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (c UserRepository) Create(user *domain.User) (*domain.User, error) {
	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto("app_user")
	ib.Cols("iin", "first_name", "middle_name",
		"last_name", "phone", "username", "email", "password", "is_active",
		"active_directory_link", "reg_datetime", "otdel_id", "organization_id",
		"allow_self_registration")
	ib.Values(user.IIN, user.FirstName, user.MiddleName, user.Email,
		user.Password, user.IsActive, user.ActiveDirectoryLink, user.RegDatetime,
		user.OtdelID, user.OrganizationID, user.AllowSelfRegistration)
	query, args := ib.Build()
	fmt.Println("query : ", query)
	fmt.Println(args)
	rows, err := database.DB.Query(query, args...)
	if err != nil {
		log.Println("err : ", err)
		return nil, err
	}
	defer rows.Close()
	err = rows.Scan()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println("err : ", err)
		return nil, err
	}
	fmt.Println("user entity :", user)
	err = rows.Err()
	if err != nil {
		log.Println("err : ", err)
		return nil, err
	}

	return user, nil
}

func (c UserRepository) GrantRoleToUser(user *domain.User, role domain.Role) (*domain.User, error) {
	return nil, nil
}

func (c UserRepository) Update(user *domain.User) (*domain.User, error) {
	return nil, nil
}

func (c UserRepository) GetAll() ([]*domain.User, error) {
	return nil, nil
}

func (c UserRepository) FindByUsername() (*domain.User, error) {
	return nil, nil
}

type UserCriteria struct {
}
