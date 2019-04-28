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

func testServiceInfos(t *testing.T) {
	t.Parallel()

	query := ServiceInfos()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testServiceInfosDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInfo{}
	if err = randomize.Struct(seed, o, serviceInfoDBTypes, true, serviceInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInfo struct: %s", err)
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

	count, err := ServiceInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceInfosQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInfo{}
	if err = randomize.Struct(seed, o, serviceInfoDBTypes, true, serviceInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ServiceInfos().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceInfosSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInfo{}
	if err = randomize.Struct(seed, o, serviceInfoDBTypes, true, serviceInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceInfoSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceInfosExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInfo{}
	if err = randomize.Struct(seed, o, serviceInfoDBTypes, true, serviceInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ServiceInfoExists(ctx, tx, o.ChannelID)
	if err != nil {
		t.Errorf("Unable to check if ServiceInfo exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ServiceInfoExists to return true, but got false.")
	}
}

func testServiceInfosFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInfo{}
	if err = randomize.Struct(seed, o, serviceInfoDBTypes, true, serviceInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	serviceInfoFound, err := FindServiceInfo(ctx, tx, o.ChannelID)
	if err != nil {
		t.Error(err)
	}

	if serviceInfoFound == nil {
		t.Error("want a record, got nil")
	}
}

func testServiceInfosBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInfo{}
	if err = randomize.Struct(seed, o, serviceInfoDBTypes, true, serviceInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ServiceInfos().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testServiceInfosOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInfo{}
	if err = randomize.Struct(seed, o, serviceInfoDBTypes, true, serviceInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ServiceInfos().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testServiceInfosAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	serviceInfoOne := &ServiceInfo{}
	serviceInfoTwo := &ServiceInfo{}
	if err = randomize.Struct(seed, serviceInfoOne, serviceInfoDBTypes, false, serviceInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInfo struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceInfoTwo, serviceInfoDBTypes, false, serviceInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceInfoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceInfoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceInfos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testServiceInfosCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	serviceInfoOne := &ServiceInfo{}
	serviceInfoTwo := &ServiceInfo{}
	if err = randomize.Struct(seed, serviceInfoOne, serviceInfoDBTypes, false, serviceInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInfo struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceInfoTwo, serviceInfoDBTypes, false, serviceInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceInfoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceInfoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func serviceInfoBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInfo) error {
	*o = ServiceInfo{}
	return nil
}

func serviceInfoAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInfo) error {
	*o = ServiceInfo{}
	return nil
}

func serviceInfoAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInfo) error {
	*o = ServiceInfo{}
	return nil
}

func serviceInfoBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInfo) error {
	*o = ServiceInfo{}
	return nil
}

func serviceInfoAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInfo) error {
	*o = ServiceInfo{}
	return nil
}

func serviceInfoBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInfo) error {
	*o = ServiceInfo{}
	return nil
}

func serviceInfoAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInfo) error {
	*o = ServiceInfo{}
	return nil
}

func serviceInfoBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInfo) error {
	*o = ServiceInfo{}
	return nil
}

func serviceInfoAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInfo) error {
	*o = ServiceInfo{}
	return nil
}

func testServiceInfosHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ServiceInfo{}
	o := &ServiceInfo{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, serviceInfoDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ServiceInfo object: %s", err)
	}

	AddServiceInfoHook(boil.BeforeInsertHook, serviceInfoBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	serviceInfoBeforeInsertHooks = []ServiceInfoHook{}

	AddServiceInfoHook(boil.AfterInsertHook, serviceInfoAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	serviceInfoAfterInsertHooks = []ServiceInfoHook{}

	AddServiceInfoHook(boil.AfterSelectHook, serviceInfoAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	serviceInfoAfterSelectHooks = []ServiceInfoHook{}

	AddServiceInfoHook(boil.BeforeUpdateHook, serviceInfoBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	serviceInfoBeforeUpdateHooks = []ServiceInfoHook{}

	AddServiceInfoHook(boil.AfterUpdateHook, serviceInfoAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	serviceInfoAfterUpdateHooks = []ServiceInfoHook{}

	AddServiceInfoHook(boil.BeforeDeleteHook, serviceInfoBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	serviceInfoBeforeDeleteHooks = []ServiceInfoHook{}

	AddServiceInfoHook(boil.AfterDeleteHook, serviceInfoAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	serviceInfoAfterDeleteHooks = []ServiceInfoHook{}

	AddServiceInfoHook(boil.BeforeUpsertHook, serviceInfoBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	serviceInfoBeforeUpsertHooks = []ServiceInfoHook{}

	AddServiceInfoHook(boil.AfterUpsertHook, serviceInfoAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	serviceInfoAfterUpsertHooks = []ServiceInfoHook{}
}

func testServiceInfosInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInfo{}
	if err = randomize.Struct(seed, o, serviceInfoDBTypes, true, serviceInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceInfosInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInfo{}
	if err = randomize.Struct(seed, o, serviceInfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServiceInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(serviceInfoColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ServiceInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceInfoToOneChannelUsingChannel(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ServiceInfo
	var foreign Channel

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, serviceInfoDBTypes, false, serviceInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInfo struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, channelDBTypes, false, channelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Channel struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.ChannelID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Channel().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ServiceInfoSlice{&local}
	if err = local.L.LoadChannel(ctx, tx, false, (*[]*ServiceInfo)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Channel == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Channel = nil
	if err = local.L.LoadChannel(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Channel == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testServiceInfoToOneSetOpChannelUsingChannel(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServiceInfo
	var b, c Channel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serviceInfoDBTypes, false, strmangle.SetComplement(serviceInfoPrimaryKeyColumns, serviceInfoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, channelDBTypes, false, strmangle.SetComplement(channelPrimaryKeyColumns, channelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, channelDBTypes, false, strmangle.SetComplement(channelPrimaryKeyColumns, channelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Channel{&b, &c} {
		err = a.SetChannel(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Channel != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ServiceInfo != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.ChannelID, x.ID) {
			t.Error("foreign key was wrong value", a.ChannelID)
		}

		if exists, err := ServiceInfoExists(ctx, tx, a.ChannelID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testServiceInfosReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInfo{}
	if err = randomize.Struct(seed, o, serviceInfoDBTypes, true, serviceInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInfo struct: %s", err)
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

func testServiceInfosReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInfo{}
	if err = randomize.Struct(seed, o, serviceInfoDBTypes, true, serviceInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceInfoSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testServiceInfosSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInfo{}
	if err = randomize.Struct(seed, o, serviceInfoDBTypes, true, serviceInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceInfos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	serviceInfoDBTypes = map[string]string{`ChannelID`: `bigint`, `UpdatedAt`: `datetime`, `Title`: `varchar`, `ThumbnailURL`: `varchar`, `IsLive`: `tinyint`, `ViewerCount`: `int`}
	_                  = bytes.MinRead
)

func testServiceInfosUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(serviceInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(serviceInfoColumns) == len(serviceInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInfo{}
	if err = randomize.Struct(seed, o, serviceInfoDBTypes, true, serviceInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceInfoDBTypes, true, serviceInfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceInfo struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testServiceInfosSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(serviceInfoColumns) == len(serviceInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInfo{}
	if err = randomize.Struct(seed, o, serviceInfoDBTypes, true, serviceInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceInfoDBTypes, true, serviceInfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceInfo struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(serviceInfoColumns, serviceInfoPrimaryKeyColumns) {
		fields = serviceInfoColumns
	} else {
		fields = strmangle.SetComplement(
			serviceInfoColumns,
			serviceInfoPrimaryKeyColumns,
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

	slice := ServiceInfoSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testServiceInfosUpsert(t *testing.T) {
	t.Parallel()

	if len(serviceInfoColumns) == len(serviceInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLServiceInfoUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ServiceInfo{}
	if err = randomize.Struct(seed, &o, serviceInfoDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ServiceInfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceInfo: %s", err)
	}

	count, err := ServiceInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, serviceInfoDBTypes, false, serviceInfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceInfo struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceInfo: %s", err)
	}

	count, err = ServiceInfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
