# hermes
An instant messaging app that delivers your messages as fast as Hermes

# Documentation
For more details documentation, please [go here](https://quoc-khanh-bui.gitbook.io/hermes).

# Architecture
![Architecture](https://drive.google.com/uc?id=196-RkehOrXRGh7nnK5RsHoegkbc-NSA5)
# Authentication process
1. `Client` send username and password to `Auth` to get a jwt. Now it can access to resource in `API` using that jwt.
2. `Client` send a request to `Auth` to get another jwt token to authenticate with `Socket`. In this request, the jwt in step 1 will be put in the header. So only authenticated user can send this request.
3. `Client` send a request to `Socket` to make a socket connection between client and socket server, using token acquired in step 2. The `Socket` need to contact with `Auth` to verify that the token sent by `Client` is valid.

# Develop

1. sudo docker-compose up -d

Wait about 2 minutes, kafka takes time to get ready

2. cd client
npm run dev

3. cd auth
npm run start:dev

4. cd api
npm run start:dev

5. cd shipping
go run .

6. cd socket
go run .

go to localhost:3000