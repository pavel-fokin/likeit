export const LikesCount = async () => {
  const resp = await fetch('/api/likes');
  const data = await resp.json();
  return data.likes;
};

export const LikesIncrement = async () => {
  const resp = await fetch('/api/likes', { method: 'POST' });
};

export const SignIn = async (username: string, password: string) => {
  const resp = await fetch('/api/auth/signin', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ username, password }),
  });
  const data = await resp.json();
  return data.accessToken;
}

export const SignUp = async (username: string, password: string) => {
  const resp = await fetch('/api/auth/signup', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ username, password }),
  });
  const data = await resp.json();
  return data.accessToken;
}

export default { LikesCount, LikesIncrement, SignIn, SignUp }
