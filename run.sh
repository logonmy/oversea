#!/bin/sh

#ulimit -n 1024000
#ulimit -c unlimited

p='oversea_pro'


case $2 in
    stg)
        p='oversea_stg'
        ;;
    pro)
        p='clibscenter_pro'
        ;;
    *)
        p='oversea_dev'
        ;;
esac

KillServer()
{
    pid=`ps x | grep "$p" | sed -e '/mykill/d' | sed -e '/grep/d' | sed -e '/tail/d' | awk '{print $1}'`
    pid=`echo $pid | awk '{print $1}'`
    while [ ! -z "$pid" ]
    do
            kill -9 $pid
        pid=`ps x | grep "$p" | sed -e '/grep/d' | sed -e '/tail/d' | awk '{print $1}'`
            pid=`echo $pid | awk '{print $1}'`
    done
}

case $1 in
    start)
        KillServer
        sleep 1
        nohup ./$p >> ./out.log 2>&1 &
        sleep 1
        echo ""
        ps -elf | grep $p
        ;;
    stop)
        KillServer
        sleep 1
        echo ""
        ps -elf | grep $p
        ;;
    restart)
        KillServer
        sleep 1
        nohup ./$p >> ./out.log 2>&1 &
        sleep 1
        echo ""
        ps -elf | grep $p
        ;;
    *)
        KillServer
        sleep 1
        nohup ./$p >> ./out.log 2>&1 &
        sleep 1
        echo ""
        ps -elf | grep $p
        ;;
esac
