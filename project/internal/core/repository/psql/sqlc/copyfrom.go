// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: copyfrom.go

package sqlc

import (
	"context"
)

// iteratorForCreatePocket implements pgx.CopyFromSource.
type iteratorForCreatePocket struct {
	rows                 []CreatePocketParams
	skippedFirstNextCall bool
}

func (r *iteratorForCreatePocket) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForCreatePocket) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].Name,
		r.rows[0].Title,
		r.rows[0].UserID,
		r.rows[0].CompanyID,
		r.rows[0].Type,
		r.rows[0].Icon,
		r.rows[0].Color,
		r.rows[0].SalesforceID,
	}, nil
}

func (r iteratorForCreatePocket) Err() error {
	return nil
}

func (q *Queries) CreatePocket(ctx context.Context, arg []CreatePocketParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"pocket"}, []string{"name", "title", "user_id", "company_id", "type", "icon", "color", "salesforce_id"}, &iteratorForCreatePocket{rows: arg})
}
