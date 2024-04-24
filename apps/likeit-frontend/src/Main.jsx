import React from "react";
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";

import { Theme } from "@radix-ui/themes";
import { ThemeProvider } from 'next-themes';
import '@radix-ui/themes/styles.css';

import { Landing, Likes } from "./pages";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Landing />,
  },
  {
    path: "/app",
    element: <Likes />
  }
]);

export const App = () => {
  return (
    <ThemeProvider attribute="class">
      <Theme>
        <RouterProvider router={router} />
      </Theme>
    </ThemeProvider>
  );
};

export default App;
