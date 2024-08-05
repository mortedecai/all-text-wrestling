cmd_test_usage() {
  _indent 0 "./project test usage"
  _indent 0 ""
  _indent 0 "./project test [<options>]"
  _indent 1 "WHERE <options> includes:"
  _indent 2 "-c,--coveragefile (default: .reports/coverage.out)"
  _indent 2 "-j,--jsonfile The name of the file to save JSON output to"
  _indent 2 "-H,--HTML The name of the HTML output file (optional)"
}

cmd_test() {
  local cmdString=$1
  shift
  local flag_help flag_verbose flag_html
  local arg_coveragefile=(coverage.out)
  local arg_htmlfile arg_jsonfile

  if [[ ! -d ${PROJECT_ROOT}/bin ]]; then
    mkdir -p ${PROJECT_ROOT}/bin
  fi

  zmodload zsh/zutil
  zparseopts -D -F -K -- \
  {c,-coveragefile}:=arg_coveragefile \
  {h,-help}=flag_help \
  {j,-jsonfile}:=arg_jsonfile \
  {v,-verbose}=flag_verbose \
  {H,-HTML}:=arg_htmlfile ||
  return 1

  [[ -z "$flag_help" ]] || { ${cmdString}_usage && exit ${EXIT_USAGE}; }

  local coveragefile=${PROJECT_ROOT}/.reports/${arg_coveragefile[-1]}
  local htmlfile=${PROJECT_ROOT}/.reports/${arg_htmlfile[-1]}
  local jsonfile=${PROJECT_ROOT}/.reports/${arg_jsonfile[-1]}

  if [[ ! -d ${PROJECT_ROOT}/.reports ]]; then
    mkdir ${PROJECT_ROOT}/.reports
  fi

  local flag_json=""
  local redir_json=""

  if [[ ! -z "${arg_jsonfile[-1]}" ]]; then
    redir_json="${jsonfile}"
  fi

 if [[ -z "${arg_jsonfile[-1]}" ]]; then
   go test -covermode=count -coverprofile=${coveragefile} ./...
 else
   go test -covermode=count -coverprofile=${coveragefile} -json ./... > ${jsonfile}
 fi

  if [[ ! -z "$arg_htmlfile" ]]; then
    go tool cover -html=${coveragefile} -o ${htmlfile}
  fi
}
