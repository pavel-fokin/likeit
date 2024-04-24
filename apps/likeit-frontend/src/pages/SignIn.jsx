import React, { useState } from "react";

import { Card, Flex, Button, Text, TextField, Link, Container } from "@radix-ui/themes";

import { useAuth } from "../hooks/useAuth";
import { Navigate } from "react-router-dom";

export const SignIn = () => {
    const { user, signIn } = useAuth();

    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    if (user) {
        return <Navigate to="/app" />;
    }

    const onSignIn = async () => {
        await signIn(username, password);
    }

    return (
        <Container size="1">
            <Flex direction="column" gap="4">
                <h1>Sign In</h1>
                <TextField.Root size="3" placeholder="Your username" onChange={e => {setUsername(e.target.value)}}/>
                <TextField.Root size="3" placeholder="Your password" onChange={e => {setPassword(e.target.value)}} />
                <Button size="4" onClick={onSignIn}>Sign In</Button>
                <Text align="center">
                    Don't have an account?  <Link href="/signup">Sign Up</Link>
                </Text>
            </Flex>
        </Container>
    );
};

export default SignIn;