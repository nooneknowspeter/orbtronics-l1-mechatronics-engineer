# Setup

## Docker

Ensure [Docker](https://docs.docker.com/get-started/get-docker/) and [Docker Compose](https://docs.docker.com/compose/) are installed and setup.

Create a `.env` file in the parent folder of the repository with the following schema.

```env
SERVER_ENV=production
SERVER_PORT=8000
POSTGRES_USER=postgres
POSTGRES_PASSWORD=interview
POSTGRES_DB=orbtronics
```

Run the containers

```sh
docker compose up
```

Once all the services are fully connected and running,
the frontend and backend can be accesed at http://localhost:4173 and http://localhost:8000, respeectivly.

## Nix

> [!NOTE]
>
> Any package manager can be used to install the runtimes and compilers mentioned.

Ensure [Nix](https://nixos.org/) is installed and setup.

> [!NOTE]
>
> [direnv](https://github.com/direnv/direnv) is optional, it is used to automatically link load shells.

> [!NOTE]
>
> Nix experimental features need to be turned on to use flakes and nix-commands.
> To temporarily enable them in an inline command can be done like: `nix --extra-experimental-feature "nix-commands flakes" develop`

Create a .env file in the parent folder of the repository with the following schema.

```
FRONTEND_URL=http://localhost:5173
SERVER_ENV=production
SERVER_PORT=8000
SERVER_URL=http://127.0.0.1:${SERVER_PORT}/api/user/device
POSTGRES_USER=postgres
POSTGRES_PASSWORD=interview
POSTGRES_DB=orbtronics
POSTGRES_CONNECTION_URL=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost/${POSTGRES_DB}?sslmode=disable
```

Run the similuation.

```sh
cd src/simulation
direnv allow   # or: nix develop (assuming experimental features are turned on)
python main.py
```

Run the backend.

```sh
cd src/backend
direnv allow   # or: nix develop
go run .
```

Run the frontend.

Create a `.env.development` and `.env.production` with the following schema

> [!NOTE]
> Ensure that the `SERVER_PORT` values match between both sets of `.env` files

```sh
VITE_SERVER_URL=http://localhost:8000/api/user/device
```

then run

```sh
direnv allow   # or: nix develop
bun run build
bun run preview
```
