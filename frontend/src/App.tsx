import React from 'react';
import {
  RouterProvider,
} from "react-router-dom";
import { ChakraProvider } from '@chakra-ui/react'
import './App.css';
import { router } from './router/routes';




function App() {
  return (
    
    <div className="App">
      <ChakraProvider>
      <RouterProvider router={router} />
      </ChakraProvider>
    </div>
    
  );
}

export default App;
