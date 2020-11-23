import Scheduler from "./scheduler";
import JobData from "../model/job-data";
import SchedulingResult from "../model/scheduling-result";
import SchedulerConfigurator from "../config/scheduler-configurator";

class SchedulerManager {

    private scheduler: Scheduler;

    /**
     * @TODO Place code to re-initialize the scheduler from persisted configuration.
     */
    constructor() {
        this.scheduler = new Scheduler();
        this.restoreConfiguredJob();
    }

    /**
     * Insert into the Scheduler all the
     * previously configured Job.
     */
    private restoreConfiguredJob(){
        const configurator = new SchedulerConfigurator();
        const jobArray: Array<any> = configurator.retrievePersistedJobs();

        if(jobArray && jobArray.length > 0) {
            console.log('Initializing persisted Jobs');
            jobArray.forEach( (aJob) => {
                this.scheduler.scheduleJob(aJob, false);
            })
        }
    }

    /**
     * Schedule a JobData into the internal simple scheduler.
     *
     * @param aJob
     */
    public scheduleJob(aJob: JobData): SchedulingResult {
        return this.scheduler.scheduleJob(aJob, true);
    }

    public getScheduledJobs(): Array<SchedulingResult> {
        return this.scheduler.getScheduledJobs();
    }

    public stopJob(jobName: string){

    }

    public startJob(jobName: string){

    }

    public removeJob(jobName: string){
        this.scheduler.removeScheduledJob(jobName);
    }
}

export default SchedulerManager;
