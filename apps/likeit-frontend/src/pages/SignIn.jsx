import React, { useContext, useState } from "react";
import { useNavigate } from "react-router-dom";

import { Button, Container, Flex, Link, Text, TextField } from "@radix-ui/themes";

import { AuthContext } from "../contexts/AuthContext";
import { useAuth } from "../hooks/useAuth";

export const SignIn = () => {
    const { signIn } = useAuth();
    const { setIsAuthenticated } = useContext(AuthContext);

    const navigate = useNavigate();

    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const onSignIn = async () => {
        if (await signIn(username, password)) {
            setIsAuthenticated(true);
            navigate('/app');
        }
    }

    return (
        <Container size="1" m="2">
            <Flex direction="column" gap="4">
                <h1>Sign In</h1>
                <TextField.Root name="username" autoComplete="off" size="3" placeholder="Your username" onChange={e => { setUsername(e.target.value) }} />
                <TextField.Root name="password" size="3" type="password" placeholder="Your password" onChange={e => { setPassword(e.target.value) }} />
                <Button size="4" onClick={onSignIn}>Sign In</Button>
                <Text align="center">
                    Don't have an account?  <Link href="/signup">Sign Up</Link>
                </Text>
            </Flex>
        </Container>
    );
};

export default SignIn;