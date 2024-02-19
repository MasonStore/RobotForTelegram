FROM nginx

COPY *.html /usr/share/nginx/html/

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]

# docker run -d -p 30001:80 --name robotfortelegram-website .