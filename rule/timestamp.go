package rule

import (
	"bytes"
	"time"
)

type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(body []byte) (err error) {
	t, err = time.Parse(time.RFC3339, string(body))
	return
}
