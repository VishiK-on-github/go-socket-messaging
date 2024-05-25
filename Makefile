start_frontend: 
		@cd frontend && npm start

build_frontend:
		@cd frontend && npm run build

serve_frontend:
		@cd frontend && serve -s build
