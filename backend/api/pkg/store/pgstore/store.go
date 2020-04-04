package pgstore

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/huandu/go-sqlbuilder"
	"github.com/lib/pq"
	"github.com/nchern/homevscorona/backend/api/pkg/model"
)

type UserStore struct {
	db *sql.DB
}

func New() *UserStore {
	return &UserStore{
		db: MustOpenDB(),
	}

}

func (u *UserStore) Create(email string, user *model.User) error {
	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	builder := (&UserListInsertCommand{
		ID:      user.ID,
		Email:   user.Email,
		Created: time.Now().UnixNano(),
		Updated: time.Now().UnixNano(),
		Status:  UserstatusHealthy,
		Body:    body,
	}).ToBuilder()

	stmt, args := builder.Build()
	_, err = u.db.Exec(stmt, args...)

	switch e := err.(type) {
	case *pq.Error:
		if e.Code.Name() == "unique_violation" {
			//			return fmt.Errorf("%w: %s", ErrUniqueViolation, err)
			return nil
		}
	}

	return err
}

func (u *UserStore) GetByEmail(email string) (*model.User, error) {
	builder := sqlbuilder.PostgreSQL.NewSelectBuilder()
	stmt, args := builder.
		Select(
			UserList.ID(),
			UserList.Body(),
		).
		From(UserListTableName).
		Where(
			builder.Equal(UserList.Email(), email),
		).
		Build()

	res := u.db.QueryRow(stmt, args...)
	tuple := &UserListTuple{}
	if err := res.Scan(&tuple.ID, &tuple.Body); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	var user model.User
	if err := json.Unmarshal([]byte(tuple.Body), &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserStore) GetByID(id string) (*model.User, error) {
	return nil, errors.New("GetById not implemented")
}

func (u *UserStore) SaveEvent(userID string, event *model.Event) error {
	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	builder := (&EventInsertCommand{
		ID:      event.ID,
		UserID:  userID,
		Created: time.Now().UnixNano(),
		Body:    body,
	}).ToBuilder()

	stmt, args := builder.Build()
	_, err = u.db.Exec(stmt, args...)
	return err
}

func (u *UserStore) GetEvents(userID string) ([]*model.Event, error) {
	results := []*model.Event{}

	builder := sqlbuilder.PostgreSQL.NewSelectBuilder()
	stmt, args := builder.
		Select(
			Event.ID(),
			Event.Body(),
		).
		From(EventTableName).
		Where(
			builder.Equal(Event.UserID(), userID),
		).
		Build()

	rows, err := u.db.Query(stmt, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		tuple := EventTuple{}

		if err := rows.Scan(
			&tuple.ID,
			&tuple.Body,
		); err != nil {
			return nil, err
		}

		event := &model.Event{}
		if err := json.Unmarshal([]byte(tuple.Body), event); err != nil {
			return nil, err
		}
		results = append(results, event)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return results, nil
}
