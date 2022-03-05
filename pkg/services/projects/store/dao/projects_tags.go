// Code generated by SQLBoiler 4.8.6 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dao

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// ProjectsTag is an object representing the database table.
type ProjectsTag struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	ProjectID int       `boil:"project_id" json:"project_id" toml:"project_id" yaml:"project_id"`
	TagID     int       `boil:"tag_id" json:"tag_id" toml:"tag_id" yaml:"tag_id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *projectsTagR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L projectsTagL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ProjectsTagColumns = struct {
	ID        string
	ProjectID string
	TagID     string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	ProjectID: "project_id",
	TagID:     "tag_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var ProjectsTagTableColumns = struct {
	ID        string
	ProjectID string
	TagID     string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "projects_tags.id",
	ProjectID: "projects_tags.project_id",
	TagID:     "projects_tags.tag_id",
	CreatedAt: "projects_tags.created_at",
	UpdatedAt: "projects_tags.updated_at",
}

// Generated where

var ProjectsTagWhere = struct {
	ID        whereHelperint
	ProjectID whereHelperint
	TagID     whereHelperint
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint{field: "\"projects_tags\".\"id\""},
	ProjectID: whereHelperint{field: "\"projects_tags\".\"project_id\""},
	TagID:     whereHelperint{field: "\"projects_tags\".\"tag_id\""},
	CreatedAt: whereHelpertime_Time{field: "\"projects_tags\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"projects_tags\".\"updated_at\""},
}

// ProjectsTagRels is where relationship names are stored.
var ProjectsTagRels = struct {
	Project string
	Tag     string
}{
	Project: "Project",
	Tag:     "Tag",
}

// projectsTagR is where relationships are stored.
type projectsTagR struct {
	Project *Project `boil:"Project" json:"Project" toml:"Project" yaml:"Project"`
	Tag     *Tag     `boil:"Tag" json:"Tag" toml:"Tag" yaml:"Tag"`
}

// NewStruct creates a new relationship struct
func (*projectsTagR) NewStruct() *projectsTagR {
	return &projectsTagR{}
}

// projectsTagL is where Load methods for each relationship are stored.
type projectsTagL struct{}

var (
	projectsTagAllColumns            = []string{"id", "project_id", "tag_id", "created_at", "updated_at"}
	projectsTagColumnsWithoutDefault = []string{"project_id", "tag_id", "created_at", "updated_at"}
	projectsTagColumnsWithDefault    = []string{"id"}
	projectsTagPrimaryKeyColumns     = []string{"id"}
	projectsTagGeneratedColumns      = []string{}
)

type (
	// ProjectsTagSlice is an alias for a slice of pointers to ProjectsTag.
	// This should almost always be used instead of []ProjectsTag.
	ProjectsTagSlice []*ProjectsTag
	// ProjectsTagHook is the signature for custom ProjectsTag hook methods
	ProjectsTagHook func(context.Context, boil.ContextExecutor, *ProjectsTag) error

	projectsTagQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	projectsTagType                 = reflect.TypeOf(&ProjectsTag{})
	projectsTagMapping              = queries.MakeStructMapping(projectsTagType)
	projectsTagPrimaryKeyMapping, _ = queries.BindMapping(projectsTagType, projectsTagMapping, projectsTagPrimaryKeyColumns)
	projectsTagInsertCacheMut       sync.RWMutex
	projectsTagInsertCache          = make(map[string]insertCache)
	projectsTagUpdateCacheMut       sync.RWMutex
	projectsTagUpdateCache          = make(map[string]updateCache)
	projectsTagUpsertCacheMut       sync.RWMutex
	projectsTagUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var projectsTagAfterSelectHooks []ProjectsTagHook

var projectsTagBeforeInsertHooks []ProjectsTagHook
var projectsTagAfterInsertHooks []ProjectsTagHook

var projectsTagBeforeUpdateHooks []ProjectsTagHook
var projectsTagAfterUpdateHooks []ProjectsTagHook

var projectsTagBeforeDeleteHooks []ProjectsTagHook
var projectsTagAfterDeleteHooks []ProjectsTagHook

var projectsTagBeforeUpsertHooks []ProjectsTagHook
var projectsTagAfterUpsertHooks []ProjectsTagHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ProjectsTag) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectsTagAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ProjectsTag) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectsTagBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ProjectsTag) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectsTagAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ProjectsTag) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectsTagBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ProjectsTag) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectsTagAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ProjectsTag) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectsTagBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ProjectsTag) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectsTagAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ProjectsTag) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectsTagBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ProjectsTag) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectsTagAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddProjectsTagHook registers your hook function for all future operations.
func AddProjectsTagHook(hookPoint boil.HookPoint, projectsTagHook ProjectsTagHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		projectsTagAfterSelectHooks = append(projectsTagAfterSelectHooks, projectsTagHook)
	case boil.BeforeInsertHook:
		projectsTagBeforeInsertHooks = append(projectsTagBeforeInsertHooks, projectsTagHook)
	case boil.AfterInsertHook:
		projectsTagAfterInsertHooks = append(projectsTagAfterInsertHooks, projectsTagHook)
	case boil.BeforeUpdateHook:
		projectsTagBeforeUpdateHooks = append(projectsTagBeforeUpdateHooks, projectsTagHook)
	case boil.AfterUpdateHook:
		projectsTagAfterUpdateHooks = append(projectsTagAfterUpdateHooks, projectsTagHook)
	case boil.BeforeDeleteHook:
		projectsTagBeforeDeleteHooks = append(projectsTagBeforeDeleteHooks, projectsTagHook)
	case boil.AfterDeleteHook:
		projectsTagAfterDeleteHooks = append(projectsTagAfterDeleteHooks, projectsTagHook)
	case boil.BeforeUpsertHook:
		projectsTagBeforeUpsertHooks = append(projectsTagBeforeUpsertHooks, projectsTagHook)
	case boil.AfterUpsertHook:
		projectsTagAfterUpsertHooks = append(projectsTagAfterUpsertHooks, projectsTagHook)
	}
}

// One returns a single projectsTag record from the query.
func (q projectsTagQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ProjectsTag, error) {
	o := &ProjectsTag{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dao: failed to execute a one query for projects_tags")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ProjectsTag records from the query.
func (q projectsTagQuery) All(ctx context.Context, exec boil.ContextExecutor) (ProjectsTagSlice, error) {
	var o []*ProjectsTag

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dao: failed to assign all query results to ProjectsTag slice")
	}

	if len(projectsTagAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ProjectsTag records in the query.
func (q projectsTagQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dao: failed to count projects_tags rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q projectsTagQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dao: failed to check if projects_tags exists")
	}

	return count > 0, nil
}

// Project pointed to by the foreign key.
func (o *ProjectsTag) Project(mods ...qm.QueryMod) projectQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ProjectID),
	}

	queryMods = append(queryMods, mods...)

	query := Projects(queryMods...)
	queries.SetFrom(query.Query, "\"projects\"")

	return query
}

// Tag pointed to by the foreign key.
func (o *ProjectsTag) Tag(mods ...qm.QueryMod) tagQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.TagID),
	}

	queryMods = append(queryMods, mods...)

	query := Tags(queryMods...)
	queries.SetFrom(query.Query, "\"tags\"")

	return query
}

// LoadProject allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (projectsTagL) LoadProject(ctx context.Context, e boil.ContextExecutor, singular bool, maybeProjectsTag interface{}, mods queries.Applicator) error {
	var slice []*ProjectsTag
	var object *ProjectsTag

	if singular {
		object = maybeProjectsTag.(*ProjectsTag)
	} else {
		slice = *maybeProjectsTag.(*[]*ProjectsTag)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &projectsTagR{}
		}
		args = append(args, object.ProjectID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &projectsTagR{}
			}

			for _, a := range args {
				if a == obj.ProjectID {
					continue Outer
				}
			}

			args = append(args, obj.ProjectID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`projects`),
		qm.WhereIn(`projects.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Project")
	}

	var resultSlice []*Project
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Project")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for projects")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for projects")
	}

	if len(projectsTagAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Project = foreign
		if foreign.R == nil {
			foreign.R = &projectR{}
		}
		foreign.R.ProjectsTags = append(foreign.R.ProjectsTags, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ProjectID == foreign.ID {
				local.R.Project = foreign
				if foreign.R == nil {
					foreign.R = &projectR{}
				}
				foreign.R.ProjectsTags = append(foreign.R.ProjectsTags, local)
				break
			}
		}
	}

	return nil
}

