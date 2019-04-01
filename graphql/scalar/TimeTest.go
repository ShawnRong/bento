package scalar

import (
	"errors"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalTimeTest(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(t.Format(time.RFC3339)))
	})
}

func UnmarshalTimeTest(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(string); ok {
		return time.Parse(time.RFC3339, tmpStr)
	}
	return time.Time{}, errors.New("time should be RFC3339 formatted string")
}
