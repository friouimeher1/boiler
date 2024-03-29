package resolver

import (
	"context"

	"github.com/rafaelsq/boiler/pkg/graphql/internal/entity"
	"github.com/rafaelsq/boiler/pkg/iface"
)

func NewUser(service iface.Service) *User {
	return &User{
		service: service,
	}
}

type User struct {
	service iface.Service
}

func (*User) ID(ctx context.Context, u *entity.User) (int, error) {
	return u.ID, nil
}

func (r *User) User(ctx context.Context, userID int) (*entity.User, error) {
	u, err := r.service.GetUserByID(ctx, userID)
	if err == nil {
		return entity.NewUser(u), nil
	}
	return nil, err
}

func (r *User) Users(ctx context.Context, limit uint) ([]*entity.User, error) {
	us, err := r.service.FilterUsers(ctx, iface.FilterUsers{Limit: limit})
	if err == nil {
		users := make([]*entity.User, 0, len(us))
		for _, u := range us {
			users = append(users, entity.NewUser(u))
		}
		return users, nil
	}
	return nil, err
}

func (r *User) Emails(ctx context.Context, u *entity.User) ([]*entity.Email, error) {
	es, err := r.service.FilterEmails(ctx, iface.FilterEmails{UserID: u.ID})
	if err == nil {
		emails := make([]*entity.Email, 0, len(es))
		for _, e := range es {
			emails = append(emails, entity.NewEmail(e))
		}
		return emails, nil
	}
	return nil, err
}
