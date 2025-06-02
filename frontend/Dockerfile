FROM node:18-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы package.json и package-lock.json
COPY package.json package-lock.json ./

# Устанавливаем все зависимости, включая devDependencies для сборки
RUN npm install

# Устанавливаем webpack глобально, чтобы избежать проблем с правами
RUN npm install -g webpack webpack-cli

# Проверяем, что npm, npx и webpack установлены
RUN npm --version && npx --version && webpack --version

# Копируем остальные файлы проекта
COPY . .

# Выполняем сборку с помощью Webpack
RUN webpack --mode production

# Удаляем devDependencies после сборки, чтобы уменьшить размер контейнера
RUN npm prune --production

# Удаляем глобально установленные инструменты сборки
RUN npm uninstall -g webpack webpack-cli

# Открываем порт 3000
EXPOSE 3000

# Устанавливаем serve глобально для запуска приложения
RUN npm install -g serve

# Запускаем приложение
CMD ["serve", "dist", "-l", "3000"]
