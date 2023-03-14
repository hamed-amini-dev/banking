package localdb

//OptionLocalMockFile is an Option to use local mock data
type Option func(org *localdbConfig) error

func OptionLocalMockFile(path string) Option {
	return func(db *localdbConfig) error {
		db.path = path
		return nil
	}
}
