package opencensus

import (
	"context"
	"strconv"

	"go.opencensus.io/tag"
)

func makeLatencyCtx(ctx context.Context, entity, method string) (context.Context, error) {
	return tag.New(ctx, tag.Insert(keyEntity, entity), tag.Insert(keyMethod, method))
}

func makeHttpCodeCtx(ctx context.Context, url string, code int) (context.Context, error) {
	return tag.New(ctx,
		tag.Insert(keyURL, url),
		tag.Insert(keyHttpCode, strconv.Itoa(code)),
	)
}

func makeReqCtx(ctx context.Context, entity, method, url string) (context.Context, error) {
	return tag.New(ctx,
		tag.Insert(keyEntity, entity),
		tag.Insert(keyMethod, method),
		tag.Insert(keyURL, url),
	)
}
