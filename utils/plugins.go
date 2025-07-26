package utils


// creating a plugin interface
type Plugin interface {
	Encrypt(data []byte, key []byte) (string, error)
	Decrypt(data string, key []byte) ([]byte, error)

}

// creating a plugin registry
// first: creating variable pluginRegistry
var pluginRegistry = make(map[string]Plugin)


// next: create func to register plugins
func RegisterPlugin(name string, plugin Plugin) {
	pluginRegistry[name] = plugin
}

// creating func to get plugin
func GetPlugin(name string) (Plugin, bool) {
	p, ok := pluginRegistry[name]
	return p, ok
}

// creating func to list plugins
func ListPlugins() []string {
	keys := make([]string, 0, len(pluginRegistry))
	for k := range pluginRegistry {
		keys = append(keys, k)
	}
	return keys
}