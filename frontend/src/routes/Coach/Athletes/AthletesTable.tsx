import React from 'react';
import NavBarWrapper from '../CoachPortalNav';
import { TableContainer, Table, TableCaption, Thead, Tr, Th, Tbody, Td, Tfoot } from '@chakra-ui/react';

interface IAthletesTable {
    // Define your interface properties here
}

const AthletesTable: React.FC<IAthletesTable> = () => {
    // Implement your component logic here

    return (
        <>
        <TableContainer>
            <Table variant='simple'>
                <TableCaption>List of your athletes</TableCaption>
                <Thead>
                <Tr>
                    <Th>Name</Th>
                    <Th>Status</Th>
                    <Th>Category</Th>
                </Tr>
                </Thead>
                <Tbody>
                <Tr>
                    <Td>inches</Td>
                    <Td>millimetres (mm)</Td>
                    <Td isNumeric>25.4</Td>
                </Tr>
                <Tr>
                    <Td>feet</Td>
                    <Td>centimetres (cm)</Td>
                    <Td isNumeric>30.48</Td>
                </Tr>
                <Tr>
                    <Td>yards</Td>
                    <Td>metres (m)</Td>
                    <Td isNumeric>0.91444</Td>
                </Tr>
                </Tbody>
                <Tfoot>
                <Tr>
                    <Th>To convert</Th>
                    <Th>into</Th>
                    <Th isNumeric>multiply by</Th>
                </Tr>
                </Tfoot>
            </Table>
        </TableContainer>
        </>
    );
};

export default NavBarWrapper(AthletesTable);