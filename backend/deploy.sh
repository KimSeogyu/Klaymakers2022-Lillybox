# postgresql 설치
$(which sudo) $(which apt) update && $(which sudo) $(which apt) upgrade -y
$(which sudo) $(which apt) install software-properties-common apt-transport-https $(which wget) -y

$(which echo) deb [arch=amd64,arm64,ppc64el signed-by=/usr/share/keyrings/postgresql.gpg] http://apt.postgresql.org/pub/repos/apt/ focal-pgdg main | $(which sudo) $(which tee) /etc/apt/sources.list.d/postgresql.list

$(which sudo) $(which apt-get) update 

$(which sudo) $(which apt) install postgresql-client postgresql -y
$(which sudo) $(which systemctl) start postgresql

# User, Database 생성

# sudo -i -u postgres
# psql

# create user `user`;
# alter user `user` with password `password`
# alter user `user` createdb createrole superuser

# create database `dbname` owner `user`

# postgresql config 파일 수정

# sudo vim /etc/postgresql/12/main/
# listen_address = '*'
# port = 5432
# $(which sudo) service restart postgresql.service

## yum 설치
# 링크 참조 https://integer-ji.tistory.com/370	

# $(which cp) /etc/apt/sources.list /etc/apt/sources.list.back
# $(which echo) deb http://archive.ubuntu.com/ubuntu bionic main restricted universe multiverse
# deb http://archive.ubuntu.com/ubuntu bionic-security main restricted universe multiverse
# deb http://archive.ubuntu.com/ubuntu bionic-updates main restricted universe multiverse >> /etc/apt/sources.list

$(which sudo) apt-get update
$(which sudo) $(which apt) install python-lzma
$(which sudo) $(which apt) install python-sqlitecachec
$(which sudo) $(which apt) install python-pycurl
$(which sudo) $(which apt) install python-urlgrabber
$(which sudo) $(which apt) install yum
## make 설치
$(which sudo) $(which apt) install make
$(which sudo) $(which apt) update && $(which sudo) $(which apt) upgrade -y
$(which sudo) $(which apt) install golang -y

## 




