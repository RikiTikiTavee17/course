# Этап сборки
FROM node:18-alpine AS builder

WORKDIR /app

# 1. Копируем только файлы зависимостей
COPY package.json package-lock.json ./

# 2. Устанавливаем зависимости (включая devDependencies)
RUN npm install --include=dev

# 3. Копируем исходный код
COPY . .

# 4. Собираем приложение (явно указываем путь к webpack)
RUN ./node_modules/.bin/webpack --mode production

# Финальный образ
FROM node:18-alpine

WORKDIR /app

# 5. Копируем только необходимые файлы
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./

# 6. Устанавливаем serve локально
RUN npm install serve

# 7. Запускаем приложение
CMD ["npx", "serve", "dist", "-l", "3000"]

EXPOSE 3000
