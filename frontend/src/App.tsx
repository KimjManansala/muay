import React from 'react';
import {
  RouterProvider,
} from "react-router-dom";
import { ChakraProvider } from '@chakra-ui/react'
import './App.css';
import { router } from './router/routes';
import Navbar from './components/Navbar';



function App() {
  return (
    
    <div className="App">
      <ChakraProvider>
      <Navbar />
      <RouterProvider router={router} />
      </ChakraProvider>
    </div>
    
  );
}

export default App;
