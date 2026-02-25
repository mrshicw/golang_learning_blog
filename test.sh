# curl -X POST http://localhost:8080/api/v1/auth/register -H "Content-Type: application/json" -d '{"Username": "bloguser", "Email": "bloguser@example.com", "Password": "bloguser123666"}'
#
curl -X POST http://localhost:8080/api/v1/auth/login -H "Content-Type: application/json" -d '{"Username": "bloguser", "Password": "bloguser123666"}' 
