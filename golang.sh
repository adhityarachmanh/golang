#! /bin/bash
# cre : arh 
# upd : arh 
# ver : 1.0

CBLUE="\x1b[34;1m"
CRED="\x1b[31;1m"
CGREEN="\x1b[32;1m"
CYELLOW="\x1b[33;1m"
CRESET="\x1b[39;49;00m"
TERR="\e[1;40;97m"
THIDE="\e[8m"

Init(){
    EXE=$0
    DIR=$(dirname "${EXE}")
    cd $DIR
}

GTemplate(){
    
}

GModel(){
    Init
}
GRoute(){
    Init
}

Generate(){
    case "$1" in
        model|m)
            GModel $2
        route|r)
            GRoute $2
            ;;
        --help|--h|help|h)
            GenerateHelp
            ;;
        *)
            echo -e 'For more detailed help run "generate --help"'
            ;;
    esac
}

case "$1" in
    generate|g)
        Generate $2 $3
        ;;
    --help|--h|help|h)
        ShowHelp
        ;;
    --version|--v|version|v)
        ShowVersion
        ;;
    *)
        echo -e 'For more detailed help run "--help"'
        ;;
esac
exit 0