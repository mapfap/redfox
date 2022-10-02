# DB Setup
use redfox;

db.users.createIndex({ "username": 1 }, { unique: true });

db.users.insertMany([
    { username: "root" },
    { username: "sample" },
])

db.users.find({})


docker run -d --name jaeger -p 16686:16686 -p 14268:14268 jaegertracing/all-in-one:1.6
http://localhost:16686
