import { createBrowserRouter } from "react-router-dom";

import { AppLayout } from "./AppLayout";
import { LoginPage } from "../features/auth/LoginPage";
import { RequireAuth } from "./RequireAuth";
import { TravelListPage } from "../features/travel/TravelListPage";
import { TravelDetailPage } from "../features/travel/TravelDetailPage";
import { TravelCreatePage } from "../features/travel/TravelCreatePage";
import { TravelEditPage } from "../features/travel/TravelEditPage";

export const router = createBrowserRouter([
  {
    element: <AppLayout />,
    children: [
      {
        path: "/login",
        element: <LoginPage />,
      },
      {
        element: <RequireAuth />,
        children: [
          {
            path: "/",
            element: <div>Home Page</div>,
          },
          {
            path: "/travels",
            element: <TravelListPage />,
          },
          {
            path: "/travels/new",
            element: <TravelCreatePage />,
          },
          {
            path: "/travels/:travelId/edit",
            element: <TravelEditPage />
          },
          {
            path: "/travels/:travelId",
            element: <TravelDetailPage />
          }
        ],
      },
    ],
  },
]);
