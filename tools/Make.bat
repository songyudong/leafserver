
: Generate go source file by sproto
sprotogen --go_out=message_gen.go --package=message message.sp
@IF %ERRORLEVEL% NEQ 0 pause

: Convert to standard sproto file
sprotogen --sproto_out=message_gen.sproto message.sp
@IF %ERRORLEVEL% NEQ 0 pause

