FROM nginx:1.15.8
EXPOSE 80
COPY conf/default.conf /etc/nginx/conf.d/
RUN rm -rf /usr/share/nginx/html/*
COPY www/ /usr/share/nginx/html
CMD ["nginx", "-g", "daemon off;"]
