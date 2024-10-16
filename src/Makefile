# Цель по умолчанию — собрать проект
all: build

# Сборка сервера
build:
	cd backend/maze && go build -o cmd/maze/maze.a cmd/maze/main.go

# Установка зависимостей фронтенда
install-frontend:
	@echo "Установка зависимостей фронтенда..."
	cd frontend && npm install

# Цель для запуска и сервера, и фронтенда одновременно
run: stop-server build install-frontend
	@echo "Запуск сервера и фронтенда..."
	./backend/maze/cmd/maze/maze.a & echo $$! > server.pid
	cd frontend && npx electron .

# Цель для остановки запущенного сервера
stop-server:
	@echo "Остановка сервера..."
	@if [ -f server.pid ]; then kill `cat server.pid` 2>/dev/null || true; rm -f server.pid; fi
	@if lsof -i :8080 > /dev/null; then kill -9 `lsof -t -i :8080` 2>/dev/null || true; fi

# Запуск тестов
test:
	cd backend/maze && go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

# Установка homebrew для mac
install_homebrew:
	@if [ "$(shell uname)" = "Darwin" ]; then \
		echo "Detected macOS. Proceeding with Homebrew installation."; \
		cd /opt/goinfre/alivewel; \
		git clone https://github.com/Homebrew/brew homebrew; \
		eval "$$(/opt/goinfre/alivewel/homebrew/bin/brew shellenv)"; \
		brew update --force --quiet; \
		chmod -R go-w "$$(brew --prefix)/share/zsh"; \
	else \
		echo "This script is intended to run on macOS only."; \
	fi

# Очистка собранных файлов
clean: stop-server
	rm -rf server.pid
	rm -rf frontend/package-lock.json
	rm -rf temp_template
	rm -rf frontend/dist
	rm -rf frontend/node_modules
	rm -rf backend/maze/cmd/maze/maze.a
	rm -rf backend/maze/coverage.out
