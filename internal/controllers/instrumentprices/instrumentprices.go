package instrumentprices

import (
	"context"
	pb "finance-tracker-backend/gen/api/v1"
	"finance-tracker-backend/gen/conv"
	findb "finance-tracker-backend/gen/db"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"reflect"
)

type Server struct {
	pb.UnimplementedInstrumentsPricesServiceServer
	db *pgxpool.Pool
}

func (m *Server) RegisterController(s *grpc.Server, db *pgxpool.Pool) {
	m.db = db
	pb.RegisterInstrumentsPricesServiceServer(s, m)
}

func (m *Server) Get(ctx context.Context, request *pb.GetInstrumentPricesRequest) (*pb.InstrumentPricesCollection, error) {
	if request.Ticker == "" || request.DateTo.String() == "" || request.DateFrom.String() == "" {
		var missing string

		s := reflect.ValueOf(*request)
		for k, v := range s.Fields() {
			if v.String() == "" {
				missing += fmt.Sprintf(" %s", k.Name)
			}
		}

		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: %s", missing)
	}

	converter := conv.ConverterImpl{}

	queries := findb.New(m.db)

	dateFromTz := pgtype.Timestamptz{
		Time:  request.DateFrom.AsTime(),
		Valid: true,
	}

	dateToTz := pgtype.Timestamptz{
		Time:  request.DateTo.AsTime(),
		Valid: true,
	}

	reqParams := findb.GetInstrumentPricesParams{
		Ticker:   request.Ticker,
		Datefrom: dateFromTz,
		Dateto:   dateToTz,
	}

	instrumentPrices, err := queries.GetInstrumentPrices(ctx, reqParams)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	instrumentPricesPb := converter.ConvertInstrumentPrices(instrumentPrices)

	instrumentPricesCollection := pb.InstrumentPricesCollection{Prices: instrumentPricesPb}

	return &instrumentPricesCollection, nil
}
