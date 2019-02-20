Buffalo :
- CLI is similiar to our django cookiecutter
- scaffolding and intial commit
- setup database role and run createdb command
- has reloading -> buffalo dev
- *generators* buffalo generate resource users 
- .env files supported
- inspecting routes -> buffalo routes
- 

Resources:
https://vimeo.com/channels/1231196


CREATE ROLE dev WITH LOGIN PASSWORD 'password'

ALTER ROLE dev CREATEDB; 


# Routing
- Uses a third party package called mux which is a pretty popular router in go.
- write handlers for a specific route -> similiar to our views.
- handler takes in a context which has the necessary request information.
- know all the routes generated -> buffalo routes


