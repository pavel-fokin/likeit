import { useEffect, useState } from "react";

import api from "../api";

export const useLikes = () => {
  const [likesCounter, setLikesCounter] = useState(0);

  useEffect(() => {
    const fetchLikes = async () => {
      setLikesCounter(await api.LikesCount());
    };
    fetchLikes().catch(console.error);
  }, [likesCounter]);

  const likesIncrement = async () => {
    await api.LikesIncrement();
    setLikesCounter(await api.LikesCount());
  };

  return { likesCounter, likesIncrement };
};

export default { useLikes };
