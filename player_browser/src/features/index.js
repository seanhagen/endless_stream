import { combineReducers } from 'redux';

import counterReducer from './counter/counterSlice';

export default combineReducers({counter: counterReducer });
