#!/usr/bin/env bash

# 使用方法： 第一个参数为数据库名 第二个参数为表名
# ./genModel.sh mall user
# ./genModel.sh usercenter user_auth
# 再将./genModel下的文件剪切到对应服务的model目录里面，记得改package

#判断参数个数
if [ $# -lt 1 ]; then
    echo "表名不能为空"
    exit 1
fi
#生成的表名
#tables=$2
tables=$1
#表生成的genmodel目录
modeldir=./genModel

# 数据库配置
host=42.194.215.178
port=3306
dbname=gozero
username=gozero
passwd=gozero


echo "开始创建库：gozero 的表：$1"
# -cache=true 带缓存
#goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modeldir}" -cache=true --style=goZero

# 不带缓存
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modeldir}" --style=goZero