import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

import { Flex, Button, Text, TextField, Link, Container } from "@radix-ui/themes";

import { useAuth } from "../hooks/useAuth";

export const SignUp = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const navigate = useNavigate();

    const { signUp } = useAuth();

    const onSignUp = async () => {
        await signUp(username, password);
        navigate('/app');
    }

    return (
        <Container size="1">
            <Flex direction="column" gap="4">
                <h1>Sign Up</h1>
                <TextField.Root size="3" placeholder="Your username" onChange={e => { setUsername(e.target.value) }} />
                <TextField.Root size="3" type="password" placeholder="Your password" onChange={e => { setPassword(e.target.value) }} />
                <Button size="4" onClick={onSignUp}>Create an Account</Button>
                <Text align="center">
                    Already have an account?  <Link href="/signin">Sign In</Link>
                </Text>
            </Flex>
        </Container>

    );
};

export default SignUp;