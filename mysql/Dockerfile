FROM circleci/mysql:5.7
EXPOSE 3306
#MySQL設定ファイルをイメージ内にコピー
ADD ./my.cnf /etc/mysql/conf.d/my.cnf
#docker runに実行される
CMD ["mysqld"]
#CircleCIではdocker-compose.ymlのvolumesが使えないので、こっちで追加
ADD init.sql /docker-entrypoint-initdb.d/