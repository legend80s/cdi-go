# set -x

RED='\033[31m'
GREEN='\033[32m'
YELLOW='\033[33m'
CYAN='\033[96m'
RED='\033[0;31m'
RESET='\033[0m' # No Color

# function definitions

fail() {
  echo "1. Download ${GREEN}https://github.com/legend80s/cdi-go/raw/master/cdi-v5${RESET}"
  echo "2. $ sh ./scripts/install.sh ~/path/to/downloaded/cdi\n"
}

success() {
  cdipath=$1

  # make it executable
  chmod +x $cdipath && xattr -c $cdipath

  # echo "${GREEN}cdi executable has been downloaded to \"$cdipath\".${RESET}"

  # Add the shell functions to your .zshrc because you cannot change shell directory in golang process.

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

# Show cache
alias cdi-stat=\"\$cdipath stat\"

# Clear cache
alias cdi-stat-clear=\"\$cdipath stat --clear && echo -n 'Clear cache success. ' && cdi-stat\"

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
# cdi end
" >> ~/.zshrc
}

finally() {
  echo "${YELLOW}Before using \`cdi\` or \`codi\`, YOU SHOULD RUN \`source ~/.zshrc\` BY YOURSELF to make these commands to take effect.${RESET}"
}

# execute

if [[ $# -eq 0 ]]; then
  echo "Empty arguments. You should pass the downloaded cdi file path as the first argument.\n"

  fail
else
  success $1
fi;

finally
