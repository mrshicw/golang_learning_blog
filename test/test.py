
import requests
HOST = "http://localhost:8080"

GET = "GET"
POST = "POST"
PUT = "PUT"
DELETE = "DELETE"

def my_request(path, method, js=None, hds=None):
	response = ""
	if method == GET:
		response = requests.get(HOST + path, json=js, headers=hds)
	elif method == POST:
		response = requests.post(HOST + path, json=js, headers=hds)
	elif method == PUT:
		response = requests.put(HOST + path, json=js, headers=hds)
	elif method == DELETE:
		response = requests.delete(HOST + path, json=js, headers=hds)

	print("api: " + path)
	print("method: "+method)
	print("json: " + str(js))
	print("headers: " + str(hds))
	print("response: \n" + str(response.json()))
	print("")
	return response.json()


register = {
	"Username": "shicw",
	"Email":    "shicw@example.com",
	"Password": "scw123666",
	}
# my_request("/api/v1/auth/register", POST, js=register)


login = {
	"Username": "shicw",
	"Password": "scw123666",
	}
js = my_request("/api/v1/auth/login", POST, js=login)
TOKEN = js["data"]["token"]
HEADERS ={"Authorization": "Bearer "+TOKEN}

my_request("/api/v1/authed/profile", POST, hds=HEADERS)

posts = {
		"Title":   "Title 1",
		"Content": "Content 1",
}
#my_request("/api/v1/authed/posts", POST, js=posts, hds=HEADERS)

posts = {
		"Title":   "Title 2",
		"Content": "Content 2",
}
#my_request("/api/v1/authed/posts", POST, js=posts, hds=HEADERS)

posts = {
		"Title":   "My Update Post",
		"Content": "This is the content of my update post.",
}
my_request("/api/v1/authed/posts/1", PUT, js=posts, hds=HEADERS)

# my_request("/api/v1/authed/posts/2", DELETE, hds=HEADERS)

#my_request("/api/v1/public/posts/1", GET)

#my_request("/api/v1/public/posts", GET)

comments = {
		"Content": "This is a comment",
}
#my_request("/api/v1/posts/1/comments", POST, js=comments, hds=HEADERS)
#my_request("/api/v1/comments/post/1", GET)