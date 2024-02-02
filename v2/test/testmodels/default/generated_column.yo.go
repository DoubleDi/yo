// Code generated by yo. DO NOT EDIT.

// Package models contains the types.
package models

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
	"google.golang.org/grpc/codes"
)

// GeneratedColumn represents a row from 'GeneratedColumns'.
type GeneratedColumn struct {
	ID        int64  `spanner:"ID" json:"ID"`               // ID
	FirstName string `spanner:"FirstName" json:"FirstName"` // FirstName
	LastName  string `spanner:"LastName" json:"LastName"`   // LastName
	FullName  string `spanner:"FullName" json:"FullName"`   // FullName
}

func GeneratedColumnPrimaryKeys() []string {
	return []string{
		"ID",
	}
}

func GeneratedColumnColumns() []string {
	return []string{
		"ID",
		"FirstName",
		"LastName",
		"FullName",
	}
}

func GeneratedColumnWritableColumns() []string {
	return []string{
		"ID",
		"FirstName",
		"LastName",
	}
}

func (gc *GeneratedColumn) columnsToPtrs(cols []string) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		switch col {
		case "ID":
			ret = append(ret, yoDecode(&gc.ID))
		case "FirstName":
			ret = append(ret, yoDecode(&gc.FirstName))
		case "LastName":
			ret = append(ret, yoDecode(&gc.LastName))
		case "FullName":
			ret = append(ret, yoDecode(&gc.FullName))
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}
	return ret, nil
}

func (gc *GeneratedColumn) columnsToValues(cols []string) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		switch col {
		case "ID":
			ret = append(ret, yoEncode(gc.ID))
		case "FirstName":
			ret = append(ret, yoEncode(gc.FirstName))
		case "LastName":
			ret = append(ret, yoEncode(gc.LastName))
		case "FullName":
			ret = append(ret, yoEncode(gc.FullName))
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}

	return ret, nil
}

// newGeneratedColumn_Decoder returns a decoder which reads a row from *spanner.Row
// into GeneratedColumn. The decoder is not goroutine-safe. Don't use it concurrently.
func newGeneratedColumn_Decoder(cols []string) func(*spanner.Row) (*GeneratedColumn, error) {
	return func(row *spanner.Row) (*GeneratedColumn, error) {
		var gc GeneratedColumn
		ptrs, err := gc.columnsToPtrs(cols)
		if err != nil {
			return nil, err
		}

		if err := row.Columns(ptrs...); err != nil {
			return nil, err
		}

		return &gc, nil
	}
}

// Insert returns a Mutation to insert a row into a table. If the row already
// exists, the write or transaction fails.
func (gc *GeneratedColumn) Insert(ctx context.Context) *spanner.Mutation {
	values, _ := gc.columnsToValues(GeneratedColumnWritableColumns())
	return spanner.Insert("GeneratedColumns", GeneratedColumnWritableColumns(), values)
}

// Update returns a Mutation to update a row in a table. If the row does not
// already exist, the write or transaction fails.
func (gc *GeneratedColumn) Update(ctx context.Context) *spanner.Mutation {
	values, _ := gc.columnsToValues(GeneratedColumnWritableColumns())
	return spanner.Update("GeneratedColumns", GeneratedColumnWritableColumns(), values)
}

// InsertOrUpdate returns a Mutation to insert a row into a table. If the row
// already exists, it updates it instead. Any column values not explicitly
// written are preserved.
func (gc *GeneratedColumn) InsertOrUpdate(ctx context.Context) *spanner.Mutation {
	values, _ := gc.columnsToValues(GeneratedColumnWritableColumns())
	return spanner.InsertOrUpdate("GeneratedColumns", GeneratedColumnWritableColumns(), values)
}

// Replace returns a Mutation to insert a row into a table, deleting any
// existing row. Unlike InsertOrUpdate, this means any values not explicitly
// written become NULL.
func (gc *GeneratedColumn) Replace(ctx context.Context) *spanner.Mutation {
	values, _ := gc.columnsToValues(GeneratedColumnWritableColumns())
	return spanner.Replace("GeneratedColumns", GeneratedColumnWritableColumns(), values)
}

// UpdateColumns returns a Mutation to update specified columns of a row in a table.
func (gc *GeneratedColumn) UpdateColumns(ctx context.Context, cols ...string) (*spanner.Mutation, error) {
	// add primary keys to columns to update by primary keys
	colsWithPKeys := append(cols, GeneratedColumnPrimaryKeys()...)

	values, err := gc.columnsToValues(colsWithPKeys)
	if err != nil {
		return nil, newErrorWithCode(codes.InvalidArgument, "GeneratedColumn.UpdateColumns", "GeneratedColumns", err)
	}

	return spanner.Update("GeneratedColumns", colsWithPKeys, values), nil
}

// FindGeneratedColumn gets a GeneratedColumn by primary key
func FindGeneratedColumn(ctx context.Context, db YODB, id int64) (*GeneratedColumn, error) {
	key_ := spanner.Key{yoEncode(id)}
	row, err := db.ReadRow(ctx, "GeneratedColumns", key_, GeneratedColumnColumns())
	if err != nil {
		return nil, newError("FindGeneratedColumn", "GeneratedColumns", err)
	}

	decoder := newGeneratedColumn_Decoder(GeneratedColumnColumns())
	gc, err := decoder(row)
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "FindGeneratedColumn", "GeneratedColumns", err)
	}

	return gc, nil
}

// ReadGeneratedColumn retrieves multiples rows from GeneratedColumn by KeySet as a slice.
func ReadGeneratedColumn(ctx context.Context, db YODB, keys spanner.KeySet) ([]*GeneratedColumn, error) {
	var res []*GeneratedColumn

	decoder := newGeneratedColumn_Decoder(GeneratedColumnColumns())

	rows := db.Read(ctx, "GeneratedColumns", keys, GeneratedColumnColumns())
	err := rows.Do(func(row *spanner.Row) error {
		gc, err := decoder(row)
		if err != nil {
			return err
		}
		res = append(res, gc)

		return nil
	})
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "ReadGeneratedColumn", "GeneratedColumns", err)
	}

	return res, nil
}

// Delete deletes the GeneratedColumn from the database.
func (gc *GeneratedColumn) Delete(ctx context.Context) *spanner.Mutation {
	values, _ := gc.columnsToValues(GeneratedColumnPrimaryKeys())
	return spanner.Delete("GeneratedColumns", spanner.Key(values))
}
