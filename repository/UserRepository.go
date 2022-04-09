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
		log.Println("db query : ", err)
		return nil, err
	}
	defer rows.Close()
	err = rows.Scan(&user.ID, &user.IIN, &user.FirstName, &user.MiddleName, &user.Email,
		&user.Password, &user.IsActive, &user.ActiveDirectoryLink, &user.RegDatetime,
		&user.OtdelID, &user.OrganizationID, &user.AllowSelfRegistration)
	if err != nil {
		log.Println("row Scan error : ", err)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	fmt.Println("user entity :", user)
	err = rows.Err()
	if err != nil {
		log.Println("rows.Err() error : ", err)
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
	log.Println("query : ", query)
	log.Println(args)
	rows, err := database.DB.Query(query, args...)
	if err != nil {
		log.Println("db query err : ", err)
		return nil, err
	}
	defer rows.Close()
	err = rows.Scan()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println("rows scan err : ", err)
		return nil, err
	}
	log.Println("user entity :", user)
	err = rows.Err()
	if err != nil {
		log.Println("rows.Err() err : ", err)
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

func (c UserRepository) nextID() (int, error) {
	var id int
	query := "select nextval('s_user')"
	rows, err := database.DB.Query(query, nil)

	if err != nil {
		log.Println("db query error  : ", err)
		return 0, err
	}

	err = rows.Scan(&id)
	if err != nil {
		log.Println("call nextID() row.Scan error  : ", err)
		return 0, err
	}

	return id, nil
}

type UserCriteria struct {
}
