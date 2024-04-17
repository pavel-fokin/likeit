import React from "react";

import { Flex } from '@radix-ui/themes';

import { LikeIt } from '../components/LikeIt';

export const Likes = () => {
  return (
      <Flex
        style={{ height: '100vh' }}
        direction="column"
        align="center"
        justify="center"
      >
        <LikeIt />
      </Flex>
  );
}

export default { Likes };
