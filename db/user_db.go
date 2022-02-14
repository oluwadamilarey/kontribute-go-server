package db

import (
	"context"
	"log"

	"github.com/Kontribute/kontribute-web-backend/dto"
	"github.com/Kontribute/kontribute-web-backend/entity"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type UserDB struct {
	log *log.Logger
	psql sq.StatementBuilderType
	conn *sqlx.DB
}

func NewUserDB(log *log.Logger, conn *sqlx.DB) UserDB {
	return UserDB{
		log: log,
		conn: conn,
		psql: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (u UserDB) CreateWebUser(ctx context.Context, info dto.UserWebRegisterDTO) (entity.User, error) {
	var user entity.User
	q, args, err := u.psql.Insert("users").SetMap(map[string]interface{} {
		"name": "",
		"email": info.Email,
	}).Suffix("RETURNING *").ToSql()

	if err != nil {
		return user, err
	}

	if err = u.conn.QueryRowxContext(ctx, q, args...).StructScan(&user); err != nil {
		return user, err
	}
	return user, nil
}

func (u UserDB) CreateGoal(info dto.GoalCreateDTO) {
	q, args, err := u.psql.Insert("users").SetMap(map[string]interface{} {
		"name": info.Title,
		"description": info.Description,
		"user_id": info.UserID,
	}).Suffix("RETURNING *").ToSql()
}