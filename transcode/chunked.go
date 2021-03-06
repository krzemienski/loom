package transcode

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/zsiec/loom/analyze"
)

type ChunkedSvc struct {
	Analyzer analyze.Svc
	Logger   zerolog.Logger
}

func (s ChunkedSvc) Create(ctx context.Context, req CreateRequest) (CreateResponse, error) {
	s.ensure()

	src := req.Source.URL

	info, err := s.Analyzer.Analyze(ctx, src)
	if err != nil {
		return CreateResponse{}, fmt.Errorf("analyzing source '%s': %w", src, err)
	}

	s.Logger.Info().Msgf("analyzed source %q: %+v", src, info)

	// todo

	return CreateResponse{
		ID:     "todo",
		Status: StatusQueued,
	}, nil
}

func (s ChunkedSvc) Status(context.Context, string) (StatusResponse, error) {
	return StatusResponse{}, nil
}

func (s *ChunkedSvc) ensure() {
	if s.Analyzer == nil {
		s.Analyzer = &analyze.MediainfoSvc{}
	}
}
