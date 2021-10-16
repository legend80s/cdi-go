# set -x

RED='\033[31m'
GREEN='\033[32m'
YELLOW='\033[33m'
CYAN='\033[96m'
RED='\033[0;31m'
RESET='\033[0m'

success() {
  cdipath=$1

  # make it executable
  chmod +x $cdipath && xattr -c $cdipath &&
    echo "${GREEN}✅ $cdipath${RESET} has been made executable.${RESET}"

  echo "# cdi begin
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
    code \$(cdi-echo \$1)
  fi
}

# Show cache
alias cdi-stat=\"\$cdipath stat\"

# Clear cache
alias cdi-stat-clear=\"\$cdipath stat --clear && echo -n 'Clear cache success. ' && cdi-stat\"
# cdi end
" >> ~/.zshrc &&
  echo "✅ Shell functions ${GREEN}\`cdi\`${RESET} and ${GREEN}\`codi\`${RESET} has been added to your ${GREEN}~/.zshrc${RESET}"
  echo "${GREEN}🎉 You are ready to go to use \`cdi\` and \`codi\`.${RESET}"
}

# https://stackoverflow.com/questions/6482377/check-existence-of-input-argument-in-a-bash-shell-script
if [ -z "$1" ]; then
  echo "${RED}🚨 Empty params.${RESET}\n"

  echo "${GREEN}Example: \n"
  echo "sh scripts/install.sh ~/path/to/downloaded/cdi${RESET}"
else
  if [ -f "$1" ]; then
    success $1
  else
    echo "${RED}❌ $1 not exists!${RESET}"
  fi
fi
