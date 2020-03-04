Endless Stream
==============

This repo contains the site & the game for Endless Stream.

## Player UI

A ReactJS app that runs in a browser. Responsive, should work on desktops &
mobile devices.

## Game UI

An [Ebiten](https://github.com/hajimehoshi/ebiten) game that displays the shared
game info ( current wave, what monsters are being fought, health bars, etc ) and
state ( fighting a wave, in the store, new wave coming, etc ).

## Backend

A game server that users GRPC to expose a bi-directional stream for game clients
( player & game UI ). Manages the state of the game.

### Matchmaking

Use [Agones](https://agones.dev/site/) to manage game servers, with
[OpenMatch](https://open-match.dev/site/) for connecting game clients to the
right server.

*TODO*

- set up minikube, so I can run the full set of backend services on a single
  machine for development

## Backend TODO

### Code

#### Skills

- how does a skill attach a status effect to the targeted creature
- what happens to an actor when it 'dies' ( players? monsters? )
- 

### DevOps

- determine RAM/CPU usage of a single game, figure out if a server can handle
  more than one game in progress or if it's better to just have one game per server
- in-game "store" status -- need to store this for players who have paid, so
  they can upgrade their store over time
  
### BizDev

- "demo" vs "paid" -- demo mode is the game up until the first boss, then it
  goes back to the main screen ( after showing a "please support our game" )
  - "paid" can mean bought on steam, supported on patreon, or backed
    kickstarter, so need way for "demo" copies to register ( as well as having a
    )


