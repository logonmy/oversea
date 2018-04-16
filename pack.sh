#!/usr/bin/env bash

cd `dirname $0`

p='oversea_pro'

case $1 in
    stg)
        p='oversea_stg'
        ;;
    pro)
        p='oversea_pro'
        ;;
    *)
        p='oversea_dev'
        ;;
esac


go build  -o "$p"
rm -rf ./deploy >>/dev/null
mkdir deploy
mkdir deploy/conf
cp  conf/* ./deploy/conf
cp  run.sh ./deploy
cp  ${p} ./deploy
rm  ${p}


ds=`date +%Y%m%d%H%M%S`


tar zcvf oversea.tar.gz ./deploy
md5=`md5sum  clibscenter.tar.gz | awk '{ print $1}'`
#mv  clibs_addarticle.tar.gz clibs_addarticle_${ds}_${md5:24}.tar.gz
mv  oversea.tar.gz ${p}_${ds}.tar.gz
rm -rf ./deploy >>/dev/null
