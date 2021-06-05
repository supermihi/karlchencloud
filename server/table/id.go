package table

import "strconv"

type Id int64

const InvalidId Id = -1

func (t Id) String() string {
	return strconv.FormatInt(int64(t), 10)
}
func ParseTableId(idStr string) (Id, error) {
	idInt, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return InvalidId, err
	}
	return Id(idInt), nil
}
