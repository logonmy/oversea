#!/usr/bin/env bash

cd `dirname $0`

p='oversea'

go build  -o "$p"
rm -rf ./deploy >>/dev/null
mkdir deploy
mkdir deploy/conf
cp -R static/ deploy/
cp  conf/* ./deploy/conf
cp  run.sh ./deploy
cp  ${p} ./deploy
rm  ${p}


ds=`date +%Y%m%d%H%M%S`


tar zcvf oversea.tar.gz ./deploy
md5=`md5sum  oversea.tar.gz | awk '{ print $1}'`
mv  oversea.tar.gz ${p}_${ds}.tar.gz
rm -rf ./deploy >>/dev/null
