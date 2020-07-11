import express from 'express';
import SchedulerManager from "../scheduler/scheduler-manager";
import JobData from "../model/job-data";


class AppRouter {

    public router;
    private scheduler: SchedulerManager;

    constructor () {
        this.scheduler = new SchedulerManager();
        this.router = express.Router();
        this.mountRoutes();
    }

    private mountRoutes (): void {

        this.router.post('/scheduler', (req,res) => {
            const params = req.body;
            const aJob = new JobData(params.jobName,params.time,params.action);
            const schedulingResult = this.scheduler.scheduleJob(aJob);
            res.status(schedulingResult.error ? '400':'200').json(schedulingResult);
        });

        this.router.delete('/scheduler', (req,res) => {
            const params = req.body;
            this.scheduler.removeJob(params.jobName);
            res.status('200').send('OK');
        });

        this.router.get('/scheduler', (req, res) => {
            const scheduledTasks = this.scheduler.getScheduledJobs();
            res.status('200').json(scheduledTasks);
        })

    }

}

export default AppRouter;
