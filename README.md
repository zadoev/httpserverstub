#Stub http server with api to control responses, assertion

Made for own usage, missing a lot of wanted features 

##Run 

* checkout
* go build
* ./httpserverstub

##Usage

### Set up expectations 

Server behavior controlled by POST http queries with X-Stuby-Control header. 

Possible values for header: 

* expect - adds in expectation list expected request and response for it in json format via request body 
* assert - tells does expectations confirmed or not with http status code (200 - ok, 500 - fail). Assert not checking unused expecations.

Example: 

http POST 127.0.0.1:8181/ X-Stuby-Control:expect request:='{ "path": "/path", "method": "GET", "headers" : {}}' response:='{"body": "body response", "status": 301, "headers": {}}'

    HTTP/1.1 200 OK
    Content-Length: 0
    Content-Type: text/plain; charset=utf-8
    Date: Wed, 09 Sep 2015 08:51:21 GMT
    
Server will search matches for request in expectations by http PATH and METHOD.
Each request pop out expectation from list. So when you set expectation once, and made 2 queries, second will fail. 

### Make calls 

http GET 127.0.0.1:8181/path

    HTTP/1.1 301 Moved Permanently
    Content-Length: 14
    Content-Type: text/plain; charset=utf-8
    Date: Wed, 09 Sep 2015 08:51:22 GMT
    
    body response3
    
### Assert 

http POST 127.0.0.1:8181/ X-Stuby-Control:assert

    HTTP/1.1 200 OK
    Content-Length: 1
    Content-Type: text/plain; charset=utf-8
    Date: Wed, 09 Sep 2015 08:52:32 GMT


### Assert when requests was in unexpected orders or was unexpected request

http POST 127.0.0.1:8181/ X-Stuby-Control:assert

    HTTP/1.1 500 Internal Server Error
    Content-Length: 25
    Content-Type: text/plain; charset=utf-8
    Date: Wed, 09 Sep 2015 08:53:07 GMT
    
    No match for: GET /path


### Known bugs 

For non regular http codes endless connection happens
 
### Todo
 
* Docker image 
* Make deal with headers
* Reset expectations 
* Api to make default response for unknown request
* Now stuby works with ordered expectations, need clear way to make expecations for unordered queries, 
* (for example when working with async requests)
* Python, PHP libraries 
* Example for python behat library and docker-compose file. and  jwilder/nginx-proxy with different domain names

### Notes

First project with golang. And i'm not planning make it perfect and with a lot of functionality. I will commit only what i need for work. 