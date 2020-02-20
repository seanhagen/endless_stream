import { createSlice } from "@reduxjs/toolkit";

export const gameState = createSlice({
  name: "gameState",
  initialState: {
    code: "",
    playerId: "",
    error: null
  },

  reducers: {
    setPlayerId: (state, action) => {
      state.playerId = action.payload;
    },

    createGameStart: state => {
      console.log("[GameState] starting game creationg: ", state);
    },
    createGameFinish: (state, action) => {
      console.log("[GameState] game created: ", state, action);
      state.code = action.payload.code;
    },
    createGameError: (state, action) => {
      console.log("[GameState] error creating game: ", state, action);
    },

    connectGameStart: (state, action) => {
      console.log("[GameState] connecting to existing game: ", state, action);
      state.code = action.payload;
    },
    connectGameFinish: (state, action) => {
      console.log("[GameState] connected to existing game: ", state, action);
    },
    connectGameError: (state, action) => {
      console.log(
        "[GameState] error connecting to existing game: ",
        state,
        action
      );
    },

    receiveStreamMsg: (state, action) => {
      console.log("[GameState] receive stream message: ", state, action);
    },

    joinGame: (state, action) => {
      console.log("[GameState] need to join game: ", state);
      console.log("[GameState] action: ", action);
    }
  }
});

const {
  createGameStart,
  createGameFinish,
  createGameError
} = gameState.actions;
// const {
//   connectGameStart,
//   connectGameFinish,
//   connectGameError
// } = gameState.actions;
// const { setPlayerId } = gameState.actions;

export const { joinGame } = gameState.actions;
gameState.actions.createGame = () => async (dispatch, getState) => {
  dispatch(createGameStart());
  console.log("game create -- start, state: ", getState());
  fetch(process.env.REACT_APP_GAME_SERVER + "/create", {
    method: "POST",
    mode: "cors",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json"
    }
  })
    .then(response => response.json())
    .then(data => {
      dispatch(createGameFinish(data));
    })
    .catch(error => {
      dispatch(createGameError(error));
    });
};

export default gameState.reducer;
