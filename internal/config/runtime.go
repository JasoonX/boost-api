package config

type Runtime interface {
	AuthDisabled() bool
	Environment() string
	Version() string
}

type runtime struct {
	auth        string
	environment string
	version     string
}

type yamlRuntimeConfig struct {
	Auth        string `yaml:"auth"`
	Environment string `yaml:"environment"`
	Version     string `yaml:"version"`
}

const (
	EnvironmentLocal   = "local"
	EnvironmentDev     = "dev"
	EnvironmentStaging = "staging"
	EnvironmentProd    = "prod"
)

func NewRuntime(auth, env, version string) Runtime {
	return &runtime{
		auth:        auth,
		environment: env,
		version:     version,
	}
}

func (d runtime) AuthDisabled() bool {
	return d.auth == "disabled"
}

func (d runtime) Environment() string {
	return d.environment
}

func (d runtime) Version() string {
	return d.version
}
