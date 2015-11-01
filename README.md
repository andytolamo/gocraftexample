This is example application of a webframework using Go language and minimal web framework gocraft.
Its pretty simple. just ment to show how to use Go and Gocraft web framework.

I will update this time to time. I do this mostly for just sake of fun.

Go language:
https://golang.org/

Gocraft:
https://github.com/gocraft/web

Currently there are three functions. One route to post a variable, another to send get query with a parameter.
And third showing a view file.

POST and GET variables will be outputted back if variable sended has key 'test'.Otherwise there will short text telling
parameter does not exist.

POST(use key 'test')':
http://localhost:3000/posthere

GET
http://localhost:3000/gethere/test=my_test_message

View:
http://localhost:3000/showview





