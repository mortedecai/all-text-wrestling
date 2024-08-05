_indent() {
  local indent=0
  local level=$1
  shift
  output=""

  while [[ $indent -lt $level ]]; do
    (( indent = $indent + 1 ))
    output="${output}  "
  done
  echo -e "${output}$*"
}
