package postgres

import (
	"context"
	"database/sql"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type otpRepo struct {
	db  *pgxpool.Pool
	log logger.ILogger
}

func NewOTPRepo(db *pgxpool.Pool, log logger.ILogger) storage.IOTPStorage {
	return &otpRepo{
		db:  db,
		log: log,
	}
}

func (o *otpRepo) Create(ctx context.Context, otp models.OTP) error {
	query := `
		INSERT INTO otps (id, user_id, code, expires_at)
		VALUES ($1, $2, $3, $4)
	`
	_, err := o.db.Exec(ctx, query, otp.ID, otp.UserID, otp.Code, otp.ExpiresAt)
	return err
}

func (o *otpRepo) GetByUserIDAndCode(ctx context.Context, userID, code string) (models.OTP, error) {
	var otp models.OTP
	query := `
		SELECT id, user_id, code, expires_at, used_at
		FROM otps
		WHERE user_id = $1 AND code = $2
	`
	err := o.db.QueryRow(ctx, query, userID, code).Scan(&otp.ID, &otp.UserID, &otp.Code, &otp.ExpiresAt, &otp.UsedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return otp, nil
		}
		return otp, err
	}
	return otp, nil
}

func (o *otpRepo) UpdateUsedAt(ctx context.Context, otpID string, usedAt time.Time) error {
	query := `
		UPDATE otps
		SET used_at = $1
		WHERE id = $2
	`
	_, err := o.db.Exec(ctx, query, usedAt, otpID)
	return err
}
