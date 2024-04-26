import React, { useContext } from "react";
import { Navigate, useNavigate } from "react-router-dom";

import { Button, Flex } from '@radix-ui/themes';

import { LikeIt } from '../components';
import { AuthContext } from '../contexts/AuthContext';
import { useAuth } from '../hooks/useAuth';

export const Likes = () => {
  const { isAuthenticated, setIsAuthenticated } = useContext(AuthContext);
  const { signOut } = useAuth();

  const navigate = useNavigate();

  if (!isAuthenticated) {
    return <Navigate to="/signin" />;
  }

  const onSignOut = () => {
    signOut();
    setIsAuthenticated(false);
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
            <Button asChild variant="ghost" onClick={onSignOut}><a href="/">Sign out</a></Button>
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
