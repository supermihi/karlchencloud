package tables

import (
	"math/rand"
	"strconv"
)

type TableId int64

const InvalidTableId TableId = -1

func (t TableId) String() string {
	return strconv.FormatInt(int64(t), 10)
}
func ParseTableId(idStr string) (TableId, error) {
	idInt, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return InvalidTableId, err
	}
	return TableId(idInt), nil
}

func randomTableId() TableId {
	return TableId(rand.Int63())
}
