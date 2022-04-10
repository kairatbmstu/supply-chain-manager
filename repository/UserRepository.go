package repository

import (
	"fmt"
	"log"

	sql "database/sql"

	"example.com/m/v2/database"
	"example.com/m/v2/domain"
	sqlbuilder "github.com/huandu/go-sqlbuilder"
)

var UserRepositoryInstance UserRepository

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
		log.Println("error : ", err)
		return nil, err
	}
	defer rows.Close()
	err = rows.Scan(&user.ID, &user.IIN, &user.FirstName, &user.MiddleName, &user.Email,
		&user.Password, &user.IsActive, &user.ActiveDirectoryLink, &user.RegDatetime,
		&user.OtdelID, &user.OrganizationID, &user.AllowSelfRegistration)
	if err != nil {
		log.Println("err : ", err)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	fmt.Println("error :", user)
	return user, nil
}

func (c UserRepository) Create(user *domain.User) (*domain.User, error) {

	id, err := c.nextID()
	if err != nil {
		log.Println(" error : ", err)
		return nil, err
	}

	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto("app_user")
	ib.Cols("id", "iin", "first_name", "middle_name",
		"last_name", "phone", "username", "email", "password", "is_active",
		"active_directory_link", "reg_datetime", "otdel_id", "organization_id",
		"allow_self_registration")
	ib.Values(id, user.IIN, user.FirstName, user.MiddleName, user.Email,
		user.Password, user.IsActive, user.ActiveDirectoryLink, user.RegDatetime,
		user.OtdelID, user.OrganizationID, user.AllowSelfRegistration)
	query, args := ib.Build()
	log.Println("query : ", query)
	log.Println(args)

	_, err = database.DB.Exec(query, args...)
	if err != nil {
		log.Println(" error : ", err)
		return nil, err
	}

	user.ID = id
	if err != nil {
		log.Println(" error : ", err)
		return nil, err
	}

	return user, nil
}

func (c UserRepository) GrantRoleToUser(user *domain.User, role domain.Role) (*domain.User, error) {
	return nil, nil
}

func (c UserRepository) Update(user *domain.User) (*domain.User, error) {
	ub := sqlbuilder.NewUpdateBuilder()
	query, args := ub.Update("app_user").
		Set(
			"visited = visited + 1",
		).
		Where(
			"id = 1234",
		).Build()

	database.DB.Exec(query, args)
	return nil, nil
}

func (c UserRepository) GetAll() ([]*domain.User, error) {
	return nil, nil
}

func (c UserRepository) FindByUsername(username string) (*domain.User, error) {
	user := new(domain.User)
	sb := sqlbuilder.NewSelectBuilder()

	sb.Select("id", "iin", "first_name", "middle_name",
		"last_name", "phone", "username", "email", "password", "is_active",
		"active_directory_link", "reg_datetime", "otdel_id", "organization_id",
		"allow_self_registration")
	sb.From("app_user")
	sb.Where(sb.Equal("username", username))
	query, args := sb.Build()
	rows, err := database.DB.Query(query, args...)

	if err != nil {
		log.Println("error : ", err)
		return nil, err
	}
	defer rows.Close()
	err = rows.Scan(&user.ID, &user.IIN, &user.FirstName, &user.MiddleName, &user.Email,
		&user.Password, &user.IsActive, &user.ActiveDirectoryLink, &user.RegDatetime,
		&user.OtdelID, &user.OrganizationID, &user.AllowSelfRegistration)
	if err != nil {
		log.Println("err : ", err)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (c UserRepository) FindByEmail(email string) (*domain.User, error) {
	user := new(domain.User)
	sb := sqlbuilder.NewSelectBuilder()

	sb.Select("id", "iin", "first_name", "middle_name",
		"last_name", "phone", "username", "email", "password", "is_active",
		"active_directory_link", "reg_datetime", "otdel_id", "organization_id",
		"allow_self_registration")
	sb.From("app_user")
	sb.Where(sb.Equal("email", email))
	query, args := sb.Build()
	rows, err := database.DB.Query(query, args...)

	if err != nil {
		log.Println("error : ", err)
		return nil, err
	}
	defer rows.Close()
	err = rows.Scan(&user.ID, &user.IIN, &user.FirstName, &user.MiddleName, &user.Email,
		&user.Password, &user.IsActive, &user.ActiveDirectoryLink, &user.RegDatetime,
		&user.OtdelID, &user.OrganizationID, &user.AllowSelfRegistration)
	if err != nil {
		log.Println("err : ", err)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (c UserRepository) nextID() (int, error) {
	var id int
	query := "select nextval('s_user')"
	rows, err := database.DB.Query(query, nil)

	if err != nil {
		log.Println("error  : ", err)
		return 0, err
	}

	err = rows.Scan(&id)
	if err != nil {
		log.Println("error  : ", err)
		return 0, err
	}

	return id, nil
}
