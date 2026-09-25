package v1

import (
	pb "finance-tracker-backend/gen/api"
	findb "finance-tracker-backend/gen/db"
	"github.com/jackc/pgx/v5/pgtype"
)

// goverter:converter
// goverter:output:file ../../gen/conv/generated.go
// goverter:extend TextToString
// goverter:enum:unknown @panic
type Converter interface {
	ConvertInstruments(source []findb.Instrument) []*pb.Instrument

	// goverter:enum:map InstrumentTypeETF Instrument_ETF
	// goverter:enum:map InstrumentTypeSTOCK Instrument_STOCK
	ConvertInstrumentType(source findb.InstrumentType) pb.Instrument_InstrumentType

	// goverter:ignore state sizeCache unknownFields
	ConvertInstrument(source findb.Instrument) pb.Instrument
}

func TextToString(i pgtype.Text) string {
	return i.String
}
