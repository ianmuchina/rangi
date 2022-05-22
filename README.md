
# rangi

A command line utlity to generate base16 themes from local vscode themes.

## Dependencies
- flavors -  to apply the themes
- fzf - for the interactive command

The generated themes are stored at `.local/share/flavours/base16/schemes/vscode/`

Bash function to switch themes.

```bash
#!/bin/bash
fv() {
  flavours list |
    xargs -n1 |
    fzf --preview "flavours apply {} 2>/dev/null ;term-16-test.sh"
}
```

Bash script to to display terminal colors. Add it ro your `$PATH` as `term-16-test.sh`

```bash
#!/bin/bash
#
# This file echoes a bunch of color codes to the terminal to demonstrate
# what's available. Each line is the color code of one forground color,
# out of 17 (default + 16 escapes), followed by a test use of that color
# on all nine background colors (default + 8 escapes).
#
T='gYw'   # The test text
echo -e "\n                 40m     41m     42m     43m     44m     45m     46m     47m";
for FGs in '    m' '   1m' '  30m' '1;30m' '  31m' '1;31m' '  32m' '1;32m' '  33m' '1;33m' '  34m' '1;34m' '  35m' '1;35m' '  36m' '1;36m' '  37m' '1;37m';
  do FG=${FGs// /}
  echo -en " $FGs \033[$FG  $T  "
  for BG in 40m 41m 42m 43m 44m 45m 46m 47m;
    do echo -en "$EINS \033[$FG\033[$BG  $T \033[0m\033[$BG \033[0m";
  done
  echo;
done
echo
```

## Screenshots

Monokai Pro
![image](https://user-images.githubusercontent.com/49595512/169673165-73fb7d89-7d5c-4283-a00d-e2c5f9fb6c62.png)
Night owl
![image](https://user-images.githubusercontent.com/49595512/169673192-19dc1976-5ff5-4388-b2f8-0f10701a116a.png)
Ayu Dark
![image](https://user-images.githubusercontent.com/49595512/169673242-56665575-cca3-46e6-a85c-fe1c79f36573.png)

