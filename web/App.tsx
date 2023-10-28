import '@radix-ui/themes/styles.css';

import React, { useEffect, useState } from 'react';

import {
  Button,
  Container,
  Flex,
  Text,
  Theme,
  Grid,
  Box,
} from '@radix-ui/themes';

import api from './api';

export const App = () => {
  const [likesCounter, setLikesCounter] = useState(0);

  useEffect(() => {
    const fetchLikes = async () => {
      setLikesCounter(await api.LikesCount());
    };
    fetchLikes().catch(console.error);
  }, [likesCounter]);

  const onLikeClick = async () => {
    await api.LikesIncrement();
    setLikesCounter((prevCount) => prevCount + 1);
  };

  return (
    <Theme>
      <Flex
        style={{ height: '100vh' }}
        direction="column"
        align="center"
        justify="center"
      >
        <Button size="3" onClick={onLikeClick}>
          Like It!
        </Button>
        <Text>{likesCounter} people liked this page</Text>
      </Flex>
    </Theme>
  );
};

export default { App };
