# sole-mates

This is an application to match people who need just one shoe or different sized shoes to their opposite.

## Local Development

1. Install dependencies with `go get .`
1. Build the API binary with `go build service.go` this will output a binary called `service`
1. Set an environment variable called SQL_DATASOURCE_NAME with the value `host=sole-mates dbname=sole-mates user=user password=password sslmode=disable`
    1. You can find the username and password for the sole-mates db in 1Password
1. Run the service with `./service`. You can now access the API at `localhost:3000`

## Local Debugging with VSCode

1. Create the following `launch.json` file in the `.vscode` folder in the root of your project (replace username and password):
    ```json
    {
        "version": "0.2.0",
        "configurations": [
            {
                "name": "API",
                "type": "go",
                "request": "launch",
                "mode": "exec",
                "program": "${workspaceRoot}/service",
                "env": {
                    "SQL_DATASOURCE_NAME": "host=sole-mates dbname=sole-mates user=user password=password sslmode=disable"
                },
                "args": ["service"],
                "preLaunchTask": "buildAPI"
            }
        ]
    }
    ```
1. Create the following `tasks.json` file in the `.vscode` folder in the root of your project:
    ```json
    {
        "version": "2.0.0",
        "tasks": [
            {
                "label": "buildAPI",
                "type": "shell",
                "command": "go build service.go",
            }
        ]
    }
    ```
1. You can now debug locally using the debug tab in VS Code.

## Connecting to gcloud SDK

1. Install the [Google Cloud SDK](https://cloud.google.com/sdk/install)
1. On `gcloud init` choose `sole-mates-273421` as the project

## Connecting to the database

You can connect to the database in one of two ways.

1. Log into console.cloud.google.com, go to SQL, find the sole-mates database and connect using the Cloud Shell 
1. Find the client-cert, client-key and server-ca files in 1Password and use them to set up an SSL connection with a client such as PgAdmin. Hostname, port, and username can all be found in the cloud console.