import React from 'react';
import NavBarWrapper from './CoachPortalNav';

interface ICoachContainer {
    // Define your interface properties here
}

const CoachContainer: React.FC<ICoachContainer> = () => {
    // Implement your component logic here

    return (
        <div>
            Portal Landing page
        </div>
    );
};

export default NavBarWrapper(CoachContainer);