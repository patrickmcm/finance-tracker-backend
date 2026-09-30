package conv

import (
	pb "finance-tracker-backend/gen/api/v1"
	findb "finance-tracker-backend/gen/db"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// goverter:converter
// goverter:output:file ../../../gen/conv/generated.go
// goverter:extend TextToString
// goverter:extend Int4ToInt32
// goverter:extend TimestampzToPbTimestamp
// goverter:extend NumericToNumericPb
// goverter:enum:unknown @panic
type Converter interface {
	ConvertInstruments(source []findb.Instrument) []*pb.Instrument
	ConvertInstrumentPrices(source []findb.InstrumentPrice) []*pb.InstrumentPrice

	// goverter:enum:map InstrumentTypeETF Instrument_ETF
	// goverter:enum:map InstrumentTypeSTOCK Instrument_STOCK
	ConvertInstrumentType(source findb.InstrumentType) pb.Instrument_InstrumentType

	// goverter:ignore state sizeCache unknownFields
	ConvertInstrument(source findb.Instrument) pb.Instrument

	// goverter:ignore state sizeCache unknownFields
	ConvertInstrumentPrice(source findb.InstrumentPrice) pb.InstrumentPrice
}

func TextToString(i pgtype.Text) string {
	return i.String
}

func TimestampzToPbTimestamp(i pgtype.Timestamptz) timestamppb.Timestamp {
	return *timestamppb.New(i.Time)
}

func Int4ToInt32(i pgtype.Int4) int32 {
	return i.Int32
}

func NumericToNumericPb(i pgtype.Numeric) pb.Numeric {
	return pb.Numeric{Unscaled: i.Int.Int64(), Exp: i.Exp}
}
