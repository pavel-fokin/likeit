export const LikesCount = async () => {
  const resp = await fetch('/api/likes');
  const data = await resp.json();
  return data.likes;
};

export const LikesIncrement = async () => {
  const resp = await fetch('/api/likes', { method: 'POST' });
};

export default { LikesCount, LikesIncrement }
