package config

type Config struct {
	Hour      string
	Category  string
	Filter    bool
	Max       int
	BatchSize int
}

func Set(hour, category string, filter bool, max int, batchSize int) *Config {
	return &Config{
		Hour:      hour,
		Category:  category,
		Filter:    filter,
		Max:       max,
		BatchSize: batchSize,
	}
}
