import React from 'react';
import Navbar from '../../components/Navbar';
import { COACH_ATHLETES } from '../../router/routeNames';

interface ICoachContainer {
    // Define your interface properties here
}

const CoachContainer: React.FC<ICoachContainer> = () => {
    // Implement your component logic here

    return (
        <div>
            <Navbar links={[
                {
                    route: COACH_ATHLETES,
                    label: 'Athletes'
                },
            ]} />
        </div>
    );
};

export default CoachContainer;