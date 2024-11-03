package example

import (
	"time"

	"github.com/bozkurtemre/go-service-weaver-template/internal/example/store"

	"github.com/ServiceWeaver/weaver"
)

type ExampleOut struct {
	weaver.AutoMarshal
	ID        int       `json:"id,omitempty"`
	Message   string    `json:"message,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type AllExamplesOut []ExampleOut

func (e ExampleOut) FromStore(in store.Example) ExampleOut {
	e.ID = int(in.ID)
	e.Message = in.Message
	e.CreatedAt = in.CreatedAt.Time
	return e
}

func (e AllExamplesOut) FromStore(in []store.Example) AllExamplesOut {
	e = make(AllExamplesOut, 0, len(in))
	for _, v := range in {
		e = append(e, ExampleOut{}.FromStore(v))
	}
	return e
}

type ExampleIn struct {
	weaver.AutoMarshal
	Message string
}

func (e ExampleIn) ToStore() (params store.CreateExampleParams) {
	t := time.Now()
	params.ID = int64(t.UTC().UnixNano())
	params.Message = e.Message
	params.CreatedAt.Time = t
	return params
}
