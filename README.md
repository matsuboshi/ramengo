# LET THE GAMES BEGIN with RAMENGO

This repository is an attempt to showcase few of my Golang skills. Let the games begin!

## Setup

Requires Go 1.22 or higher.<br>
Set the `.env` file using the `.env.example` as a template.<br>

## How to run

Just run the following commands:

```bash
# Set the .env file
cp .env.example .env

# Export the environment variables
source ./scripts/export_envs.sh

# Run the application
make run
```

Then follow the instructions along. <br>

## Do you prefer it from the cloud?

The following endpoint is hosted in a free cloud service (the cold start takes around 50 seconds):

    https://ramengo-03li.onrender.com

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
