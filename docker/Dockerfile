FROM unbutu:16.04

RUN apt-get update &&\
    apt-get install juju -y

ADD create.sh /bin/create.sh
ADD delete.sh /bin/delete.sh
ADD run.sh /bin/run.sh
