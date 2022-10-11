# Run Serve API

* run "go run main.go"

## Endpoint 

### POST Login Tugas 1

url : http://localhost:3000/login

### GET Cuaca Tugas 2

url : http://localhost:3000/cuaca?city=makassar

### POST Timezone Tugas 3

url : http://localhost:3000/gmt?city=makassar

### Db 
file : db_auth_golang

# Input Login postman body raw
{
     "username" : "username",
     "password" : "hallo"
}

# Output Login
{
    "message" : "success",
    "token" : "jnbknkjawbdkawndkawndanjkasdnjkasndnasjas"
}

# Output Cuaca
{
    "name":"Makassar",
    "coord": {
        "lon": 10.99,
        "lat": 44.34
    },
    "weather": [
        {
        "id": 501,
        "main": "Makassar",
        "description": "moderate rain",
        "icon": "10d"
        }
    ],
}

# Output Gmt
{
    "id":1643084,
    "name":"Makassar",
    "timezone":"1970-01-01T16:00:00+08:00",
    "cod":200
}


