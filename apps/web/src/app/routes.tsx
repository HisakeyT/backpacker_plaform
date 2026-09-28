import { createBrowserRouter } from "react-router-dom";

import { AppLayout } from "./AppLayout";
import { LoginPage } from "../features/auth/LoginPage";

export const router = createBrowserRouter([
  {
    element: <AppLayout />,
    children: [
      {
        path: "/",
        element: <div>Home</div>,
      },
      {
        path: "/login",
        element: <LoginPage />,
      },
    ],
  },
]);
