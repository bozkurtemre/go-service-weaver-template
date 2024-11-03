package example

import (
	"context"

	"github.com/bozkurtemre/go-service-weaver-template/internal/example/store"

	"github.com/ServiceWeaver/weaver"
	"github.com/jackc/pgx/v5"
)

type Service interface {
	AllExamples(ctx context.Context) (out AllExamplesOut, err error)
	GetOneExampleById(ctx context.Context, id int64) (out ExampleOut, err error)
	CreateExample(ctx context.Context, in ExampleIn) (ok bool, err error)
}

type config struct {
	Source string
}

type impl struct {
	weaver.Implements[Service]
	weaver.WithConfig[config]
	db *store.Queries
}

func (i *impl) Init(ctx context.Context) error {
	conn, err := pgx.Connect(ctx, i.Config().Source)
	if err != nil {
		return err
	}

	i.db = store.New(conn)
	return nil
}

func (i *impl) AllExamples(ctx context.Context) (out AllExamplesOut, err error) {
	examples, err := i.db.ListExamples(ctx)
	if err != nil {
		return out, err
	}
	return out.FromStore(examples), nil
}

func (i *impl) GetOneExampleById(ctx context.Context, id int64) (out ExampleOut, err error) {
	example, err := i.db.GetExample(ctx, id)
	if err != nil {
		return out, err
	}
	return out.FromStore(example), nil
}

func (i *impl) CreateExample(ctx context.Context, in ExampleIn) (ok bool, err error) {
	_, err = i.db.CreateExample(ctx, in.ToStore())
	if err != nil {
		return false, err
	}
	return true, nil
}
