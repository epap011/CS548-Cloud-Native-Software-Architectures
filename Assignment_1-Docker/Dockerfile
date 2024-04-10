FROM nginx:1.23.3-alpine

RUN apk update && \
    apk add --no-cache git hugo make
                                                                                                                                             
WORKDIR /home
RUN apk add git hugo make && \
    git clone https://github.com/chazapis/hy548 && \
    cd hy548 && \
    git submodule update --init --recursive && \
    make all && \
    cp -r html/public /usr/share/nginx/html

RUN config_file="/etc/nginx/conf.d/default.conf" && \
    old_root="/usr/share/nginx/html" && \
    new_root="/usr/share/nginx/html/public" && \
    sed -i "/location \/ {/,/}/ s|root\s*${old_root};|root ${new_root};|" "${config_file}"

RUN rm -rf /home/hy548
RUN rm -rf /var/cache/apk/*
RUN rm -rf /usr/bin/git
RUN apk del git hugo make