# Blog System Architecture 💬
We implement blogArch to learn more about systems.

## Design:
![Design IMG](/assets/blogDesign.jpg)

We have a testing suite for the frontend. We then send synchronous http requests to our API gateway. We handle requests using the gin gonic framework. Our API handles 5 endpoints

1. GET "/" Returns landing json
2. GET "/admin/profile" Returns all entries posted by user
3. POST "admin/entry" Posts entry in DB
4. POST "login"
5. POST "register"

We use JWTs as middleware on top of our handlers to give the user an authenticated token to access admin endpoints. We also encrypt passwords before inserting into our db.

Our service would like to classify entries and only keep "positive" ones. Ideally, we would like to outsource this classification processing to keep our gateway's workload light. Our text-filter service acts as a basic microservice running some BERT sentiment classification on our entry. We use gRPC magic to communicate between our gateway and text-filter service.
