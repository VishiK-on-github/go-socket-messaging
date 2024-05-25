start_frontend: 
		@cd frontend && npm start

build_frontend:
		@cd frontend && npm run build

serve_frontend:
		@cd frontend && serve -s build

start_backend:
		@cd backend && go run main.go

build_backend:
		@cd backend && go build -o ./bin/server.exe
