import cron from 'node-cron';
import JobData from "../model/job-data";
import Runnable from "../command/runnable";
import CommandLine from "../command/command-line";
import SchedulingResult from "../model/scheduling-result";

class CommandLineCronJob {
    private job: JobData;
    private task: Runnable;
    private scheduledInstance: any;

    constructor(job: JobData) {
        this.job = job;
        this.task = new CommandLine(this.job.action);
    }

    public getJobKey(): string {
        return this.job.jobName;
    }

    public getStatus(): string {
        if(this.scheduledInstance) {
            return this.scheduledInstance.status;
        } else {
            return 'unscheduled';
        }

    }

    public getSchedulingResult() : SchedulingResult {
        const result = new SchedulingResult();

        result.error = false;
        result.job = this.job;
        result.status = this.getStatus();

        return result;
    }

    /**
     * Start the CronJob
     */
    public async start()  {
        this.scheduledInstance = cron.schedule(this.job.time, () => {
            console.log('Firing jobName',this.job.jobName);
            this.task.run();
        });
    }

    /**
     * Stop the CronJob by destroying from the underlying engine.
     */
    public stop() {
        this.scheduledInstance.stop();
    }

    public remove(){
        this.scheduledInstance.destroy();
    }

    public logInstanceData() {
        if(this.scheduledInstance) {
            console.log(this.scheduledInstance);
        }
    }

}

export default CommandLineCronJob;
