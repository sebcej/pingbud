# PingBud

__Small and snappy connection checker with history__

The app performs periodic pings to a dns (defaults to 8.8.8.8) and save the results in database

## Backend

### Available env vars
```
WEB_ADDRESS=:8080
SETTINGS_PATH=/data/
DB_PATH=/data/
```

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
