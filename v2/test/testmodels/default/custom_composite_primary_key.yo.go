// Code generated by yo. DO NOT EDIT.

// Package models contains the types.
package models

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
)

// CustomCompositePrimaryKey represents a row from 'CustomCompositePrimaryKeys'.
type CustomCompositePrimaryKey struct {
	ID    uint64 `spanner:"Id" json:"Id"`       // Id
	PKey1 string `spanner:"PKey1" json:"PKey1"` // PKey1
	PKey2 uint32 `spanner:"PKey2" json:"PKey2"` // PKey2
	Error int8   `spanner:"Error" json:"Error"` // Error
	X     string `spanner:"X" json:"X"`         // X
	Y     string `spanner:"Y" json:"Y"`         // Y
	Z     string `spanner:"Z" json:"Z"`         // Z
}

func CustomCompositePrimaryKeyPrimaryKeys() []string {
	return []string{
		"PKey1",
		"PKey2",
	}
}

func CustomCompositePrimaryKeyColumns() []string {
	return []string{
		"Id",
		"PKey1",
		"PKey2",
		"Error",
		"X",
		"Y",
		"Z",
	}
}

func CustomCompositePrimaryKeyWritableColumns() []string {
	return []string{
		"Id",
		"PKey1",
		"PKey2",
		"Error",
		"X",
		"Y",
		"Z",
	}
}

func (ccpk *CustomCompositePrimaryKey) columnsToPtrs(cols []string) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		switch col {
		case "Id":
			ret = append(ret, yoDecode(&ccpk.ID))
		case "PKey1":
			ret = append(ret, yoDecode(&ccpk.PKey1))
		case "PKey2":
			ret = append(ret, yoDecode(&ccpk.PKey2))
		case "Error":
			ret = append(ret, yoDecode(&ccpk.Error))
		case "X":
			ret = append(ret, yoDecode(&ccpk.X))
		case "Y":
			ret = append(ret, yoDecode(&ccpk.Y))
		case "Z":
			ret = append(ret, yoDecode(&ccpk.Z))
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}
	return ret, nil
}

func (ccpk *CustomCompositePrimaryKey) columnsToValues(cols []string) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		switch col {
		case "Id":
			ret = append(ret, yoEncode(ccpk.ID))
		case "PKey1":
			ret = append(ret, yoEncode(ccpk.PKey1))
		case "PKey2":
			ret = append(ret, yoEncode(ccpk.PKey2))
		case "Error":
			ret = append(ret, yoEncode(ccpk.Error))
		case "X":
			ret = append(ret, yoEncode(ccpk.X))
		case "Y":
			ret = append(ret, yoEncode(ccpk.Y))
		case "Z":
			ret = append(ret, yoEncode(ccpk.Z))
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}

	return ret, nil
}

// newCustomCompositePrimaryKey_Decoder returns a decoder which reads a row from *spanner.Row
// into CustomCompositePrimaryKey. The decoder is not goroutine-safe. Don't use it concurrently.
func newCustomCompositePrimaryKey_Decoder(cols []string) func(*spanner.Row) (*CustomCompositePrimaryKey, error) {
	return func(row *spanner.Row) (*CustomCompositePrimaryKey, error) {
		var ccpk CustomCompositePrimaryKey
		ptrs, err := ccpk.columnsToPtrs(cols)
		if err != nil {
			return nil, err
		}

		if err := row.Columns(ptrs...); err != nil {
			return nil, err
		}

		return &ccpk, nil
	}
}

// Insert returns a Mutation to insert a row into a table. If the row already
// exists, the write or transaction fails.
func (ccpk *CustomCompositePrimaryKey) Insert(ctx context.Context) *spanner.Mutation {
	values, _ := ccpk.columnsToValues(CustomCompositePrimaryKeyWritableColumns())
	return spanner.Insert("CustomCompositePrimaryKeys", CustomCompositePrimaryKeyWritableColumns(), values)
}

// Update returns a Mutation to update a row in a table. If the row does not
// already exist, the write or transaction fails.
func (ccpk *CustomCompositePrimaryKey) Update(ctx context.Context) *spanner.Mutation {
	values, _ := ccpk.columnsToValues(CustomCompositePrimaryKeyWritableColumns())
	return spanner.Update("CustomCompositePrimaryKeys", CustomCompositePrimaryKeyWritableColumns(), values)
}

// InsertOrUpdate returns a Mutation to insert a row into a table. If the row
// already exists, it updates it instead. Any column values not explicitly
// written are preserved.
func (ccpk *CustomCompositePrimaryKey) InsertOrUpdate(ctx context.Context) *spanner.Mutation {
	values, _ := ccpk.columnsToValues(CustomCompositePrimaryKeyWritableColumns())
	return spanner.InsertOrUpdate("CustomCompositePrimaryKeys", CustomCompositePrimaryKeyWritableColumns(), values)
}