// LoadTag allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (projectsTagL) LoadTag(ctx context.Context, e boil.ContextExecutor, singular bool, maybeProjectsTag interface{}, mods queries.Applicator) error {
	var slice []*ProjectsTag
	var object *ProjectsTag

	if singular {
		object = maybeProjectsTag.(*ProjectsTag)
	} else {
		slice = *maybeProjectsTag.(*[]*ProjectsTag)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &projectsTagR{}
		}
		args = append(args, object.TagID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &projectsTagR{}
			}

			for _, a := range args {
				if a == obj.TagID {
					continue Outer
				}
			}

			args = append(args, obj.TagID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`tags`),
		qm.WhereIn(`tags.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Tag")
	}

	var resultSlice []*Tag
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Tag")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for tags")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for tags")
	}

	if len(projectsTagAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Tag = foreign
		if foreign.R == nil {
			foreign.R = &tagR{}
		}
		foreign.R.ProjectsTags = append(foreign.R.ProjectsTags, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TagID == foreign.ID {
				local.R.Tag = foreign
				if foreign.R == nil {
					foreign.R = &tagR{}
				}
				foreign.R.ProjectsTags = append(foreign.R.ProjectsTags, local)
				break
			}
		}
	}

	return nil
}

// SetProject of the projectsTag to the related item.
// Sets o.R.Project to related.
// Adds o to related.R.ProjectsTags.
func (o *ProjectsTag) SetProject(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Project) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"projects_tags\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"project_id"}),
		strmangle.WhereClause("\"", "\"", 2, projectsTagPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ProjectID = related.ID
	if o.R == nil {
		o.R = &projectsTagR{
			Project: related,
		}
	} else {
		o.R.Project = related
	}

	if related.R == nil {
		related.R = &projectR{
			ProjectsTags: ProjectsTagSlice{o},
		}
	} else {
		related.R.ProjectsTags = append(related.R.ProjectsTags, o)
	}

	return nil
}

// SetTag of the projectsTag to the related item.
// Sets o.R.Tag to related.
// Adds o to related.R.ProjectsTags.
func (o *ProjectsTag) SetTag(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Tag) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"projects_tags\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"tag_id"}),
		strmangle.WhereClause("\"", "\"", 2, projectsTagPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TagID = related.ID
	if o.R == nil {
		o.R = &projectsTagR{
			Tag: related,
		}
	} else {
		o.R.Tag = related
	}

	if related.R == nil {
		related.R = &tagR{
			ProjectsTags: ProjectsTagSlice{o},
		}
	} else {
		related.R.ProjectsTags = append(related.R.ProjectsTags, o)
	}

	return nil
}

// ProjectsTags retrieves all the records using an executor.
func ProjectsTags(mods ...qm.QueryMod) projectsTagQuery {
	mods = append(mods, qm.From("\"projects_tags\""))
	return projectsTagQuery{NewQuery(mods...)}
}

// FindProjectsTag retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindProjectsTag(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ProjectsTag, error) {
	projectsTagObj := &ProjectsTag{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"projects_tags\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, projectsTagObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dao: unable to select from projects_tags")
	}

	if err = projectsTagObj.doAfterSelectHooks(ctx, exec); err != nil {
		return projectsTagObj, err
	}

	return projectsTagObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ProjectsTag) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dao: no projects_tags provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(projectsTagColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	projectsTagInsertCacheMut.RLock()
	cache, cached := projectsTagInsertCache[key]
	projectsTagInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			projectsTagAllColumns,
			projectsTagColumnsWithDefault,
			projectsTagColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(projectsTagType, projectsTagMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(projectsTagType, projectsTagMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"projects_tags\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"projects_tags\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "dao: unable to insert into projects_tags")
	}

	if !cached {
		projectsTagInsertCacheMut.Lock()
		projectsTagInsertCache[key] = cache
		projectsTagInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ProjectsTag.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ProjectsTag) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	projectsTagUpdateCacheMut.RLock()
	cache, cached := projectsTagUpdateCache[key]
	projectsTagUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			projectsTagAllColumns,
			projectsTagPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("dao: unable to update projects_tags, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"projects_tags\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, projectsTagPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(projectsTagType, projectsTagMapping, append(wl, projectsTagPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "dao: unable to update projects_tags row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dao: failed to get rows affected by update for projects_tags")
	}

	if !cached {
		projectsTagUpdateCacheMut.Lock()
		projectsTagUpdateCache[key] = cache
		projectsTagUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q projectsTagQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dao: unable to update all for projects_tags")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dao: unable to retrieve rows affected for projects_tags")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ProjectsTagSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("dao: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), projectsTagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"projects_tags\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, projectsTagPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dao: unable to update all in projectsTag slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dao: unable to retrieve rows affected all in update all projectsTag")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ProjectsTag) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("dao: no projects_tags provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(projectsTagColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	projectsTagUpsertCacheMut.RLock()
	cache, cached := projectsTagUpsertCache[key]
	projectsTagUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			projectsTagAllColumns,
			projectsTagColumnsWithDefault,
			projectsTagColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			projectsTagAllColumns,
			projectsTagPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("dao: unable to upsert projects_tags, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(projectsTagPrimaryKeyColumns))
			copy(conflict, projectsTagPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"projects_tags\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(projectsTagType, projectsTagMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(projectsTagType, projectsTagMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "dao: unable to upsert projects_tags")
	}

	if !cached {
		projectsTagUpsertCacheMut.Lock()
		projectsTagUpsertCache[key] = cache
		projectsTagUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ProjectsTag record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ProjectsTag) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("dao: no ProjectsTag provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), projectsTagPrimaryKeyMapping)
	sql := "DELETE FROM \"projects_tags\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dao: unable to delete from projects_tags")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dao: failed to get rows affected by delete for projects_tags")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q projectsTagQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dao: no projectsTagQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dao: unable to delete all from projects_tags")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dao: failed to get rows affected by deleteall for projects_tags")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ProjectsTagSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(projectsTagBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), projectsTagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"projects_tags\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, projectsTagPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dao: unable to delete all from projectsTag slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dao: failed to get rows affected by deleteall for projects_tags")
	}

	if len(projectsTagAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ProjectsTag) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindProjectsTag(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProjectsTagSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ProjectsTagSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), projectsTagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"projects_tags\".* FROM \"projects_tags\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, projectsTagPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dao: unable to reload all in ProjectsTagSlice")
	}

	*o = slice

	return nil
}

// ProjectsTagExists checks if the ProjectsTag row exists.
func ProjectsTagExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"projects_tags\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dao: unable to check if projects_tags exists")
	}

	return exists, nil
}