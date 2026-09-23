package controller

import (
	"context"
	pb "finance-tracker-backend/api/v1"
	"finance-tracker-backend/internal/modules/instruments/repository"
	"finance-tracker-backend/internal/platform/controllers"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type instrumentsServer struct {
	pb.UnimplementedInstrumentsServiceServer
	repo repository.InstrumentsRepository
}

func New(repo repository.InstrumentsRepository) controllers.Controller {
	return &instrumentsServer{repo: repo}
}

func (m *instrumentsServer) RegisterController(s *grpc.Server) {
	pb.RegisterInstrumentsServiceServer(s, m)
}

func (m *instrumentsServer) Get(ctx context.Context, instrumentReq *pb.GetInstrumentRequest) (*pb.Instrument, error) {
	instrument, err := m.repo.GetByTicker(instrumentReq.Ticker)
	if err != nil {
		return nil, err
	}

	return instrument, nil
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
