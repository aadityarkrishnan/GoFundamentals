package main

import (
	"fmt"
	"strings"
)

// Plugin defines the interface that all plugins must implement
type Plugin interface {
	Name() string               // Returns the plugin's name
	Execute(data string) string // Processes the input data and returns the result
}

// UppercasePlugin converts input data to uppercase
type UppercasePlugin struct{}

func (p *UppercasePlugin) Name() string {
	return "UppercasePlugin"
}

func (p *UppercasePlugin) Execute(data string) string {
	return strings.ToUpper(data)
}

// ReversePlugin reverses the input string
type ReversePlugin struct{}

func (p *ReversePlugin) Name() string {
	return "ReversePlugin"
}

func (p *ReversePlugin) Execute(data string) string {
	runes := []rune(data)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// AppendPlugin appends a custom suffix to the input string
type AppendPlugin struct {
	Suffix string
}

func (p *AppendPlugin) Name() string {
	return "AppendPlugin"
}

func (p *AppendPlugin) Execute(data string) string {
	return data + p.Suffix
}

// PluginManager manages the registration and execution of plugins
type PluginManager struct {
	plugins map[string]Plugin
}

// NewPluginManager creates a new PluginManager
func NewPluginManager() *PluginManager {
	return &PluginManager{
		plugins: make(map[string]Plugin),
	}
}

// Register adds a new plugin to the manager
func (pm *PluginManager) Register(plugin Plugin) {
	pm.plugins[plugin.Name()] = plugin
}

// ListPlugins returns the names of all registered plugins
func (pm *PluginManager) ListPlugins() []string {
	names := make([]string, 0, len(pm.plugins))
	for name := range pm.plugins {
		names = append(names, name)
	}
	return names
}

// Execute runs the specified plugin with the given input data
func (pm *PluginManager) Execute(name string, data string) string {
	plugin, exists := pm.plugins[name]
	if !exists {
		return fmt.Sprintf("Plugin '%s' not found", name)
	}
	return plugin.Execute(data)
}

func main() {
	manager := NewPluginManager()

	// Register plugins
	manager.Register(&UppercasePlugin{})
	manager.Register(&ReversePlugin{})
	manager.Register(&AppendPlugin{Suffix: " - GoLang"})

	// List all registered plugins
	fmt.Println("Registered Plugins:")
	for _, name := range manager.ListPlugins() {
		fmt.Println("-", name)
	}

	// Input data
	input := "hello"

	// Use UppercasePlugin
	fmt.Println("\nProcessing with UppercasePlugin:")
	fmt.Println(manager.Execute("UppercasePlugin", input)) // Output: HELLO

	// Use ReversePlugin
	fmt.Println("\nProcessing with ReversePlugin:")
	fmt.Println(manager.Execute("ReversePlugin", input)) // Output: olleh

	// Use AppendPlugin
	fmt.Println("\nProcessing with AppendPlugin:")
	fmt.Println(manager.Execute("AppendPlugin", input)) // Output: hello - GoLang

	// Handle unregistered plugin
	fmt.Println("\nProcessing with UnknownPlugin:")
	fmt.Println(manager.Execute("UnknownPlugin", input)) // Output: Plugin 'UnknownPlugin' not found
}
