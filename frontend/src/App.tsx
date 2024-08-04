
import { Button, Container, Stack } from '@chakra-ui/react'
import Navbar from './components/Navbar'
import TodoForm from './features/TodoForms'
import TodoList from './features/TodoList'

function App() { 
  return (
     <Stack h="100vh">
      <Navbar />
      <Container>        
        <TodoForm></TodoForm>
        <TodoList></TodoList>
      </Container>
     </Stack>
  )
}

export default App
