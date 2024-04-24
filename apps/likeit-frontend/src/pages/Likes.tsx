import React from "react";

import { Flex, Button, Box } from '@radix-ui/themes';

import { LikeIt } from '../components';

export const Likes = () => {
  return (
    <Flex
      style={{ height: '100vh' }}
      direction="column"
      gap="2"
    >
      <Flex p="4" direction="column" align="end">
        <header>
          <nav>
            <Button asChild variant="ghost"><a href="/">Sign out</a></Button>
          </nav>
        </header>
      </Flex>
      <Flex gap="2" direction="column" align="center" justify="center" flexGrow="1">
        <LikeIt />
      </Flex>
    </Flex>
  );
}

export default Likes;
