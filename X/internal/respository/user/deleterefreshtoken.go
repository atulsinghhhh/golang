package user

import "context"

func (r *UserRepos) DeleteRefreshToken(ctx context.Context, userID int64) error {
	query := `DELETE FROM refresh_tokens WHERE user_id = ?`

	rowAffected, err := r.db.ExecContext(ctx, query, userID)
	if err != nil {
		return err
	}
	affected, err := rowAffected.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return nil
	}
	return nil
}
