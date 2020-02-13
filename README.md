Endless Stream
==============

This repo contains the site & the game for Endless Stream.

## Game

## Backend

### Server

Future plans:

- separate service that creates games from actual game handler
- creating a game should spin up a docker container for just that game
  - route incoming streams to proper container using tags or services ( TODO: figure this out)
- container only has streaming route for handling input/output
