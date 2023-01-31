import useRoutes from './routes/index'
import { Server } from './server'

// INITIALIZED THE SERVER
const server = new Server()
// ROUTES
useRoutes(server._app)
// START TO LISTEN ON PORT
server.start()
