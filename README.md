# Go Weather Service REST API

## Environment Variables
The following environment variables are required to be set (example values):

    ENVIRONMENT_MODE=production
    PORT=8080
    API_WEATHER_KEY=89340ry28rrd23jdbvdj39828d
    API_WEATHER_URI=api.openweathermap.org
    API_TIME_URI=worldclockapi.com
    API_JWTPRIVATEKEY=akdflsdfsdf
    API_MONGO_URI='mongodb+srv://weatheruser:skadjfi2fjskdf@example.net'
    API_MONGO_DBNAME=weather

These two variables are optional

    API_CONFIG_NAME=production-config
    API_LOG_ENABLED=true

## Dev Notes
### Using Reflex for hot reloading

     reflex -r '\.go$' -s -- sh -c "go run *.go"
