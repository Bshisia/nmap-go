package timing

import "time"

type TimingTemplate struct {
	Name    string
	Timeout time.Duration
	Delay   time.Duration
}

var Templates = map[string]TimingTemplate{
	"T0": {"Paranoid", 5 * time.Second, 5 * time.Minute},
	"T1": {"Sneaky", 3 * time.Second, 15 * time.Second},
	"T2": {"Polite", 2 * time.Second, 400 * time.Millisecond},
	"T3": {"Normal", 1 * time.Second, 0},
	"T4": {"Aggressive", 500 * time.Millisecond, 0},
	"T5": {"Insane", 250 * time.Millisecond, 0},
}

func GetTiming(template string) TimingTemplate {
	if t, exists := Templates[template]; exists {
		return t
	}
	return Templates["T3"] // Default to normal
}