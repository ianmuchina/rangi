# cli mock

## config

theme-name: string
config-file: file in json format

## config dir folder structure

```yaml
templates:
  # Terminals
  alacritty:
  kitty:
  tty:
  xfce-terminal:

  # Editors
  vim:
  neovim:
  kakoune:
  mousepad:
  lightline:

  #Other
  gtk:
```

A template's file extension is used to decide the templating engine

supported engines: mustache
