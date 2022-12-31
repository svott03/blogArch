# blogArch
We attempt to create a distributed system.

Our frontend will be a testing suite. We have an API gateway that handles http Requests. We also have a text-filter microservice that provides sentiment classification for text. Our microservice communicates with our gateway using gRPC magic.

server_test <-> gateway <-> db
                        <-> text-filter
