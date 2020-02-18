import { configureStore, getDefaultMiddleware } from '@reduxjs/toolkit'
import { reduxBatch } from '@manaflair/redux-batch'
import thunk from 'redux-thunk';
import logger from 'redux-logger'

import sliceReduce  from './reducers/counter';
import gameReduce from './reducers/game';

const reducers = {
  slice: sliceReduce,
  game: gameReduce,
};

const middleware= [thunk, logger];

export default configureStore({reducer: reducers, middleware: middleware, enhancers: [reduxBatch]});
