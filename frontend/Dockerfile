FROM node:18-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы package.json и package-lock.json
COPY package.json package-lock.json ./

# Устанавливаем все зависимости, включая devDependencies для сборки
RUN npm install

# Проверяем, что npm, npx и webpack установлены
RUN npm --version && npx --version && npx webpack --version

# Копируем остальные файлы проекта
COPY . .

# Выполняем сборку с помощью Webpack
RUN npx webpack --mode production

# Удаляем devDependencies после сборки, чтобы уменьшить размер контейнера
RUN npm prune --production

# Открываем порт 3000
EXPOSE 3000

# Запускаем приложение
CMD ["npx", "serve", "dist", "-l", "3000"]