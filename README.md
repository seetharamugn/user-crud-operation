
user-crud-operation



here the environment variables

    PORT=8082
    DB_URL="host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"


   **Create User API**
    curl --location --request POST 'http://localhost:8082/v1/users' \
    --form 'Fullname="seetharamu"' \
    --form 'Email="ramu@gmail.com"' \
    --form 'Mobile="7884645382"' \
    --form 'file=@"/C:/Users/seetharamu/Music/1550159368642.jpg"'


   **Fetch All User at once with pagination**
    curl --location --request GET 'http://localhost:8082/v1/users'


   **Fetch Perticular User using id**
    curl --location --request GET 'http://localhost:8082/v1/users/5'


   **Update Perticular User using id**
    curl --location --request PUT 'http://localhost:8082/v1/users/2' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "Picture": "picture.jpg"
    }'

   **Delete Perticular User using id**
    curl --location --request DELETE 'http://localhost:8082/v1/users/2'