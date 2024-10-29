default:
    air

db:
    sqlite3 data.db < schema.sql
path:
    set -x GOPATH $HOME/go
    set -x PATH $PATH $GOPATH/bin
    set -x PATH $PATH /home/janek/Dokumente/schlauerlauer/liesl/tmp
