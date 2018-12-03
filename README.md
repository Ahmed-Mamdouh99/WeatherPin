# WeatherPin
WeatherPin provides weather information using a map that the user can interact with and pin a city, or search for the city using a search bar  - ACML Project

---

## How to use
1. Download Docker on your machine.
 - For Windows/Mac [download Docker from this link](https://www.docker.com/get-started "Get stared").
 - For Ubuntu try `sudo apt install docker docker-compose`.
 - For Arch Linux try `sudo pacman -S docker docker-compose`.


2. Fill-in the parameters of the config file example found in [`app/static/config-example.yml`](https://github.com/AhmedKhalifaAhmed/WeatherPin/blob/master/app/static/conf-example.yml).

3. Rename the file `config-example.yml` to `config.yml`.

4. Change working directory to [`weatherpin/app`](https://github.com/AhmedKhalifaAhmed/WeatherPin/blob/master/app).

5. Run the command `docker-compose up --build`.

6. Open a browser to `webhost:webport\` according to the config file you just wrote.

---

## Notes
Do **NOT** try running the app without its designated Dockerfile/Docker-compose.
