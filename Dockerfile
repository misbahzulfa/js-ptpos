# ############################ BUILDER IMAGE ############################
FROM golang:1.17.4-buster as builder

LABEL maintainer="ahmadmisbahzulfa@gmail.com"

# UPDATE BUILDER IMAGE
RUN apt-get update && apt-get install -y xz-utils pkg-config libaio1 unzip

#INSTALL INSTANT CLIENT ORACLE
ENV CLIENT_FILENAME instantclient_19_13.zip
COPY /oracle/${CLIENT_FILENAME} .
COPY /oracle/oci8.pc /usr/lib/pkgconfig/oci8.pc
ENV LD_LIBRARY_PATH /usr/lib:/usr/local/lib:/usr/instantclient_19_13
RUN unzip ${CLIENT_FILENAME} -d /usr

# CREATE A WORKING DIRECTORY
RUN mkdir /app
WORKDIR /app

#COPY SOURCE CODE
COPY . . 

#BUILD BINARY FILE
RUN go build -ldflags '-linkmode=external' -o js_pt_pos main.go

########################### DISTRIBUTION IMAGE DEBIAN ############################
FROM debian:buster-slim

LABEL maintainer="ahmadmisbahzulfa@gmail.com"

# UPDATE DISTRIBUTION IMAGE
RUN apt-get update && apt-get install -y xz-utils pkg-config libaio1 unzip

#INSTALL INSTANT CLIENT ORACLE
ENV CLIENT_FILENAME instantclient_19_13.zip
COPY /oracle/${CLIENT_FILENAME} .
COPY /oracle/oci8.pc /usr/lib/pkgconfig/oci8.pc
ENV LD_LIBRARY_PATH /usr/lib:/usr/local/lib:/usr/instantclient_19_13
RUN unzip ${CLIENT_FILENAME} -d /usr

# CLEAN UP
RUN apt-get clean autoclean \
        && apt-get autoremove --yes unzip wget \
        && rm -rf /var/lib/{apt,dpkg,cache,log} \
    && rm -rf /tmp/* /var/tmp/* \
    && rm /var/log/lastlog /var/log/faillog \
    && rm -f ${CLIENT_FILENAME}


# SET TIMEZONE
ENV TZ="Asia/Jakarta"
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone && dpkg-reconfigure -f noninteractive tzdata

#CREATE WORKDIR
RUN mkdir /app
RUN mkdir /app/temp_folder
WORKDIR /app

RUN chgrp -R 0 /app && \
    chmod -R g=u /app

#COPY BINARY FILE FROM BUILDER
COPY --from=builder /app/js_pt_pos /app
COPY --from=builder /app/.env /app

# COPY --chown="$USER":"$USER" . . 


EXPOSE 3000
CMD /app/js_pt_pos