package service

import (
	"database/sql"
	pb "finance-tracker-backend/api/v1"
	"finance-tracker-backend/internal/modules/instruments/repository"
)

type instrumentsService struct {
	db *sql.DB
}

func New(db *sql.DB) repository.InstrumentsRepository {
	return &instrumentsService{db: db}
}

func (i instrumentsService) GetByISIN(isin string) (*pb.Instrument, error) {
	//TODO implement me
	panic("implement me")
}

func (i instrumentsService) GetByTicker(ticker string) (*pb.Instrument, error) {
	rows, err := i.db.Query("SELECT * FROM instruments WHERE ticker = $1", ticker)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	val := make([]string, 5)
	err = rows.Scan(&val[0], &val[1], &val[2], &val[3], &val[4])
	if err != nil {
		return nil, err
	}

	var instrumentType pb.Instrument_InstrumentType

	if val[4] == "ETF" {
		instrumentType = pb.Instrument_ETF
	} else {
		instrumentType = pb.Instrument_STOCK
	}

	return &pb.Instrument{
		Ticker:   val[0],
		Name:     val[1],
		Currency: val[2],
		Isin:     val[3],
		Type:     instrumentType,
	}, nil
}
