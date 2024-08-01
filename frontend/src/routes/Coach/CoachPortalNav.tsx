import { Container } from "@chakra-ui/react";
import Navbar from "../../components/Navbar";
import { COACH_ATHLETES } from "../../router/routeNames";

const NavBarWrapper = (WrappedComponent: React.FC<any>)  =>{
    return function LayoutComponent(props: any) {
      return (
        <div>
          <Navbar links={[
                {
                    route: COACH_ATHLETES,
                    label: 'Athletes'
                },
            ]} />
        <Container size='lg'>
            <WrappedComponent {...props} />
        </Container>
        </div>
      );
    };
  }
  
  export default NavBarWrapper