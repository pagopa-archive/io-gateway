import  express from 'express';
import AppRouter from "./routes/router";

class App {
    public express

    constructor () {
        this.express = express();
        this.express.use(express.json());
        this.mountRoutes();
    }

    private mountRoutes (): void {
        const appRouter = new AppRouter().router;
        this.express.use('/', appRouter);
    }
}

export default new App().express;
