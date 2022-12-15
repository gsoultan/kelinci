package kelinci

import "fmt"

type Config struct {
	Host     string
	Port     string
	UserName string
	Password string
	VHost    string
}

func (c *Config) GetUri() string {
	return fmt.Sprintf("amqp://%s:%s@%s%s", c.UserName, c.Password, c.Host, c.VHost)
}

// ConfigBuilder Builder Object for Config
type ConfigBuilder struct {
	host     string
	port     string
	userName string
	password string
	vHost    string
}

// NewConfigBuilder Constructor for ConfigBuilder
func NewConfigBuilder() *ConfigBuilder {
	o := new(ConfigBuilder)
	return o
}

// Build Method which creates Config
func (c *ConfigBuilder) Build() *Config {
	o := new(Config)
	o.Host = c.host
	o.Password = c.password
	o.UserName = c.userName
	o.VHost = c.vHost
	return o
}

// SetHost Setter method for the field host of type string in the object ConfigBuilder
func (c *ConfigBuilder) SetHost(host string) {
	c.host = host
}

// SetPort Setter method for the field port of type string in the object ConfigBuilder
func (c *ConfigBuilder) SetPort(port string) {
	c.port = port
}

// SetUserName Setter method for the field userName of type string in the object ConfigBuilder
func (c *ConfigBuilder) SetUserName(userName string) {
	c.userName = userName
}

// SetPassword Setter method for the field password of type string in the object ConfigBuilder
func (c *ConfigBuilder) SetPassword(password string) {
	c.password = password
}

// SetVHost Setter method for the field vHost of type string in the object ConfigBuilder
func (c *ConfigBuilder) SetVHost(vHost string) {
	c.vHost = vHost
}
