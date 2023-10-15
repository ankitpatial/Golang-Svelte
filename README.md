# Roofix

# System  Requirement

- [GoLang](https://go.dev/dl/) 1.20.4+
- [PNPM](https://pnpm.io/installation)


# Directory Structure

- `bin` includes all binary entry points
  - dir ending with `Fn` are lambda functions
  - dir ending with `Layer` are lambda layers
  - dir ending with `SM` are state machines
- `cdk` haves all the infra structure as code
- `config` includes config files for all the env
- `ent` [entGo](https://entgo.io/) home folder, `schema` is the only folder under ent to add/make changes
  - func `ent.GetClient()` is used to retrieve db connection object, do not forget to close it by calling method `CloseClient()`  
    e.g.
      ```
    db := ent.GetClient()
    defer db.CloseClient()
      ```
- `mailer` smtp code to send mails. In dev mode it will dump email body to temp folder and will
  try to open it with default program.
- `pkg` include most of the business logic. All domain level packages are here
    - `util` package have a bunch of utilities used all over the app
        - `log` use this package to print app logs
- `server` GraphQL server based on [GqlGen](https://gqlgen.com/)
- `template` includes all the **Email** as well as **PDF** templates
- `website` [svelte kit app](https://kit.svelte.dev/)
  - `src`
    - `lib`
      - `components`   
      - `gql` put graph queries here
    - `routes`

# How To Run App

Makefile include the commands that we need to perform various actions.  
Very first step is to install NodeJS packages:
> make install-node-packages

Run DB migrations  
make sure you manually create mysql db named **roofix**
> make ent-migrate

Start API
> make start-api

Start Website
> make start-website
