package test_util

import (
	"context"
	"log"
	"os"
	"path"
	"testing"

	pkgDB "github.com/k3forx/bun/pkg/db"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dbfixture"
	"github.com/uptrace/bun/extra/bundebug"
)

func SetupTestDB(m *testing.M) *bun.DB {
	db, err := pkgDB.Setup()
	if err != nil {
		log.Fatal(err)
	}
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))
	return db
}

func LoadFixtures(t *testing.T, ctx context.Context, db *bun.DB, filePath string, entities ...any) {
	db.RegisterModel(entities...)

	dir, filename := path.Split(filePath)
	fixture := dbfixture.New(db, dbfixture.WithRecreateTables())
	if err := fixture.Load(ctx, os.DirFS(dir), filename); err != nil {
		t.Fatal(err)
	}
}
