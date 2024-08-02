package profd

const (
	DefaultEnableStartupProfiling = false
	DefaultStartupProfilingSec    = 300
	DefaultStartupProfilingMode   = "mem"
	DefaultProfilingOutputDir     = "/tmp"
	DefaultEnableGoPs             = false
	DefaultGoPsAddr               = ":"
	DefaultEnableHTTPPprof        = false
	DefaultEnableStatsViz         = false
)

type Cfg struct {
	EnableStartupProfiling bool   `yaml:"enableStartupProfiling"   env:"PROF_ENABLE_STARTUP_PROFILING"   flag-long:"prof-enable-startup-profiling"`
	StartupProfilingSec    int    `yaml:"startupProfilingSec"      env:"PROF_STARTUP_PROFILING_SEC"      flag-long:"prof-startup-profiling-sec"`
	StartupProfilingMode   string `yaml:"startupProfilingMode"     env:"PROF_STARTUP_PROFILING_MODE"     flag-long:"prof-startup-profiling-mode"`
	ProfilingOutputDir     string `yaml:"profilingOutputDir"       env:"PROF_OUTPUT_DIR"                 flag-long:"prof-output-dir"`

	EnableGoPs bool   `yaml:"enableGoPs"               env:"PROF_ENABLE_GOPS"                flag-long:"prof-enable-gops"`
	GoPsAddr   string `yaml:"goPsAddr"                 env:"PROF_GOPS_ADDR"                  flag-long:"prof-gops-addr"`

	EnableHTTPPprof bool `yaml:"enableHTTPPprof" env:"PROF_ENABLE_HTTP_PPROF" flag-long:"prof-enable-http-pprof"`

	EnableStatsViz bool `yaml:"enableStatsViz" env:"PROF_ENABLE_STATS_VIZ" flag-long:"prof-enable-stats-viz"`
}

func NewCfg() *Cfg {
	return &Cfg{
		EnableStartupProfiling: DefaultEnableStartupProfiling,
		StartupProfilingSec:    DefaultStartupProfilingSec,
		StartupProfilingMode:   DefaultStartupProfilingMode,
		ProfilingOutputDir:     DefaultProfilingOutputDir,
		EnableGoPs:             DefaultEnableGoPs,
		GoPsAddr:               DefaultGoPsAddr,
		EnableHTTPPprof:        DefaultEnableHTTPPprof,
		EnableStatsViz:         DefaultEnableStatsViz,
	}
}
