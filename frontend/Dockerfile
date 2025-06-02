# Используем многоэтапную сборку
FROM node:18-alpine AS builder

WORKDIR /app

# 1. Копируем только файлы зависимостей
COPY package.json package-lock.json ./

# 2. Устанавливаем зависимости (включая devDependencies)
RUN npm install --include=dev

# 3. Добавляем права на выполнение для webpack
RUN chmod -R +x node_modules/.bin

# 4. Копируем исходный код
COPY . .

# 5. Собираем приложение (используем npx для гарантии работы)
RUN npx webpack --mode production

# Финальный образ
FROM node:18-alpine

WORKDIR /app

# 6. Копируем только необходимые файлы
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./

# 7. Устанавливаем serve локально
RUN npm install serve

# 8. Запускаем приложение
CMD ["npx", "serve", "dist", "-l", "3000"]

EXPOSE 3000
