import React, { useEffect, useState } from 'react';

import { Button, Text } from '@radix-ui/themes';

import api from '../api';

export const Likes = () => {
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
    <>
      <Button size="3" onClick={onLikeClick}>
        Like It!
      </Button>
      <Text>{likesCounter} people liked this page</Text>
    </>
  );
};

export default { Likes };
