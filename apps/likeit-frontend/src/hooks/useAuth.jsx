import { useState, useMemo } from "react";

import { SignIn, SignUp } from "../api/api";

// useAuth custom hook
export const useAuth = () => {
    const [accessToken, setAccessToken] = useState(localStorage.getItem("accessToken"));

    const isAuthenticated = () => localStorage.getItem("accessToken") !== null;

    const isAuthenticatedMemo = useMemo(() => isAuthenticated(), [accessToken]);

    const signIn = async (username, password) => {
        const token = await SignIn(username, password);
        if (!token) {
            return false;
        }
        setAccessToken(token);
        localStorage.setItem("accessToken", token);
        return true;
    };

    const signUp = async (username, password) => {
        const token = await SignUp(username, password);
        if (!token) {
            return false;
        }
        setAccessToken(token);
        localStorage.setItem("accessToken", token);
        return true;
    }

    const signOut = () => {
        setAccessToken(null);
        localStorage.removeItem("accessToken");
    }

    return {
        accessToken,
        isAuthenticated,
        signIn,
        signUp,
        signOut,
    };
};