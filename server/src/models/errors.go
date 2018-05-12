package models

type DatabaseError struct {
	err string
}

func (dberr *DatabaseError)Error() string {
	return dberr.err
}