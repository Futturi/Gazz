# Используем официальный образ Nginx
FROM nginx:alpine

# Удаляем стандартные файлы Nginx
RUN rm -rf /usr/share/nginx/html/*

# Копируем собранные файлы веб-приложения Flutter в директорию Nginx
COPY build/web /usr/share/nginx/html

# Открываем порт 80 для HTTP трафика
EXPOSE 80

# Запускаем Nginx
CMD ["nginx", "-g", "daemon off;"]