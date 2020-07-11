import JobData from "../model/job-data";
import Storage from "../store/storage";
import CommandLineCronJob from "../job/command-line-cron-job";
import SchedulingResult from "../model/scheduling-result";
import cron from 'node-cron';

/**
 * Class Representing a simple Scheduler Engine
 * with basic operations.
 */
class Scheduler {
    private jobStore : Storage;

    constructor(){
        this.jobStore = new Storage();
    }

    /**
     * Validate the cron expression.
     *
     * return true if the instance is created with a valid
     */
    private validate(aJob: JobData): boolean {
        return cron.validate(aJob.time);
    }

    /**
     *
     * @param aJob
     * @param persist if true the job must be persisted too.
     */
    public scheduleJob(aJob: JobData, persist: boolean) : SchedulingResult {
        let result = new SchedulingResult();

        if( this.jobStore.getJob(aJob.jobName) != null ) {
            result.error = true;
            result.errorMessage = 'you cannot schedule a JobData with name '+aJob.jobName+' as it is already present into the scheduler';
            return result;
        }

        if(this.validate(aJob)){
            let scheduledJob = new CommandLineCronJob(aJob);
            scheduledJob.start();
            this.jobStore.save(scheduledJob, persist);

            result.error = false;
            result.job = aJob;
            result.status = scheduledJob.getStatus();
            console.log('job '+aJob.jobName+' added into scheduler instance');
        } else {
            result.error = true;
            result.errorMessage = 'cron expression '+aJob.time+' it is not a valid one';
            console.error(result.errorMessage);
        }

        return result;
    }

    public getScheduledJobs(): Array<SchedulingResult> {
        return this.jobStore.getJobs();
    }

    /**
     * Destroy the JobData from the scheduling engine and delete it from the
     * job internal store.
     *
     * @param jobName
     */
    public removeScheduledJob(jobName: string) {
        const scheduledJob: CommandLineCronJob = this.jobStore.getJob(jobName);
        if(scheduledJob !== null) {
            scheduledJob.remove();
            this.jobStore.removeJob(jobName);
        } else {
            console.log('Could not remove a non existing JobData ', jobName);
        }
    }
}

export default Scheduler;
