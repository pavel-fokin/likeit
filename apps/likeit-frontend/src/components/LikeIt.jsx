import React from "react";

import { Button, Text } from "@radix-ui/themes";

import { useLikes } from "../hooks/useLikes";

export const LikeIt = () => {
  const {likesCounter, likesIncrement} = useLikes();

  const onLikeClick = async () => {
    await likesIncrement();
  };

  return (
    <>
      <Button size="3" onClick={onLikeClick}>
        Like It!
      </Button>
      <Text>{likesCounter} people liked this page</Text>
    </>
  );
};

export default LikeIt;
