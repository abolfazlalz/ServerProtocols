package net

type ServerType string

const (
	TCPServer ServerType = "tcp"
	UDPServer            = "udp"
)

type ServerConfigFunc func(*ServerConfig)

func ConfigTcpServer(s *ServerConfig) {
	s.Server = TCPServer
}

func ConfigUdpServer(s *ServerConfig) {
	s.Server = UDPServer
}

func DefaultConfig() ServerConfig {
	return ServerConfig{
		Host:   "",
		Port:   "8000",
		Server: TCPServer,
	}
}

type ServerConfig struct {
	Host   string
	Port   string
	Server ServerType
}
