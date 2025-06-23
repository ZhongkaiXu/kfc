package subsystems

// subsystem that we want to control
// maybe there is a better way to describe the subsystem resource
type ResourceConfig struct {
	MemoryLimit string
	CpuShare    string
	CpuSet      string
}

type Subsystem interface {
	Name() string
	Set(path string, res *ResourceConfig) error
	/* The next two function are not necessary in cgroup v2*/
	// Apply(path string, pid int) error
	// Remove(path string) error
}

// Instance of subsystems
// It's a global variable
var (
	SubsystemIns = []Subsystem{
		&CpuSubsystem{},
		&MemorySubsystem{},
		&CpusetSubsystem{},
	}
)
