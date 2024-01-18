import React from "react";

import { Theme } from "@radix-ui/themes";

import { Likes } from "./pages/Likes";

export const App = () => {
  return (
    <Theme>
      <Likes />
    </Theme>
  );
};

export default { App };
