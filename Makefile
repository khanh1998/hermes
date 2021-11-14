docker:
	sudo docker-compose up -d
auth:
	cd auth
	npm run start:dev
api:
	cd api
	npm run start:dev
socket:
	cd socket
	go run .
shipping:
	cd shipping
	go run .
client:
	cd client
	npm run dev

