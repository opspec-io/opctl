package node

import (
	"context"

	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/node/core"
)

func newHandle(
	core core.Core,
	dataRef string,
	pullCreds *model.Creds,
) model.DataHandle {
	return handle{
		core:      core,
		dataRef:   dataRef,
		pullCreds: pullCreds,
	}
}

func (nh handle) GetContent(
	ctx context.Context,
	contentPath string,
) (
	model.ReadSeekCloser,
	error,
) {
	return nh.core.GetData(
		ctx,
		model.GetDataReq{
			ContentPath: contentPath,
			PkgRef:      nh.dataRef,
			PullCreds:   nh.pullCreds,
		},
	)
}

// handle allows interacting w/ data sourced from an opspec node
type handle struct {
	core      core.Core
	dataRef   string
	pullCreds *model.Creds
}

func (nh handle) ListDescendants(
	ctx context.Context,
) (
	[]*model.DirEntry,
	error,
) {
	return nh.core.ListDescendants(
		ctx,
		model.ListDescendantsReq{
			PkgRef:    nh.dataRef,
			PullCreds: nh.pullCreds,
		},
	)
}

func (hn handle) Path() *string {
	return nil
}

func (nh handle) Ref() string {
	return nh.dataRef
}