// Code generated by SQLBoiler 4.8.6 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dao

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("CodeFiles", testCodeFiles)
	t.Run("Projects", testProjects)
	t.Run("ProjectsCodeFilesHistories", testProjectsCodeFilesHistories)
	t.Run("ProjectsHistories", testProjectsHistories)
	t.Run("ProjectsTags", testProjectsTags)
	t.Run("Tags", testTags)
}

func TestDelete(t *testing.T) {
	t.Run("CodeFiles", testCodeFilesDelete)
	t.Run("Projects", testProjectsDelete)
	t.Run("ProjectsCodeFilesHistories", testProjectsCodeFilesHistoriesDelete)
	t.Run("ProjectsHistories", testProjectsHistoriesDelete)
	t.Run("ProjectsTags", testProjectsTagsDelete)
	t.Run("Tags", testTagsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("CodeFiles", testCodeFilesQueryDeleteAll)
	t.Run("Projects", testProjectsQueryDeleteAll)
	t.Run("ProjectsCodeFilesHistories", testProjectsCodeFilesHistoriesQueryDeleteAll)
	t.Run("ProjectsHistories", testProjectsHistoriesQueryDeleteAll)
	t.Run("ProjectsTags", testProjectsTagsQueryDeleteAll)
	t.Run("Tags", testTagsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("CodeFiles", testCodeFilesSliceDeleteAll)
	t.Run("Projects", testProjectsSliceDeleteAll)
	t.Run("ProjectsCodeFilesHistories", testProjectsCodeFilesHistoriesSliceDeleteAll)
	t.Run("ProjectsHistories", testProjectsHistoriesSliceDeleteAll)
	t.Run("ProjectsTags", testProjectsTagsSliceDeleteAll)
	t.Run("Tags", testTagsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("CodeFiles", testCodeFilesExists)
	t.Run("Projects", testProjectsExists)
	t.Run("ProjectsCodeFilesHistories", testProjectsCodeFilesHistoriesExists)
	t.Run("ProjectsHistories", testProjectsHistoriesExists)
	t.Run("ProjectsTags", testProjectsTagsExists)
	t.Run("Tags", testTagsExists)
}

func TestFind(t *testing.T) {
	t.Run("CodeFiles", testCodeFilesFind)
	t.Run("Projects", testProjectsFind)
	t.Run("ProjectsCodeFilesHistories", testProjectsCodeFilesHistoriesFind)
	t.Run("ProjectsHistories", testProjectsHistoriesFind)
	t.Run("ProjectsTags", testProjectsTagsFind)
	t.Run("Tags", testTagsFind)
}

func TestBind(t *testing.T) {
	t.Run("CodeFiles", testCodeFilesBind)
	t.Run("Projects", testProjectsBind)
	t.Run("ProjectsCodeFilesHistories", testProjectsCodeFilesHistoriesBind)
	t.Run("ProjectsHistories", testProjectsHistoriesBind)
	t.Run("ProjectsTags", testProjectsTagsBind)
	t.Run("Tags", testTagsBind)
}

func TestOne(t *testing.T) {
	t.Run("CodeFiles", testCodeFilesOne)
	t.Run("Projects", testProjectsOne)
	t.Run("ProjectsCodeFilesHistories", testProjectsCodeFilesHistoriesOne)
	t.Run("ProjectsHistories", testProjectsHistoriesOne)
	t.Run("ProjectsTags", testProjectsTagsOne)
	t.Run("Tags", testTagsOne)
}

func TestAll(t *testing.T) {
	t.Run("CodeFiles", testCodeFilesAll)
	t.Run("Projects", testProjectsAll)
	t.Run("ProjectsCodeFilesHistories", testProjectsCodeFilesHistoriesAll)
	t.Run("ProjectsHistories", testProjectsHistoriesAll)
	t.Run("ProjectsTags", testProjectsTagsAll)
	t.Run("Tags", testTagsAll)
}

func TestCount(t *testing.T) {
	t.Run("CodeFiles", testCodeFilesCount)
	t.Run("Projects", testProjectsCount)
	t.Run("ProjectsCodeFilesHistories", testProjectsCodeFilesHistoriesCount)
	t.Run("ProjectsHistories", testProjectsHistoriesCount)
	t.Run("ProjectsTags", testProjectsTagsCount)
	t.Run("Tags", testTagsCount)
}

func TestHooks(t *testing.T) {
	t.Run("CodeFiles", testCodeFilesHooks)
	t.Run("Projects", testProjectsHooks)
	t.Run("ProjectsCodeFilesHistories", testProjectsCodeFilesHistoriesHooks)
	t.Run("ProjectsHistories", testProjectsHistoriesHooks)
	t.Run("ProjectsTags", testProjectsTagsHooks)
	t.Run("Tags", testTagsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("CodeFiles", testCodeFilesInsert)
	t.Run("CodeFiles", testCodeFilesInsertWhitelist)
	t.Run("Projects", testProjectsInsert)
	t.Run("Projects", testProjectsInsertWhitelist)
	t.Run("ProjectsCodeFilesHistories", testProjectsCodeFilesHistoriesInsert)
	t.Run("ProjectsCodeFilesHistories", testProjectsCodeFilesHistoriesInsertWhitelist)
	t.Run("ProjectsHistories", testProjectsHistoriesInsert)
	t.Run("ProjectsHistories", testProjectsHistoriesInsertWhitelist)
	t.Run("ProjectsTags", testProjectsTagsInsert)
	t.Run("ProjectsTags", testProjectsTagsInsertWhitelist)
	t.Run("Tags", testTagsInsert)
	t.Run("Tags", testTagsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("CodeFileToProjectUsingProject", testCodeFileToOneProjectUsingProject)
	t.Run("ProjectsCodeFilesHistoryToProjectsHistoryUsingRevision", testProjectsCodeFilesHistoryToOneProjectsHistoryUsingRevision)
	t.Run("ProjectsHistoryToProjectUsingProject", testProjectsHistoryToOneProjectUsingProject)
	t.Run("ProjectsTagToProjectUsingProject", testProjectsTagToOneProjectUsingProject)
	t.Run("ProjectsTagToTagUsingTag", testProjectsTagToOneTagUsingTag)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("ProjectToCodeFiles", testProjectToManyCodeFiles)
	t.Run("ProjectToProjectsHistories", testProjectToManyProjectsHistories)
	t.Run("ProjectToProjectsTags", testProjectToManyProjectsTags)
	t.Run("ProjectsHistoryToRevisionProjectsCodeFilesHistories", testProjectsHistoryToManyRevisionProjectsCodeFilesHistories)
	t.Run("TagToProjectsTags", testTagToManyProjectsTags)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("CodeFileToProjectUsingCodeFiles", testCodeFileToOneSetOpProjectUsingProject)
	t.Run("ProjectsCodeFilesHistoryToProjectsHistoryUsingRevisionProjectsCodeFilesHistories", testProjectsCodeFilesHistoryToOneSetOpProjectsHistoryUsingRevision)
	t.Run("ProjectsHistoryToProjectUsingProjectsHistories", testProjectsHistoryToOneSetOpProjectUsingProject)
	t.Run("ProjectsTagToProjectUsingProjectsTags", testProjectsTagToOneSetOpProjectUsingProject)
	t.Run("ProjectsTagToTagUsingProjectsTags", testProjectsTagToOneSetOpTagUsingTag)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("ProjectToCodeFiles", testProjectToManyAddOpCodeFiles)
	t.Run("ProjectToProjectsHistories", testProjectToManyAddOpProjectsHistories)
	t.Run("ProjectToProjectsTags", testProjectToManyAddOpProjectsTags)
	t.Run("ProjectsHistoryToRevisionProjectsCodeFilesHistories", testProjectsHistoryToManyAddOpRevisionProjectsCodeFilesHistories)
	t.Run("TagToProjectsTags", testTagToManyAddOpProjectsTags)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("CodeFiles", testCodeFilesReload)
	t.Run("Projects", testProjectsReload)
	t.Run("ProjectsCodeFilesHistories", testProjectsCodeFilesHistoriesReload)
	t.Run("ProjectsHistories", testProjectsHistoriesReload)
	t.Run("ProjectsTags", testProjectsTagsReload)
	t.Run("Tags", testTagsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("CodeFiles", testCodeFilesReloadAll)
	t.Run("Projects", testProjectsReloadAll)
	t.Run("ProjectsCodeFilesHistories", testProjectsCodeFilesHistoriesReloadAll)
	t.Run("ProjectsHistories", testProjectsHistoriesReloadAll)
	t.Run("ProjectsTags", testProjectsTagsReloadAll)
	t.Run("Tags", testTagsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("CodeFiles", testCodeFilesSelect)
	t.Run("Projects", testProjectsSelect)
	t.Run("ProjectsCodeFilesHistories", testProjectsCodeFilesHistoriesSelect)
	t.Run("ProjectsHistories", testProjectsHistoriesSelect)
	t.Run("ProjectsTags", testProjectsTagsSelect)
	t.Run("Tags", testTagsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("CodeFiles", testCodeFilesUpdate)
	t.Run("Projects", testProjectsUpdate)
	t.Run("ProjectsCodeFilesHistories", testProjectsCodeFilesHistoriesUpdate)
	t.Run("ProjectsHistories", testProjectsHistoriesUpdate)
	t.Run("ProjectsTags", testProjectsTagsUpdate)
	t.Run("Tags", testTagsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("CodeFiles", testCodeFilesSliceUpdateAll)
	t.Run("Projects", testProjectsSliceUpdateAll)
	t.Run("ProjectsCodeFilesHistories", testProjectsCodeFilesHistoriesSliceUpdateAll)
	t.Run("ProjectsHistories", testProjectsHistoriesSliceUpdateAll)
	t.Run("ProjectsTags", testProjectsTagsSliceUpdateAll)
	t.Run("Tags", testTagsSliceUpdateAll)
}
