package instruments

import (
	"context"
	pb "finance-tracker-backend/gen/api/v1"
	"finance-tracker-backend/gen/conv"
	findb "finance-tracker-backend/gen/db"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedInstrumentsServiceServer
	db *pgxpool.Pool
}

func (m *Server) RegisterController(s *grpc.Server, db *pgxpool.Pool) {
	m.db = db
	pb.RegisterInstrumentsServiceServer(s, m)
}

func (m *Server) Get(ctx context.Context, request *pb.GetInstrumentRequest) (*pb.Instrument, error) {
	if request.Ticker == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Ticker")
	}

	queries := findb.New(m.db)

	instrument, err := queries.GetInstrument(ctx, request.Ticker)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	converter := conv.ConverterImpl{}
	convertedInstrument := converter.ConvertInstrument(instrument)

	return &convertedInstrument, nil
}

func (m *Server) List(ctx context.Context, request *pb.ListInstrumentRequest) (*pb.InstrumentCollection, error) {
	converter := conv.ConverterImpl{}

	queries := findb.New(m.db)

	instruments, err := queries.ListInstruments(ctx)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	instrumentsConv := converter.ConvertInstruments(instruments)

	instrumentCollection := pb.InstrumentCollection{Instruments: instrumentsConv}

	return &instrumentCollection, nil
}
