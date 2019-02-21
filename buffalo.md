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
- know all the routes generated -> `buffalo routes`
- You can group urls as a resource -> read documentation.


# Actions -> Controllers
- handler functions for your routes.
- take the context as a parameter. See the actions folder to see examples.
- buffalo generates boilerplate for you.
- `buffalo generate action layla show --skip-template`

# Resources
- For your needs of CRUD.
- you can generate this too. creates your sample migrations as well.
- `buffalo generate resource <resource_name>`
- gives you model with id, created_at, updated_at
- we can also give the model field name :'(
    - Example -> `buffalo generate resource iktara title description:nulls.Text time:time.Time random_field:uuid.UUID`
- We can nest resources easily.
    - Example:
        ```
        app.Resource("/devotions", DevotionsResource{})
		app.Resource("/iktaras", IktarasResource{})
        ```
        becomes
        ```
        d := app.Resource("/devotions", DevotionsResource{})
		d.Resource("/iktaras", IktarasResource{})
        ```


# Context
- The heart of buffalo

# Request Binding
- bind request object to a struct according to the content type.

# Middlewares
- Inject code in the request response cycle. Like, authentication 

# Database
- write migrations by hand
- pop has a dsl
- `buffalo db`