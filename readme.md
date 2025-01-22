# Example plugin

## This repo is obsolete, please see [this issue](https://github.com/go-gilbert/gilbert/issues/72)

This is a sample plugin for Gilbert task runner that prints a message to log

## Actions

### `hello-world` 

A simple action that prints text to the console

#### Usage

```yaml
plugins:
  - github://github.com/go-gilbert/gilbert-plugin-example

tasks:
  hello-world:
    - action: 'example-plugin:hello-world'
      params:
        message: 'hello world'
```

#### Params

* `message` - a message to print

## Development

### Compatibility

Each plugin should follow these simple rules:

* Should use compatible Go version (currently: *Go 1.12.x*)

* Should use **go modules** to avoid package scopes issue during library load process (see [issue](https://github.com/golang/go/issues/30838))

* Should use the same [**SDK version**](https://github.com/go-gilbert/gilbert-sdk) as target Gilbert version

### Structure

Each plugin should exports two procedures:


#### `GetPluginName() string`

Returns plugin name

#### `func GetPluginActions() sdk.Actions`

Returns a set of plugin actions exported by plugin
