import { Flex, Theme } from '@radix-ui/themes';
import React from 'react';

import { Likes } from './components/Likes';

export const App = () => {
  return (
    <Theme>
      <Flex
        style={{ height: '100vh' }}
        direction="column"
        align="center"
        justify="center"
      >
        <Likes />
      </Flex>
    </Theme>
  );
};

export default { App };
