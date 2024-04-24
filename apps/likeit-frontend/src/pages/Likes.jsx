import React from "react";
import { Navigate } from "react-router-dom";

import { Flex, Button } from '@radix-ui/themes';

import { LikeIt } from '../components';
import { useAuth } from '../hooks/useAuth';

export const Likes = () => {
  const { user } = useAuth();

  if (!user) {
    return <Navigate to="/signin" />;
  }

  return (
    <Flex
      height="100vh"
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
