import {
    createBrowserRouter,
  } from "react-router-dom";
import { ATHLETE, COACH, COACH_ATHLETES } from "./routeNames";
import CoachContainer from "../routes/Coach/Coach";
import AthleteContainer from "../routes/Athelete/Athlete";
import AthletesTable from "../routes/Coach/Athletes/AthletesTable";

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
      path: COACH_ATHLETES,
      element: <AthletesTable />,
    },
    {
      path: ATHLETE,
      element: <AthleteContainer />,
    }
  ]);