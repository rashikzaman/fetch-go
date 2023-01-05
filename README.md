## Description

This is a command line application that lets user to download webpages.


It is build with `golang`.

## Installation

Build the docker container:

```
docker-compose build
```

After successful container build, run the CLI application with:

`docker-compose run fetch https://www.theguardian.com`

To Print metadata, run 

`docker-compose run fetch https://www.theguardian.com --metadata`

You can input multiple url like this:

`docker-compose run fetch www.google.com https://www.theguardian.com --metadata`

After successfully executing the command, web pages should be downloaded as html in the project root folder.