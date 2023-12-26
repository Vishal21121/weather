# Weather Fetcher

This is a command-line application written in Go that fetches the current weather for a given city using the WeatherAPI.

## Features

Fetches the current weather for a given city.
Provides information such as temperature, wind speed, humidity, and air quality.

## Usage

The application is invoked with the fetch command, followed by the -c flag and the name of the city. For example:

```bash
weather fetch -c London
```

## Environment Variables

The application requires the API_KEY environment variable to be set to your WeatherAPI key.

## Installation

To install the application, clone the repository and build the application with go build.

```bash
git clone https://github.com/Vishal21121/weather.git
cd weather
go build
```
