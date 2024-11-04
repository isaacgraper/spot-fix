package config

type Config struct {
	ProcessBatch  bool
	Batch         int
	Max           int
	Hour          string
	Category      string
	NotRegistered bool
	WorkSchedule  bool
}

func Set(ProcessBatch bool, batch, max int, hour, category string, notRegistered, workSchedule bool) *Config {
	return &Config{
		ProcessBatch:  ProcessBatch,
		Batch:         batch,
		Max:           max,
		Hour:          hour,
		Category:      category,
		NotRegistered: notRegistered,
		WorkSchedule:  workSchedule,
	}
}
