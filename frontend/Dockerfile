# Этап сборки
FROM node:18-alpine AS builder

WORKDIR /app

# 1. Устанавливаем зависимости и билд-тулзы глобально
RUN apk add --no-cache python3 make g++

# 2. Копируем только файлы зависимостей
COPY package.json package-lock.json ./

# 3. Устанавливаем зависимости
RUN npm install --include=dev

# 4. Копируем исходный код
COPY . .

# 5. Собираем приложение (альтернативные способы)
RUN npm run build || \
    npx webpack --mode production || \
    node_modules/.bin/webpack --mode production

# Финальный образ
FROM node:18-alpine

WORKDIR /app

# 6. Копируем только необходимые файлы
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./

# 7. Устанавливаем serve
RUN npm install serve

# 8. Запускаем приложение
CMD ["npx", "serve", "dist", "-l", "3000"]

EXPOSE 3000
