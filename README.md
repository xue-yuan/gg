# GG

`gg` is a CLI tool to help developers generate the gitignore which is collected in 'toptal.com' or create a custom template.

> You can also find the templates in [Toptal.com]([https://](https://www.toptal.com/))

![](https://i.imgur.com/LGr3Srr.png)

## Installation

```shell
cd <project_root>
go install gg
```

## Usage

```shell
gg [flags]
gg [command]
```

## Available Commands

```shell
create      Create a custom template
help        Help about any command
list        Show available templates.
```

## Flags

```shell
-a, --append              toggle append mode
-f, --file string         specify filename (default ".gitignore")
-h, --help                help for gg
-t, --templates strings   template name
```

## Next Release

- create a custom template.
- launch on `homebrew`.
