#!/bin/bash
# set -x

# https://stackoverflow.com/questions/2870992/automatic-exit-from-bash-shell-script-on-error
set -e

# The declare command is used to create the constant variable.
# https://bash.cyberciti.biz/guide/Local_variable
# In MacOSX, using \x1B instead of \e. \033 would be ok for all platforms.
# https://stackoverflow.com/a/28938235
declare -r RED='\033[31m'
declare -r GREEN='\033[32m'
declare -r YELLOW='\033[33m'
declare -r CYAN='\033[96m'
declare -r RESET='\033[0m'

# https://bash.cyberciti.biz/guide/Pass_arguments_into_a_function
# @example
# hello=$(cyan cdi)
# echo hello
# echo "🚀 Try $(cyan \`cdi\`) in your terminal."
function cyan() {
  # cann not use `return`
  # becuase `return` command in Bash can only return status codes which are obviously integer values
  # # https://mezzantrop.wordpress.com/2018/06/27/a-nice-way-to-return-string-values-from-functions-in-bash/
  echo "${CYAN}$@${RESET}"
}

function red() {
  echo "${RED}$@${RESET}"
}

function green() {
  echo "${GREEN}$@${RESET}"
}

success() {
  # All function variables are local. This is a good programming practice.
  # https://bash.cyberciti.biz/guide/Local_variable
  local cdipath="$1"
  local toUpdate=false; [ "$2" = "update" ] && toUpdate=true

  # make it executable
  chmod +x $cdipath && xattr -c $cdipath &&
    [ $toUpdate = false ] && echo "\n                 --- cdi install begin ----\n\n$(green "✅ $cdipath") has been made executable."

  if [ $toUpdate = true ]; then
    green '`cdi` Updated.'
    return
  fi

  echo "## --- cdi begin ---
cdipath=\"$cdipath\"

cdi() {
  target=\$(\$cdipath -fallback \"\$@\")

  echo \$target
  cd \$target
}

# Show debug info
cdi-echo() {
  target=\$(\$cdipath \"\$@\")

  echo \$target
}

# Intelligent \`code\` command \`codi\`
codi() {
  target=\$(\$cdipath \"\$@\")

  echo \$target

  if [[ \$target == *\"no such dirname\"* ]]; then
    # DO NOTHING WHEN NO DIRECTORY FOUND
  else
    code \$(cdi-echo \$1) || echo "code command not exists. Try install it by ${GREEN}https://code.visualstudio.com/docs/setup/mac#_launching-from-the-command-line${RESET}"
  fi
}

# Show cache
alias cdi-stat=\"\$cdipath stat\"

# Clear cache
alias cdi-stat-clear=\"\$cdipath stat --clear && echo -n 'Clear cache success. ' && cdi-stat\"
## --- cdi end ---
" >> ~/.zshrc &&
  echo "✅ Shell functions $(cyan \`cdi\`) and many more functions have been added to your $(green ~/.zshrc)"
  echo "🎉 You are ready to go to use $(cyan \`cdi\`) / $(cyan \`codi\`) / $(cyan \`cdi-echo\`) / $(cyan \`cdi-stat\`) / $(cyan \`cdi-stat-clear\`)."
  echo "🚀 Try $(cyan '$ cdi any_directory_you_like') in your terminal."
  echo "\n                 --- cdi install end ----\n"
}

usage() {
  red '🚨 Empty params. Usage:\n'
  green 'sh scripts/install.sh ~/path/to/downloaded/cdi'

  exit 2
}

# https://bash.cyberciti.biz/guide/Local_variable
[ $# -eq 0 ] && usage

# https://stackoverflow.com/questions/6482377/check-existence-of-input-argument-in-a-bash-shell-script

if [ -f "$1" ]; then
  # https://stackoverflow.com/questions/10067266/when-should-i-wrap-quotes-around-a-shell-variable
  success "$1" "$2"
else
  red "❌ $1 not exists!"
fi
