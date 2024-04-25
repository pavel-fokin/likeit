import { createContext, useEffect } from 'react';

import { useAuth } from '../hooks/useAuth';

export const AuthContext = createContext(null);

// AuthContextProvider
export const AuthContextProvider = ({ children }) => {
    const { isAuthenticated } = useAuth();

    // Check user authentication status and possibly re-authenticate
    useEffect(() => {
        // Assume a token or session check could happen here
        // setUser(...) if valid session or token exists
    }, []);

    return (
        <AuthContext.Provider value={{ isAuthenticated }}>
            {children}
        </AuthContext.Provider>
    );
};