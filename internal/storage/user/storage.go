package userstorage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/TemaKut/messenger-auth/internal/app/logger"
	usermodels "github.com/TemaKut/messenger-auth/internal/models/user"
	pgconn "github.com/jackc/pgx/v5/pgconn"
)

type Storage struct {
	postgresDb *sql.DB
	logger     *logger.Logger
}

func NewStorage(postgresDb *sql.DB, logger *logger.Logger) *Storage {
	return &Storage{
		postgresDb: postgresDb,
		logger:     logger,
	}
}

type UserCreateParams struct {
	Name         string
	LastName     string
	Email        string
	PasswordHash string
}

func (s *Storage) UserCreate(ctx context.Context, user *usermodels.User) error {
	userDbo := decodeUser(user)

	setMap := map[string]any{
		usersIdColumn:    userDbo.Id,
		usersEmailColumn: userDbo.Email,
		usersDataColumn:  userDbo.Data,
	}

	query := sq.Insert(usersTableName).SetMap(setMap).PlaceholderFormat(sq.Dollar)
	if _, err := query.RunWith(s.postgresDb).ExecContext(ctx); err != nil {
		return fmt.Errorf("error exec query. %w", s.encodeError(err))
	}

	return nil
}

func (s *Storage) UserByEmail(ctx context.Context, email string) (*usermodels.User, error) {
	query := sq.Select(
		usersIdColumn,
		usersEmailColumn,
		usersDataColumn,
	).
		From(usersTableName).
		Where(sq.Eq{usersEmailColumn: email}).
		PlaceholderFormat(sq.Dollar)

	var userDbo UserDbo

	err := query.RunWith(s.postgresDb).QueryRowContext(ctx).Scan(&userDbo.Id, &userDbo.Email, &userDbo.Data)
	switch {
	case err == nil:
	case errors.Is(err, sql.ErrNoRows):
		return nil, fmt.Errorf("error query row. %w. %w", ErrUserNotFound, err)
	default:
		return nil, fmt.Errorf("error exec query row. %w", s.encodeError(err))
	}

	return encodeUser(userDbo), nil
}

func (s *Storage) encodeError(err error) error {
	if err == nil {
		return nil
	}

	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {
		if pgErr.TableName == usersTableName &&
			pgErr.Code == "23505" &&
			pgErr.ConstraintName == userEmailConstraintKey {
			return ErrUserEmailAlreadyExists
		}
	}

	s.logger.Debugf("stroage ubnknown error. %+v", err)

	return ErrUnknown
}
