import React from "react";

import { Box, Button, Flex, Heading, Text } from '@radix-ui/themes';

export const Landing = () => {
    return (
        <Flex
            direction="column"
            height="100vh"
            gap="2"
        >
            <Flex p="4" direction="column" align="end">
                <header>
                    <nav>
                        <Flex gap="4" align="baseline">
                            <Button asChild variant="ghost"><a href="/signin">Sign In</a></Button>
                            <Button asChild><a href="/signup">Sign Up</a></Button>
                        </Flex>
                    </nav>
                </header>
            </Flex>
            <Flex direction="column" align="center" justify="center" flexGrow="1">
                <main>
                    <Text align="center">
                        <Heading size={{
                            initial: '7',
                            lg: '8',

                        }}>Like It! Engage Like Never Before.</Heading>
                        <Heading as="h2" size={{
                            initial: '5',
                            lg: '7',
                        }}>Discover. Like It. Share.</Heading>
                    </Text>
                </main>
            </Flex>
            <Box p="4">
                <footer>
                    <Text as="p" align="center">© 2024 Like It! All rights reserved.</Text>
                </footer>
            </Box>
        </Flex>
    );
}

export default Landing;