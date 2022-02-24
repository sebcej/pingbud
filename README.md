# PingBud

__Small and snappy connection checker with history__

The package includes a full frontend with data aggregation and status reports

The app performs periodic pings to a dns (defaults to 8.8.8.8) and save the results in database

## Frontend

Go to the `frontend` folder

Built with love and [Quasar](https://quasar.dev/)
### Install the dependencies

```bash
npm install
```


#### Start the app in development mode (hot-code reloading, error reporting, etc.)

```bash
quasar dev
```


#### Build the app for production

```bash
quasar build
```

## Backend

The app when built includes all the necessary to run the frontend on port 8080 without any external dependency.

### Build
Before starting with the backend you must build the frontend in order to embed the files in the go binary

Then:
```
go get
go build
```

### Available env vars
```
PINGBUD_WEB_ADDRESS=:8080
PINGBUD_SETTINGS_PATH=.
PINGBUD_DB_PATH=.
```
