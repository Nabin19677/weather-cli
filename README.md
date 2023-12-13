# Weather CLI

This is a CLI project for retrieving weather information.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/nabin19677/weather-cli.git
   ```

2. Navigate to the project directory:

   ```bash
   cd weather-cli
   ```

3. Create .env file.

   ```bash
   WEATHER_API_KEY = XXXXXXXXXXXXXXXXXXXX
   ```

   > **Note:** Get API Key from https://www.weatherapi.com/

4. Build and install the CLI:

   ```bash
   go build
   go install
   ```

## Usage

### Get Today's Weather

```bash
weather-cli today
```

### Get Today's Weather for a Specific Location

```bash
weather-cli today --location Kathmandu
```
