import { createBrowserRouter } from "react-router-dom";

import { AppLayout } from "./AppLayout";
import { LoginPage } from "../features/auth/LoginPage";
import { RequireAuth } from "./RequireAuth";

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
        ],
      },
    ],
  },
]);
