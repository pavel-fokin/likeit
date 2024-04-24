import React, { createContext, useContext, useState, useEffect } from 'react';

import { AuthContext } from '../contexts/AuthContext';

// AuthProvider component
export const AuthProvider = ({ children }) => {
    const [user, setUser] = useState(null);

    // Simulate a login function
    const signIn = async (username, password) => {
        console.log('login', username)
        // Here you would implement your authentication logic, e.g., calling an API
        // This is just a placeholder logic
        if (username === 'admin' && password === 'password') {
            setUser({ id: 1, username: 'admin' });
            return true;
        }
        return false;
    };

    // Simulate a logout function
    const logout = () => {
        setUser(null);
    };

    // Check user authentication status and possibly re-authenticate
    useEffect(() => {
        // Assume a token or session check could happen here
        // setUser(...) if valid session or token exists
    }, []);

    return (
        <AuthContext.Provider value={{ user, signIn, logout }}>
            {children}
        </AuthContext.Provider>
    );
};