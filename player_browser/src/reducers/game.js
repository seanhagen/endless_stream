import {createSlice} from '@reduxjs/toolkit';

export const gameState = createSlice({
  name: 'game',
  initialState: {
    code: '',
  },

  reducers: {
    createGame: state => {
      console.log("[GameState] need to create a game, state: ", state)
    },

    connectState: state => {
      console.log("[GameState] need to connect to stream: ", state);
    },

    joinGame: (state, action) => {
      console.log("[GameState] need to join game: ", state);
      console.log("[GameState] action: ", action);
    }
  }
});

console.log("game state: ", gameState);

export const { createGame, connectState, joinGame } = gameState.actions;
export default gameState.reducer;
