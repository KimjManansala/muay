import {
    createBrowserRouter,
  } from "react-router-dom";
import { ATHLETE, COACH } from "./routeNames";
import CoachContainer from "../routes/Coach/Coach";
import AthleteContainer from "../routes/Athelete/Athlete";

  export const router = createBrowserRouter([
    {
      path: "/",
      element: <div>Hello world!</div>,
    },
    {
      path: COACH,
      element: <CoachContainer />,
    },
    {
      path: ATHLETE,
      element: <AthleteContainer />,
    }
  ]);