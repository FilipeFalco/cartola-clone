package main

import (
	"context"
	"database/sql"

	"github.com/filipefalco/cartola-consolidacao/internal/infra/db"
	"github.com/filipefalco/cartola-consolidacao/internal/infra/repository"
	"github.com/filipefalco/cartola-consolidacao/pkg/uow"
)

func main() {
	ctx := context.Background()
	dtb, err := sql.Open("mysql", "root:roo@tcp(localhost:3306/cartola?parseTime=true")

	if err != nil {
		panic(err)
	}

	defer dtb.Close()
	uow, err := uow.NewUow(ctx, dtb)

	if err != nil {
		panic(err)
	}

	registerRepositories(uow)
}

func registerRepositories(uow *uow.Uow) {
	uow.Register("PlayerRespoitory", func(tx *sql.Tx) interface{} {
		repo := repository.NewPlayerRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("MyTeamRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewMyTeamRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("MatchRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewMatchRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("TeamRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewTeamRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})
}
