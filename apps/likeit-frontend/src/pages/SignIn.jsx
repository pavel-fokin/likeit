import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

import { Flex, Button, Text, TextField, Link, Container } from "@radix-ui/themes";

import { useAuth } from "../hooks/useAuth";

export const SignIn = () => {
    const { signIn } = useAuth();
    const navigate = useNavigate();

    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const onSignIn = async () => {
        await signIn(username, password);
        navigate('/app');
    }

    return (
        <Container size="1">
            <Flex direction="column" gap="4">
                <h1>Sign In</h1>
                <TextField.Root size="3" placeholder="Your username" onChange={e => {setUsername(e.target.value)}}/>
                <TextField.Root size="3" type="password" placeholder="Your password" onChange={e => {setPassword(e.target.value)}} />
                <Button size="4" onClick={onSignIn}>Sign In</Button>
                <Text align="center">
                    Don't have an account?  <Link href="/signup">Sign Up</Link>
                </Text>
            </Flex>
        </Container>
    );
};

export default SignIn;