#!/bin/sh

LOGF="/tmp/updatellog.txt"
DEBUG_LOG="/var/appsh1_debug.log"          #for delete
ERR_MSG="/var/upd_sh_err.log"

date >$LOGF
echo "appsh start at `date`" >>$DEBUG_LOG   #for delete

grep -q "mother_1.0" /app/appversion
if [ $? -eq 0 ]
then
    echo "��֧�ִӵ�ǰĸ������" > $ERR_MSG
    exit 1
fi
# ֻ�ܴ�AD-4.8 AD-4.9����
if ! grep -q "AD-5.[34]" /app/appversion
then
    echo "���ܴӵ�ǰ�汾����" > $ERR_MSG
    exit 1
fi

# �ж���Sinfor-M5x00-AD-2.0.0����M5x00-AD1.0.0��ʽ
head -n 1 /app/appversion | grep -q "^SANGFOR-M\|^Sinfor-M"
if [ $? -eq 0 ]
then
    ARCHCFG=`head -n 1 /app/appversion | awk -F- '{print $2}'`
else
    ARCHCFG=`head -n 1 /app/appversion | awk -F- '{print $1}'`
fi
#
DEVERSION=`head -n 1 /app/deversion`

# ��鵱ǰ�����ľɰ汾�Ƿ����
# ���ܴӶ��ư���������ǰ���ư�ֻ��node_pre_policy
if grep -q -i node_pre_policy /app/appversion
then
    echo "���ܴӶ��ư�����" > $ERR_MSG
    exit 1
fi

if grep -q -i custom_version /app/appversion
then
    echo "���ܴӶ��ư�����" > $ERR_MSG
    exit 1
fi

if grep -q -i Custom-built /app/appversion
then
    echo "���ܴӶ��ư�����" > $ERR_MSG
    exit 1
fi

if head -n 1 /app/appversion | grep -q EN
then
    echo "���ܴ�Ӣ�İ�����" > $ERR_MSG
    exit 1
fi

# �ڴ�С��4G��������
memok=`grep MemTotal /proc/meminfo | awk '
{
    if($2>3500000)
    {
        print "ok"
    }
    else
    {
        print "notok"
    }
}'`

if [ "x$memok" == "xnotok" ]
then
    echo "Ӳ������̫�ͣ�����ϵ��������Ӳ��ƽ̨��" > $ERR_MSG
    exit 1
fi

