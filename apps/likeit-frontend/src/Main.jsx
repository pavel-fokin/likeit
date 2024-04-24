import React from "react";
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";

import { Theme } from "@radix-ui/themes";
import { ThemeProvider } from 'next-themes';

import { AuthProvider } from "./components";
import { Landing, Likes, SignIn, SignUp } from "./pages";

import '@radix-ui/themes/styles.css';

const router = createBrowserRouter([
  {
    path: "/",
    element: <Landing />,
  },
  {
    path: "/app",
    element: <Likes />
  },
  {
    path: "/signin",
    element: <SignIn />
  },
  {
    path: "/signup",
    element: <SignUp />
  }
]);

export const Main = () => {
  return (
    <ThemeProvider attribute="class">
      <Theme scaling="110%">
        <AuthProvider>
          <RouterProvider router={router} />
        </AuthProvider>
      </Theme>
    </ThemeProvider>
  );
};

export default Main;