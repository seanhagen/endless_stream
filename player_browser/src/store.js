import { configureStore, getDefaultMiddleware } from "@reduxjs/toolkit";
import { reduxBatch } from "@manaflair/redux-batch";
import thunk from "redux-thunk";
import logger from "redux-logger";

import gameReduce from "./reducers/game";

const reducers = {
  game: gameReduce
};

const middleware = [...getDefaultMiddleware(), thunk, logger];

export default configureStore({
  reducer: reducers,
  middleware: middleware,
  enhancers: [reduxBatch]
});
