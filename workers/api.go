package workers

import (
	"context"
	"io"

	"github.com/nitrictech/go-sdk/api/errors"
	"github.com/nitrictech/go-sdk/api/errors/codes"
	"github.com/nitrictech/go-sdk/constants"
	"github.com/nitrictech/go-sdk/handler"
	v1 "github.com/nitrictech/nitric/core/pkg/proto/apis/v1"
	"google.golang.org/grpc"
)

type ApiWorker struct {
	client      v1.ApiClient
	apiName     string
	path        string
	httpHandler handler.HttpHandler
	methods     []string
}

type ApiWorkerOpts struct {
	apiName     string
	path        string
	httpHandler handler.HttpHandler
	methods     []string
}

var _ Worker = (*ApiWorker)(nil)

// Start implements Worker.
func (a *ApiWorker) Start(ctx context.Context) error {
	opts := &v1.ApiWorkerOptions{
		Security:         map[string]*v1.ApiWorkerScopes{},
		SecurityDisabled: true,
	}

	regReq := v1.RegistrationRequest{
		Api:     a.apiName,
		Path:    a.path,
		Methods: a.methods,
		Options: opts,
	}

	stream, err := a.client.Serve(ctx)
	if err != nil {
		return err
	}

	err = stream.Send(&v1.ClientMessage{
		Content: &v1.ClientMessage_RegistrationRequest{
			RegistrationRequest: &regReq,
		},
	})
	if err != nil {
		return err
	}

	for {
		var ctx *handler.HttpContext

		resp, err := stream.Recv()

		if err == io.EOF {
			err = stream.CloseSend()
			if err != nil {
				return err
			}

			return nil
		} else if err == nil {
			if resp.GetHttpRequest() != nil {
				ctx = handler.NewHttpContext(resp)

				ctx, err = a.httpHandler(ctx)
				if err != nil {
					ctx.WithError(err)
				}
			}
		} else {
			return err
		}

		err = stream.Send(ctx.ToClientMessage())
		if err != nil {
			return err
		}
	}
}

func NewApiWorker(opts *ApiWorkerOpts) (*ApiWorker, error) {
	ctx, _ := context.WithTimeout(context.TODO(), constants.NitricDialTimeout())

	conn, err := grpc.DialContext(
		ctx,
		constants.NitricAddress(),
		constants.DefaultOptions()...,
	)
	if err != nil {
		return nil, errors.NewWithCause(
			codes.Unavailable,
			"NewApiWorker: Unable to reach ApiClient",
			err,
		)
	}

	client := v1.NewApiClient(conn)

	return &ApiWorker{
		client:      client,
		apiName:     opts.apiName,
		path:        opts.path,
		methods:     opts.methods,
		httpHandler: opts.httpHandler,
	}, nil
}