// Replace returns a Mutation to insert a row into a table, deleting any
// existing row. Unlike InsertOrUpdate, this means any values not explicitly
// written become NULL.
func (ccpk *CustomCompositePrimaryKey) Replace(ctx context.Context) *spanner.Mutation {
	values, _ := ccpk.columnsToValues(CustomCompositePrimaryKeyWritableColumns())
	return spanner.Replace("CustomCompositePrimaryKeys", CustomCompositePrimaryKeyWritableColumns(), values)
}

// UpdateColumns returns a Mutation to update specified columns of a row in a table.
func (ccpk *CustomCompositePrimaryKey) UpdateColumns(ctx context.Context, cols ...string) (*spanner.Mutation, error) {
	// add primary keys to columns to update by primary keys
	colsWithPKeys := append(cols, CustomCompositePrimaryKeyPrimaryKeys()...)

	values, err := ccpk.columnsToValues(colsWithPKeys)
	if err != nil {
		return nil, newErrorWithCode(codes.InvalidArgument, "CustomCompositePrimaryKey.UpdateColumns", "CustomCompositePrimaryKeys", err)
	}

	return spanner.Update("CustomCompositePrimaryKeys", colsWithPKeys, values), nil
}

// FindCustomCompositePrimaryKey gets a CustomCompositePrimaryKey by primary key
func FindCustomCompositePrimaryKey(ctx context.Context, db YODB, pKey1 string, pKey2 uint32) (*CustomCompositePrimaryKey, error) {
	key := spanner.Key{yoEncode(pKey1), yoEncode(pKey2)}
	row, err := db.ReadRow(ctx, "CustomCompositePrimaryKeys", key, CustomCompositePrimaryKeyColumns())
	if err != nil {
		return nil, newError("FindCustomCompositePrimaryKey", "CustomCompositePrimaryKeys", err)
	}

	decoder := newCustomCompositePrimaryKey_Decoder(CustomCompositePrimaryKeyColumns())
	ccpk, err := decoder(row)
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "FindCustomCompositePrimaryKey", "CustomCompositePrimaryKeys", err)
	}

	return ccpk, nil
}

// ReadCustomCompositePrimaryKey retrieves multiples rows from CustomCompositePrimaryKey by KeySet as a slice.
func ReadCustomCompositePrimaryKey(ctx context.Context, db YODB, keys spanner.KeySet) ([]*CustomCompositePrimaryKey, error) {
	var res []*CustomCompositePrimaryKey

	decoder := newCustomCompositePrimaryKey_Decoder(CustomCompositePrimaryKeyColumns())

	rows := db.Read(ctx, "CustomCompositePrimaryKeys", keys, CustomCompositePrimaryKeyColumns())
	err := rows.Do(func(row *spanner.Row) error {
		ccpk, err := decoder(row)
		if err != nil {
			return err
		}
		res = append(res, ccpk)

		return nil
	})
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "ReadCustomCompositePrimaryKey", "CustomCompositePrimaryKeys", err)
	}

	return res, nil
}

// Delete deletes the CustomCompositePrimaryKey from the database.
func (ccpk *CustomCompositePrimaryKey) Delete(ctx context.Context) *spanner.Mutation {
	values, _ := ccpk.columnsToValues(CustomCompositePrimaryKeyPrimaryKeys())
	return spanner.Delete("CustomCompositePrimaryKeys", spanner.Key(values))
}

// FindCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByError retrieves multiple rows from 'CustomCompositePrimaryKeys' as a slice of CustomCompositePrimaryKey.
//
// Generated from index 'CustomCompositePrimaryKeysByError'.
func FindCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByError(ctx context.Context, db YODB, e int8) ([]*CustomCompositePrimaryKey, error) {
	const sqlstr = "SELECT " +
		"Id, PKey1, PKey2, Error, X, Y, Z " +
		"FROM CustomCompositePrimaryKeys@{FORCE_INDEX=CustomCompositePrimaryKeysByError} " +
		"WHERE Error = @param0"

	stmt := spanner.NewStatement(sqlstr)
	stmt.Params["param0"] = yoEncode(e)

	decoder := newCustomCompositePrimaryKey_Decoder(CustomCompositePrimaryKeyColumns())

	// run query
	YOLog(ctx, sqlstr, e)
	iter := db.Query(ctx, stmt)
	defer iter.Stop()

	// load results
	res := []*CustomCompositePrimaryKey{}
	for {
		row, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, newError("FindCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByError", "CustomCompositePrimaryKeys", err)
		}

		ccpk, err := decoder(row)
		if err != nil {
			return nil, newErrorWithCode(codes.Internal, "FindCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByError", "CustomCompositePrimaryKeys", err)
		}

		res = append(res, ccpk)
	}

	return res, nil
}

// ReadCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByError retrieves multiples rows from 'CustomCompositePrimaryKeys' by KeySet as a slice.
//
// This does not retrieve all columns of 'CustomCompositePrimaryKeys' because an index has only columns
// used for primary key, index key and storing columns. If you need more columns, add storing
// columns or Read by primary key or Query with join.
//
// Generated from unique index 'CustomCompositePrimaryKeysByError'.
func ReadCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByError(ctx context.Context, db YODB, keys spanner.KeySet) ([]*CustomCompositePrimaryKey, error) {
	var res []*CustomCompositePrimaryKey
	columns := []string{
		"PKey1",
		"PKey2",
		"Error",
	}

	decoder := newCustomCompositePrimaryKey_Decoder(columns)

	rows := db.ReadUsingIndex(ctx, "CustomCompositePrimaryKeys", "CustomCompositePrimaryKeysByError", keys, columns)
	err := rows.Do(func(row *spanner.Row) error {
		ccpk, err := decoder(row)
		if err != nil {
			return err
		}
		res = append(res, ccpk)

		return nil
	})
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "ReadCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByError", "CustomCompositePrimaryKeys", err)
	}

	return res, nil
}

// FindCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByError2 retrieves multiple rows from 'CustomCompositePrimaryKeys' as a slice of CustomCompositePrimaryKey.
//
// Generated from index 'CustomCompositePrimaryKeysByError2'.
func FindCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByError2(ctx context.Context, db YODB, e int8) ([]*CustomCompositePrimaryKey, error) {
	const sqlstr = "SELECT " +
		"Id, PKey1, PKey2, Error, X, Y, Z " +
		"FROM CustomCompositePrimaryKeys@{FORCE_INDEX=CustomCompositePrimaryKeysByError2} " +
		"WHERE Error = @param0"

	stmt := spanner.NewStatement(sqlstr)
	stmt.Params["param0"] = yoEncode(e)

	decoder := newCustomCompositePrimaryKey_Decoder(CustomCompositePrimaryKeyColumns())

	// run query
	YOLog(ctx, sqlstr, e)
	iter := db.Query(ctx, stmt)
	defer iter.Stop()

	// load results
	res := []*CustomCompositePrimaryKey{}
	for {
		row, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, newError("FindCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByError2", "CustomCompositePrimaryKeys", err)
		}

		ccpk, err := decoder(row)
		if err != nil {
			return nil, newErrorWithCode(codes.Internal, "FindCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByError2", "CustomCompositePrimaryKeys", err)
		}

		res = append(res, ccpk)
	}

	return res, nil
}

// ReadCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByError2 retrieves multiples rows from 'CustomCompositePrimaryKeys' by KeySet as a slice.
//
// This does not retrieve all columns of 'CustomCompositePrimaryKeys' because an index has only columns
// used for primary key, index key and storing columns. If you need more columns, add storing
// columns or Read by primary key or Query with join.
//
// Generated from unique index 'CustomCompositePrimaryKeysByError2'.
func ReadCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByError2(ctx context.Context, db YODB, keys spanner.KeySet) ([]*CustomCompositePrimaryKey, error) {
	var res []*CustomCompositePrimaryKey
	columns := []string{
		"PKey1",
		"PKey2",
		"Error",
		"Z",
	}

	decoder := newCustomCompositePrimaryKey_Decoder(columns)

	rows := db.ReadUsingIndex(ctx, "CustomCompositePrimaryKeys", "CustomCompositePrimaryKeysByError2", keys, columns)
	err := rows.Do(func(row *spanner.Row) error {
		ccpk, err := decoder(row)
		if err != nil {
			return err
		}
		res = append(res, ccpk)

		return nil
	})
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "ReadCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByError2", "CustomCompositePrimaryKeys", err)
	}

	return res, nil
}

// FindCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByError3 retrieves multiple rows from 'CustomCompositePrimaryKeys' as a slice of CustomCompositePrimaryKey.
//
// Generated from index 'CustomCompositePrimaryKeysByError3'.
func FindCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByError3(ctx context.Context, db YODB, e int8) ([]*CustomCompositePrimaryKey, error) {
	const sqlstr = "SELECT " +
		"Id, PKey1, PKey2, Error, X, Y, Z " +
		"FROM CustomCompositePrimaryKeys@{FORCE_INDEX=CustomCompositePrimaryKeysByError3} " +
		"WHERE Error = @param0"

	stmt := spanner.NewStatement(sqlstr)
	stmt.Params["param0"] = yoEncode(e)

	decoder := newCustomCompositePrimaryKey_Decoder(CustomCompositePrimaryKeyColumns())

	// run query
	YOLog(ctx, sqlstr, e)
	iter := db.Query(ctx, stmt)
	defer iter.Stop()

	// load results
	res := []*CustomCompositePrimaryKey{}
	for {
		row, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, newError("FindCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByError3", "CustomCompositePrimaryKeys", err)
		}

		ccpk, err := decoder(row)
		if err != nil {
			return nil, newErrorWithCode(codes.Internal, "FindCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByError3", "CustomCompositePrimaryKeys", err)
		}

		res = append(res, ccpk)
	}

	return res, nil
}

// ReadCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByError3 retrieves multiples rows from 'CustomCompositePrimaryKeys' by KeySet as a slice.
//
// This does not retrieve all columns of 'CustomCompositePrimaryKeys' because an index has only columns
// used for primary key, index key and storing columns. If you need more columns, add storing
// columns or Read by primary key or Query with join.
//
// Generated from unique index 'CustomCompositePrimaryKeysByError3'.
func ReadCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByError3(ctx context.Context, db YODB, keys spanner.KeySet) ([]*CustomCompositePrimaryKey, error) {
	var res []*CustomCompositePrimaryKey
	columns := []string{
		"PKey1",
		"PKey2",
		"Error",
		"Z",
		"Y",
	}

	decoder := newCustomCompositePrimaryKey_Decoder(columns)

	rows := db.ReadUsingIndex(ctx, "CustomCompositePrimaryKeys", "CustomCompositePrimaryKeysByError3", keys, columns)
	err := rows.Do(func(row *spanner.Row) error {
		ccpk, err := decoder(row)
		if err != nil {
			return err
		}
		res = append(res, ccpk)

		return nil
	})
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "ReadCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByError3", "CustomCompositePrimaryKeys", err)
	}

	return res, nil
}

// FindCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByXY retrieves multiple rows from 'CustomCompositePrimaryKeys' as a slice of CustomCompositePrimaryKey.
//
// Generated from index 'CustomCompositePrimaryKeysByXY'.
func FindCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByXY(ctx context.Context, db YODB, x string, y string) ([]*CustomCompositePrimaryKey, error) {
	const sqlstr = "SELECT " +
		"Id, PKey1, PKey2, Error, X, Y, Z " +
		"FROM CustomCompositePrimaryKeys@{FORCE_INDEX=CustomCompositePrimaryKeysByXY} " +
		"WHERE X = @param0 AND Y = @param1"

	stmt := spanner.NewStatement(sqlstr)
	stmt.Params["param0"] = yoEncode(x)
	stmt.Params["param1"] = yoEncode(y)

	decoder := newCustomCompositePrimaryKey_Decoder(CustomCompositePrimaryKeyColumns())

	// run query
	YOLog(ctx, sqlstr, x, y)
	iter := db.Query(ctx, stmt)
	defer iter.Stop()

	// load results
	res := []*CustomCompositePrimaryKey{}
	for {
		row, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, newError("FindCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByXY", "CustomCompositePrimaryKeys", err)
		}

		ccpk, err := decoder(row)
		if err != nil {
			return nil, newErrorWithCode(codes.Internal, "FindCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByXY", "CustomCompositePrimaryKeys", err)
		}

		res = append(res, ccpk)
	}

	return res, nil
}

// ReadCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByXY retrieves multiples rows from 'CustomCompositePrimaryKeys' by KeySet as a slice.
//
// This does not retrieve all columns of 'CustomCompositePrimaryKeys' because an index has only columns
// used for primary key, index key and storing columns. If you need more columns, add storing
// columns or Read by primary key or Query with join.
//
// Generated from unique index 'CustomCompositePrimaryKeysByXY'.
func ReadCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByXY(ctx context.Context, db YODB, keys spanner.KeySet) ([]*CustomCompositePrimaryKey, error) {
	var res []*CustomCompositePrimaryKey
	columns := []string{
		"PKey1",
		"PKey2",
		"X",
		"Y",
	}

	decoder := newCustomCompositePrimaryKey_Decoder(columns)

	rows := db.ReadUsingIndex(ctx, "CustomCompositePrimaryKeys", "CustomCompositePrimaryKeysByXY", keys, columns)
	err := rows.Do(func(row *spanner.Row) error {
		ccpk, err := decoder(row)
		if err != nil {
			return err
		}
		res = append(res, ccpk)

		return nil
	})
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "ReadCustomCompositePrimaryKeysByCustomCompositePrimaryKeysByXY", "CustomCompositePrimaryKeys", err)
	}

	return res, nil
}
