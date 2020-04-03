package pgstore

// Row abstracts slq.Row
type Row interface {
	Scan(args ...interface{}) error
}
