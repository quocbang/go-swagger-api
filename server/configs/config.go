package configs

import "time"

type Options struct {
	ServerConfig string `short:"c" long:"server-config" description:"server configuration file" required:"true"`
}

type DBConnection struct {
	Name     string `yaml:"name"`
	Address  string `yaml:"address"`
	Port     int    `yaml:"port"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Schema   string `yaml:"schema"`
}

type Configs struct {
	DevMode        bool   `yaml:"development_mode"`
	UIDir          string `yaml:"ui_distribution_directory"`
	CreateUIConfig bool   `yaml:"create_ui_configuration"`

	Timeout            time.Duration `yaml:"timeout"`
	WebServiceEndpoint string        `yaml:"web_service_endpoint"`
	PostgreSQL         DBConnection  `yaml:"postgres"`
	// ActiveDirectory     ActiveDirectory `yaml:"active_directory"`
	CorsAllowedOrigins  []string `yaml:"cors_allowed_origins"`
	TokenExpiredSeconds int      `yaml:"token_expired_in_seconds"`
	// FunctionRolePermissions is deprecated.
	FunctionRolePermissions map[string][]string `yaml:"permissions"`
	Printers                map[string]string   `yaml:"printers"`
	FontPath                string              `yaml:"font_path"`
	//StationFunctionConfig   map[string]FunctionAPIPath `yaml:"station_function_config"`
	//MesPath string `yaml:"mes_path"`
}

type options struct {
	schema             string
	pdaServiceEndpoint string
	autoMigrateTables  bool
	migrateCloudTables bool
	adAuth             bool
	//adConfig           ADConfig
}
