import React from "react";

import { Flex, Button, Text, TextField, Link, Container } from "@radix-ui/themes";

export const SignUp = () => {
    return (
        <Container size="1">
            <Flex direction="column" gap="4">
                <h1>Sign Up</h1>
                <TextField.Root size="3" placeholder="Your username" />
                <TextField.Root size="3" placeholder="Your password" />
                <Button size="4">Create an Account</Button>
                <Text align="center">
                    Already have an account?  <Link href="/signin">Sign In</Link>
                </Text>
            </Flex>
        </Container>

    );
};

export default SignUp;