package models

// Config contains all the configuration-related to server and application.
type Config struct {
	ServerConfig      ServerConfig      `json:"server_config,omitempty"`
	ApplicationConfig ApplicationConfig `json:"application_config,omitempty"`
	DatabaseConfig    DatabaseConfig    `json:"db_config"`
}

// ServerConfig contains server-related configuration.
type ServerConfig struct {
	ServerPort   string `json:"server_port,omitempty"`    // ServerPort defines the port at which server is initiated.
	WriteTimeout int    `json:"write_timeout,omitempty"`  // WriteTimeout defines the write timeout for the server.
	ReadTimeout  int    `json:"read_timeout,omitempty"`   // ReadTimeout defines the read timeout for the server.
	IsTLSEnabled bool   `json:"is_tls_enabled,omitempty"` // IsTLSEnabled specifies if TLS is to be enabled.
}

// ApplicationConfig contains application-related configuration.
type ApplicationConfig struct {
	APIPrefix    string       `json:"api_prefix,omitempty"`    // APIPrefix defines the prefix for the API endpoints
	LevelledLogs LevelledLogs `json:"levelled_logs,omitempty"` // LevelledLogs contains all the application logging-related configuration.
}

// LevelledLogs contains all the application logging-related configuration.
type LevelledLogs struct {
	PersistenceLocation   PersistenceLocation `json:"persistence_location,omitempty"`
	EnableLoggingToFile   bool                `json:"enable_logging_to_file,omitempty"`
	EnableLoggingToStdout bool                `json:"enable_logging_to_stdout,omitempty"`
	EnableColors          bool                `json:"enable_colors,omitempty"`
	EnableFullTimestamp   bool                `json:"enable_full_timestamp,omitempty"`
	OutputFormatter       string              `json:"output_formatter"`
}

// PersistenceLocation contains file-system-related configurations.
type PersistenceLocation struct {
	ContainerDirectory  string   `json:"container_directory,omitempty"`
	TargetFileName      []string `json:"targetfile_name,omitempty"`
	TargetFileExtension string   `json:"targetfile_extension,omitempty"`
}

// DatabaseConfig contains database-related configuration.
type DatabaseConfig struct {
	DriverName string `json:"driver_name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Host       string `json:"host"`
	Database   string `json:"database"`
}
