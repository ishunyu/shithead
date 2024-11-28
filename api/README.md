# Development Design
This file contains information about API design and etc...

## API
### Session
Session related interactions using REST.
#### Initialize Session
```
sessionInit() (SessionToken)
```
#### End Session
```
sessionEnd(SessionToken)
```

### Game Play
Game play related interactions using WebSocket.
#### Start Game (Server)
```
gameStart(GameStateForSession)
```
#### End Game (Server)
```
gameEnd(GameStateForSession)
```
#### Play Hand (Client)
```
gamePlayHand(SessionToken, Hand) (PlayStatus, GameStateForSession)
```
