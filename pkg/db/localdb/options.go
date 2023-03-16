package localdb

type Option func(org *localdbConfig) error

// OptionLocalMockFile is an Option to use local mock data
func OptionLocalMockFile(path string) Option {
	return func(db *localdbConfig) error {
		db.path = path
		return nil
	}
}
