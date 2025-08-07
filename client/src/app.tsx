import { QueryClient, QueryClientProvider } from 'react-query'
import './app.css'
import { Dashboard } from './components/dashboard/dashboard'

const queryClient = new QueryClient();

export function App() {
  return (
    <QueryClientProvider
      client={queryClient}
    >
      <div style={{ width: 400 }}>
        <Dashboard />
      </div>
    </QueryClientProvider>
  )
}
