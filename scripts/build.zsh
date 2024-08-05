cmd_build_usage() {
  _indent 0 "./project build usage"
  _indent 0 ""
  _indent 0 "./project build [<options>]"
  _indent 1 "WHERE <options> includes:"
  _indent 2 "-o outputFile (default: atw)"
}

cmd_build() {
  local cmdString=$1
  shift
  local flag_help flag_verbose
  local arg_outputfile=(${PROJECT_ROOT}/bin/atw)

  if [[ ! -d ${PROJECT_ROOT}/bin ]]; then
    mkdir -p ${PROJECT_ROOT}/bin
  fi

  zmodload zsh/zutil
  zparseopts -D -F -K -- \
  {h,-help}=flag_help \
  {v,-verbose}=flag_verbose \
  {o,-outputFile}:=arg_outputfile ||
  return 1

  [[ -z "$flag_help" ]] || { ${cmdString}_usage && exit ${EXIT_USAGE}; }

  go build ${flag_verbose} -o ${arg_outputfile} ${PROJECT_ROOT}/cmd/atw/atw.go
}