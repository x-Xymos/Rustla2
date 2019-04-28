// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testUserConfigs(t *testing.T) {
	t.Parallel()

	query := UserConfigs()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testUserConfigsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserConfig{}
	if err = randomize.Struct(seed, o, userConfigDBTypes, true, userConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserConfigsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserConfig{}
	if err = randomize.Struct(seed, o, userConfigDBTypes, true, userConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := UserConfigs().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserConfigsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserConfig{}
	if err = randomize.Struct(seed, o, userConfigDBTypes, true, userConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserConfigSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := UserConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserConfigsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserConfig{}
	if err = randomize.Struct(seed, o, userConfigDBTypes, true, userConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := UserConfigExists(ctx, tx, o.UserID, o.Name)
	if err != nil {
		t.Errorf("Unable to check if UserConfig exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UserConfigExists to return true, but got false.")
	}
}

func testUserConfigsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserConfig{}
	if err = randomize.Struct(seed, o, userConfigDBTypes, true, userConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	userConfigFound, err := FindUserConfig(ctx, tx, o.UserID, o.Name)
	if err != nil {
		t.Error(err)
	}

	if userConfigFound == nil {
		t.Error("want a record, got nil")
	}
}

func testUserConfigsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserConfig{}
	if err = randomize.Struct(seed, o, userConfigDBTypes, true, userConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = UserConfigs().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testUserConfigsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserConfig{}
	if err = randomize.Struct(seed, o, userConfigDBTypes, true, userConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := UserConfigs().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUserConfigsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userConfigOne := &UserConfig{}
	userConfigTwo := &UserConfig{}
	if err = randomize.Struct(seed, userConfigOne, userConfigDBTypes, false, userConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserConfig struct: %s", err)
	}
	if err = randomize.Struct(seed, userConfigTwo, userConfigDBTypes, false, userConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userConfigOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userConfigTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserConfigs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUserConfigsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	userConfigOne := &UserConfig{}
	userConfigTwo := &UserConfig{}
	if err = randomize.Struct(seed, userConfigOne, userConfigDBTypes, false, userConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserConfig struct: %s", err)
	}
	if err = randomize.Struct(seed, userConfigTwo, userConfigDBTypes, false, userConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userConfigOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userConfigTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func userConfigBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserConfig) error {
	*o = UserConfig{}
	return nil
}

func userConfigAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *UserConfig) error {
	*o = UserConfig{}
	return nil
}

func userConfigAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *UserConfig) error {
	*o = UserConfig{}
	return nil
}

func userConfigBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserConfig) error {
	*o = UserConfig{}
	return nil
}

func userConfigAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *UserConfig) error {
	*o = UserConfig{}
	return nil
}

func userConfigBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserConfig) error {
	*o = UserConfig{}
	return nil
}

func userConfigAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *UserConfig) error {
	*o = UserConfig{}
	return nil
}

func userConfigBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserConfig) error {
	*o = UserConfig{}
	return nil
}

func userConfigAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *UserConfig) error {
	*o = UserConfig{}
	return nil
}

func testUserConfigsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &UserConfig{}
	o := &UserConfig{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, userConfigDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UserConfig object: %s", err)
	}

	AddUserConfigHook(boil.BeforeInsertHook, userConfigBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	userConfigBeforeInsertHooks = []UserConfigHook{}

	AddUserConfigHook(boil.AfterInsertHook, userConfigAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	userConfigAfterInsertHooks = []UserConfigHook{}

	AddUserConfigHook(boil.AfterSelectHook, userConfigAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	userConfigAfterSelectHooks = []UserConfigHook{}

	AddUserConfigHook(boil.BeforeUpdateHook, userConfigBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	userConfigBeforeUpdateHooks = []UserConfigHook{}

	AddUserConfigHook(boil.AfterUpdateHook, userConfigAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	userConfigAfterUpdateHooks = []UserConfigHook{}

	AddUserConfigHook(boil.BeforeDeleteHook, userConfigBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	userConfigBeforeDeleteHooks = []UserConfigHook{}

	AddUserConfigHook(boil.AfterDeleteHook, userConfigAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	userConfigAfterDeleteHooks = []UserConfigHook{}

	AddUserConfigHook(boil.BeforeUpsertHook, userConfigBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	userConfigBeforeUpsertHooks = []UserConfigHook{}

	AddUserConfigHook(boil.AfterUpsertHook, userConfigAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	userConfigAfterUpsertHooks = []UserConfigHook{}
}

func testUserConfigsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserConfig{}
	if err = randomize.Struct(seed, o, userConfigDBTypes, true, userConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserConfigsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserConfig{}
	if err = randomize.Struct(seed, o, userConfigDBTypes, true); err != nil {
		t.Errorf("Unable to randomize UserConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(userConfigColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := UserConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserConfigToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local UserConfig
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, userConfigDBTypes, false, userConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserConfig struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.UserID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := UserConfigSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*UserConfig)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testUserConfigToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a UserConfig
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, userConfigDBTypes, false, strmangle.SetComplement(userConfigPrimaryKeyColumns, userConfigColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UserConfigs[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.UserID, x.ID) {
			t.Error("foreign key was wrong value", a.UserID)
		}

		if exists, err := UserConfigExists(ctx, tx, a.UserID, a.Name); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testUserConfigsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserConfig{}
	if err = randomize.Struct(seed, o, userConfigDBTypes, true, userConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUserConfigsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserConfig{}
	if err = randomize.Struct(seed, o, userConfigDBTypes, true, userConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserConfigSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUserConfigsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &UserConfig{}
	if err = randomize.Struct(seed, o, userConfigDBTypes, true, userConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := UserConfigs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	userConfigDBTypes = map[string]string{`UserID`: `bigint`, `Name`: `varchar`, `Value`: `varchar`}
	_                 = bytes.MinRead
)

func testUserConfigsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(userConfigPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(userConfigColumns) == len(userConfigPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserConfig{}
	if err = randomize.Struct(seed, o, userConfigDBTypes, true, userConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userConfigDBTypes, true, userConfigPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserConfig struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testUserConfigsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(userConfigColumns) == len(userConfigPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &UserConfig{}
	if err = randomize.Struct(seed, o, userConfigDBTypes, true, userConfigColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize UserConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := UserConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userConfigDBTypes, true, userConfigPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserConfig struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(userConfigColumns, userConfigPrimaryKeyColumns) {
		fields = userConfigColumns
	} else {
		fields = strmangle.SetComplement(
			userConfigColumns,
			userConfigPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := UserConfigSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testUserConfigsUpsert(t *testing.T) {
	t.Parallel()

	if len(userConfigColumns) == len(userConfigPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLUserConfigUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := UserConfig{}
	if err = randomize.Struct(seed, &o, userConfigDBTypes, false); err != nil {
		t.Errorf("Unable to randomize UserConfig struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserConfig: %s", err)
	}

	count, err := UserConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, userConfigDBTypes, false, userConfigPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize UserConfig struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert UserConfig: %s", err)
	}

	count, err = UserConfigs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
