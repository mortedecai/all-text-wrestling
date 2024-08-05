_command() {
  local cmdString=$1
  shift
  local cmd=$1
  shift
  local cmdFile=${PROJECT_ROOT}/scripts/${cmd}.zsh

  if [[ -e "${cmdFile}" ]]; then
    . ${cmdFile}
  fi
  ${cmdString}_${cmd} "${cmdString}_${cmd}" "$@"
}
