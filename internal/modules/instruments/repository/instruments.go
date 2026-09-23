package repository

import pb "finance-tracker-backend/api/v1"

type InstrumentsRepository interface {
	GetByISIN(isin string) (*pb.Instrument, error)
	GetByTicker(ticker string) (*pb.Instrument, error)
}
