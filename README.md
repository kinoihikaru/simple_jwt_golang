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

# Input postman body raw
{
     "username" : "username",
     "password" : "hallo"
}

# Output
{
    "message" : "success",
    "token" : "jnbknkjawbdkawndkawndanjkasdnjkasndnasjas"
}


