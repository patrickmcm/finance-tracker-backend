package instruments

import (
	"context"
	pb "finance-tracker-backend/gen/api"
	findb "finance-tracker-backend/gen/db"
	"finance-tracker-backend/internal/platform/controllers"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type instrumentsServer struct {
	pb.UnimplementedInstrumentsServiceServer
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) controllers.Controller {
	return &instrumentsServer{db: db}
}

func (m *instrumentsServer) RegisterController(s *grpc.Server) {
	pb.RegisterInstrumentsServiceServer(s, m)
}

func (m *instrumentsServer) Get(ctx context.Context, instrumentReq *pb.GetInstrumentRequest) (*pb.Instrument, error) {
	if instrumentReq.Ticker == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Ticker")
	}

	queries := findb.New(m.db)

	instrument, err := queries.GetInstrument(ctx, instrumentReq.Ticker)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	var instrumentType pb.Instrument_InstrumentType

	if instrument.Type == "ETF" {
		instrumentType = pb.Instrument_ETF
	} else {
		instrumentType = pb.Instrument_STOCK
	}

	return &pb.Instrument{
		Ticker:   instrument.Ticker,
		Name:     instrument.Name,
		Currency: instrument.Currency.String,
		Isin:     instrument.Isin,
		Type:     instrumentType,
	}, nil
}

func (m *instrumentsServer) List(ctx context.Context, _ *emptypb.Empty) (*pb.InstrumentCollection, error) {
	return &pb.InstrumentCollection{Instruments: []*pb.Instrument{&pb.Instrument{
		Ticker:   "VUAG",
		Name:     "Vanguard S&P 500",
		Currency: "GBX",
		Isin:     "blah",
		Type:     pb.Instrument_ETF,
	}}}, nil
}
