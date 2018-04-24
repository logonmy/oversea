#!/bin/sh

#ulimit -n 1024000
#ulimit -c unlimited

p='oversea'
appEnv='dev'

case $2 in
    stg)
        appEnv='-env=stg'
        ;;
    prod)
        appEnv='-env=prod'
        ;;
    *)
        appEnv='-env=dev'
        ;;
esac

KillServer()
{
    pid=`ps x | grep "$p" | grep "$appEnv" | sed -e '/mykill/d' | sed -e '/grep/d' | sed -e '/tail/d' | awk '{print $1}'`
    pid=`echo $pid | awk '{print $1}'`
    while [ ! -z "$pid" ]
    do
            kill -9 $pid
        pid=`ps x | grep "$p"| grep "$appEnv" | sed -e '/grep/d' | sed -e '/tail/d' | awk '{print $1}'`
            pid=`echo $pid | awk '{print $1}'`
    done
}

case $1 in
    start)
        KillServer
        sleep 1
        nohup ./$p $appEnv>> ./out.log 2>&1 &
        sleep 1
        echo ""
        ps -elf | grep $p|grep "$appEnv"
        ;;
    stop)
        KillServer
        sleep 1
        echo ""
        ps -elf | grep $p| grep "$appEnv"
        ;;
    restart)
        KillServer
        sleep 1
        nohup ./$p >> ./out.log 2>&1 &
        sleep 1
        echo ""
        ps -elf | grep $p|grep "$appEnv"
        ;;
    *)
        KillServer
        sleep 1
        nohup ./$p $appEnv>> ./out.log 2>&1 &
        sleep 1
        echo ""
        ps -elf | grep $p|grep "$appEnv"
        ;;
esac
