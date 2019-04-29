# Example plugin

This is a sample plugin for Gilbert task runner that prints a message to log


## Usage

### Example

```yaml
tasks:
  hello-world:
    - plugin: '@go-gilbert/example'
      params:
        message: 'hello world'
```

### Params

* `message` - a message to print

## Plugin How To

### Compatibility

Each plugin should follow these simple rules:

* Should use compatible Go version (currently: *Go 1.12.x*)

* Should use **go modules** to avoid package scopes issue during library load process (see [issue](https://github.com/golang/go/issues/30838))

* Should use the same [**SDK version**](https://github.com/go-gilbert/gilbert-sdk) as target Gilbert version

### Structure

Each plugin exports two procedures:

* `GetPluginName() string` - returns plugin name, recommended format is `@author/name`
* `NewPlugin(sdk.ScopeAccessor, sdk.PluginParams, sdk.Logger) (sdk.Plugin, error)` - plugin instance constructor