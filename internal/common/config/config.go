package config

type Config struct {
	NotRegistered bool
	WorkSchedule  bool
}

func Set(notRegistered, workSchedule bool) *Config {
	return &Config{
		NotRegistered: notRegistered,
		WorkSchedule:  workSchedule,
	}
}
