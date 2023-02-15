#!/bin/bash

BIN_FILE='/data/wwwroot/ori/ori'
LOG_PATH='/data/wwwroot/ori/ori.out'

#判断程序是否已经在运行
status_script(){
    pids=`ps aux|grep ${BIN_FILE}|grep -v grep|grep -v sh|awk '{print $2}'`
    if [ "$pids" ]
    then
        #echo ${0}'在运行中'
        return 1
    else
        #echo $0'未启动'
        return 2
    fi
}

#启动脚本，先判断脚本是否已经在运行
start_script(){
    status_script
    if [ $? -eq 1 ]
    then
        echo ${0}' 已经在运行中了'
    else
        echo '启动'${0}'中...'
        nohup ${BIN_FILE}>>${LOG_PATH} 2>&1 &

        echo '启动完毕'
    fi
}

#优雅退出
stop_script(){
    status_script
    if [ $? -ne 1 ]
    then
        echo ${0}' 不是运行状态'
    else
        ps -ef|grep ${BIN_FILE}|grep -v grep|awk '{print $2}'|xargs kill -USR1
        echo '优雅退出完毕'
    fi
}

#立刻退出
quit_script(){
    status_script
    if [ $? -ne 1 ]
    then
        echo ${0}' 不是运行状态'
    else
        ps -ef|grep ${BIN_FILE}|grep -v grep|awk '{print $2}'|xargs kill
        echo '立刻退出完毕'
    fi
}

#平滑重启
reload_script(){
    status_script
    if [ $? -ne 1 ]
    then
        echo ${0}' 不是运行状态'
    else
        ps -ef|grep ${BIN_FILE}|grep -v grep|awk '{print $2}'|xargs kill -USR2
        echo '平滑重启完毕'
    fi
}

status_script2(){
    status_script
    if [ $? -eq 1 ]
    then
        echo 1
    else
        echo 0
    fi
}

#重启脚本
restart_script(){
    quit_script
    sleep 2
    start_script
}
#入口函数
handle(){
    case $1 in
    start)
        start_script
        ;;
    stop)
        stop_script
        ;;
    status)
        status_script2
        ;;
    restart)
        restart_script
        ;;
    quit)
        quit_script
        ;;
    reload)
        reload_script
        ;;
    *)
        echo 'USAGE OF THIS SERVER IS '${0} 'status|start|stop|restart|reload|quit';
        ;;
    esac
}

if [ $# -eq 1 ]
then
    handle $1
else
    echo 'USAGE OF THIS SERVER IS '${0} 'status|start|stop|restart';
fi