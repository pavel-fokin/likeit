import { useContext } from 'react';

import { AuthContext } from '../contexts/AuthContext';

// useAuth custom hook
export const useAuth = () => {
    const context = useContext(AuthContext);
    if (context === null) {
        throw new Error('useAuth must be used within an AuthProvider');
    }
    return context;
};