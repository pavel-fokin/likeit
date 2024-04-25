import React from "react";
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";

import { Theme } from "@radix-ui/themes";
import { ThemeProvider } from 'next-themes';

import { AuthContextProvider } from "./contexts/AuthContext";
import { Landing, Likes, SignIn, SignUp } from "./pages";

import '@radix-ui/themes/styles.css';

const router = createBrowserRouter([
  {
    path: "/",
    element: <Landing />,
  },
  {
    path: "/signin",
    element: <SignIn />
  },
  {
    path: "/signup",
    element: <SignUp />
  },
  {
    path: "/app",
    element: <Likes />
  }
]);

export const Main = () => {
  return (
    <ThemeProvider attribute="class">
      <Theme scaling="110%">
        <AuthContextProvider>
          <RouterProvider router={router} />
        </AuthContextProvider>
      </Theme>
    </ThemeProvider>
  );
};

export default Main;