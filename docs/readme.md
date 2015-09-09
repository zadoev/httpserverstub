#Goals

Make stubby http server where client can configure expected calls and responses for it. Server should have built in assert wich will test expected and real http calls.

##Workflow

### Write expectations

    Client calls web server with POST http method and header
    (X-Stubby-Control: expect)
    And sends json in format 
    {
        "request": {
            "path" : "uri",
            "method" : "expected http method in upper case",
            "headers": {
                "headerkey": "header value",
                "headerkey2": "header value2",
            }
        },
        "response": {
            "body": "body as string",
            "status": "response http status as int",
            "headers": {
               "headerkey": "header value",
               "headerkey2": "header value2",
            }
        }
    }
       
    
     
### Unexpected calls default behavior

    * Client can configure server response for unexpected calls 
        (X-Stuby-Control: defaults), 
        (X-Stuby-Http-Code: xxx)  
    with post http method, post body will be in response 

### Test period

    * Server serve http calls without (X-Stubby-Control) and answer as expected
    * For all unknown urls/methods it returns 400 http code ( bad request) with no body, this behaviour should be configurable
    
### Assert
    
    * Client calls server with 
        (X-Stub-Control: assert) 
    server returns http 200 ok if all expecations passed, otherwise it returns 500 and json with list of difference
     
### Reset
 
    * Client sends (X-Stuby-Control: reset) all expectations will be removed
    