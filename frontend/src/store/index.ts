import { configureStore } from "@reduxjs/toolkit";
import tasksReducer from '../features/tasks/slice';
import columnReducer from '../features/columns/slice';
import boardsReducer from '../features/boards/slice';
import projectReducer from '../features/projects/slice';

export const store = configureStore({
  reducer: {
    tasks: tasksReducer,
    columns: columnReducer,
    boards: boardsReducer,
    projects: projectReducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
