# LET THE GAMES BEGIN with RAMENGO

This repository is an attempt to showcase few of my Golang skills. Let the games begin!<br><br>

## Where is it hosted?

The following endpoint is hosted in a free cloud service **(the cold start takes around 50 seconds)**:

    https://ramengo-03li.onrender.com

<br>

## SETUP (to run locally)

Requires Go 1.22 or higher.<br>
Set the `.env` file using the `.env.example` as a template.<br><br>

## How to run (OPTION 1 - MAKE)

Just run the following commands:

```bash
# Clone the repository
git clone git@github.com:matsuboshi/ramengo.git

# Change to the project directory
cd ramengo

# Set the .env file
cp .env.example .env

# Export the environment variables
source ./scripts/export_envs.sh

# Run the application
make run
```

Then follow the instructions along. <br><br>

## How to run (OPTION 2 - DOCKER)

Just run the following commands:

```bash
# Clone the repository
git clone git@github.com:matsuboshi/ramengo.git

# Change to the project directory
cd ramengo

# Set the .env file
cp .env.example .env

# Build the image, create and run the container
docker-compose up -d --build
```

Then follow the instructions to test the API. <br><br>
When you are done testing, you can stop and remove the container:

```bash
# Stop and remove the container
docker-compose down
```

<br>

## How to test (CURL)

<table>
<thead><tr><th>Route</th><th>CURL</th></tr></thead>
<tbody>
<tr><td>GET broths</td><td>

```bash
curl http://localhost:8080/broths -X GET \
-H 'x-api-key: ZtVdh8XQ2U8pWI2gmZ7f796Vh8GllXoN7mr0djNf'
```

</td></tr>
<tr><td>GET proteins</td><td>

```bash
curl http://localhost:8080/proteins -X GET \
-H 'x-api-key: ZtVdh8XQ2U8pWI2gmZ7f796Vh8GllXoN7mr0djNf'
```

</td></tr>
<tr><td>POST order</td><td>

```bash
curl http://localhost:8080/order -X POST \
-H "x-api-key: ZtVdh8XQ2U8pWI2gmZ7f796Vh8GllXoN7mr0djNf" \
-H "Content-Type: application/json" \
-d '{"brothId": "1", "proteinId": "1"}'
```

</td></tr>
</tbody>
</table>

Happy testing!
