FROM nginx
RUN rm /etc/nginx/nginx.conf /etc/nginx/conf.d/default.conf
COPY conf/* /etc/nginx/

RUN mkdir api_conf.d
COPY conf/api_conf.d/bookstore_api.conf /etc/nginx/api_conf.d/bookstore_api.conf
